package handlers

import (
	"blockchain/chainS"
	"encoding/json"
	"io"
	"net/http"
)

func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	byte, err := json.MarshalIndent(chainS.Blockchain, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(byte))
}
