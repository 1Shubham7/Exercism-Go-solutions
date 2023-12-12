package raindrops

import (
    "fmt"
)

func Convert(number int) string {
    a:= ""
    if number % 3 == 0 {
        a = a + "Pling"
    }
    if number % 5 == 0{
        a = a + "Plang"
    }
    if number % 7 == 0{
        a = a + "Plong"
    }
	if (number % 3 != 0 && number % 5 != 0 && number % 7 != 0){
        num := fmt.Sprint(number)
        return num
    }
	return (a)
}
