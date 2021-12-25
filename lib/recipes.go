package lib

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"strings"
)

type (
	Recipe struct {
		Name        string       `yaml:"Name"`
		Ingredients []Ingredient `yaml:"Ingredients"`
		MealTypes   []string     `yaml:"MealTypes"`
	}
	Ingredient struct {
		Name   string `yaml:"Name"`
		Amount string `yaml:"Amount,omitempty"`
	}
)

//GetRecipes is a simple wrapper func in case I ever change to a database backend in the future
func GetRecipes() []Recipe {
	return getRecipesFromFile()
}

func getRecipesFromFile() []Recipe {
	file, err := ioutil.ReadFile("recipes.yaml")
	if err != nil {
		panic(err)
	}

	var recipes []Recipe
	err = yaml.Unmarshal(file, &recipes)
	if err != nil {
		panic(err)
	}

	return recipes
}

func PrintRecipes(recipes []Recipe) {
	for _, recipe := range recipes {
		recipe.PrintRecipe()
	}
}

func (recipe Recipe) PrintRecipe() {
	fmt.Println("- Name:", recipe.Name)
	fmt.Println("MealTypes:")
	for _, mealType := range recipe.MealTypes {
		fmt.Println("\t", mealType)
	}
	fmt.Println("Ingredients:")
	for _, ingredient := range recipe.Ingredients {
		fmt.Println("\t-", ingredient.Name)
		if ingredient.Amount != "" {
			fmt.Println("\t\tAmount:", ingredient.Amount)
		}
	}
	fmt.Println()
}

func PickRandomRecipe(recipes []Recipe) Recipe {
	//rand is seeded in the root command PersistentPreRun
	i := rand.Intn(len(recipes))

	return recipes[i]
}

//PickRecipeByIngredient does assume that ingredientQuery is not a blank string
func PickRecipeByIngredient(recipes []Recipe, ingredientsQuery []string) Recipe {
	var matchingRecipes []Recipe
	for _, recipe := range recipes {
		matching := 0
		for _, ingredient := range recipe.Ingredients {
			for _, ingredientQuery := range ingredientsQuery {
				if strings.Contains(strings.ToLower(ingredient.Name), strings.ToLower(ingredientQuery)) {
					matching++
				}
			}
		}
		//Only add the recipe to matching recipes if all the ingredientsQuery match
		if matching == len(ingredientsQuery) {
			matchingRecipes = append(matchingRecipes, recipe)
		}
	}

	var i int
	if len(matchingRecipes) == 0 {
		return Recipe{}
	} else {
		i = rand.Intn(len(matchingRecipes))
	}

	return matchingRecipes[i]
}
