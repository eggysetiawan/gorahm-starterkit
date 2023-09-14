package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func writeResponse(w http.ResponseWriter, header string, code int, data interface{}) {
	w.Header().Add("Content-Type", header)

	w.WriteHeader(code)

	if header == "application/xml" {
		if err := xml.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}

	}
}
