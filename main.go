package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

var month int
var year int

func init() {
	now := time.Now()
	flag.IntVar(&month, "month", int(now.Month()), "the month to start and end the search")
	flag.IntVar(&year, "year", now.Year(), "the year of the month to display")
	flag.Parse()
}

func main() {
	// Create a map to store the relevantDays
	relevantDays := make(map[int]struct{})

	// Read datetime stamps from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		// Parse datetime stamp
		layout := "2006-01-02"
		date, err := time.Parse(layout, line)
		if err != nil {
			fmt.Println("Error parsing datetime:", err)
			continue
		}

		// If the month doesn't match don't add to relevent days
		if date.Month() != time.Month(month) {
			continue
		}

		// If the year doesn't match don't add to relevent days
		if date.Year() != year {
			continue
		}

		// Store the day of the date as a key in the map
		relevantDays[date.Day()] = struct{}{}
	}

	// Get the first day of the month
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	// Get the last day of the month
	lastDay := firstDay.AddDate(0, 1, -1)

	// Get the weekday of the first day
	weekday := firstDay.Weekday()

	// Determine the number of days in the month
	numDays := lastDay.Day()

	fmt.Printf("    %s %d\n", firstDay.Month().String(), year)
	fmt.Println("Su Mo Tu We Th Fr Sa")

	// Calculate the number of leading spaces before the first day
	for i := 0; i < int(weekday); i++ {
		fmt.Print("   ")
	}

	// Iterate over the days of the month
	for day := 1; day <= numDays; day++ {
		// Check if date is present in the map
		if _, ok := relevantDays[day]; ok {
			color.New(color.FgGreen).Printf("%2d ", day)
		} else {
			fmt.Printf("%2d ", day)
		}

		// Move to the next line after every Saturday
		if (weekday+time.Weekday(day))%7 == 0 {
			fmt.Println()
		}
	}

	fmt.Println()
}
