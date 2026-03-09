package response

import "net/http"

func Error(w http.ResponseWriter, status int, message string, errors interface{}) {
	resp := Response{
		Success: false,
		Message: message,
		Errors:  errors,
	}

	JSON(w, status, resp)
}
