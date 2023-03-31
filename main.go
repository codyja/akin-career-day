package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Title         string
	ContainerName string
	BG_Color      string
}

func main() {
	template := template.Must(template.ParseFiles("./templates/index.html"))

	fs := http.FileServer(http.Dir("./static/css/"))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:         "Akin Career Day",
			BG_Color:      os.Getenv("BG_COLOR"),
			ContainerName: os.Getenv("CONTAINER_NAME"),
		}
		template.Execute(w, data)
	})

	bindPort := os.Getenv("BIND_PORT")
	if bindPort == "" {
		log.Fatal("BIND_PORT environment variable is not set")
	}

	fmt.Printf("Starting server at port %s\n", bindPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", bindPort), nil); err != nil {
		log.Fatal(err)
	}
}
