package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(lasagna []string, time int) int {
    if time == 0 {
        time = 2
    }
    return (len(lasagna)* time)
}

// TODO: define the 'Quantities()' function
func Quantities(lasagna []string) (amountNoodles int, amountSauce float64){
    var countSauce, countNoodles int
    for _, value := range lasagna{
        if value == "sauce"{
            countSauce++
        }else if value == "noodles"{
        	countNoodles++
        }
    }
	amountNoodles = 50 * countNoodles
    amountSauce = 0.2 * float64(countSauce)
    return amountNoodles, amountSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string){
   var prev string
	for _, friendListValue := range friendList {
		var found bool
		// fmt.Println(friendListValue)
		for _, myListValue := range myList {
			if myListValue == friendListValue {
				found = true
				break
			}
		}
		if found == false {
			for index, myListValue := range myList {
				if myListValue == "?" || myListValue == prev {
					// fmt.Println(friendListValue)
					prev = friendListValue
					myList[index] = friendListValue
				}
			}
		}
	}
} 

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64{
   ret := make([]float64, len(quantities))
	for index, value := range quantities {
		dv := float64(scale) / 2
		ret[index] = value * dv
	}
	return ret
}
