package main

import ("fmt" 
        "errors"
		"bufio"
		"os"
		"strings"
		"strconv"
	)

func convertTemp (celcius float64, temp string) (float64, error) {
	if temp == "1" {
		return (celcius * 9 / 5 ) + 32 , nil 
	} else if temp == "2" {
		return celcius + 273.15, nil
	} else if temp == "3" {
		return celcius *  4 / 5, nil
	} else {
      return 0, errors.New("pilihan tidak valid")
	}
}

func endApp () {
  fmt.Println("program berakhir")
  err := recover()
  fmt.Println("terjadi panic", err)
}

func main () {
	// celciusTemp := 25

	defer endApp()

	reader := bufio.NewReader(os.Stdin) 

	fmt.Println("--- Konverter Suhu Celsius ---")

	fmt.Println(" 1 : Fahrenheit ")
	fmt.Println(" 2 : Kelvin ")
	fmt.Println(" 3 : Reamur ")

	var celsiusInput string
	var celsiusValue float64
	var err error

	celsiusInput, _ = reader.ReadString('\n')           
	celsiusInput = strings.TrimSpace(celsiusInput)      
	celsiusValue, err = strconv.ParseFloat(celsiusInput, 64) 


	var choice string
	
	choice, _ = reader.ReadString('\n')
	choice = strings.TrimSpace(choice) 
	fmt.Println("choice", choice)
	



	tempConverted, err := convertTemp(celsiusValue, choice)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("nilai konversi adalah", tempConverted)
	}

	// _, err = convertTemp(float64(celsiusValue), 99)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	panic(err)
	// }
}