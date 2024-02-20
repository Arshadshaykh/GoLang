package module2

import (
	"fmt"
	"yourmodulepath/module1" // Import module1
)

// Function in module2 that calls code from module1
func Greet() {
	fmt.Println("Greetings from module2!")

	// Call SayHello from module1
	module1.SayHello()
}
