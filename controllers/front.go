// route a network request to a correct controller
package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	// handle any of /users/ or /user address
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

}

func encodeResponseAsJSON(data interface{},w io.Writer){
	enc := json.NewEncoder(w)
	enc.Encode(data)
}