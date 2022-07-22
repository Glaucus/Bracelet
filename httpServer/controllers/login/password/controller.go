package password

import (
	"github.com/brianvoe/sjwt"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var request Request
	json.Unmarshal(reqBody, &request)

	// Set Claims
	claims := sjwt.New()
	claims.Set("username", request.Username)
	claims.Set("user_id", -1)

	// Generate jwt
	secretKey := []byte(os.Getenv("JWT_SIGNING_KEY"))
	jwt := claims.Generate(secretKey)

	json.NewEncoder(w).Encode(Response{Token: jwt})
}
