# challenge-goapi

## Getting started

Install Go https://go.dev/dl

## Usage Steps

- Clone the project with the following command

```
git clone https://git.enigmacamp.com/enigma-20/ridho-mahendra/challenge-goapi.git
```

- Get required module with this command

```
go get github.com/lib/pq
```

```
go get github.com/gin-gonic/gin
```

- Execute the sql query on db directory (DDL.sql & DML.sql)

- Adjust the parameter host, port, user etc on main.go for database connection.

```
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "enigma_laundry"
)
```

- Run the program with the following command

```
go run .
```

### Customer API

#### Get Customer all

Request :

- Method : GET
- Endpoint : `api/customer`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
{
  "data": [
    {
      "id": 1,
      "name": "string",
      "phone": "string",
      "address": "string"
    },
    {
      "id": 2,
      "name": "String",
      "phone": "String",
      "address": "String"
    }
  ]
}
```

#### Get Customer by Name

Request :

- Method : GET
- Endpoint : `api/customer/:name`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
{
  "data": [
    {
      "id": 2,
      "name": "string",
      "phone": "string",
      "address": "string"
    },
    {
      "id": 3,
      "name": "String",
      "phone": "String",
      "address": "String"
    }
  ]
}
```

#### Add Customer

Request :

- Method : POST
- Endpoint : `api/customer`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
{
    "id": 9,
    "name": "string",
    "phone": "string",
    "address": "string"
}
```

#### Update Customer

Request :

- Method : PUT
- Endpoint : `api/customer/:id`
- Header :
  - Accept : application/json
- Body
```json
{
    "name": "string",
    "phone": "string",
    "address": "string"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
    "id": 9,
    "name": "string",
    "phone": "string",
    "address": "string"
}
```

#### Delete Customer

Request :

- Method : DELETE
- Endpoint : `api/customer/:id`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
{
    "id": 9,
    "status": "string"
}
```

### Transaction Laundry API

#### Add Transaction

Request :

- Method : POST
- Endpoint : `api/transaction`
- Header :
  - Accept : application/json
- Body :
```json
{
    "unit": "string",
    "amount": 2,
    "date_in": "2023-01-20",
    "date_out": "2023-01-23",
    "id_customer": 0,
    "id_service": 9
}
```

Response :

- Status : 201 Created
- Body :

```json
{
    "data": {
        "id": 9,
        "unit": "string",
        "amount": 2,
        "date_in": "2023-01-20",
        "date_out": "2023-01-23",
        "id_customer": 0,
        "id_service": 9
    },
    "status": "insert success"
}
```

#### Get Transaction All

Request :

- Method : GET
- Endpoint : `api/transaction`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
[
    {
        "id": 2,
        "unit": "Kg",
        "amount": 2,
        "date_in": "0001-01-01T00:00:00Z",
        "date_out": "0001-01-01T00:00:00Z",
        "id_customer": 1,
        "id_service": 3
    },
    {
        "id": 3,
        "unit": "Ton",
        "amount": 1,
        "date_in": "2023-11-20T00:00:00Z",
        "date_out": "2023-11-22T00:00:00Z",
        "id_customer": 2,
        "id_service": 3
    }
]
```

#### Get Transaction by id

Request :

- Method : GET
- Endpoint : `api/transaction?id=`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
[
    {
        "id": 9,
        "unit": "string",
        "amount": 2,
        "date_in": "2023-01-20T00:00:00Z",
        "date_out": "2023-01-23T00:00:00Z",
        "id_customer": 0,
        "id_service": 9
    }
]
```

#### Delete Transaction by id

Request :

- Method : DELETE
- Endpoint : `api/service/:id`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
{
    "id": 9,
    "status": "delete success"
}
```

### Service API

#### Add Service

Request :

- Method : POST
- Endpoint : `api/service`
- Header :
  - Accept : application/json
- Body :
```json
{
    "service": "string",
    "price": 15000
}
```

Response :

- Status : 201 Created
- Body :

```json
{
    "data": {
        "id": 13,
        "service": "string",
        "price": 15000
    },
    "status": "insert success"
}
```

#### Get Service All

Request :

- Method : GET
- Endpoint : `api/service`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
[
    {
        "id": 3,
        "service": "string",
        "price": 10000
    },
    {
        "id": 5,
        "service": "string",
        "price": 4000
    }
]
```

#### Get Service by name of service

Request :

- Method : GET
- Endpoint : `api/service?service=`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
[
    {
        "id": 3,
        "service": "string",
        "price": 10000
    },
    {
        "id": 5,
        "service": "string",
        "price": 4000
    }
]
```

#### Delete Service by id

Request :

- Method : DELETE
- Endpoint : `api/service/:id`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
{
    "id": 8,
    "status": "delete success"
}
```