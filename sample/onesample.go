package sample

import "fmt"

func GetALoop() {
	for i := 0; i < 5; i++ {
		fmt.Println("iteration number ", i)
	}
}
