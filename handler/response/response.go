package response

type responseBody struct {
	Error string `json:"error"`
}

func buildErrorResponse(message string) responseBody {
	return responseBody{
		Error: message,
	}
}
