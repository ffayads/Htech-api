# Htech-api

## Requerimientos

* go version go1.13.5 linux/amd64
* GO111MODULE="auto"

## Instalacion
````
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/secure
go get github.com/appleboy/gin-jwt
go get github.com/jinzhu/gorm
go get github.com/go-sql-driver/mysql
go get github.com/jinzhu/gorm
go get github.com/namsral/flag
go get golang.org/x/crypto/bcrypt
````

* Migracion 

La migracion se realiza en el archivo application.go en la funcion "Migrate"

## Datos necesarios para iniciar
````
Insert into type_documents value (1,"2020-04-04 18:00:00",NULL,NULL,"CC","Cedula de Ciudadania");
Insert into type_documents value (2,"2020-04-04 18:00:00",NULL,NULL,"CE","Cedula de Extranjeria");
Insert into type_documents value (3,"2020-04-04 18:00:00",NULL,NULL,"TI","Tarjeta de Identidad");
Insert into type_documents value (4,"2020-04-04 18:00:00",NULL,NULL,"NIT","NIT");
````

## Postman
Archivo con extencion .json para importar los servicios y consumir los endpoints
````
{
	"id": "6478bd60-a189-81e5-866b-172ec14e98ba",
	"name": "HTech",
	"description": "",
	"order": [
		"e1edf8fc-a70d-5c7b-f61c-6ce7988dbe2d",
		"f0283ceb-6488-b939-1ee7-0d39f3be3349",
		"914fe138-613c-4eed-a11a-814969065830",
		"6f04ea64-bd24-69ca-efd6-e3d707896a66",
		"01c79043-6bde-faed-cb79-717642688765",
		"633ddf2f-516b-5c3f-3b7e-a11a239aa879",
		"03da35d2-a5f1-974b-ebcc-d91a9ff1a8ee"
	],
	"folders": [],
	"folders_order": [],
	"timestamp": 1586305412011,
	"owner": "5553868",
	"public": false,
	"requests": [
		{
			"id": "01c79043-6bde-faed-cb79-717642688765",
			"headers": "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZlbGl4ZmF5YWRAZ21haWwuY29tIiwiZXhwIjoxNTg2Mzg2MjE4LCJpZCI6MSwib3JpZ19pYXQiOjE1ODYzODI2MTh9.NwwHEgbJJf7h_JwUHMskNtXwZvaXltluxyfp8GnlQ9o\nPostman-Token: 3aec9f7b-bdfa-4d09-af11-1f80b3e4f43d\ncache-control: no-cache\nContent-Type: application/x-www-form-urlencoded\n",
			"headerData": [
				{
					"key": "Authorization",
					"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZlbGl4ZmF5YWRAZ21haWwuY29tIiwiZXhwIjoxNTg2Mzg2MjE4LCJpZCI6MSwib3JpZ19pYXQiOjE1ODYzODI2MTh9.NwwHEgbJJf7h_JwUHMskNtXwZvaXltluxyfp8GnlQ9o",
					"description": "",
					"enabled": true
				},
				{
					"key": "Postman-Token",
					"value": "3aec9f7b-bdfa-4d09-af11-1f80b3e4f43d",
					"description": "",
					"enabled": true
				},
				{
					"key": "cache-control",
					"value": "no-cache",
					"description": "",
					"enabled": true
				},
				{
					"key": "Content-Type",
					"value": "application/x-www-form-urlencoded",
					"description": "",
					"enabled": true
				}
			],
			"url": "localhost:5000/auth/cards/getCards",
			"queryParams": [],
			"preRequestScript": null,
			"pathVariables": {},
			"pathVariableData": [],
			"method": "POST",
			"data": [
				{
					"key": "card_id",
					"value": "2",
					"description": "",
					"type": "text",
					"enabled": false
				},
				{
					"key": "clients_id",
					"value": "1",
					"description": "",
					"type": "text",
					"enabled": true
				}
			],
			"dataMode": "urlencoded",
			"version": 2,
			"tests": null,
			"currentHelper": "normal",
			"helperAttributes": {},
			"time": 1586384451310,
			"name": "Get Cards",
			"description": "",
			"collectionId": "6478bd60-a189-81e5-866b-172ec14e98ba",
			"responses": []
		},
		{
			"id": "03da35d2-a5f1-974b-ebcc-d91a9ff1a8ee",
			"headers": "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZlbGl4ZmF5YWRAZ21haWwuY29tIiwiZXhwIjoxNTg2MzgxMTc0LCJpZCI6MSwib3JpZ19pYXQiOjE1ODYzNzc1NzR9.lW6EeamWhyDOrnytl4czOT9DeSJajxKbsPv88-FuI5Q\nPostman-Token: 3aec9f7b-bdfa-4d09-af11-1f80b3e4f43d\ncache-control: no-cache\nContent-Type: application/json\n",
			"headerData": [
				{
					"key": "Authorization",
					"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZlbGl4ZmF5YWRAZ21haWwuY29tIiwiZXhwIjoxNTg2MzgxMTc0LCJpZCI6MSwib3JpZ19pYXQiOjE1ODYzNzc1NzR9.lW6EeamWhyDOrnytl4czOT9DeSJajxKbsPv88-FuI5Q",
					"description": "",
					"enabled": true
				},
				{
					"key": "Postman-Token",
					"value": "3aec9f7b-bdfa-4d09-af11-1f80b3e4f43d",
					"description": "",
					"enabled": true
				},
				{
					"key": "cache-control",
					"value": "no-cache",
					"description": "",
					"enabled": true
				},
				{
					"key": "Content-Type",
					"value": "application/json",
					"description": "",
					"enabled": true
				}
			],
			"url": "localhost:5000/auth/cards/updateCards",
			"queryParams": [],
			"pathVariables": {},
			"pathVariableData": [],
			"preRequestScript": null,
			"method": "POST",
			"collectionId": "6478bd60-a189-81e5-866b-172ec14e98ba",
			"data": [],
			"dataMode": "raw",
			"name": "Update Cards",
			"description": "",
			"descriptionFormat": "html",
			"time": 1586384516199,
			"version": 2,
			"responses": [],
			"tests": null,
			"currentHelper": "normal",
			"helperAttributes": {},
			"rawModeData": "{ \n\t\"id\": 1,\n  \"name\": \"Felix Fayad S\",\n  \"email\": \"felixfayad@gmail.com\",\n  \"phone\": \"3005359910\"\n}"
		},
		{
			"id": "633ddf2f-516b-5c3f-3b7e-a11a239aa879",
			"headers": "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZlbGl4ZmF5YWRAZ21haWwuY29tIiwiZXhwIjoxNTg2Mzg2NTc3LCJpZCI6MSwib3JpZ19pYXQiOjE1ODYzODI5Nzd9.7t8Zd7K9maEKVe-T6o-Ej4VLoI5KmLNd4MXseIZ7IFM\nPostman-Token: 3aec9f7b-bdfa-4d09-af11-1f80b3e4f43d\ncache-control: no-cache\nContent-Type: application/json\n",
			"headerData": [
				{
					"key": "Authorization",
					"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZlbGl4ZmF5YWRAZ21haWwuY29tIiwiZXhwIjoxNTg2Mzg2NTc3LCJpZCI6MSwib3JpZ19pYXQiOjE1ODYzODI5Nzd9.7t8Zd7K9maEKVe-T6o-Ej4VLoI5KmLNd4MXseIZ7IFM",
					"description": "",
					"enabled": true
				},
				{
					"key": "Postman-Token",
					"value": "3aec9f7b-bdfa-4d09-af11-1f80b3e4f43d",
					"description": "",
					"enabled": true
				},
				{
					"key": "cache-control",
					"value": "no-cache",
					"description": "",
					"enabled": true
				},
				{
					"key": "Content-Type",
					"value": "application/json",
					"description": "",
					"enabled": true
				}
			],
			"url": "localhost:5000/auth/cards/createCards",
			"queryParams": [],
			"preRequestScript": null,
			"pathVariables": {},
			"pathVariableData": [],
			"method": "POST",
			"data": [],
			"dataMode": "raw",
			"version": 2,
			"tests": null,
			"currentHelper": "normal",
			"helperAttributes": {},
			"time": 1586384453714,
			"name": "Create Cards",
			"description": "",
			"collectionId": "6478bd60-a189-81e5-866b-172ec14e98ba",
			"responses": [],
			"rawModeData": "{ \n  \"number\": \"1234 5678 9012 3456\",\n  \"expiry\": \"10/04\",\n  \"cvv\": \"557\",\n  \"clients_id\": 1\n}"
		},
		{
			"id": "6f04ea64-bd24-69ca-efd6-e3d707896a66",
			"headers": "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZlbGl4ZmF5YWRAZ21haWwuY29tIiwiZXhwIjoxNTg2MzgxMTc0LCJpZCI6MSwib3JpZ19pYXQiOjE1ODYzNzc1NzR9.lW6EeamWhyDOrnytl4czOT9DeSJajxKbsPv88-FuI5Q\nPostman-Token: 3aec9f7b-bdfa-4d09-af11-1f80b3e4f43d\ncache-control: no-cache\nContent-Type: application/json\n",
			"headerData": [
				{
					"key": "Authorization",
					"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZlbGl4ZmF5YWRAZ21haWwuY29tIiwiZXhwIjoxNTg2MzgxMTc0LCJpZCI6MSwib3JpZ19pYXQiOjE1ODYzNzc1NzR9.lW6EeamWhyDOrnytl4czOT9DeSJajxKbsPv88-FuI5Q",
					"description": "",
					"enabled": true
				},
				{
					"key": "Postman-Token",
					"value": "3aec9f7b-bdfa-4d09-af11-1f80b3e4f43d",
					"description": "",
					"enabled": true
				},
				{
					"key": "cache-control",
					"value": "no-cache",
					"description": "",
					"enabled": true
				},
				{
					"key": "Content-Type",
					"value": "application/json",
					"description": "",
					"enabled": true
				}
			],
			"url": "localhost:5000/auth/clients/updateClient",
			"queryParams": [],
			"preRequestScript": null,
			"pathVariables": {},
			"pathVariableData": [],
			"method": "POST",
			"data": [],
			"dataMode": "raw",
			"version": 2,
			"tests": null,
			"currentHelper": "normal",
			"helperAttributes": {},
			"time": 1586382649417,
			"name": "Update Client",
			"description": "",
			"collectionId": "6478bd60-a189-81e5-866b-172ec14e98ba",
			"responses": [],
			"rawModeData": "{ \n\t\"id\": 1,\n  \"name\": \"Felix Fayad S\",\n  \"email\": \"felixfayad@gmail.com\",\n  \"phone\": \"3005359910\"\n}"
		},
		{
			"id": "914fe138-613c-4eed-a11a-814969065830",
			"headers": "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZlbGl4ZmF5YWRAZ21haWwuY29tIiwiZXhwIjoxNTg2MzY5NDk4LCJpZCI6MSwib3JpZ19pYXQiOjE1ODYzNjU4OTh9.E4AFrqmC1gLdvYrlq_U1ceppxHxPzcwOonNh0-pCtqc\nPostman-Token: 3aec9f7b-bdfa-4d09-af11-1f80b3e4f43d\ncache-control: no-cache\nContent-Type: application/json\n",
			"headerData": [
				{
					"key": "Authorization",
					"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZlbGl4ZmF5YWRAZ21haWwuY29tIiwiZXhwIjoxNTg2MzY5NDk4LCJpZCI6MSwib3JpZ19pYXQiOjE1ODYzNjU4OTh9.E4AFrqmC1gLdvYrlq_U1ceppxHxPzcwOonNh0-pCtqc",
					"description": "",
					"enabled": true
				},
				{
					"key": "Postman-Token",
					"value": "3aec9f7b-bdfa-4d09-af11-1f80b3e4f43d",
					"description": "",
					"enabled": true
				},
				{
					"key": "cache-control",
					"value": "no-cache",
					"description": "",
					"enabled": true
				},
				{
					"key": "Content-Type",
					"value": "application/json",
					"description": "",
					"enabled": true
				}
			],
			"url": "localhost:5000/auth/clients/createClient",
			"queryParams": [],
			"preRequestScript": null,
			"pathVariables": {},
			"pathVariableData": [],
			"method": "POST",
			"data": [],
			"dataMode": "raw",
			"version": 2,
			"tests": null,
			"currentHelper": "normal",
			"helperAttributes": {},
			"time": 1586367512694,
			"name": "Create Clients",
			"description": "",
			"collectionId": "6478bd60-a189-81e5-866b-172ec14e98ba",
			"responses": [],
			"rawModeData": "{ \n  \"name\": \"Felix Fayad\",\n  \"document\": \"1070618746\",\n  \"type_document\": \"CC\",\n  \"email\": \"felixfayad@gmail.com\",\n  \"phone\": \"3005359910\"\n}"
		},
		{
			"id": "e1edf8fc-a70d-5c7b-f61c-6ce7988dbe2d",
			"headers": "",
			"headerData": [],
			"url": "localhost:5000/login",
			"queryParams": [],
			"preRequestScript": null,
			"pathVariables": {},
			"pathVariableData": [],
			"method": "POST",
			"data": [
				{
					"key": "email",
					"value": "felixfayad@gmail.com",
					"description": "",
					"type": "text",
					"enabled": true
				},
				{
					"key": "password",
					"value": "$Password123",
					"description": "",
					"type": "text",
					"enabled": true
				}
			],
			"dataMode": "params",
			"version": 2,
			"tests": null,
			"currentHelper": "normal",
			"helperAttributes": {},
			"time": 1586309638008,
			"name": "Login",
			"description": "",
			"collectionId": "6478bd60-a189-81e5-866b-172ec14e98ba",
			"responses": []
		},
		{
			"id": "f0283ceb-6488-b939-1ee7-0d39f3be3349",
			"headers": "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZlbGl4ZmF5YWRAZ21haWwuY29tIiwiZXhwIjoxNTg2MzY5NDk4LCJpZCI6MSwib3JpZ19pYXQiOjE1ODYzNjU4OTh9.E4AFrqmC1gLdvYrlq_U1ceppxHxPzcwOonNh0-pCtqc\nPostman-Token: 3aec9f7b-bdfa-4d09-af11-1f80b3e4f43d\ncache-control: no-cache\nContent-Type: application/x-www-form-urlencoded\n",
			"headerData": [
				{
					"key": "Authorization",
					"value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZlbGl4ZmF5YWRAZ21haWwuY29tIiwiZXhwIjoxNTg2MzY5NDk4LCJpZCI6MSwib3JpZ19pYXQiOjE1ODYzNjU4OTh9.E4AFrqmC1gLdvYrlq_U1ceppxHxPzcwOonNh0-pCtqc",
					"description": "",
					"enabled": true
				},
				{
					"key": "Postman-Token",
					"value": "3aec9f7b-bdfa-4d09-af11-1f80b3e4f43d",
					"description": "",
					"enabled": true
				},
				{
					"key": "cache-control",
					"value": "no-cache",
					"description": "",
					"enabled": true
				},
				{
					"key": "Content-Type",
					"value": "application/x-www-form-urlencoded",
					"description": "",
					"enabled": true
				}
			],
			"url": "localhost:5000/auth/clients/getClients",
			"queryParams": [],
			"preRequestScript": null,
			"pathVariables": {},
			"pathVariableData": [],
			"method": "POST",
			"data": [
				{
					"key": "document",
					"value": "",
					"description": "",
					"enabled": true
				},
				{
					"key": "type_document",
					"value": "",
					"description": "",
					"enabled": true
				}
			],
			"dataMode": "urlencoded",
			"version": 2,
			"tests": null,
			"currentHelper": "normal",
			"helperAttributes": {},
			"time": 1586367509769,
			"name": "Get Clients",
			"description": "",
			"collectionId": "6478bd60-a189-81e5-866b-172ec14e98ba",
			"responses": []
		}
	]
}
````
