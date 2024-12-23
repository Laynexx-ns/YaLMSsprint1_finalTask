# YALMS Calculator (Sprint1)


>[!NOTE]
> В конце подарочек)


## Как запустить?

#### склонируйте репозиторий:
ssh:
~~~
git clone git@github.com:Laynexx-ns/gocalc_LMS24.git
~~~
https:
~~~
 git clone https://github.com/Laynexx-ns/gocalc_LMS24.git
~~~


#### введите эту команду 
~~~
go run .\cmd\main.go
 ~~~

## Архитектура проекта:

- *root/cmd* - основной `go` файл, в котором я создаю сам сервер (библиотека gin) 
- root/internal/api:
	- /handlers: основной handler и middleware для проверки валидности запроса и создании самого response
	- /models: шуточная папка с очень важной информацией
	- /services: содержит функцию, которая напрямую обращается к calc с уже обработанным expression
- root/pkg:
	- /calc: содержит сам калькулятор и его тесты
	- /errors: кастомные ошибки (no usages, lol)


### Где находятся тесты?

* Для сервера - `root/internal/test/app_test.go
* Для калькулятора - `root/pkg/calc/calc_test.go

___

>[!WARNING]
>Сам калькулятор работает не идеально.
>Например: он не сможет правильно, либо вообще обработать запрос по типу - "2 + (-1)"
>

___

## Как проверить работоспособность? 


### Отправить запрос через Insomnia/postman/httpie:


1. Выбери такой тип запроса:
`POST - http://localhost:8080/api/v1/calculate`
2. Добавь в body типа JSON что-то вроде такого:
 `{
  "expression": "2+2*2+(3+54/2)"
}
`
3. Получи response `{
	"result": 36
}`

>[!TIP]
>httpie - не самое популярное решение. Однако он имеет самый красивый, интуитивный и  удобный интерфейс. Так же радует наличие ИИ возможностей и удобной сортировки запросов
>


### Curl
Пример запроса curl:
~~~shell
curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"(2+2)*2\"}"

{"result":8}
~~~




# Примеры 


### Пример 1

req
~~~
{
  "expression" : "2+3/2"
}
~~~
res
~~~
{
  "result": 3.5
}
~~~

### Пример 2

req
~~~
{
  "expression" : "2+3/2/0"
}
~~~
res
~~~
422
{
  "error": "Dividing by zero"
}
~~~

### Пример 3

req
~~~
{
  "expression" : "2+(3/2"
}
~~~
res
~~~
422
{
  "error": "Incorrect brackets"
}
~~~

### Пример 4

req
~~~
{
  "expression" : 2+3/2"
}
~~~
res
~~~
{
  "error": "Invalid JSON format"
}
~~~



> [!CAUTION]
> Примеры неправильных запросов curl:
> 
> 1.`curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"(2+2)*2\)"}"
{"error":"Invalid JSON format"}` (400)
> 
> 2.`curl --location "http://localhost:8080/api/v1/calculate" --header "Content-type: application/json" --data "{\"expression\" : \"2+2*2)\"}"
{"error":"Incorrect brackets"}` (422)
>



# Подарок

Вот репозиторий с огромным количеством бесплатных API:
https://github.com/public-apis/public-apis