package synapsemodule

import "fmt"

func Greeting(hour int) {
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}
	fmt.Println(15)
}
