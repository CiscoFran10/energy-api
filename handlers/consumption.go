package handlers

import (
	"encoding/json"
	"net/http"
)


type Consumption struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Value int `json:"value"`
}

func getConsumption(w http.ResponseWriter, r *http.Request) {
	consumption := []Consumption{
		{
			ID:    1,
			Name:  "Jan",
			Value: 200,
		},
		{
			ID:    2,
			Name:  "Fev",
			Value: 300,
		},
		{
			ID:    3,
			Name:  "Mar",
			Value: 400,
		},
		{
			ID:    4,
			Name:  "Abr",
			Value: 400,
		},
		{
			ID:    5,
			Name:  "Mai",
			Value: 100,
		},
		{
			ID:    6,
			Name:  "Jun",
			Value: 500,
		},
		{
			ID:    7,
			Name:  "Jul",
			Value: 300,
		},
		{
			ID:    8,
			Name:  "Ago",
			Value: 800,
		},
		{
			ID:    9,
			Name:  "Set",
			Value: 400,
		},
		{
			ID:    10,
			Name:  "Out",
			Value: 200,
		},
		{
			ID:    11,
			Name:  "Nov",
			Value: 600,
		},
		{
			ID:    12,
			Name:  "Dez",
			Value: 100,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consumption)
}
