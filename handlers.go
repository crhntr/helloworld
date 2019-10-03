package helloworld

import (
	"fmt"
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Hello, world!")
}
