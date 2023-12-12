package collatzconjecture

import (
    "fmt"
)

func CollatzConjecture(n int) (int, error) {
    var count int = 0
    for(n>1){
        if (n%2==0){
            n=n/2
        } else {
        	n=3*n+1
        }
    count++
    
    }
    if (n<1){
        return -1,fmt.Errorf("Error")
    }
	return count,nil
    
}
