package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {

	// Color names in ANSI codes
	yellow := "\033[33m"
	brightRed := "\033[91;1m"
	lineColor := "\033[97m"
	neonBlue := "\033[38;5;75;1m"
	reset := "\033[0m" // Reset color to default
	brightGreen := "\033[92m"

	// Input parameter for filtering by topic, this is optional
	var selectedTopic string
	flag.StringVar(&selectedTopic, "topic", "", "Specify the topic for random problem selection")
	flag.Parse()

	// Opening excel file
	excelFileName := "LC.xlsx"
	xlFile, err := excelize.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("Error opening Excel file: %v\n", err)
		return
	}

	// Getting sheet name
	sheetName := xlFile.GetSheetName(1)

	// Getting all Rows
	rows := xlFile.GetRows(sheetName)

	problemIndex, linkIndex, topicIndex, solvedIndex := -1, -1, -1, -1

	// Find the indices of the columns with headers "PROBLEM," "LINK," "TOPICS" and "Times Solved"
	headerRow := rows[0]
	for i, cell := range headerRow {
		text := cell
		switch text {
		case "PROBLEM":
			problemIndex = i
		case "LINK":
			linkIndex = i
		case "TOPIC":
			topicIndex = i
		case "Times Solved":
			solvedIndex = i
		}
	}

	// If any of the columns are not found, handle the error
	if problemIndex == -1 || linkIndex == -1 || topicIndex == -1 || solvedIndex == -1 {
		fmt.Println("One or more headers not found.")
		return
	}

	// Filter problems based on the specified topic
	filteredRows := make([][]string, 0)
	for _, r := range rows[1:] { // Start from the second row to skip headers
		if selectedTopic == "" || strings.EqualFold(r[topicIndex], selectedTopic) {
			filteredRows = append(filteredRows, r)
		}
	}
	if len(filteredRows) == 0 {
		fmt.Println("No problems found for the specified topic.")
		return
	}

	var solvedInput string

	// Retrieve the values from the random row
	randomRow := rand.Intn(len(filteredRows))
	row := filteredRows[randomRow]
	problem := row[problemIndex]
	link := row[linkIndex]
	topics := row[topicIndex]
	solvedStr := row[solvedIndex]

	// if no value is set in Times Solved, make it zero so that we can update it to 1
	if solvedStr == "" {
		solvedStr = "0"
	}
	solved, err := strconv.Atoi(solvedStr)
	if err != nil {
		fmt.Printf("Error parsing SOLVED field: %v\n", err)
		return
	}

	linkText := "\033[92m 🔗 Open Problem\033[0m" // Bright green link text

	fmt.Printf("%s─────────────────────────────────────────────────%s\n", lineColor, reset)
	fmt.Printf("  🚀 %sSolve this%s\n", brightRed, reset)
	fmt.Printf("%s─────────────────────────────────────────────────%s\n", lineColor, reset)
	fmt.Printf("  📚 Problem: %s%s%s                     \n", yellow, problem, reset)
	fmt.Printf("  🏷️  Topic: %s%s%s                       \n", neonBlue, topics, reset)
	fmt.Printf("%s─────────────────────────────────────────────────%s\n", lineColor, reset)
	fmt.Printf("  ┌──────────────────┐  \n")
	fmt.Printf("  │\033]8;;%s\033\\%s\033]8;;\033\\  │\n", link, linkText)
	fmt.Printf("  └──────────────────┘  \n")
	fmt.Printf("%s─────────────────────────────────────────────────%s\n", lineColor, reset)

	// Ask user if they have solved the problem
	fmt.Printf("\n%sDid you solve it? (Yes/No): %s", neonBlue, reset)
	fmt.Scan(&solvedInput)
	solvedInput = strings.ToLower(solvedInput)

	// Getting actual row number to update solved count
	var rowNumber int = -1
	for i, row := range rows {
		if row[linkIndex] == link {
			rowNumber = i
			break
		}
	}
	if rowNumber == -1 {
		fmt.Println("No row found with LINK =")
		return
	}

	// if user solved it, update the count. Otherwise dont
	if solvedInput == "yes" {
		solved++
		// Update the Excel file with the new Times Solved count
		newVal := excelize.ToAlphaString(solvedIndex) + strconv.Itoa(rowNumber+1)
		xlFile.SetCellValue(sheetName, newVal, solved)
		fmt.Printf("%sSolved count updated, keep it up!%s\n", brightGreen, reset)
		// Save the updated Excel file
		if err := xlFile.Save(); err != nil {
			fmt.Printf("Error saving Excel file: %v\n", err)
			return
		}
	} else {
		fmt.Printf("%sSolved count not updated%s\n", brightRed, reset)
	}

}
