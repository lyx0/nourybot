package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type banphraseResponse struct {
	Banned        bool          `json:"banned"`
	InputMessage  string        `json:"input_message"`
	BanphraseData banphraseData `json:"banphrase_data"`
}

type banphraseData struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Phrase    string `json:"phrase"`
	Length    int    `json:"length"`
	Permanent bool   `json:"permanent"`
}

// CheckMessage checks if a message contains
// banphrased content.
// If a message is allowed it returns
// false, "okay"
// If a message is not allowed it returns:
// true, "[banphrased] monkaS"

func CheckMessage(text string) (bool, string) {
	log.Info("fn CheckMessage")

	reqBody, err := json.Marshal(map[string]string{
		"message": text,
	})
	if err != nil {
		log.Error(err)
	}

	log.Info(reqBody)

	resp, err := http.Post("https://pajlada.pajbot.com/api/v1/banphrases/test", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	var responseObject banphraseResponse
	json.Unmarshal(body, &responseObject)

	reason := responseObject.BanphraseData.Name

	log.Info("Bancheck: ", responseObject.Banned)
	log.Info("Reason: ", reason)

	log.Info("Bancheck: ", responseObject.Banned)
	if responseObject.Banned {
		return true, fmt.Sprintf("Banphrased, reason: %s", reason)
	} else {
		return false, "okay"
	}

	return true, "couldnt check"
}
