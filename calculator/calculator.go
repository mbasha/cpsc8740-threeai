package calculator

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type CalculatorData struct {
	Display string
}

var calcTemplate *template.Template

func init() {
	var err error
	calcTemplate, err = template.ParseFiles("templates/calculator.html")
	if err != nil {
		fmt.Printf("Error parsing calculator template: %v\n", err)
	}
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/calculator", handleCalculatorPage)
	mux.HandleFunc("/calculator/calculate", handleCalculate)
}

func handleCalculatorPage(w http.ResponseWriter, r *http.Request) {
	data := CalculatorData{Display: "0"}
	calcTemplate.Execute(w, data)
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	num1Str := r.FormValue("num1")
	num2Str := r.FormValue("num2")
	operation := r.FormValue("operation")

	if num1Str == "" || num2Str == "" || operation == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid numbers", http.StatusBadRequest)
		return
	}

	var result float64
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
			return
		}
		result = num1 / num2
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%.2f", result)
}
