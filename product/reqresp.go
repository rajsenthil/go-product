package product

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type (
	GetProductRequest struct {
		id int32 `json:"id"`
	}

	GetProductResponse struct {
		Product Product `json:"product"`
	}
)

func decodeIdReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetProductRequest
	vars := mux.Vars(r)

	i, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		log.Println("Query param for product id is invalid")
		return nil, err
	}
	req = GetProductRequest{
		id: int32(i),
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.Println("Response: ", response)
	//json := json.NewEncoder(w).Encode(response)
	//log.Println("JSON Response: ", json)
	//return json
	return json.NewEncoder(w).Encode(response)
}
