package model

type Data struct {
	ProjectRunner ProjectRunner     `json:"projectrunner"`
	LintMessages  LintMessages      `json:"lintmessages"`
	Metadata      Metadata          `json:"metadata"`
	Score         Score             `json:"score"`
	Errors        map[string]string `json:"errors,omitempty"`
}
