package password

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var request Request
	json.Unmarshal(reqBody, &request)

	json.NewEncoder(w).Encode(Response{Token: "random-token"})
}
