package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello errcheck")

}

// missingCheckOnDeferClose demonstrates not checking the err condition in a `defer` statement
func missingCheckOnDeferClose() {
	f, err := os.Open("CreatedFile.txt")
	if err != nil {
		return
	}
	defer f.Close() // <-- error! We didn't use the err value
	fmt.Printf("Opened a file but havent closed it yet -- what if the defer close didnt work?")
}

// missingCheckOnOSStat demonstrates assigning err to discard (_)
// requires use of the [-blank] flag
func missingCheckOnOSStat() {
	f, _ := os.Stat("non_existant_file.txt") // <-- error! Discarded
	fmt.Printf("[missingCheckOnOSStat] got file %s", f.Name())
}

// missingCheckOnReassignment demonstrates that errcheck isn't perfect!
// this error is a re-assignment followed by non-use of the `err` variable
// This is covered by a different tool, `https://staticcheck.io/`
func missingCheckOnReassignment() {
	f, err := os.Create("CreatedFile.txt")
	if err != nil {
		fmt.Printf("Created file")
	}
	f, err = os.Open("SomeDifferentFile.txt") // <-- error! Reassigned and Unchecked
	doAThing(f)
	fmt.Printf("We've used err already, so the compiler doesn't complain... but this err is actually a brand new err! And we didn't check it. Whoops.")
}

// missingReturnCheck demonstrates using a command that returns an error and not using it at all
func missingReturnCheck() {
	f, err := os.Open("SomeFile")
	if err != nil {
		return
	}
	doAThing(f)
	f.Chdir() // <-- error! Unchecked
}

// incorrectTypeAssertion demonstrates ignoring the results of a type assertion
// requires use of the [-assert] flag
func incorrectTypeAssertion(obj interface{}) {
	i := obj.(int) // <-- error! Unchecked
	i++
}

// ignoreTheseErrors deonstrates that by using the [-ignore] flag, we can ignore certain methods if we don't care about their error return values
func ignoreTheseErrors() {
	f, err := os.Open("SomeFile.txt")
	if err != nil {
		return
	}
	f.Write([]byte{}) // <-- would be an error, but we ignored the Write methods
}

// doAThing is a placeholder method so that the compiler thinks we've used `f`
func doAThing(obj interface{}) {

}
