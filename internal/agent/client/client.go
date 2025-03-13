package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"

	cfg "github.com/arhefr/Yandex-Go/config"
	"github.com/arhefr/Yandex-Go/internal/agent/models"
	"github.com/arhefr/Yandex-Go/internal/agent/service"
	orchestrator "github.com/arhefr/Yandex-Go/internal/orchestrator/models"
	Err "github.com/arhefr/Yandex-Go/pkg/errors"

	log "github.com/sirupsen/logrus"
)

func RunWorkers(cfg *cfg.AgentConfig) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	url := fmt.Sprintf("http://localhost:%s%s", cfg.Port, cfg.Path)
	for i := 1; i <= cfg.AgentsValue; i++ {
		wg.Add(1)
		go Worker(cfg.AgentPeriodicity, cfg.OperationTime, wg, url)
	}

	log.Infof("%d Workers starting on: %s", runtime.NumGoroutine()-1, url)
}

func Worker(tick time.Duration, operation_time models.OperationTime, wg *sync.WaitGroup, url string) {
	defer wg.Done()
	for {
		time.Sleep(tick)
		start := time.Now()

		task, err := fetchTask(url)
		if err != nil {
			if err == Err.IncorrectJSON {
				log.Warn(err)
			}
			continue
		}

		res := service.MakeTask(task, operation_time)
		req := models.Response{ID: task.ID, Result: res}
		log.Debug(task.ID, res, time.Since(start))

		buf := bytes.NewBuffer([]byte{})
		if err := json.NewEncoder(buf).Encode(req); err != nil {
			log.Warn(Err.IncorrectJSON)
		}

		if _, err := http.Post(url, "application/json", buf); err != nil {
			log.Warn(Err.CannotConnect)
		}
	}
}

func fetchTask(url string) (*orchestrator.Task, error) {
	task := new(orchestrator.Task)

	resp, err := http.Get(url)
	if err != nil {
		return task, Err.CannotConnect
	}

	if resp.StatusCode != http.StatusOK {
		return task, Err.NotFoundTask
	}

	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return task, Err.IncorrectJSON
	}

	return task, nil
}
