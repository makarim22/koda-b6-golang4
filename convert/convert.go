package convert

import "errors"

func ConvertTemp (celcius float64, temp string) (float64, error) {
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
