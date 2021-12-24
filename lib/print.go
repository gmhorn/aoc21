package lib

import (
	"fmt"
	"time"
)

// PrintSolution prints
func PrintSolution(name string, soln func() (int, error)) {
	start := time.Now()
	defer func() {
		fmt.Printf("(Took %s)\n", time.Since(start))
	}()

	fmt.Printf("%s:\n", name)

	ans, err := soln()
	if err != nil {
		fmt.Printf("Error ocurred: %v\n", err)
		return
	}
	fmt.Println(ans)
}
