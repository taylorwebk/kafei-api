package structs

// Response Base Structure for responses
type Response struct {
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}
