package main

import (
	"net/http"
)

func main() {
	//repopath := "/~public"
	http.ListenAndServe(":3006", http.FileServer(http.Dir("/Users/dell/Desktop/go-workspace/public")))
}
