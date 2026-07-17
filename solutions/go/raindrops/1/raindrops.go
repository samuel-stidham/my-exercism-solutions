package raindrops

import (
	"strconv"
)

func isDivisibleBy(number int, divisor int) bool {
    return number % divisor == 0
}

func Convert(number int) string {
    var result string
    result = ""

	if isDivisibleBy(number, 3) {
        result += "Pling"
    }

    if isDivisibleBy(number, 5) {
        result += "Plang"
    }

    if isDivisibleBy(number, 7) {
        result += "Plong"
    }
    
    if result == "" {
        result = strconv.Itoa(number)
    }
    
	return result
}
