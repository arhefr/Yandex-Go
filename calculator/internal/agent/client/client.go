package client

import (
	"bytes"
	"calculator/internal/agent/models"
	"calculator/internal/agent/service"
	models_orchestrator "calculator/internal/orchestrator/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func RunWorkers(config models.Config, port string) {
	for i := 1; i <= config.AgentsValue; i++ {
		go Worker(config.AgentPeriodicity, config.OperationTime, fmt.Sprintf("http://localhost:%s%s", port, config.Path))
	}

	log.Infof("%d workers starting working on localhost:%s", config.AgentsValue, port)
}

func Worker(tick time.Duration, operation_time models.OperationTime, url string) {
	client := &http.Client{}
	for {
		time.Sleep(tick)
		start := time.Now()

		resp, _ := client.Get(url)
		if resp.StatusCode != http.StatusOK {
			continue
		}

		var task models_orchestrator.Task
		json.NewDecoder(resp.Body).Decode(&task)

		res := service.MakeTask(task, operation_time)

		log.Debug(task, res, time.Since(start))
		buf := bytes.NewBuffer([]byte{})
		json.NewEncoder(buf).Encode(models.Response{ID: task.ID, Result: res})
		client.Post(url, "application-json", buf)
	}

}
