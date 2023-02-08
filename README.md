# Delos Far App
## List of environment to use
To use this program, you will need:
- go (v19.5)
- Docker
- PostgreSQL (for non docker run)

This program is created with hexagonal architecture.

## Run with Docker
To run this program you will need to do :
- Set value `dsn` variable in `main.go` :
    - My value is : `dsn := "host=db user=postgres password=developmentdb dbname=delos_farm port=5432 sslmode=disable TimeZone=Asia/Jakarta"`
    - You must change it according to what you want. **BUT**, you mus change the docker-compose.yml too 
    - **NOTE** : I tried using .env variables, but always return `too many parameters`
- Go to root dir of the project
- Execute command `docker-compose up -d`
- Wait until installation is complete
- The program should be work properly in port `9090`
## Run with go run command
To run this program you will need to do :
- Set value `dsn` variable in `main.go` :
    - My value is : `dsn := "host=localhost user=postgres password=development dbname=delos_farm port=5432 sslmode=disable TimeZone=Asia/Jakarta"`
    - You must change it according to your local database
    - **NOTE** : I tried using .env variables, but always return `too many parameters`
- Make sure dabase is already created in your PostgreSQL
- Go to root dir of the project
- Execute command `go run main.go`
- The program should be work properly in port `9090`

## List of API
There are 3 sub-groups APIs :
- Farm (`/v1/farm`)
    - (GET) `/v1/farm/:id` get Farm data by its ID
    - (GET) `/v1/farm` get list of all Farms
    - (POST) `/v1/farm` create new Farm data
    - (PUT) `/v1/farm` update Farm data
    - (DELETE) `/v1/farm/:id` soft delete Farm data
        - Notes : if Farm is delete, all ponds inside the farm will also deleted

- Pond (`/v1/pond`)
    - (GET) `/v1/pond/:id` get Pond data by its ID
    - (GET) `/v1/pond` get list of all Ponds
    - (GET) `/v1/pond/farm/:id` get Ponds data by its Farm ID
    - (POST) `/v1/pond` create new Pond data
    - (PUT) `/v1/pond` update Pond data
    - (DELETE) `/v1/pond/:id` soft delete Farm data By its ID
    - (DELETE) `/v1/pond/farm/:id` soft delete Farm data By Farm ID
- Monitoring (`/v1/monitor`)
    - (GET) `/v1/monitor` get all monitoring data

You can find the postman collection inside `postman_collection` directory

## User Agent
to add user agent, you must add header `agent` in every request

<p align="center">
  <img src="./postman_collection/agent.png">
</p>

---------------------------------------------------------------------------
- This project is created By Ahmad Mujahid A
- For Delos Task Assignment [Backend]
- I just familiarize my self with Golang for just about 2 months. (I'm proficient in Java, Kotlin, Python, OOP concept, microservice arch, hexagonal arc, etc) It will be great if you give me a chance.