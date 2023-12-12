package hamming

import (
    "fmt"
)


func Distance(a, b string) (int, error) {
    var i,diff int;
    if (len(a)!=len(b)){
        return -1,fmt.Errorf("Error")
    }

    for i = 0;i<len(a); i++{
            if a[i]!=b[i]{
                diff++
            }
        }   
	

return diff,nil
}
