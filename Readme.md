## This is a sample implementation of golang grpc services

This code contains both the client and server implementation of grpc services in golang.
There are 2 services implemented in this code:
1.  **GreetService** - This service contains 3 methods:
    - Greet - This method takes a name and returns a greeting message. It is a unary RPC.
    - GreetManyTimes - This method takes a name and returns a stream of greeting messages. It is a server streaming RPC.
    - LongGreet - This method takes a stream of names and returns a single greeting message. It is a client streaming RPC.
    - GreetEveryone - This method takes a stream of names and returns a stream of greeting messages. It is a bidirectional streaming RPC.
    - GreetWithDeadline - This method takes a name and returns a greeting message. It is a unary RPC with a deadline. If the deadline is exceeded, the server will return an error.
2. **CalulatorService** - This service contains 4 methods:
    - Sum - This method takes two numbers and returns their sum. It is a unary RPC.
    - PrimeNumberDecomposition - This method takes a number and returns a stream of prime factors. It is a server streaming RPC.
    - ComputeAverage - This method takes a stream of numbers and returns their average. It is a client streaming RPC.
    - FindMaximum - This method takes a stream of numbers and returns the maximum number. It is a bidirectional streaming RPC.
3. **BlogService** - This service contains 5 methods:
    - CreateBlog - This method takes a blog and returns the blog with an ID. It is a unary RPC.
    - ReadBlog - This method takes an ID and returns the blog. It is a unary RPC.
    - ListBlog - This method takes a filer and returns a stream of blogs. It is a server streaming RPC.


## Building the code 

To build the code, you need to have the following installed:
- Go 1.16 or later
- Protocol Buffers 3.15 or later
- Go plugins for Protocol Buffers

Before building the go binary, make sure the certs are available , to generate certs run the following command
- make certs

This should generate ssl certs under ssl folder

To build the code, run the following commands:
- `make greet` - This will build the greet service
- `make calculator` - This will build the calculator service

## Running the code

To run the code, you need to run the following commands
```
/bin/greet/server
``` 
This will start the greet server

```
/bin/greet/client
```
This will start the greet client (run this on seperate terminal)

```
/bin/calculator/server ## This will start the calculator server
/bin/calculator/client ## This will start the calculator client (run this on seperate terminal)
```

### blog service

To run the blog serivce, you need to run the following dependencies
- MongoDB
- MongoExpress (optional)
- Docker / Podman compose
- grpcui (optional) - for viewing the response in browser

Install Docker compose /podman compose and then run the following command

```sh 
    make compose-up
```

This will start the mongo and mongo express server. You can access the mongo express server at http://localhost:8081
To run the blog service, run the following commands

```sh
    make blog
    /bin/blog/server
```
This will start the blog server

To run the blog client, run the following commands

```sh
  brew install grpcui
  grpcui -insecure localhost:50051
```
This will start the blog client and you can access the client at browser 



