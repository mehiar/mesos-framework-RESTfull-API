# Mesos Framework RESTful API

This code repository provides a simple RESTful API for mesos written in Go. It is based on the code description in the [The New Stack](http://thenewstack.io/make-a-restful-json-api-go/) article, but was modified to log and track Docker container requests.


## Running the Code
1. Change your working directory
```sh
cd Path/To/mesos-framework-restful-api/src
```
2. Get dependent packages:
```sh
go get
```
3. Run codes\\\\\\\\

```sh
# Run codes
go run *.go
```

## Using the RESTful API
Now that the codes are running, if you enter the address: 'localhost:8080' in your web browser, you should recive a welcome page. 

To submit a container request, execute the following from another terminal window on the host where the codes are running:
```sh
curl -H "Content-Type: application/json" -d '{"name":"docker-task-1","status":"submitted", "cpus":1.5, "mem":256}' http://localhost:8080/container-requests
```

The list of submitted containers can be found under: 'localhost:8080/container-requests'

Information about the docker request whose ID is 1 can be obtained under: 'localhost:8080/container-requests/1'
 


