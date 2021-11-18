package api

import (
	"encoding/json"
	"net/http"
)

type Recipe struct {
	Name         string        `json:"name"`
	Ingredients  []Ingredient  `json:"ingredients"`
	Instructions []Instruction `json:"instructions"`
}

type Ingredient struct {
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type Instruction struct {
	Description string `json:"description"`
	Order       int    `json:"order"`
}

func ListRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := Recipe{
		Name: "Spinach Pesto",
		Ingredients: []Ingredient{
			{
				Name:     "Spinach",
				Quantity: 150,
				Unit:     "grams",
			},
		},
		Instructions: []Instruction{
			{
				Description: "In a large frying pan, fry the spinach for 3 to 4 minutes",
				Order:       1,
			},
		},
	}

	json.NewEncoder(w).Encode(recipes)
}
