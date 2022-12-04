package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thanhpk/randstr"
	"math/rand"
	"net/http"
)

func PrintStylesFromFile(resi http.ResponseWriter, reqi *http.Request) {
	resi.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(resi, "<style>")
	fmt.Fprintf(resi, ".the-grid{\n            "+
		"display: inline-grid;\n            "+
		"grid-template-columns: 250px 250px 250px;\n        }")
	fmt.Fprintf(resi, "</style>")

	fmt.Fprintf(resi, "<style>")
	fmt.Fprintf(resi, ".dot-0 {\n            "+
		"height: 30vh;\n            "+
		"width: 30vh;\n            "+
		"border-radius: 50%%;\n            "+
		"background-color: #bbb\n"+
		"display: inline-block;\n            "+
		"justify-content: center;\n            "+
		"justify-items: center;\n            "+
		"text-align: center;\n        }")
	fmt.Fprintf(resi, "</style>")

	fmt.Fprintf(resi, "<style>")
	fmt.Fprintf(resi, ".dot-1 {\n            "+
		"height: 30vh;\n            "+
		"width: 30vh;\n            "+
		"display: inline-block;\n            "+
		"text-align: center;\n        }")

	fmt.Fprintf(resi, "</style>")

	fmt.Fprintf(resi, "<h1>Welcome to GoMosaic</h1>")

}
func PrintCircleHTML(resi http.ResponseWriter, reqi *http.Request) {
	letters := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
		"T", "U", "V", "W", "X", "Y", "Z",
	}
	fmt.Fprintf(resi, "<div class=\"the-grid\">")
	fmt.Fprintf(resi, "<div class='dot-%d'style=\"background-color:%s\">%s</div>", rand.Intn((2-1)+1), randstr.Hex(3), letters[rand.Intn(len(letters))])
	fmt.Fprintf(resi, "<div class='dot-%d'style=\"background-color:%s\">%s</div>", rand.Intn((2-1)+1), randstr.Hex(3), letters[rand.Intn(len(letters))])
	fmt.Fprintf(resi, "<div class='dot-%d'style=\"background-color:%s\">%s</div>", rand.Intn((2-1)+1), randstr.Hex(3), letters[rand.Intn(len(letters))])
	fmt.Fprintf(resi, "<div class='dot-%d'style=\"background-color:%s\">%s</div>", rand.Intn((2-1)+1), randstr.Hex(3), letters[rand.Intn(len(letters))])
	fmt.Fprintf(resi, "<div class='dot-%d'style=\"background-color:%s\">%s</div>", rand.Intn((2-1)+1), randstr.Hex(3), letters[rand.Intn(len(letters))])
	fmt.Fprintf(resi, "<div class='dot-%d'style=\"background-color:%s\">%s</div>", rand.Intn((2-1)+1), randstr.Hex(3), letters[rand.Intn(len(letters))])
	fmt.Fprintf(resi, "<div class='dot-%d'style=\"background-color:%s\">%s</div>", rand.Intn((2-1)+1), randstr.Hex(3), letters[rand.Intn(len(letters))])
	fmt.Fprintf(resi, "<div class='dot-%d'style=\"background-color:%s\">%s</div>", rand.Intn((2-1)+1), randstr.Hex(3), letters[rand.Intn(len(letters))])
	fmt.Fprintf(resi, "<div class='dot-%d'style=\"background-color:%s\">%s</div>", rand.Intn((2-1)+1), randstr.Hex(3), letters[rand.Intn(len(letters))])
	fmt.Fprintf(resi, "</div>")
	fmt.Fprintf(resi, "<body>")
	fmt.Fprintf(resi, "<button onclick=\"window.location.reload();\">Randomize</button>")
	fmt.Fprintf(resi, "</body>")
}
func HttpRequestHandlerWithMosaic(resi http.ResponseWriter, reqi *http.Request) {
	resi.Header().Set("Content-Type", "text/html")
	PrintStylesFromFile(resi, reqi)
	PrintCircleHTML(resi, reqi)

	//fmt.Fprintf(resi, "<div class='dot-%d'>W</div>", rand.Intn(2))

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/Mosaic", HttpRequestHandlerWithMosaic).Methods("GET")
	err := http.ListenAndServe(":80", router)
	if err != nil {
		errors.New("ERROR")

	}
}
