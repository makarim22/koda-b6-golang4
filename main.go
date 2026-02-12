package main

import (
	"bufio"
	"fmt"
	"golang-4/convert"
	"os"
	"strconv"
	"strings"
)

func endApp () {
  fmt.Println("program berakhir")
  err := recover()
  fmt.Println("terjadi panic", err)
}

func main () {
	defer endApp()

	reader := bufio.NewReader(os.Stdin) 

	fmt.Println("--- Konverter Suhu Celsius ---")

	fmt.Println("--- Input nilai celcius ---")

	var celsiusInput string
	var celsiusValue float64
	var err error

	celsiusInput, _ = reader.ReadString('\n')           
	celsiusInput = strings.TrimSpace(celsiusInput)      
	celsiusValue, err = strconv.ParseFloat(celsiusInput, 64) 

	fmt.Println("--- Masukkan pilihan suhu ---")
	fmt.Println(" 1 : Fahrenheit ")
	fmt.Println(" 2 : Kelvin ")
	fmt.Println(" 3 : Reamur ")

	var choice string
	
	choice, _ = reader.ReadString('\n')
	choice = strings.TrimSpace(choice) 
	fmt.Println("choice", choice)
	



	tempConverted, err := convert.ConvertTemp(celsiusValue, choice)
	if err != nil {
		// fmt.Println("Error", err)
		panic(err)
	} else {
		fmt.Println("nilai konversi adalah", tempConverted)
	}

	// _, err = convertTemp(float64(celsiusValue), 99)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	panic(err)
	// }
}