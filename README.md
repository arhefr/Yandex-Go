# Установка
### Требования:

- [**Golang**](https://go.dev/doc/install)
- [**Docker**](https://www.docker.com/products/docker-desktop/)
- [**Docker Compose**](https://docs.docker.com/compose/install/)
- [**Git**](https://git-scm.com/downloads)

### Клонирование репозитория:
``` bash
git clone https://github.com/arhefr/Yandex-Go
```
``` bash
cd Yandex-Go
```

# Запуск:
- ### Запуск с помощью Docker Compose:
``` bash
docker compose up
```

- ### Запуск руками:
``` bash
set PORT = 8080
```
``` bash
set COMPUTING_POWER = 5
```
``` bash
set AGENT_PERIODICITY_MS=1000
```
- В разных терминалах выполните следующие комманды:
``` bash
go run agent/cmd/main.go
```
``` bash
go run orch/cmd/main.go
```
