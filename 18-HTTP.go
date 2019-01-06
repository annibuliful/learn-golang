package main
import (
	"net/http"
  "fmt"
)

func hello(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(res,"<h1>Hello</h1>")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9000", nil)
}
