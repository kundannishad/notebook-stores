package utils

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		return
	}

	defer r.Body.Close()

	// Unmarshal the request body into x
	if err := json.Unmarshal(body, x); err != nil {
		fmt.Println("Error in unmarshaling JSON body request:", err)
	}
}
