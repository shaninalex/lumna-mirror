package domain

type ApiResponse struct {
	Status   bool     `json:"status"`
	Data     any      `json:"data"`
	Messages []string `json:"messages,omitempty"`
}

func NewApiResponse(data any) *ApiResponse {
	return &ApiResponse{
		Status: true,
		Data:   data,
	}
}
