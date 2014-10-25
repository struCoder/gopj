package app

import "net/http"
import "io"
import "fmt"

func sayHello(res http.ResponseWriter, req *http.Request) {
	fmt.Println("here")
	io.WriteString(res, "test")
}
func Router(router map[string]func(http.ResponseWriter, *http.Request)) {
	router["/"] = sayHello
}
