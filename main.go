package main

import (
	"net/http"

	"github.com/Go/webservices/controllers"
)

func main() {

	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
