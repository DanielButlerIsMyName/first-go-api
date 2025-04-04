package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func RespondJSON(w http.ResponseWriter, result float64) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"result": result,
	})
}

func ParseOperands(r *http.Request) (float64, float64, error) {
	query := r.URL.Query()
	aStr := query.Get("a")
	bStr := query.Get("b")

	a, errA := strconv.ParseFloat(aStr, 64)
	b, errB := strconv.ParseFloat(bStr, 64)

	if errA != nil || errB != nil {
		return 0, 0, errors.New("invalid parameters: expected float64 'a' and 'b'")
	}
	return a, b, nil
}
