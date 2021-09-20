package gross

// Units store the Gross Store unit measurement
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill create a new bill
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem add item to customer bill
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitVal, ok := units[unit]
	if !ok {
		return false
	} else {
		bill[item] += unitVal
		return true
	}
}

// RemoveItem remove item from customer bill
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemQuant, ok := bill[item]
	if !ok {
		return false
	}
	unitVal, ok := units[unit]
	if !ok {
		return false
	}
	newQuant := itemQuant - unitVal
	switch {
	case newQuant < 0:
		return false
	case newQuant == 0:
		delete(bill, item)
		return true
	default:
		bill[item] = newQuant
		return true
	}
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
	itemQuant, ok := bill[item]
	return itemQuant, ok
}
