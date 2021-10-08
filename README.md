App for display  parameter from URL

This application starts an http server, and greets the user by name, which is entered as a URL parameter.
For example, following the path: `http: //0.0.0.0: 2000 / Bob`, you will see" Hello, Bob ".
To start the application, you need to start the terminal, go to the cmd directory and run the command:
 
 ````
 go run main.go
 ````

To run application tests, you need to start the terminal, go to the `internal / app` directory and run the command: 

````
go test app_test.go
````

Server config can be changed in .env file
