package delivery

type (
	Response struct {
		RC string       `json:"responseCode"`
		RD ResponseData `json:"responseData"`
	}

	ResponseData struct {
		Data         interface{}
		Description  string `json:"description"`
		ErrorMessage string `json:"error"`
	}
)
