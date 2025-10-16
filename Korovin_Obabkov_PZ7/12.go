package main

import "fmt"

func main() {
	book := RecipeBook{}

	borsh := book.CreateRecipe("Борщ", "Традиционный суп", "Супы", 120)
	borsh.AddIngredient("Свёкла", "2 шт")
	borsh.AddIngredient("Капуста", "200 г")
	borsh.AddIngredient("Картошка", "300 г")
	borsh.AddIngredient("Морковь", "100 г")
	borsh.AddStep("Сварить бульон")
	borsh.AddStep("Порезать овощи")
	borsh.AddStep("Варить 2 часа")

	salad := book.CreateRecipe("Овощной салат", "Салат", "Закуски", 15)
	salad.AddIngredient("Огурец", "2 шт")
	salad.AddIngredient("Помидор", "3 шт")
	salad.AddIngredient("Масло", "50 мл")
	salad.AddStep("Нарезать овощи")
	salad.AddStep("Заправить маслом")

	fmt.Println("Рецепты в категории 'Супы':", book.FilterByCategory("Супы"))
	fmt.Println("Рецепты с 'Огурец':", book.FilterByIngredient("Огурец"))
	fmt.Println("Самый долгий рецепт:", book.LongestRecipe().Title)
}

type Ingredient struct {
	Name   string
	Amount string
}

type Recipe struct {
	ID          int
	Title       string
	Description string
	Ingredients []Ingredient
	Steps       []string
	Category    string
	PrepTime    int
}

type RecipeBook struct {
	Recipes []Recipe
	NextID  int
}

func (rb *RecipeBook) CreateRecipe(title, desc, category string, prepTime int) *Recipe {
	recipe := Recipe{
		ID:          rb.NextID,
		Title:       title,
		Description: desc,
		Category:    category,
		PrepTime:    prepTime,
	}
	rb.Recipes = append(rb.Recipes, recipe)
	rb.NextID++
	return &rb.Recipes[len(rb.Recipes)-1]
}

func (r *Recipe) AddIngredient(name, amount string) {
	r.Ingredients = append(r.Ingredients, Ingredient{Name: name, Amount: amount})
}

func (r *Recipe) AddStep(step string) {
	r.Steps = append(r.Steps, step)
}

func (rb *RecipeBook) FilterByCategory(category string) []Recipe {
	var result []Recipe
	for _, r := range rb.Recipes {
		if r.Category == category {
			result = append(result, r)
		}
	}
	return result
}

func (rb *RecipeBook) FilterByIngredient(ingredient string) []Recipe {
	var result []Recipe
	for _, r := range rb.Recipes {
		for _, ing := range r.Ingredients {
			if ing.Name == ingredient {
				result = append(result, r)
				break
			}
		}
	}
	return result
}

func (rb *RecipeBook) LongestRecipe() *Recipe {
	if len(rb.Recipes) == 0 {
		return nil
	}
	longest := rb.Recipes[0]
	for _, r := range rb.Recipes {
		if r.PrepTime > longest.PrepTime {
			longest = r
		}
	}
	return &longest
}
