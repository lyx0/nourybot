package commands

import (
	"fmt"
)

var cm = map[string]string{
	"`": "ё",
	"~": "Ё",
	"=": "ъ",
	"+": "Ъ",
	"[": "ю",
	"]": "щ",
	`\`: "э",
	"{": "Ю",
	"}": "Щ",
	"|": "Э",
	";": "ь",
	":": "Ь",
	"'": "ж",
	`"`: "Ж",

	"q": "я",
	"w": "ш",
	"e": "е",
	"r": "р",
	"t": "т",
	"y": "ы",
	"u": "у",
	"i": "и",
	"o": "о",
	"p": "п",
	"a": "а",
	"s": "с",
	"d": "д",
	"f": "ф",
	"g": "г",
	"h": "ч",
	"j": "й",
	"k": "к",
	"l": "л",
	"z": "з",
	"x": "х",
	"c": "ц",
	"v": "в",
	"b": "б",
	"n": "н",
	"m": "м",

	"Q": "Я",
	"W": "Ш",
	"E": "Е",
	"R": "Р",
	"T": "Т",
	"Y": "Ы",
	"U": "У",
	"I": "И",
	"O": "О",
	"P": "П",
	"A": "А",
	"S": "С",
	"D": "Д",
	"F": "Ф",
	"G": "Г",
	"H": "Ч",
	"J": "Й",
	"K": "К",
	"L": "Л",
	"Z": "З",
	"X": "Х",
	"C": "Ц",
	"V": "В",
	"B": "Б",
	"N": "Н",
	"M": "М",
}

// Phonetic replaces the characters in a given message with the characters on a
// russian phonetic keyboard.
func Phonetic(message string) (string, error) {
	var ts string

	for _, c := range message {
		if _, ok := cm[string(c)]; ok {
			ts = ts + cm[string(c)]
		} else {
			ts = ts + string(c)

		}
	}

	return fmt.Sprint(ts), nil
}
