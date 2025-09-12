package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int {
            "quarter_of_a_dozen":  3,
            "half_of_a_dozen": 6,
            "dozen": 12,
        	"small_gross": 120,
        	"gross": 144,
        	"great_gross": 1728,
        }
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int {
        
    }
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    value, exists := units[unit]
    _, existsBill := bill[item]

	if !exists {
        return false
    } else if existsBill {
        bill[item] = bill[item] + value
        return true
    } else {
        bill[item] = value
        return true
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    unitValue, existsUnit := units[unit]
    itemValue, existsItem := bill[item]

    if !existsUnit || !existsItem {
        return false
    }

    newValue := itemValue - unitValue
    if newValue < 0 {
        return false
    } else if newValue == 0 {
        delete(bill, item)
    } else {
        bill[item] = newValue
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]

    if !exists {
        return 0, false
    } else {
        return value, true
    }
}
