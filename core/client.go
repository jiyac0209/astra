package core

import (
	"astraSecurity/domain"
	"astraSecurity/filemanagement"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func Push(astra *domain.Astra) error {
	//attach UUID
	astra.UUID = uuid.New().String()
	astra.Timestamp = time.Now()

	jsonData, err := json.Marshal(astra)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	//push data to file
	err = filemanagement.PushDataToFile(fmt.Sprintf("data/%s.json", astra.UUID), string(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
