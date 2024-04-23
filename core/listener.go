package core

import (
	"astraSecurity/domain"
	"astraSecurity/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Function to read files from a directory, process them, and delete them
func ProcessFiles(db *sql.DB, dirPath string) {
	// Continuously monitor the directory for changes
	for {
		// Get a list of files in the directory
		files, err := ioutil.ReadDir(dirPath)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			continue
		}

		// Iterate through each file in the directory
		for _, file := range files {
			if file.IsDir() {
				continue // Skip directories
			}

			// Read the file contents
			filePath := filepath.Join(dirPath, file.Name())
			data, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			var astra domain.Astra
			err = json.Unmarshal(data, &astra)
			if err != nil {
				fmt.Println(err)
			}

			// Process the file data (e.g., insert into database)
			err = repository.Insert(db, astra)
			if err != nil {
				fmt.Println("Error processing file data:", err)
				continue
			}

			// Delete the file
			err = os.Remove(filePath)
			if err != nil {
				fmt.Println("Error deleting file:", err)
				continue
			}

			fmt.Println("File processed and deleted:", filePath)
		}

		// Sleep for a while before checking again
		time.Sleep(5 * time.Second) // Adjust interval as needed
	}
}
