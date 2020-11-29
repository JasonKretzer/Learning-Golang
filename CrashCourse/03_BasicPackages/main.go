package main

// how to import multiple packages
import (
	"fmt"
	"math"

	// you import your local packages based on the relative path from the gopath
	"github.com/jasonkretzer/Learning-Golang/CrashCourse/03_Packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))

	//use strutil
	fmt.Println(strutil.Reverse("olleH"))
}
