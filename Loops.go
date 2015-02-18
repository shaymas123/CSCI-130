package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	//
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	//
	// x := 0
	// for x < 100 {
	//
	// 	if x >= 100 {
	// 		break
	// 	}
	//
	// 	if x%2 == 0 {
	// 		x++
	// 		continue
	// 	}
	//
	// 	fmt.Println(x)
	// 	x++
	// }

	mySlice := []int{1, 5, 15, 20, 25, 30}

	for i, currentEntry := range mySlice {
		fmt.Println(i, "-", currentEntry)
	}

	myColors := map[string]string{
		"Red":   "Sucks",
		"Green": "Sucks!",
		"Black": "Sucks",
		"Blue":  "Sweet!",
	}

	fmt.Println(myColors["Blue"])

	myColors["Blue"] = "Nevermind, it sucks"

	fmt.Println(myColors["Blue"])

}
