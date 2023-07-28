package QuestionSchema

type Description struct {
	DescriptionText        string `json:"description_text"`
	DescriptionType        string `json:"description_type"`
	DescriptionTrueString  string `json:"description_true_string"`
	DescriptionFalseString string `json:"description_false_string"`
	Priority               int64
}

type Question struct {
	Title      string
	BranchType string
}
