package endpoints

import (
	"encoding/json"
	"net/http"
)

//GetEngineVersionEndpoint Returns version of access engine
func GetEngineVersionEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("0.0")
}

//GetPortAddressEndpoint Returns the port on which the server is currently running
func GetPortAddressEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("0.0")
}
