package main
import (
	"fmt"
	"net/http"
)

func main(){
	// Declarations without meaning only for build the main.go

	/*Especified de url of the service*/
	http.HandleFunc("/radiup", func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, "Testando")
		})
	http.ListenAndServe(":8080", nil)
}