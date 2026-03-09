package response

import "net/http"

func Success(w http.ResponseWriter, status int, message string, data interface{}) {
	resp := Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	JSON(w, status, resp)
}
