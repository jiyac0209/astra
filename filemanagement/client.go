package filemanagement

import (
	"bufio"
	"fmt"
	"os"
)

// PushDataToFile writes data to a file.
func PushDataToFile(filename string, data string) error {

	directory := "data"

	// Check if the directory exists
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		// Create the directory if it doesn't exist
		err := os.Mkdir(directory, 0755) // 0755 is the default permission mode for directories
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}
	}

	// Open the file in write mode, create it if it doesn't exist, and append to it if it does
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffered writer
	writer := bufio.NewWriter(file)

	// Write data to the file
	_, err = fmt.Fprintln(writer, data)
	if err != nil {
		return err
	}

	// Flush the buffered writer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
