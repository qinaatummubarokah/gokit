package reportcsv

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	GetDataRequest struct {
		Id string `json:"id"`
	}
	GetDataResponse struct {
		Data []Csvdata `json:"data"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
