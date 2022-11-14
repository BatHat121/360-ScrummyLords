package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type InnerJson struct {
	VarTypeName int
}

var Randomarray []int

func RNGarray(n int) {
	Randomarray = make([]int, n, n)
	for i := 0; i < len(Randomarray); i++ {
		Randomarray[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Randomarray), func(i, j int) { Randomarray[i], Randomarray[j] = Randomarray[j], Randomarray[i] })
	fmt.Print(Randomarray)
}

func GetHttpRequest(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	answerjson, err := os.Open("package.json")
	if err != nil {
		fmt.Println(err)
	}
	defer answerjson.Close()
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(res, "<h1>Quiz Master</h>")
	readit, err := io.ReadAll(answerjson)
	var JsonBank Quizbankentry
	json.Unmarshal(readit, &JsonBank)
	for i := 0; i < len(Randomarray); i++ {
		fmt.Fprintf(res, "<br>Question %d : ", i+1)
		fmt.Fprintf(res, "%s", JsonBank.EntryPoint[BankID-1].BankQuestions[Randomarray[i]])
		for j := 0; i < len(Randomarray); j++ {
		}
		fmt.Fprintf(res, "<br>")
		fmt.Fprintf(res, "%s", JsonBank.EntryPoint[BankID-1].PossibleAnswers[Randomarray[i]])
	}

}
func EntryPoint2(resi http.ResponseWriter, reqi *http.Request) {
	resi.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(resi, "<h1>Hello Answers</h1>")
	answerjson, err := os.Open("package.json")
	if err != nil {
		fmt.Println("error")
	}
	defer answerjson.Close()
	readit, err := io.ReadAll(answerjson)
	var JsonBank Quizbankentry
	json.Unmarshal(readit, &JsonBank)
	//fmt.Fprintf(resi, " PRINT %s", JsonBank.EntryPoint)
	for i := 0; i < len(Randomarray); i++ {
		fmt.Fprintf(resi, "<br>Question %d : ", i+1)
		fmt.Fprintf(resi, "%s", JsonBank.EntryPoint[BankID-1].BankQuestions[Randomarray[i]])

		fmt.Fprintf(resi, "<br>Answer: ")
		fmt.Fprintf(resi, "%s", JsonBank.EntryPoint[BankID-1].CorrectAnswers[Randomarray[i]])
		fmt.Fprintf(resi, "<br>")
	}
}

var BankID int = 2
var QNum int = 3

func main() {
	var n int
	fmt.Println("Please enter the number of questions")
	fmt.Scanln(&n)
	fmt.Println("Please enter which bank you would like?")
	fmt.Scanln(&BankID)
	RNGarray(n)
	router := mux.NewRouter()
	for x := 0; x < 10; x++ {
		go router.HandleFunc("/Questions", GetHttpRequest).Methods("GET")
		go router.HandleFunc("/Answers", EntryPoint2).Methods("GET")
	}
	err := http.ListenAndServe("10.1.1.4:80", router)
	if err != nil {
		errors.New("ERROR")

	}
}

type Quizbanks struct {
	BankID          int        `json:"Bank"`
	BankQuestions   []string   `json:"Questions"`
	PossibleAnswers [][]string `json:"Possible Answers"`
	CorrectAnswers  []string   `json:"Correct Answers"`
}

type Quizbankentry struct {
	EntryPoint []Quizbanks `json:"QuestionBanks"`
}
