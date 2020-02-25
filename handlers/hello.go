package handlers

import "net/http"

func HelloHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello, World!"))
}
