package main

import ("fmt" 
        "errors"
		"bufio"
		"os"
		"strings"
		"strconv"
	)

func convertTemp (celcius float64, temp int) (float64, error) {
	if temp == 1 {
		return (celcius * 9 / 5 ) + 32 , nil 
	} else if temp == 2 {
		return celcius + 273.15, nil
	} else if temp == 3 {
		return celcius *  4 / 5, nil
	} else {
      return 0, errors.New("pilihan tidak valid")
	}
}

func main () {
	// celciusTemp := 25

	reader := bufio.NewReader(os.Stdin) // Membuat pembaca untuk input pengguna

	fmt.Println("--- Konverter Suhu Celsius ---")

	var celsiusInput string
	var celsiusValue float64
	var err error


	celsiusInput, _ = reader.ReadString('\n')           
	celsiusInput = strings.TrimSpace(celsiusInput)      
	celsiusValue, err = strconv.ParseFloat(celsiusInput, 64) 

	fahrenheit, err := convertTemp(celsiusValue, 1)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("nilai fahrenheitnya adalah", fahrenheit)
	}

	_, err = convertTemp(float64(celsiusValue), 99)
	if err != nil {
		fmt.Println("Error:", err)
	}
}