package client

import (
	"bytes"
	"calculator/internal/agent/models"
	"calculator/internal/agent/service"
	models_orchestrator "calculator/internal/orchestrator/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunWorkers(config models.Config, port string) {
	for i := 1; i <= config.AgentsValue; i++ {
		go Worker(config.AgentPeriodicity, config.OperationTime, fmt.Sprintf("http://localhost:%s%s", port, config.Path))
	}

	log.Printf("%d workers starting working on localhost:%s", config.AgentsValue, port)
}

func Worker(tick time.Duration, operation_time models.OperationTime, url string) {
	for {
		time.Sleep(tick)
		client := &http.Client{}

		resp, _ := client.Get(url)
		if resp.StatusCode != http.StatusOK {
			continue
		}

		var task models_orchestrator.Task
		json.NewDecoder(resp.Body).Decode(&task)

		res := service.MakeTask(task, operation_time)

		log.Println(task, "Result:", res)
		buf := bytes.NewBuffer([]byte{})
		json.NewEncoder(buf).Encode(models.Response{ID: task.ID, Result: res})
		client.Post(url, "application-json", buf)
	}

}
