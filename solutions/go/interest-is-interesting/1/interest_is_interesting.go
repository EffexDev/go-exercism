package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    var interest float32
	switch {
        case balance < 0.0:
        	interest = 3.213
        case balance >= 0.0 && balance < 1000.0:
        	interest = 0.5
        case balance >= 100 && balance < 5000:
        	interest = 1.621
        case balance >= 5000:
        	interest = 2.475
    }
    return interest
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := float64(InterestRate(balance)) / 100
	interest := balance * rate
    return interest
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	gain := Interest(balance)
    balance += gain
    return balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    years := 0
    
	for balance < targetBalance {
        balance = AnnualBalanceUpdate(balance)
        years++
    }
    return years
}
