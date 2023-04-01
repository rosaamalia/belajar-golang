package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struct {
	ID			int
	Name		string
	Age			int
	Division	string
}

var employees = []Employee {
	{ID: 1, Name: "Cinderella", Age: 24, Division: "Finance"},
	{ID: 2, Name: "Ariel", Age: 28, Division: "Accounting"},
	{ID: 3, Name: "Jasmine", Age: 25, Division: "IT Support"},
}

var PORT =":8080"

func main() {
	/**
	 * GET
	 */
	http.HandleFunc("/employees", getEmployees)

	/**
	 * POST
	 */
	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

/**
 * Fungsi controller
 */
func getEmployees(w http.ResponseWriter, r *http.Request) {
	// Set header content-type application/json
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// Mengkonversi data bertipe slice menjadi json
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		/**
		 * Mendapatkan data dari body
		 */
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		// convert age dari string ke int
		convertAge, err := strconv.Atoi(age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// menambahkan data baru
		newEmployee := Employee {
			ID: len(employees) + 1,
			Name: name,
			Age: convertAge,
			Division: division,
		}
		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}