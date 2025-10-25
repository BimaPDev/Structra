package main

import (
	"fmt"
	"log"
	"net/http"
)


func checkHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}


func main()	{
	port := ":8080"
	http.HandleFunc("/health", checkHealth)
	fmt.Printf("✅ API server is running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
	

}