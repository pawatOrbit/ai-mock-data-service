package model

type GenerateMockDataWithOneTableRequest struct {
	TableName string `json:"table_name"`
	NumSample int    `json:"num_samples" validate:"true"`
}

type GenerateMockDataWithOneTableResponseData struct {
	Query            string  `json:"query"`
	PromptTokens     int     `json:"prompt_tokens"`
	CompletionTokens int     `json:"completion_tokens"`
	TotalTokens      int     `json:"total_tokens"`
	TimeTaken        float64 `json:"time_taken"`
}

type GenerateMockDataWithOneTableResponse struct {
	Status int                                      `json:"status"`
	Data   GenerateMockDataWithOneTableResponseData `json:"data"`
}
