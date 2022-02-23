package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Test struct {
	Name string
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!!!")
}

func Index(w http.ResponseWriter, r *http.Request) {
	test := Test{
		"Hello World Server",
	}
	myTemplate := template.Must(template.ParseFiles("src/web_simple/index.html"))
	myTemplate.Execute(w, test)

	myJson, err := json.Marshal(test)
	if err != nil {
		panic("Fail to generate Json")
	}
	fmt.Println("Json : " + string(myJson))
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/home", Index)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Starting server...")
	HandleRequest()
}
