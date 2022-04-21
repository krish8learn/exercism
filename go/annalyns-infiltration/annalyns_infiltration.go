package annalyn

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	//panic("Please implement the CanFastAttack() function")
    if knightIsAwake == false {
        return true
    }
	return false
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	//panic("Please implement the CanSpy() function")
    if (knightIsAwake == false && archerIsAwake == true && prisonerIsAwake == true) ||(knightIsAwake == true && archerIsAwake == false && prisonerIsAwake == false) ||(knightIsAwake == true && archerIsAwake == true && prisonerIsAwake == false) ||(knightIsAwake == true && archerIsAwake == false && prisonerIsAwake == true)||(knightIsAwake == false && archerIsAwake == true && prisonerIsAwake == false) ||(knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == true)||(knightIsAwake == true && archerIsAwake == true && prisonerIsAwake == true){
        return true
    }
	return false
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	//panic("Please implement the CanSignalPrisoner() function")
    if archerIsAwake == false && prisonerIsAwake == true{
        return true
    }
	return false
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	//panic("Please implement the CanFreePrisoner() function")
    if (knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == true && petDogIsPresent == true) || (knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == false && petDogIsPresent == true) ||(knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == true && petDogIsPresent == false)||(knightIsAwake == true && archerIsAwake == false && prisonerIsAwake == false && petDogIsPresent == true)||(knightIsAwake == true && archerIsAwake == false && prisonerIsAwake == true && petDogIsPresent == true){
        return true
    }
	return false
}
