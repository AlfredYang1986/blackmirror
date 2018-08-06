package main

import (
	//"encoding/json"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmmodel/brand"
	"github.com/alfredyang1986/blackmirror/bmmodel/location"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		sjson := string(body)
		fmt.Println(sjson)

		rst, _ := jsonapi.FromJsonAPI(sjson)
		println(rst)

		//t := rst.(brand.Brand)
		t := rst.(location.Location)
		err = t.InsertBMObject()

		w.Header().Add("Content-Type", "application/json")
		err = jsonapi.ToJsonAPI(&t, w)

		//fmt.Fprintf(w, "Welcome to my website!")
	}).Methods("POST")

	r.HandleFunc("/find", func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		sjson := string(body)
		fmt.Println(sjson)

		rst, _ := jsonapi.FromJsonAPI(sjson)
		println(rst)

		t := rst.(request.Request)

		tmp := brand.Brand{}
		tmp.FindOne(t)
		fmt.Println(tmp)

		w.Header().Add("Content-Type", "application/json")
		err = jsonapi.ToJsonAPI(&tmp, w)
		fmt.Fprintf(w, "Welcome to my website!")

	}).Methods("POST")

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", r)
}
