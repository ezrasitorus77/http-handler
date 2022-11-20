package helper

import (
	"encoding/json"
	"net/http"

	"github.com/ezrasitorus77/http-handler/domain/delivery"
)

func Send(w http.ResponseWriter, httpStatus int, rc string, respData delivery.ResponseData) {
	var (
		jsonEncoder *json.Encoder
		resp        delivery.Response = delivery.Response{
			RC: rc,
			RD: respData,
		}
	)

	w.WriteHeader(httpStatus)

	jsonEncoder = json.NewEncoder(w)
	jsonEncoder.Encode(resp)
}
