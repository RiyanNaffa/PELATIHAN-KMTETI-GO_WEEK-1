package main

import "fmt"

func main() {
	fmt.Println("\n========== TEMPERATURE CONVERTER ==========\n ")
	fmt.Print("Temperature (in C): ")
	var temp_c float64
	fmt.Scan(&temp_c)
	fmt.Println("Convert to:\n1. Fahrenheit\n2. Reamur\n3. Kelvin")
	var opt int
	fmt.Scan((&opt))
	switch opt {
	case 1:
		{
			temp_f := temp_c*9/5 + 32
			fmt.Printf("The temperature is %f F", temp_f)
		}
	case 2:
		{
			temp_r := temp_c * 4 / 5
			fmt.Printf("The temperature is %f R", temp_r)
		}
	case 3:
		{
			temp_k := temp_c + 273.15
			fmt.Printf("The temperature is %f K", temp_k)
		}
	default:
		fmt.Println("Invalid option")
	}
}
