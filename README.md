# Concurrent calculator
![Слайд1](https://github.com/user-attachments/assets/a1250ba0-7909-42aa-8cdf-887946252dac)
# **Content**
  - [**Description**](#description)
  - [**Run Server**](#run-server)
  - [**Endpoints**](#endpoints)
  - [**Feedback**](#feedback)

# **Description**
>[!NOTE]
>![Слайд2](https://github.com/user-attachments/assets/264eea11-d04c-4ce0-ab9a-c407b7a23b94)
> ```mermaid
>graph LR
>Orchestrator("Orchestrator") <-----> |"Task/Result"| Agent1("Agent 1")
>Orchestrator <-----> |"Task/Result"| ...(...)
>Orchestrator <-----> |"Task/Result"| AgentN("Agent N")
>```

# **Run Server**
>[!IMPORTANT]
>### Download Golang 1.23.4+
> ``` shell
> cd calculator
> ```
> ``` shell
> go run cmd/server/main.go
> ```
> 
# **Endpoints**
> [!IMPORTANT]
>> ## POST /api/v1/calculate
>> | Request | Response |
>> | ------- | -------- |
>> | ```curl -X POST -H "Content-type:application/json" --data "{\"expression\":\"...\"}" http://localhost:8080/api/v1/calculate``` | ``` {"id": ...} ``` |
>> 
>> | Status | Response | Reason |
>> | ------ | ------ | ------ |
>> | 200 | ```{"id": ...}``` | OK |
>> | 422 | ```{"error":"cannot reading body"}``` OR ```{"error":"incorrect JSON"}``` | incorrect body request or JSON |
>> ```mermaid
>>graph LR
>>
>>Client[["Client"]] -----> |"POST Request (Expression)"| Server("Server")
>>
>>Server -----> |"Response (Id)"| Client
>>```
>> ### Example:
>> ``` shell
>> curl -X POST -H "Content-type:application/json" --data "{\"expression\":\"2+2\"}" http://localhost:8080/api/v1/calculate
>> ```
>
>> ## GET /api/v1/expressions
>> | Request | Response |
>> | ------- | -------- |
>> | ```curl localhost:8080/api/v1/expressions``` | ``` {"expressions": [{"id": ..., "status": "...", "result": "..."}, ... ]} ``` |
>>
>> | Status | Reason |
>> | ------ | ------ |
>> | 200 | ```OK``` |
>> ```mermaid
>>graph LR
>>
>>Client[["Client"]] -----> |"GET Request"| Server("Server")
>>
>>Server -----> |"Response (Array Status)"| Client
>>```
>> ### Example:
>> ``` shell
>> curl localhost:8080/api/v1/expressions
>> ```
> 
>> ## GET /api/v1/expressions/{id}
>> | Request | Response |
>> | ------- | -------- |
>> | ```curl localhost:8080/api/v1/expressions/id``` | ``` {"expression": {"id": ..., "status": "...", "result": "..."}} ``` |
>>
>> | Status | Response | Reason |
>> | ------ | ------ | ----- |
>> | 200 | ```{"expression": {"id": ..., "status": "...", "result": "..."}}``` | OK |
>> | 422 | ```{"error":"not found task"}``` | incorrect ID of math expression |
>> ```mermaid
>>graph LR
>>
>>Client[["Client"]] -----> |"GET Request (ID)"| Server("Server")
>>
>>Server -----> |"Response (Status)"| Client
>>```
>> ### Example:
>> ``` shell
>> curl localhost:8080/api/v1/expressions/id
>> ```

## **Feedback**
### [Telegram](https://t.me/arhefr)
