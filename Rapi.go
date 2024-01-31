package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	//to handle route
	"log"

	"github.com/gorilla/mux"
)

type event struct {
	/*  The json:"ID" tag indicates that when converting this struct to JSON, the field should be labeled as "ID." */
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

// new type
type allEvents []event

// new var
var events = allEvents{

	{

		ID:          "1",
		Title:       "Rest api Go",
		Description: "join me to work with go and build Rest apis",
	},
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome home ! ")
}
func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	// req body contains title id description so we have to read it from
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w,"kindly enter data with event title and description only in order")
	}
// decode bc reqbody has type of bytes 
json.Unmarshal(reqBody ,&newEvent)
events = append(events , newEvent)
// status code request and response are  succeed  201 = added 
w.WriteHeader(http.StatusCreated)
//request body
fmt.Println(reqBody)
// response body
fmt.Println(r.Body)
//send data to w (response ) 
json.NewEncoder(w).Encode(newEvent)




}
func getOneEvent(w http.ResponseWriter, r *http.Request) {
// uses the Gorilla Mux router to extract the value associated with the key "id" from the request variables (r)
	eventID :=mux.Vars(r)["id"]
	// range to iterate over events 
	//_to ignore the loop index
	for _,singleevent:=range events{
		if singleevent.ID == eventID{
			json.NewEncoder(w).Encode(singleevent)
		}
	}
}
func getAllEvents(w http.ResponseWriter, r *http.Request) {
	// send json from server to client 
json.NewEncoder(w).Encode(events)
}
func updateEvent(w http.ResponseWriter, r *http.Request) {

}
func deleteEvent(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("lesgooooo with Rest api")
	/* endpoint is a specific route or
	   URL where a service can be accessed.
	   Each endpoint is associated with a
	   particular function or resource in the API. */
	router := mux.NewRouter()
	router.HandleFunc("/", homepage)
	//add
	router.HandleFunc("/event", createEvent).Methods("POST")
	// get one event
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	// get all events
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	//update event
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	// delete event
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	/*  the code suggests that the program is attempting
	to start an HTTP server using http.ListenAndServe
	with the router router. If an error occurs while listening for incoming connections, the log.Fatal function will log an error message
	describing the nature of the error. Subsequently,
	 the program will terminate with a non-zero exit code (1). */
	log.Fatal(http.ListenAndServe(":8080", router))

}
