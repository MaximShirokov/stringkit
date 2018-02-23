package stringkit

import (
	"log"
	"strings"
	"regexp"
)

func Transliterate(str string) string {
	mapTranslit := map[string]string {
		"Ґ":"G","Ё":"YO","Є":"E","Ї":"YI","І":"I","А":"A","Б":"B","В":"V","Г":"G","Д":"D","Е":"E",
		"Ж":"ZH","З":"Z","И":"I","Й":"Y","К":"K","Л":"L","М":"M","Н":"N","О":"O","П":"P","Р":"R","С":"S",
		"Т":"T","У":"U","Ф":"F","Х":"H","Ц":"TS","Ч":"CH","Ш":"SH","Щ":"SCH","Ъ":"'","Ы":"Y","Ь":"","Э":"E",
		"Ю":"YU","Я":"YA","ЬЕ":"IE","ЬЁ":"IE","і":"i","ґ":"g","ё":"yo","№":"#","є":"e",
		"ї":"yi","а":"a","б":"b","в":"v","г":"g","д":"d","е":"e","ж":"zh","з":"z","и":"i",
		"й":"y","к":"k","л":"l","м":"m","н":"n","о":"o","п":"p","р":"r","с":"s","т":"t",
		"у":"u","ф":"f","х":"h","ц":"ts","ч":"ch","ш":"sh","щ":"sch","ъ":"'","ы":"y","ь":"",
		"э":"e","ю":"yu","я":"ya","ье":"ie","ьё":"ie"}

	strParts := strings.Split(SqueezeWhiteSpaces(str), " ")
	
	var translitStrParts []string
	var result string

	for _, part := range strParts {
		var translitPart string

		for _, char := range part {
			translitPart += mapTranslit[string(char)]
		}

		translitStrParts = append(translitStrParts, translitPart)
	}

	result = strings.Join(translitStrParts, " ")

	return result
}

func SqueezeWhiteSpaces(str string) string {
	reLeadcloseWhtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	reInsideWhtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	
    result := reLeadcloseWhtsp.ReplaceAllString(str, "")
	result = reInsideWhtsp.ReplaceAllString(result, " ")
	
	return result
}

func removeNonAlphaNum(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	
    if err != nil {
        log.Fatal(err)
	}
	
    return reg.ReplaceAllString(str, "")
}