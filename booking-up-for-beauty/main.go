package main

import (
	"fmt"
	_ "fmt"
	"time"
)

func main() {
	var year int
	var month int
	var day int
	var hour int
	var min int
	fmt.Println("Enter the appointment", AnniversaryDate().String())
	fmt.Println("Enter the appointment year")
	_, err := fmt.Scanln(&year)
	if err != nil {
		fmt.Println("Failed to read from cli", err)
		return
	}
	fmt.Println("Enter the appointment month")
	_, err = fmt.Scanln(&month)
	if err != nil {
		fmt.Println("Failed to read from cli", err)
		return
	}
	fmt.Println("Enter the appointment day")
	_, err = fmt.Scanln(&day)
	if err != nil {
		fmt.Println("Failed to read from cli", err)
		return
	}
	fmt.Println("Enter the appointment hour")
	_, err = fmt.Scanln(&hour)
	if err != nil {
		fmt.Println("Failed to read from cli", err)
		return
	}
	fmt.Println("Enter the appointment minutes")
	_, err = fmt.Scanln(&min)
	if err != nil {
		fmt.Println("Failed to read from cli", err)
		return
	}
	fmt.Printf("You entered: %d/%d/%d %d:%d\n", year, month, day, hour, min)

	input := formatInputDate(year, month, day, hour, min)

	hasPassed, err := HasPassed(input)
	if err != nil {
		fmt.Println("Failed:", err)
		return
	}
	if hasPassed {
		isAfter, err := IsAfternoonAppointment(input)
		if err != nil {
			fmt.Println("Failed:", err)
			return
		}
		if isAfter {
			fmt.Println("Appointment is on the afternoon")
		} else {
			fmt.Println("Appointment is not on the afternoon")
		}
		desc, err := Description(input)
		if err != nil {
			fmt.Println("Failed:", err)
			return
		}
		fmt.Println(desc)
	} else {
		fmt.Println("Appointment no longer valid")
		return
	}

	return
}

func formatInputDate(year, month, day, hour, min int) string {
	result := ""
	if month < 1 || month > 12 {
		panic("Month should be between 1 and 12")
	}
	if month < 10 {
		result += fmt.Sprintf("0%v", month)
	} else {
		result += fmt.Sprintf("%v", month)
	}
	if day < 1 || day > 31 {
		panic("Day should be between 1 and 31")
	}
	if day < 10 {
		result += fmt.Sprintf("-0%v", day)
	} else {
		result += fmt.Sprintf("-%v", day)
	}
	result += fmt.Sprintf("-%v ", year)
	if hour < 0 || hour > 23 {
		panic("Hour should be between 0 and 23")
	}
	if hour < 10 {
		result += fmt.Sprintf("0%v", hour)
	} else {
		result += fmt.Sprintf("%v", hour)
	}
	if min < 0 || min > 59 {
		panic("Minutes should be between 0 and 59")
	}
	if min < 10 {
		result += fmt.Sprintf(":0%v", min)
	} else {
		result += fmt.Sprintf(":%v", min)
	}

	return result
}

func AnniversaryDate() time.Time {
	var month int = 8
	var day int = 15

	return time.Date(time.Now().Year(), time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func Description(appointment string) (string, error) {
	date, err := Schedule(&ScheduleParams{date: appointment})
	if err != nil {
		return "", err
	}
	return "You have an appointment on " + date.Format("Mon Jan 2 15:04"), nil
}

func IsAfternoonAppointment(appointment string) (bool, error) {
	date, err := Schedule(&ScheduleParams{date: appointment})
	if err != nil {
		return false, err
	}
	return date.Hour() >= 12 && date.Hour() < 18, nil
}

func HasPassed(appointment string) (bool, error) {
	date, err := Schedule(&ScheduleParams{date: appointment})
	if err != nil {
		return false, err
	}
	return date.After(time.Now()), nil
}

type ScheduleParams struct {
	date   string
	layout string
}

func Schedule(params *ScheduleParams) (time.Time, error) {
	if params.layout == "" {
		params.layout = "01-02-2006 15:04"
	}
	return time.Parse(params.layout, params.date)
}
