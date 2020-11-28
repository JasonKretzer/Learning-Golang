package main

// how to import multiple packages
import (
	"fmt"
	"math"

	"github.com/jasonkretzer/Learning-Golang/CrashCourse/03_Packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))

	//use strutil
	fmt.Println(strutil.Reverse("olleH"))
}
