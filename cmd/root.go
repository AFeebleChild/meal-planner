package cmd

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/afeeblechild/meal-planner/lib"
	"github.com/spf13/cobra"
)

var (
	AllRecipes []lib.Recipe
	IngredientQuery string
	IngredientsQuery []string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "meal-planner",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())

		AllRecipes = lib.GetRecipes()

		if strings.Contains(IngredientQuery, ",") {
			splitIngredients := strings.Split(IngredientQuery, ",")
			for _, splitIngredient := range splitIngredients {
				IngredientsQuery = append(IngredientsQuery, splitIngredient)
			}
		} else {
			IngredientsQuery = append(IngredientsQuery, IngredientQuery)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", MainHandler)
		http.HandleFunc("/random/", RandomHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.meal-planner.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&IngredientQuery, "ingredient", "i", "", "ingredient(s) to search recipes for, comma separated")
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, r)
}

func RandomHandler(w http.ResponseWriter, r *http.Request) {
	randomRecipe := lib.PickRandomRecipe(AllRecipes)
	if len(randomRecipe.Ingredients) > 0 {
		t, _ := template.ParseFiles("templates/listRecipe.html")
		t.Execute(w, randomRecipe)
	}else {
		http.Redirect(w, r, "http://"+r.Host, 302)
	}
}
