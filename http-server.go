package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	fmt.Fprintf(w, "<h1>Hello, %s</h1>", name)

}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	x := r.FormValue("x")
	xnum, err := strconv.ParseFloat(x, 64)
	if err != nil {
		fmt.Println("error")
	}
	y := r.FormValue("y")
	ynum, err := strconv.ParseFloat(y, 64)
	if err != nil {
		fmt.Println("error")
	}

	fmt.Fprintf(w, "X + Y = %f", xnum+ynum)
}

func main() {
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/add", handleAdd)

	//serve *.html
	http.Handle("/www/",
		http.StripPrefix("/www/",
			http.FileServer(http.Dir("."))))

	http.ListenAndServe(":8000", nil)
}
