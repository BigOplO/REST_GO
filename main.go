package main

import{
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
}

type Movie struct{
	ID string 'json:"id"'
	Isbn string 'json:"isbn"'
	Title string 'json:"title"'
	Director *Director 'json:"director"'
}



type Director struct{
	Firstname string 'json:"firstname"'
	Lastname string 'json:"lastname"'
}

var movies []Movie

func getMovies(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break;
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovies(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(movies)
			return 
		}
	}
}


func createMovies(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewEncoder(r.Body).Decoder(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}


func updateMovie( http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			var movie Movie
			_ = json.NewEncoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies,movie)
			json.NewEncoder(w).Encode(movie)
			return 
		}
	}
}



func(){
	r := mux.NewRouter()

	movies = append(movies,Movie{ID:"1",Isbn:"789789777",Title:"Test GO",Director: &Director{Firstname:"Kevin",Lastname:"Adminn"}})
	movies = append(movies,Movie{ID:"2",Isbn:"4262382277",Title:"REST GO",Director: &Director{Firstname:"Daris",Lastname:"faris"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovies).Methods("GET")
	r.HandleFunc("/movies",createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovies).Methods("PUT")e
	r.HandleFunc("/movies/{id}",deleteMovies).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}