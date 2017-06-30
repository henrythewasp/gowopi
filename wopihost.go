package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/wopi/files/{file_id:[0-9]+}", filesHandler).Methods("GET")
	r.HandleFunc("/wopi/files/{file_id:[0-9]+}/contents", fileContentsHandler).Methods("GET")

	http.ListenAndServe(":8181", r)
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["file_id"])

	fmt.Fprintf(w, "File ID is: [%d]", id);
}

func fileContentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["file_id"])

	fmt.Fprintf(w, "Contents for File ID [%d] blah blah blah", id);
}
