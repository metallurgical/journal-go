# Introduction
Microservice implementation to handle journal data using GRPC with PHP as client and GO as server. - Mainly used in internal application -

## Add .env
Copy from example. `cp .env.example .env`

## Setup
```bash
$ go get github.com/metallurgical/journal-go
$ make # to build services files
$ make compile # to compile code into executable file
$ make clean # clean all generated files and executable file 
```