package model

const ScoreName = "score"

// Score stores the overall rank and raw score computed from criterias
type Score struct {
	Value   float64              `json:"value"`
	Rank    string               `json:"rank"`
	Details []*EvaluatorResponse `json:"details,omitempty"`
}

// EvaluatorResponse
type EvaluatorResponse struct {
	Name    string               `json:"name"`
	Score   float64              `json:"score"`
	Weight  float64              `json:"weight"`
	Desc    string               `json:"desc,omitempty"`
	Message string               `json:"msg"`
	URL     string               `json:"url,omitempty"`
	Details []*EvaluatorResponse `json:"details,omitempty"`
}
