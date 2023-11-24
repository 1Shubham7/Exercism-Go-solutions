package elon

type struct Car{
    speed int
    batteryDrain int
    battery int
    distance int
}


// TODO: define the 'Drive()' method
func Drive(speed, batteryDrain){
    meters := 
    Car {
    	speed: speed,
   		batteryDrain: batteryDrain,
  	 	battery: 100 - batteryDrain,
	    distance: speed,
    }
    
}

// TODO: define the 'DisplayDistance() string' method
func DisplayDistance(distance int)  string{
    
}

// TODO: define the 'DisplayBattery() string' method
func DisplayBattery() {
    
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
CanFinish(trackDistance int) bool{
    
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
