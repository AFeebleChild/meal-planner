package cmd

import (
	"github.com/afeeblechild/meal-planner/lib"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		recipes := lib.GetRecipes()
		//lib.PrintRecipes(recipes)

		var recipe lib.Recipe
		//Check the single IngredientQuery if anything has been passed from the flag,
		//but use IngredientsQuery when picking a recipe by ingredient
		if IngredientQuery == "" {
			recipe = lib.PickRandomRecipe(recipes)
		} else {
			recipe = lib.PickRecipeByIngredient(recipes, IngredientsQuery)
		}
		recipe.PrintRecipe()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
