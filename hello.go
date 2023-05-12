package main

// import multiple packages
import (
	"fmt"
	"strings"
)

func main() {
	// convert the string to lowercas
	lower := strings.ToLower("GOLANG STRINGS")
	fmt.Println(lower)
	// convert the string to uppercase
	upper := strings.ToUpper("golang strings")
	fmt.Println(upper)

}
