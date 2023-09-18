package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// randomize_a_file()
	for {
		// rand.Seed(seed)
		mainMenuPromptScanSelectAndBeginSelectedExercise()
	}
}

/*
We have a file: sequentialList.txt
Which, contains records which are each delimited by a blank line
We want to read the records of that file into a slice
And then, randomly write those records to a new file: randomList.txt
*/
func randomize_a_file() {
	// Open the sequentialList.txt file for reading
	sequentialFile, err := os.Open("sequentialList.txt")
	if err != nil {
		fmt.Println("Error opening the sequentialList.txt file:", err)
		return
	}
	
	// defer sequentialFile.Close()
	defer func(sequentialFile *os.File) {
		_ = sequentialFile.Close()
	}(sequentialFile)

	// Read the file, line by line, into slice: records
	var records []string
	currentRecord := ""
	scanner := bufio.NewScanner(sequentialFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Reached a blank line, consider it as the end of a record
			records = append(records, currentRecord)
			currentRecord = ""
		} else {
			// Add the line to the current record
			currentRecord += line + "\n"
		}
	}

	// After the loop, check if there's a last record to append
	if currentRecord != "" {
		records = append(records, currentRecord)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading sequentialList.txt:", err)
		return
	}

	// Shuffle the records randomly
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	// Create a new file for writing
	randomFile, err := os.Create("randomList.txt")
	if err != nil {
		fmt.Println("Error creating randomList.txt:", err)
		return
	}
	
	// defer randomFile.Close() 
	defer func(randomFile *os.File) {
		err := randomFile.Close()
		if err != nil {
			fmt.Println("Error closing file randomFile: randomList.txt")
		}
	}(randomFile)

	// Write the shuffled records to the new file
	writer := bufio.NewWriter(randomFile)
	for _, record := range records {
		_, err := writer.WriteString(record + "\n")
		if err != nil {
			fmt.Println("Error writing to randomList.txt:", err)
			return
		}
	}

	// Flush the writer and check for errors
	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

}
