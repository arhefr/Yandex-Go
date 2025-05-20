## Описание:
![image](https://github.com/user-attachments/assets/c8edc1dd-a1db-4d6a-ab70-299364d7fb0d)

## Архитектура:
![image](https://github.com/user-attachments/assets/be330bfa-38b1-4198-86c3-7060688f83c6)


## Установка:
Установите [Git](https://git-scm.com) если у вас его нет.
```
git clone https://github.com/arhefr/Yandex-Go
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
Каждый эндпоинт возвращает статус код 200, если в запросе возникла ошибка, статус код и ошибка возвращается ввиде JSON.
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
- 200 OK ```nil``` ок.
  
- 200 StatusUnprocessableEntity ```{
    "status": 422,
    "message": "error wrong JSON"
}``` некорректный JSON.

- 200 StatusUnprocessableEntity ```{
    "status": 422,
    "message": "error password must contain 8 characters or more and login must contain 3 characters or more"
}``` пароль должен содержать минимум 8 символов, а логин минимум 3.

- 200 InternalServerError ```{
    "status": 500,
    "message": "error something went wrong"
}``` непредвиденная ошибка на сервере.
  
### **POST localhost:8080/api/v1/log-in**
Вход в пользователя, возвращает JWT токен. Загружает в cookie JWT токен. 
``` curl
curl --location 'localhost:8080/api/v1/log-in' \
--header 'Content-Type: application/json' \
--data '{
    "login": "<ЛОГИН>",
    "password": "<ПАРОЛЬ>"
}'
```
- 200 OK ```{
  "token": "<JWT>"
}``` ок.
  
- 200 StatusUnprocessableEntity ```{
    "status": 422,
    "message": "error wrong JSON"
}``` некорректный JSON.

- 200 StatusUnprocessableEntity ```{
    "status": 422,
    "message": "error login not exists"
}``` несуществующий логин.

- 200 StatusUnprocessableEntity ```{
    "status": 422,
    "message": "error wrong password"
}``` неверный пароль.

- 200 InternalServerError ```{
    "status": 500,
    "message": "error something went wrong"
}``` непредвиденная ошибка на сервере.
  
### **AUTH POST localhost:8080/api/v1/calculate**
Отправляет на сервер запрос с математическим выражением, возвращает UUID запроса.
``` curl
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--header 'Cookie: Auth=<JWT TOKEN>' \
--data '{
    "expression": "(2*2)*2"
}'
```

- 200 OK ```{
  "id": "<UUID>"
}```

- 200 Unauthorized ```{
    "status": 401,
    "message": "error authentication"
}``` требуется авторизация.

- 200 StatusUnprocessableEntity ```{
    "status": 422,
    "message": "error wrong JSON"
}``` некорректный JSON.

- 200 StatusUnprocessableEntity ```{
    "status": 422,
    "message": "error expired or wrong jwt token"
}``` некорректный JWT.

- 200 InternalServerError ```{
    "status": 500,
    "message": "error something went wrong"
}``` непредвиденная ошибка на сервере.
  
### **AUTH GET localhost:8080/api/v1/expressions**
``` curl
curl --location 'localhost:8080/api/v1/expressions' \
--header 'Cookie: Auth=<JWT TOKEN>'
```

- 200 OK ```{
  "expressions": <EXPRESSIONS>
}```

- 200 Unauthorized ```{
    "status": 401,
    "message": "error expired or wrong jwt token"
}``` требуется авторизация.

- 200 StatusUnprocessableEntity ```{
    "status": 422,
    "message": "error wrong JSON"
}``` некорректный JSON.

- 200 InternalServerError ```{
    "status": 500,
    "message": "error something went wrong"
}``` непредвиденная ошибка на сервере.

  
### **AUTH GET localhost:8080/api/v1/expressions/<UUID>** 
``` curl
curl --location 'localhost:8080/api/v1/expressions/<UUID>' \
--header 'Cookie: Auth=<JWT TOKEN>'
```

- 200 OK ```{
  "expression": <EXPRESSION>
}```

- 200 Unauthorized ```{
    "status": 401,
    "message": "error expired or wrong jwt token"
}``` требуется авторизация.

- 200 StatusUnprocessableEntity ```{
    "status": 422,
    "message": "error wrong JSON"
}``` некорректный JSON.

- 200 StatusUnprocessableEntity ```{
    "status": 422,
    "message": "error wrong uuid"
}``` некорректный UUID выражения.

- 200 InternalServerError ```{
    "status": 500,
    "message": "error something went wrong"
}``` непредвиденная ошибка на сервере.

## Обратная связь:
[TG](https://t.me/arhefrr)
