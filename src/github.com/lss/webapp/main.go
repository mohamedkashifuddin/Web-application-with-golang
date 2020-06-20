package main

import (
	"net/http"
)

func main() {
	//repopath := "/~public"
	http.ListenAndServe(":3006", http.FileServer(http.Dir("/Users/Admin/Desktop/go-workspace/Web-application-with-golang/public/")))
}
