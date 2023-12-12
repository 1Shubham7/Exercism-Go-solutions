package twofer

func ShareWith(name string) string {
    a := ""
    if (name == ""){
        a= "One for you, one for me." 
    } else {
    a = "One for "+ name + ", one for me."

        }
	return a
}
