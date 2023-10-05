// The whole catbox upload functionality has been copied from
// here so that I could use it with litterbox:
// https://github.com/wabarc/go-catbox/blob/main/catbox.go <3
//
// Copyright 2021 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"go.uber.org/zap"
)

const (
	CATBOX_ENDPOINT = "https://litterbox.catbox.moe/resources/internals/api.php"
	GOFILE_ENDPOINT = "https://store1.gofile.io/uploadFile"
	KAPPA_ENDPOINT  = "https://kappa.lol/api/upload"
	YAF_ENDPOINT    = "https://i.yaf.ee/upload"
)

type Uploader struct {
	Client       *http.Client
	Time         string
	Userhash     string
	TwitchClient *twitch.Client
	Log          *zap.SugaredLogger
}

func NewUpload(destination, fileName, target string, twitchClient *twitch.Client, log *zap.SugaredLogger) {
	client := &http.Client{
		Timeout: 300 * time.Second,
	}

	ul := &Uploader{
		Client:       client,
		Time:         "24h",
		TwitchClient: twitchClient,
		Log:          log,
	}

	switch destination {
	case "catbox":
		go ul.CatboxUpload(target, fileName)
	case "yaf":
		go ul.YafUpload(target, fileName)
	case "kappa":
		go ul.KappaUpload(target, fileName)
	case "gofile":
		go ul.GofileUpload(target, fileName)

	}
}

func (ul *Uploader) CatboxUpload(target, fileName string) {
	defer os.Remove(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	defer file.Close()
	ul.TwitchClient.Say(target, "Uploading to catbox.moe... dankCircle")

	// if size := helper.FileSize(fileName); size > 209715200 {
	// 	return "", fmt.Errorf("file too large, size: %d MB", size/1024/1024)
	// }

	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		m.WriteField("reqtype", "fileupload")
		m.WriteField("time", ul.Time)
		part, err := m.CreateFormFile("fileToUpload", filepath.Base(file.Name()))
		if err != nil {
			ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			return
		}

		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()

	req, _ := http.NewRequest(http.MethodPost, CATBOX_ENDPOINT, r)
	req.Header.Add("Content-Type", m.FormDataContentType())

	resp, err := ul.Client.Do(req)
	if err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		return
	}

	reply := string(body)
	ul.TwitchClient.Say(target, fmt.Sprintf("Removing file: %s", fileName))
	ul.TwitchClient.Say(target, reply)
}

func (ul *Uploader) GofileUpload(target, path string) {
	defer os.Remove(path)
	ul.TwitchClient.Say(target, "Uploading to gofile.io... dankCircle")
	pr, pw := io.Pipe()
	form := multipart.NewWriter(pw)

	type gofileData struct {
		DownloadPage string `json:"downloadPage"`
		Code         string `json:"code"`
		ParentFolder string `json:"parentFolder"`
		FileId       string `json:"fileId"`
		FileName     string `json:"fileName"`
		Md5          string `json:"md5"`
	}

	type gofileResponse struct {
		Status string `json:"status"`
		Data   gofileData
	}

	go func() {
		defer pw.Close()

		file, err := os.Open(path) // path to image file
		if err != nil {
			ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			os.Remove(path)
			return
		}

		w, err := form.CreateFormFile("file", path)
		if err != nil {
			ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			os.Remove(path)
			return
		}

		_, err = io.Copy(w, file)
		if err != nil {
			ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			os.Remove(path)
			return
		}

		form.Close()
	}()

	req, err := http.NewRequest(http.MethodPost, GOFILE_ENDPOINT, pr)
	if err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		os.Remove(path)
		return
	}
	req.Header.Set("Content-Type", form.FormDataContentType())

	httpClient := http.Client{
		Timeout: 300 * time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		os.Remove(path)
		ul.Log.Errorln("Error while sending HTTP request:", err)

		return
	}
	defer resp.Body.Close()
	ul.TwitchClient.Say(target, "Uploaded PogChamp")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		os.Remove(path)
		ul.Log.Errorln("Error while reading response:", err)
		return
	}

	jsonResponse := new(gofileResponse)
	if err := json.Unmarshal(body, jsonResponse); err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		ul.Log.Errorln("Error while unmarshalling JSON response:", err)
		return
	}

	var reply = jsonResponse.Data.DownloadPage

	ul.TwitchClient.Say(target, fmt.Sprintf("Removing file: %s", path))
	ul.TwitchClient.Say(target, reply)
}

func (ul *Uploader) KappaUpload(target, path string) {
	defer os.Remove(path)
	ul.TwitchClient.Say(target, "Uploading to kappa.lol... dankCircle")
	pr, pw := io.Pipe()
	form := multipart.NewWriter(pw)

	type kappaResponse struct {
		Link string `json:"link"`
	}

	go func() {
		defer pw.Close()

		err := form.WriteField("name", "xd")
		if err != nil {
			ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			os.Remove(path)
			return
		}

		file, err := os.Open(path) // path to image file
		if err != nil {
			ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			os.Remove(path)
			return
		}

		w, err := form.CreateFormFile("file", path)
		if err != nil {
			ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			os.Remove(path)
			return
		}

		_, err = io.Copy(w, file)
		if err != nil {
			ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			os.Remove(path)
			return
		}

		form.Close()
	}()

	req, err := http.NewRequest(http.MethodPost, KAPPA_ENDPOINT, pr)
	if err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		os.Remove(path)
		return
	}
	req.Header.Set("Content-Type", form.FormDataContentType())

	httpClient := http.Client{
		Timeout: 300 * time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		os.Remove(path)
		ul.Log.Errorln("Error while sending HTTP request:", err)

		return
	}
	defer resp.Body.Close()
	ul.TwitchClient.Say(target, "Uploaded PogChamp")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		os.Remove(path)
		ul.Log.Errorln("Error while reading response:", err)
		return
	}

	jsonResponse := new(kappaResponse)
	if err := json.Unmarshal(body, jsonResponse); err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		ul.Log.Errorln("Error while unmarshalling JSON response:", err)
		return
	}

	var reply = jsonResponse.Link

	ul.TwitchClient.Say(target, fmt.Sprintf("Removing file: %s", path))
	ul.TwitchClient.Say(target, reply)
}
func (ul *Uploader) YafUpload(target, path string) {
	defer os.Remove(path)
	ul.TwitchClient.Say(target, "Uploading to yaf.ee... dankCircle")
	pr, pw := io.Pipe()
	form := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()

		err := form.WriteField("name", "xd")
		if err != nil {
			ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			os.Remove(path)
			return
		}

		file, err := os.Open(path) // path to image file
		if err != nil {
			ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			os.Remove(path)
			return
		}

		w, err := form.CreateFormFile("file", path)
		if err != nil {
			ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			os.Remove(path)
			return
		}

		_, err = io.Copy(w, file)
		if err != nil {
			ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
			os.Remove(path)
			return
		}

		form.Close()
	}()

	req, err := http.NewRequest(http.MethodPost, YAF_ENDPOINT, pr)
	if err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		os.Remove(path)
		return
	}
	req.Header.Set("Content-Type", form.FormDataContentType())

	httpClient := http.Client{
		Timeout: 300 * time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		os.Remove(path)
		ul.Log.Errorln("Error while sending HTTP request:", err)

		return
	}
	defer resp.Body.Close()
	ul.TwitchClient.Say(target, "Uploaded PogChamp")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ul.TwitchClient.Say(target, fmt.Sprintf("Something went wrong FeelsBadMan: %q", err))
		os.Remove(path)
		ul.Log.Errorln("Error while reading response:", err)
		return
	}

	var reply = string(body[:])

	ul.TwitchClient.Say(target, fmt.Sprintf("Removing file: %s", path))
	ul.TwitchClient.Say(target, reply)
}
