package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID          int     `json:"id"`
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Consumption float64 `json:"consumption"`
	Phone       string  `json:"phone"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			ID:          1,
			Image:       "https://images.unsplash.com/photo-1633332755192-727a05c4013d?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=880&q=80",
			Name:        "Gustavo Ferraz",
			Email:       "gustavoferraz@gmail.com",
			Consumption: 119.0,
			Phone:       "21 98135-4937",
		},
		{
			ID:          2,
			Image:       "https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=687&q=80",
			Name:        "Luiza Tomé",
			Email:       "luiza.tome@gmail.com",
			Consumption: 4.692,
			Phone:       "53 97252-2971",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}