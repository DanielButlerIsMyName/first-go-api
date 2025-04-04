package handler

import (
	"net/http"

	"first-go-api/calc"
	"first-go-api/internal"
)

func RegisterRoutes() {
	http.HandleFunc("/add", AddHandler)
	http.HandleFunc("/subtract", SubtractHandler)
	http.HandleFunc("/multiply", MultiplyHandler)
	http.HandleFunc("/divide", DivideHandler)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := utils.ParseOperands(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := calc.Add(a, b)
	utils.RespondJSON(w, result)
}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := utils.ParseOperands(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := calc.Subtract(a, b)
	utils.RespondJSON(w, result)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := utils.ParseOperands(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := calc.Multiply(a, b)
	utils.RespondJSON(w, result)
}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := utils.ParseOperands(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, divErr := calc.Divide(a, b)
	if divErr != nil {
		http.Error(w, divErr.Error(), http.StatusBadRequest)
		return
	}
	utils.RespondJSON(w, result)
}
