package exception

import (
	"encoding/json"
	"net/http"

	"github.com/nuraziz04/golang-esta/model"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {

	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := model.Response{
			Message: exception.Error,
		}

		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)

		encoder.Encode(response)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	response := model.Response{
		Message: "Internal server error",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)

	encoder.Encode(response)
}
