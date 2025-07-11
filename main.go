package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// average calculator method which accepts a dictionary as a paramter
func averageCalculator(subjects map[string]int) float64 {
	sum := 0
	num := 0

	for _, grade := range subjects {
		sum += grade
		num += 1
	}

	// if not subjects added return 0
	if num == 0 {
		return 0.0
	}

	return float64(sum) / float64(num)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ask for name
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("Name cannot be empty.")
		return
	}

	// Ask for number of subjects using a loop incase incorrect number entered
	var subs int
	for {
		fmt.Print("Enter the number of subjects you have taken: ")
		subjects, _:= reader.ReadString('\n')
		subjects = strings.TrimSpace(subjects)

		sub, err := strconv.Atoi(subjects)

		// if invalid value is entered continue the loop and ask again
		if err != nil || sub <= 0 {
			fmt.Println("Enter a valid positive number.")
			continue
		}

		subs = sub
		break
	}

	// dictionary for storing subjects Names and Grades
	subjectsDetail := make(map[string]int)

	for i := 0; i < subs; i++ {
		for {
			// ask for a space separated subject name and grade
			fmt.Printf("Enter Name and Grade for subject %v (e.g, Civics 70): ", i+1)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			// split subject name and grade by space
			parts := strings.Split(input, " ")
			if len(parts) != 2 {
				fmt.Println("Please enter both name and age, separated by a space.")
				continue
			}
			subName := parts[0]
			subGrade := parts[1]

			// validate subject name
			if subName == "" {
				fmt.Println("Subject name cannot be empty.")
				continue
			}

			// convert grade to number
			grade, err := strconv.Atoi(subGrade)
			if err != nil || grade > 100 || grade < 0 {
				fmt.Println("Invalid grade.")
				continue
			}

			// add the subject and corresponding grade to the dictionary
			subjectsDetail[subName] = grade
			break
		}
	}

	// Displaying the final results 
	fmt.Println(" ") // gap for nicer display of result
	fmt.Println("Name: ", name)
	fmt.Println("Subjects taken: ", subs)
	fmt.Println(" ") // gap for nicer display of result
	
	fmt.Printf("%-15s | %-5s\n", "Subject", "Grade")
	fmt.Println("------------------------")
	for subject, grade := range subjectsDetail {
		fmt.Printf("%-15s | %-5d\n", subject, grade)
	}
	fmt.Println(" ") // gap for nicer display of result

	fmt.Println("Average: ", averageCalculator(subjectsDetail))

	
}