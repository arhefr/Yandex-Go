![Слайд1](https://github.com/user-attachments/assets/a1250ba0-7909-42aa-8cdf-887946252dac)
# **Содержание**
>  - [**Описание**](#описание)
>  - [**Как использовать?**](#как-использовать)
>     - [Быстрый запуск](#быстрый-запуск)
>     - [Отправка запросов](#отправка-запросов)
>     - [Получение запросов](#получение-запросов)
>     - [Получение запроса](#получение-запроса)
>  - [**Обратная связь**](#обратная-связь)

# **Описание**
>[!NOTE]
>Это сервер-калькулятор, который использует worker pull. Давайте посмотрим как это работает.
>
>![Слайд2](https://github.com/user-attachments/assets/264eea11-d04c-4ce0-ab9a-c407b7a23b94)
>
>Гопхер-оркестратор принимает запросы на вычисление математичеких выражений, разбивает их на последовательные операции, которые распределяет между гопхерами-агентами. Гопхеры-агенты выполняют порученное им задание и возвращают результат гопхеру-оркестратору. Так происходит до тех пор, пока выражение не вычилится, потом гопхер-оркестратор возвращает нам результат.

# **Как использовать**
>[!WARNING]
>Убедитесь, что у вас установлен Golang 1.23.4 и выше.

>[!IMPORTANT]
>## **Быстрый запуск**
> ``` shell
> cd calculator
> ```
> ``` shell
> go run cmd/web/start.go
> ```
>## **Отправка запросов**
>### /api/v1/calculate POST
> Отправка математического выражения:
> Пример запроса:
> ``` shell
> curl -X POST -H "Content-type:application/json" --data "{\"expression\":\"2+2*2\"}" http://localhost:8080/api/v1/calculate
> ```
> Ответ сервера:
> ```
> {"id": УНИКАЛЬНЫЙ_ID_ЗАПРОСА}
> ```
>## **Получение запросов**
>### /api/v1/calculate POST
> Отправка математического выражения:
> Пример запроса:
> ``` shell
> curl -X POST -H "Content-type:application/json" --data "{\"expression\":\"2+2*2\"}" http://localhost:8080/api/v1/calculate
> ```
> Ответ сервера:
> ```
> {"id": УНИКАЛЬНЫЙ_ID_ЗАПРОСА}
> ```
> Другие возможные ответы сервера:
> ```
> {"error":"Method is not allowed"}
> ```
>## **Получение запроса**
>### /api/v1/calculate GET
> Возвращает список, принятых запросов
> Пример запроса:
> ``` shell
> curl localhost:8080/api/v1/expressions
> ```
> Ответ сервера:
> ```
> [{"id":id,"status":"","result":""}, ...]
> ```
> Другие возможные ответы сервера:
> ```
> {"error":"Method is not allowed"}
> ```
>### /api/v1/calculate/{id} GET
> Возвращает запрос по ID
> Пример запроса:
> ``` shell
> curl localhost:8080/api/v1/expressions/:id
> ```
> Ответ сервера:
> ```
> {"id":id,"status":"","result":""}
> ```
> Другие возможные ответы сервера:
> ```
> {"error":"Method is not allowed"}
> ```

# **Обратная связь**
> Немного не успеваю написать тесты и полноценный ридми, простите пожалуйста. Подождите пару дней и я всё до делаю. Telegram: ```@arhefr```
