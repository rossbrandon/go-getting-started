package main

import (
	"github.com/rosstafarian/go-getting-started/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	_ = http.ListenAndServe(":3000", nil)
}
