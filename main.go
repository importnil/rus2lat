package rus2lat

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	charsMap = map[rune]string{
		'а': "a",
		'б': "b",
		'в': "v",
		'г': "g",
		'д': "d",
		'е': "e",
		'ё': "yo",
		'ж': "zh",
		'з': "z",
		'и': "i",
		'й': "j",
		'к': "k",
		'л': "l",
		'м': "m",
		'н': "n",
		'о': "o",
		'п': "p",
		'р': "r",
		'с': "s",
		'т': "t",
		'у': "u",
		'ф': "f",
		'х': "h",
		'ц': "c",
		'ч': "ch",
		'ш': "sh",
		'щ': "shch",
		'ъ': "\"",
		'ы': "y",
		'ь': "'",
		'э': "eh",
		'ю': "yu",
		'я': "ya",
	}
)

func Raw(o string) string {
	var result bytes.Buffer
	runeString := []rune(o)
	len := utf8.RuneCountInString(o)

	for i := 0; i < len; i++ {
		v := runeString[i]
		switch char, ok := charsMap[unicode.ToLower(v)]; {
		case !ok:
			result.WriteRune(v)
		case unicode.IsUpper(v):
			if (len > i+1 && unicode.IsUpper(runeString[i+1])) || (i > 0 && unicode.IsUpper(runeString[i-1])) {
				char = strings.ToUpper(char)
			} else {
				char = strings.Title(char)
			}
			fallthrough
		default:
			result.WriteString(char)
		}
	}

	return result.String()
}

func URL(o string) string {
	var result bytes.Buffer
	runeString := []rune(o)
	len := utf8.RuneCountInString(o)

	for i := 0; i < len; i++ {
		v := runeString[i]
		switch char, _ := charsMap[unicode.ToLower(v)]; {
		case unicode.IsDigit(v):
			result.WriteRune(v)
		case v == '_' || v == '~' || (v == '-' && !unicode.IsSpace(runeString[i+1])):
			result.WriteRune(v)
		case v == 'ъ' || v == 'ь':
			continue
		case unicode.IsSpace(v):
			if len > i+1 && runeString[i+1] != '-' || unicode.IsSpace(runeString[i+1]) {
				i++
				continue
			}

			result.WriteString("-")
		case 'a' <= v && v <= 'z' || 'A' <= v && v <= 'Z':
			result.WriteRune(unicode.ToLower(v))
		default:
			result.WriteString(char)
		}
	}

	return result.String()
}
