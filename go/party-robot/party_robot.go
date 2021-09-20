package partyrobot

import (
    "fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and stands out their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %v years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbour, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %X. Your table is %v, exactly %.1f meters from here.\nYou will be sitting next to %v.", name, table, direction, distance, neighbour)
}
