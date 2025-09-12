package raindrops

import (
    "fmt"
)

func Convert(number int) string {
	drops := ""
    if number % 3 != 0 && number % 5 != 0 && number % 7 != 0{
        return fmt.Sprint(number)
    }
    if number % 3 == 0 {
    	drops += "Pling"
    }
    if number % 5 == 0 {
        drops += "Plang"
    }
    if number % 7 == 0 {
        drops += "Plong"
    }
    return drops
}
