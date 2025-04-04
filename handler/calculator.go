package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"first-go-api/calc"
)

func Calculate(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	aStr := query.Get("a")
	bStr := query.Get("b")
	op := query.Get("op") // add, sub, mul, div

	a, err1 := strconv.ParseFloat(aStr, 64)
	b, err2 := strconv.ParseFloat(bStr, 64)

	if err1 != nil || err2 != nil || op == "" {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	var result float64
	var err error

	switch op {
	case "add":
		result = calc.Add(a, b)
	case "sub":
		result = calc.Sub(a, b)
	case "mul":
		result = calc.Mul(a, b)
	case "div":
		result, err = calc.Div(a, b)
	default:
		http.Error(w, "Unsupported operation", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"result": result,
	})
}
