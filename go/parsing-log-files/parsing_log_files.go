package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	// panic("Please implement the IsValidLine function")
	requiredText := strings.Split(text, " ")
	validString := []string{
		"[TRC]",
		"[DBG]",
		"[INF]",
		"[WRN]",
		"[ERR]",
		"[FTL]",
	}
	for _, value := range validString {
		if matched := strings.EqualFold(value, requiredText[0]); matched {
			return matched
		}
	}
	return false
}

func SplitLogLine(text string) []string {
	// panic("Please implement the SplitLogLine function")
	object, err := regexp.Compile("<[^a-zA-Z\\d\\s:]*>")
	if err != nil {
		return nil
	}

	return object.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	// panic("Please implement the CountQuotedPasswords function")
	count := 0
	object, err := regexp.Compile(`".*(?i)password"`)
	if err != nil {
		return 0
	}

	for _, value := range lines {
		if object.MatchString(value) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	// panic("Please implement the RemoveEndOfLineText function")
	object, err := regexp.Compile("end-of-line\\d+")
	if err != nil {
		return ""
	}
	output := object.ReplaceAllString(text, "")
	return output
}

func TagWithUserName(lines []string) []string {

	re, err := regexp.Compile(`(?:User\s+([A-z0-9]+))`)

	if err != nil {
		panic(err)
	}

	for index, line := range lines {
		if re.MatchString(line) {
			submatch := re.FindStringSubmatch(line)
			fmt.Println(submatch)
			lines[index] = fmt.Sprintf("[USR] %s %s", submatch[1], line)
		}
	}

	return lines
	// panic("Please implement the TagWithUserName function")
}
