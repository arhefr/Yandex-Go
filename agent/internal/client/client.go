package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"

	Err "github.com/arhefr/Yandex-Go/agent/internal/errors"
	"github.com/arhefr/Yandex-Go/agent/internal/model"
	"github.com/arhefr/Yandex-Go/agent/internal/service"

	log "github.com/sirupsen/logrus"
)

const (
	URL = "/internal/task"
)

func RunWorkers(cfg *service.Config) {
	cfg.WG.Add(cfg.AgentsValue)
	for i := 1; i <= cfg.AgentsValue; i++ {
		go func(url string) {
			Worker(cfg.AgentPeriodicity, url)
			cfg.WG.Done()
		}(fmt.Sprintf("http://orch:%s%s", cfg.Port, URL))
	}

	log.Infof("%d Workers start working", runtime.NumGoroutine()-1)
	cfg.WG.Wait()
}

func Worker(tick time.Duration, url string) {
	for {
		time.Sleep(tick)

		task, err := getWork(url)
		if err != nil {
			continue
		}

		res := service.MakeTask(task)
		resp := model.Response{ID: task.ID, Sub_ID: task.Sub_ID, Result: res}

		buf := bytes.NewBuffer([]byte{})
		if err := json.NewEncoder(buf).Encode(resp); err != nil {
			log.Warn(Err.IncorrectJSON)
		}

		if _, err := http.Post(url, "application/json", buf); err != nil {
			log.Warn(Err.CannotConnect)
		}
	}
}

func getWork(url string) (*model.Task, error) {
	task := new(model.Task)

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
