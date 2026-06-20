package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main()  {
	http.HandleFunc("/status", statusHandle)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}

func statusHandle(w http.ResponseWriter, r *http.Request)  {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "code parameter is required", http.StatusBadRequest)
		return
	}
	convertCode, err := strconv.Atoi(code)
	if err != nil {
		http.Error(w, "code must be a valid ineger", http.StatusBadRequest)
		return
	}
	if convertCode < 100 || convertCode > 599 {
		http.Error(w, "code must be a valid HTTP status code (100-599)", http.StatusBadRequest)
	}else {
		w.WriteHeader(convertCode)
	}


}