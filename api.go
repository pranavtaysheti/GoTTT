package main

import (
	"net/http"
	"encoding/json"
	"log"
)

func JsonResponseHandler(w http.ResponseWriter, d any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
