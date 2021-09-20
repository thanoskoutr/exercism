package lasagna

// TODO: define the 'OvenTime()' function
func OvenTime() int {
    return 40
}

// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime(time int) int {
    return 40-time
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers int) int {
    return layers*2
}

// TODO: define the 'ElapsedTime()' function
func ElapsedTime(layers int, time int) int {
	return layers*2+time
}