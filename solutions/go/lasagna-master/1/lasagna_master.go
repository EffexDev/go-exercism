package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutes int) int {
	if minutes > 0 {
        return len(layers) * minutes
    } else {
        return len(layers) * 2
    }
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodlesCount := 0
    sauceCount := 0.0

    for _, item := range layers {
        if item == "noodles" {
            noodlesCount += 50
        } else if item == "sauce" {
            sauceCount += 0.20
        }
    }
    return noodlesCount, sauceCount
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    secretIngredient := friendsList[len(friendsList) - 1]

    myList = (myList)[:len(myList) - 1]
    myList = append(myList, secretIngredient)
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaledQuantities := []float64{}
    floatPortions := float64(portions)
    for _, i := range quantities {
        scaledQuantities = append(scaledQuantities, i * (floatPortions / 2))
    }
	return scaledQuantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
