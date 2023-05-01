package seeding

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type EmployeeIDList struct {
	Data []EmployeeID `json:"data"`
}

type EmployeeID struct {
	EmployeeID string `json:"name"`
}

func GetDetails() error {
	keepRunning := true
	startPoint := 0
	for keepRunning {
		req, err := http.NewRequest("GET", fmt.Sprintf("http://10.0.8.190:8090/api/resource/Employee?limit_start=%d", startPoint), nil)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "token 472e6bdd5f355a2:a5b7f758c8cfbf0")

		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			return err
		}

		defer response.Body.Close()
		title, _ := io.ReadAll(response.Body)
		var employeeIDList EmployeeIDList
		json.Unmarshal(title, &employeeIDList)
		if len(employeeIDList.Data) == 0 {
			keepRunning = false
		}
		for i := 0; i < len(employeeIDList.Data); i++ {
			fmt.Println(employeeIDList.Data[i].EmployeeID)

		}
		startPoint += 20
	}
	return nil
}
