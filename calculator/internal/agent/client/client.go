package client

import (
	"bytes"
	agent "calculator/internal/agent/config"
	"calculator/internal/agent/model"
	"calculator/internal/agent/service"
	"calculator/internal/orchestrator/transport/http/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunWorkers(config agent.Config) {
	for i := 1; i <= config.AgentsValue; i++ {
		log.Printf("Worker %d starting working on %s:%s", i, config.Host, config.Port)
		go Worker(config.AgentPeriodicity, config.OperationTime, fmt.Sprintf("http://%s:%s%s", config.Host, config.Port, config.Path))
	}
}

func Worker(tick time.Duration, operation_time agent.OperationTime, url string) {
	for {
		time.Sleep(tick)
		client := &http.Client{}

		got, _ := client.Get(url)
		var task models.Task
		if err := json.NewDecoder(got.Body).Decode(&task); err != nil {
			continue
		}

		res, err := service.MakeTask(task, operation_time)
		if err != nil {
			continue
		}
		log.Println(task, "Result:", res)
		buf := bytes.NewBuffer([]byte{})
		json.NewEncoder(buf).Encode(model.Response{ID: task.ID, Result: res})
		if _, err := client.Post(url, "application-json", buf); err != nil {
			continue
		}
	}

}
