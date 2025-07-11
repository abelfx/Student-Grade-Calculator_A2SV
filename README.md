# Student Grade Calculator

A simple Go console application that helps students calculate their average grade across multiple subjects.

## Features

- Prompts the student to enter their name.
- Asks for the number of subjects taken.
- For each subject, prompts for the subject name and grade.
- Validates inputs (ensures grades are between 0 and 100, and subjects are entered correctly).
- Supports multi-word subject names (e.g., "Computer Science").
- Calculates and displays the average grade with two decimal precision.
- Displays all subjects and grades in a nicely formatted table.

## Getting Started

### Prerequisites

- Go 1.16 or higher installed on your system.

### Installation

1. Clone the repository or download the source code.

2. Initialize the Go module (if not already done):

   ```
   go mod init student-grade-calculator
   ```
3, To run the application
   ```
   go run main.go
   ```
4, To test the application
   ```
   go test
   ```
