package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"
// )

// func Exercise4() {
// 	http.HandleFunc("/calculate", CalculateHandle)
// 	fmt.Println("Live")
// 	http.ListenAndServe(":8080", nil)
// }

// func CalculateHandle(w http.ResponseWriter, r *http.Request) {
// 	num1 := r.URL.Query().Get("a")
// 	var1, err := strconv.Atoi(num1)
// 	if err != nil {
// 		http.Error(w, "error", 400)
// 		return
// 	}
// 	num2 := r.URL.Query().Get("b")
// 	var2, err := strconv.Atoi(num2)
// 	if err != nil {
// 		http.Error(w, "error", 400)
// 		return
// 	}

// 	op := r.URL.Query().Get("op")
// 	if op != "add" && op != "substract" && op != "multiply" {
// 		http.Error(w, "error", 400)
// 		return
// 	}

// 	switch op {
// 	case "add":
// 		fmt.Fprint(w, var1+var2)
// 		return

// 	case "substract":
// 		fmt.Fprint(w, var2-var1)
// 		return

// 	case "multiply":
// 		fmt.Fprint(w, var1*var2)
// 		return

// 	}
// }
