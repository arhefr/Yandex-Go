![image](https://github.com/user-attachments/assets/c8edc1dd-a1db-4d6a-ab70-299364d7fb0d)
Привет! Я не успеваю допилить проект, прошу, умоляю, проверить мою работу после 14 числа. Буду крайне благодарен за понимание!

Эндпоиты как в тз.
Быстрый запуск с помощью docker compose:
```
docker compose build
```
```
docker compose up
```
Запуск руками:
- Установка переменных среды:
```
set DB_HOST=ХОСТ
set DB_PORT=ПОРТ
set DB_USER=ЮЗЕР
set DB_PASSWORD=ПАРОЛЬ
set DB_NAME=БД
set DB_MAX_ATMPS=5
set DB_DELAY_ATMPS_S=5
set COMPUTING_POWER=3
set AGENT_PERIODICITY_MS=250
set PORT=8080
```
- Запуск оркестратора в одном терминале:
```
cd orch
```
```
go run cmd/main.go
```
- Запуск агента в другом терминале:
```
cd agent
```
```
go run cmd/main.go
```
- Запуск Postresql:
  Запустите сервер postgresql в любом удобном для вас приложении.
