package main

import "fmt"

func main() {
	var _, err = validDate(0)
	if err != nil {
		fmt.Println(err)
	}

	_, err = validDate(1)
	if err != nil {
		fmt.Println(err)
	}

}

func validDate(date int) (bool, error) {
	if !(date >= 1 && date <= 31) {
		return false, fmt.Errorf(fmt.Sprintf("Incorrect date: %v", date))
	}

	return true, nil
}
