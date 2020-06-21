package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	templeteString := "Lemonade stand supply "
	t, err := template.New("title").Parse(templeteString)
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
}
