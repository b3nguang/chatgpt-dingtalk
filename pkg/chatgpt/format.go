package chatgpt

import (
	"regexp"
	"strings"
)

func formatAnswer(answer string) string {
	// 适配 deepseek r1
	answer = strings.TrimSpace(answer)

	re := regexp.MustCompile(`(?s)<think>.*?</think>`)
	answer = re.ReplaceAllString(answer, "")

	answer = strings.ReplaceAll(answer, "<think>", "")
	answer = strings.ReplaceAll(answer, "</think>", "")

	answer = strings.TrimSpace(answer)

	return answer
}
