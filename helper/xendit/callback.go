package helper_xendit

import (
	"encoding/json"
	"errors"
	"net/http"
)

func GetCallback(domain interface{}) (interface{}, error) {
	callback_otp := "BjVVRO8eKgceve38jmqm6twtK9YLjtAfk7CbJLxfiToTilHX"
	var req *http.Request
	var responseWriter http.ResponseWriter

	decoder := json.NewDecoder(req.Body)
	callbackData := domain

	err := decoder.Decode(&callbackData)
	if err != nil {
		return "empty", errors.New("internal status error")
	}

	defer req.Body.Close()

	callback, _ := json.Marshal(callbackData)

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Header().Set("x-callback-token", callback_otp)

	responseWriter.WriteHeader(200)
	return callback, nil
}
