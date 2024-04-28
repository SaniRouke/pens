package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Items []Item

type Item struct {
	Title   string
	Content string
}

func main() {
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFiles("index.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		items := Items{
			Item{
				Title:   "title1",
				Content: "content1",
			},
			Item{
				Title:   "title2",
				Content: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Eaque sit cum similique et dicta corporis fugit architecto obcaecati minima deleniti officiis maiores voluptatum quo ea molestiae, porro magnam ullam cumque?",
			},
		}

		err = tpl.ExecuteTemplate(w, "index.html", items)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
