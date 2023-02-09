package types

type TranslateWordReq struct {
	Word string `json:"word"`
}

type TranslateWordResp struct {
	CN       string   `json:"cn,omitempty"`
	English  string   `json:"english,omitempty"`
	Examples []string `json:"examples,omitempty"`
	// QueryHistory []ts
}
