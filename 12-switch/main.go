package main

import "fmt"

func dayOfweek(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid number"
	}
}

func dayOfweek2(number int) string {
	var dayOfweek string

	switch {
	case number == 1:
		dayOfweek = "Sunday"
	case number == 2:
		dayOfweek = "Monday"
	case number == 3:
		dayOfweek = "Tuesday"
	case number == 4:
		dayOfweek = "Wednesday"
	case number == 5:
		dayOfweek = "Thursday"
	case number == 6:
		dayOfweek = "Friday"
	case number == 7:
		dayOfweek = "Saturday"
	default:
		dayOfweek = "Invalid number"
	}

	return dayOfweek
}

func main() {
	dia := dayOfweek(5)
	fmt.Println(dia)

	fmt.Println("-----------")
	dia2 := dayOfweek2(7)
	fmt.Println(dia2)
}
