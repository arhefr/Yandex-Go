## Описание:
![image](https://github.com/user-attachments/assets/c8edc1dd-a1db-4d6a-ab70-299364d7fb0d)

## Архитектура:
![image](https://github.com/user-attachments/assets/be330bfa-38b1-4198-86c3-7060688f83c6)


## Установка:
Установите [Git](https://git-scm.com) если у вас его нет.
```
git clone github.com/arhefr/Yandex-Go
```
## Инструкция:
### Конфигурация:
В файле [.env](.env) если необходимо установите переменные среды.
### Запуск:
Установите [Docker](https://www.docker.com), если у вас его нет.
```
docker compose up
```

## Эндпоиты:
### **POST localhost:8080/api/v1/register** 
Регистрирует нового пользователя.
``` curl
curl --location 'localhost:8080/api/v1/register' \
--header 'Content-Type: application/json' \
--data '{
    "login": "<ЛОГИН>",
    "password": "<ПАРОЛЬ>"
}'
```
- 200 OK ```nil```
- 422 StatusUnprocessableEntity ```{"message": "error incorrect JSON"}``` некорректный JSON
- 500 InternalServerError ```{"message": "error invalid data"}``` непредвиденная ошибка на сервере
  
### **POST localhost:8080/api/v1/log-in**
Вход в пользователя, возвращает JWT токен.
``` curl
curl --location 'localhost:8080/api/v1/log-in' \
--header 'Content-Type: application/json' \
--data '{
    "login": "<ЛОГИН>",
    "password": "<ПАРОЛЬ>"
}'
```
- 200 OK ```{"token": "<JWT>"}```
- 422 StatusUnprocessableEntity ```{"message": "error incorrect JSON"}``` некорректный JSON
- 500 InternalServerError ```{"message": "error invalid data"}``` непредвиденная ошибка на сервере
  
### **AUTHENTICATION POST localhost:8080/api/v1/calculate**
Отправляет на сервер запрос с математическим выражением, возвращает UUID запроса.
``` curl
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <JWT>' \
--data '{
    "expression": "<ПРИМЕР>"
}'
```
- 200 OK ```{"id": "UUID"}```
- 422 StatusUnprocessableEntity ```{"message": "error incorrect JSON"}``` некорректный JSON
- 500 InternalServerError ```{"message": "error invalid data"}``` непредвиденная ошибка на сервере
- 511 NetworkAuthenticationRequired ```{"message": "error not authorized"}``` отсутсвие или некорректность JWT токена
  
### **AUTHENTICATION GET localhost:8080/api/v1/expressions**
``` curl
curl --location 'localhost:8080/api/v1/expressions' \
--header 'Authorization: Bearer <JWT>'
```
- 200 OK ```{"expressions": <ПРИМЕРЫ>}```
- 422 StatusUnprocessableEntity ```{"message": "error incorrect JSON"}``` некорректный JSON
- 500 InternalServerError ```{"message": "error invalid data"}``` непредвиденная ошибка на сервере
- 511 NetworkAuthenticationRequired ```{"message": "error not authorized"}``` отсутсвие или некорректность JWT токена
  
### **AUTHENTICATION GET localhost:8080/api/v1/expressions/UUID** 
``` curl
curl --location 'localhost:8080/api/v1/expressions/<UUID>' \
--header 'Authorization: Bearer <JWT>'
```
- 200 OK ```{"expression": <ПРИМЕР>}```
- 422 StatusUnprocessableEntity ```{"message": "error incorrect JSON"}``` некорректный JSON
- 500 InternalServerError ```{"message": "error invalid data"}``` непредвиденная ошибка на сервере
- 511 NetworkAuthenticationRequired ```{"message": "error not authorized"}``` отсутсвие или некорректность JWT токена

## Обратная связь:
[TG](https://t.me/arhefr)
