package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type food struct {
	Ingredients []string
	Allergens   []string
}

// Problem: https://adventofcode.com/2020/day/21
// Input: https://adventofcode.com/2020/day/21/input
func main() {
	foods, iToPa, aToPi, err := readAndParseInput("input.txt")
	if err != nil {
		panic(fmt.Errorf("readAndParseInput returned %v\n", err))
	}
	impossibleIngredients, c := computePartOne(foods, iToPa)
	fmt.Println(c)
	fmt.Println(computePartTwo(impossibleIngredients, foods, aToPi))
}

func computePartOne(foods []food, ingredientToPossibleAllergens map[string][]string) ([]string, int) {
	var impossibleIngredients []string
	count := 0
	isPossible := false

	for ingredient, possibleAllergens := range ingredientToPossibleAllergens {
		isPossible = false

		for _, possibleAllergen := range possibleAllergens {
			isPossible = isPossible || isPresentInEveryFoodContainingAllergen(foods, possibleAllergen, ingredient)
		}

		if !isPossible {
			impossibleIngredients = append(impossibleIngredients, ingredient)
		}
	}

	for _, food := range foods {
		for _, impossibleIngredient := range impossibleIngredients {
			if contains(food.Ingredients, impossibleIngredient) {
				count++
			}
		}
	}

	return impossibleIngredients, count
}

func computePartTwo(impossibleIngredients []string, foods []food, allergenToPossibleIngredients map[string][]string) string {
	allergenToIngredients := make(map[string][]string)
	matchedIngredients := make(map[string]bool)

	for allergen, _ := range allergenToPossibleIngredients {
		var allIntersectionsToDo [][]string
		for _, food := range foods {
			if contains(food.Allergens, allergen) {
				var possibleIngredients []string
				for _, i := range food.Ingredients {
					if !contains(impossibleIngredients, i) {
						possibleIngredients = append(possibleIngredients, i)
					}
				}
				allIntersectionsToDo = append(allIntersectionsToDo, possibleIngredients)
			}

			//Intersect each slice in allIntersectionsToDo
			// init ans
			if len(allIntersectionsToDo) > 0 {
				ans := intersect(allIntersectionsToDo[0], allIntersectionsToDo[0])
				for _, s := range allIntersectionsToDo[1:] {
					ans = intersect(ans, s)
				}

				allergenToIngredients[allergen] = ans
				if len(ans) == 1 {
					matchedIngredients[ans[0]] = true
				}
			}
		}
	}

	// Iteratively generate the correct allergenToIngredients by filtering out each ingredient in matchedIngredients
	for len(matchedIngredients) < len(allergenToIngredients) {
		for allergen, ingredients := range allergenToIngredients {
			for matchedIngredient := range matchedIngredients {
				if contains(ingredients, matchedIngredient) && len(ingredients) != 1 {
					allergenToIngredients[allergen] = removeElement(ingredients, matchedIngredient)
				}
			}
			if len(allergenToIngredients[allergen]) == 1 {
				matchedIngredients[allergenToIngredients[allergen][0]] = true
			}
		}
	}

	var allAllergens []string
	for allergen := range allergenToIngredients {
		allAllergens = append(allAllergens, allergen)
	}
	sort.Strings(allAllergens)

	var canonicalDangerousIngredients []string

	for _, allergen := range allAllergens {
		canonicalDangerousIngredients = append(canonicalDangerousIngredients, allergenToIngredients[allergen][0])
	}

	return strings.Join(canonicalDangerousIngredients, ",")

}

func isPresentInEveryFoodContainingAllergen(foods []food, allergen, ingredient string) bool {
	isPresent := true
	for _, food := range foods {
		if contains(food.Allergens, allergen) {
			isPresent = isPresent && contains(food.Ingredients, ingredient)
		}
	}

	return isPresent
}

func readAndParseInput(inputFilename string) ([]food, map[string][]string, map[string][]string, error) {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return nil, nil, nil, err
	}
	defer inputFile.Close()

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return nil, nil, nil, err
	}

	ingredientToPossibleAllergens := make(map[string][]string)
	allergenToPossibleIngredients := make(map[string][]string)
	var foods []food

	for _, f := range strings.Split(string(content), "\n") {
		splitOnContains := strings.Split(f, " (contains ")
		ingredients := strings.Split(splitOnContains[0], " ")
		allergens := strings.Split(splitOnContains[1], ", ")

		// Remove trailing ")"
		allergens[len(allergens)-1] = strings.TrimSuffix(allergens[len(allergens)-1], ")")
		foods = append(foods, food{ingredients, allergens})
		// Construct ingredientToPossibleAllergens map
		for _, ingredient := range ingredients {
			if _, ok := ingredientToPossibleAllergens[ingredient]; ok {
				for _, allergen := range allergens {
					if !contains(ingredientToPossibleAllergens[ingredient], allergen) {
						ingredientToPossibleAllergens[ingredient] = append(ingredientToPossibleAllergens[ingredient], allergen)
					}
				}
			} else {
				ingredientToPossibleAllergens[ingredient] = allergens
			}
		}

		// Construct allergenToPossibleIngredients map
		for _, allergen := range allergens {
			if _, ok := allergenToPossibleIngredients[allergen]; ok {
				for _, ingredient := range ingredients {
					if !contains(allergenToPossibleIngredients[allergen], ingredient) {
						allergenToPossibleIngredients[allergen] = append(allergenToPossibleIngredients[allergen], ingredient)
					}
				}
			} else {
				allergenToPossibleIngredients[allergen] = ingredients
			}
		}
	}

	return foods, ingredientToPossibleAllergens, allergenToPossibleIngredients, nil

}

// helpers (miss you generics :( )
func contains(l []string, t string) bool {
	for _, s := range l {
		if s == t {
			return true
		}
	}
	return false
}

func intersect(a []string, b []string) []string {
	set := make([]string, 0)

	for _, sa := range a {
		if contains(b, sa) {
			set = append(set, sa)
		}
	}

	return set
}

func removeElement(l []string, s string) []string {
	var filteredL []string

	for _, val := range l {
		if val != s {
			filteredL = append(filteredL, val)
		}
	}

	return filteredL
}
