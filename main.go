// Revised_hello_world.go asks the operator for input then displays different output depending on operator input.

package main

// TODO, go install gomobile@latest fails. Can I do it otherwise?
// Your IP address is 130.45.44.245 in San Antonio, Texas, United States (78213)
import (
	"fmt"
	//"fyne.io/fyne/v2/internal/driver/mobile/app"
	"log"
)

// init is run before main()
func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile) // Path and File name
	log.Println("Revised_hello_world init with path and file name")
	log.SetFlags(log.LstdFlags | log.Lshortfile) // File name
	log.Println("Revised_hello_world init with file name")
}

func main() {

	var x int
	fmt.Print("enter a number between 1-3: ")

	scan, err := fmt.Scan(&x)
	if err != nil {
		log.Println("Scan err", scan, err)
		return
	}
	if x == 1 {
		log.Println("Hello World!")
	} else if x == 2 {
		log.Println("Hello Dante!")
	} else if x == 3 {
		log.Println("Hello Dale!")
	}
}
