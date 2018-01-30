package structs

// Token ad to response content if the response needs a token
type Token struct {
	Token string      `json:"token"`
	Data  interface{} `json:"data"`
}
