# tire-change-workshop
[![CircleCI](https://circleci.com/gh/Surmus/tire-change-workshop.svg?style=svg)](https://circleci.com/gh/Surmus/tire-change-workshop)
[![codecov](https://codecov.io/gh/Surmus/tire-change-workshop/branch/master/graph/badge.svg)](https://codecov.io/gh/Surmus/tire-change-workshop)
[![Go Report Card](https://goreportcard.com/badge/github.com/surmus/tire-change-workshop)](https://goreportcard.com/report/github.com/surmus/tire-change-workshop)
[![Release](https://img.shields.io/github/release/surmus/tire-change-workshop.svg?style=flat-square)](https://github.com/surmus/tire-change-workshop/releases/latest)

Provides tire change workshop backend API's for Tire change booking practise assigment application

## General info
Project contains two independent web server applications.

Each application is responsible for providing automobile tire change times and ability to book aforementioned tire change times 
through exposed REST API interface.

## Usage
### Download server binaries 
Choose one of the following options:

#### Download Github release from https://github.com/Surmus/tire-change-workshop/releases
##### When running Windows
1. Extract win64 folder contents from downloaded release.tar.gz
2. Run application binaries:
     ```sh
     london-server.exe
     manchester-server.exe
     ```
3. Applications should be accessible from:
     Manchester tire workshop - http://localhost:9003/swagger/index.html
     London tire workshop - http://localhost:9004/swagger/index.html
##### When running Linux
1. Extract linux64 folder contents from downloaded release.tar.gz
2. Run application binaries:
     ```sh
     ./london-server
     ./manchester-server
     ```
3. Applications should be accessible from:
     Manchester tire workshop - http://localhost:9003/swagger/index.html
     London tire workshop - http://localhost:9004/swagger/index.html     
#### Using Docker
1. Run docker images
    ```sh
    $ docker run -d -p 9003:80 surmus/london-tire-workshop:latest
    $ docker run -d -p 9004:80 surmus/manchester-tire-workshop:latest
    ```
2. Applications should be accessible from:
     Manchester tire workshop - http://localhost:9003/swagger/index.html
     London tire workshop - http://localhost:9004/swagger/index.html        
#### Compile and install from source code
1. Install Golang https://golang.org/
2. Download source and compile binaries:
    ```sh
    $ go get -u github.com/surmus/tire-change-workshop/cmd/london-server
    ```
    ```sh
    $ go get -u github.com/surmus/tire-change-workshop/cmd/manchester-server
    ```
3.  Run server applications (NB: check that GOBIN env variable is set and added to the PATH)
    ```sh
    $ london-server
    $ manchester-server
    ```

4. Applications should be accessible from:
     Manchester tire workshop - http://localhost:9003/swagger/index.html
     London tire workshop - http://localhost:9004/swagger/index.html

## CLI options
```sh
$ ./london-server help
  NAME:
     london-server - London tire workshop API server
  
  USAGE:
     london-server [global options] command [command options] [arguments...]
  
  VERSION:
     v2.0.0
  
  COMMANDS:
     help, h  Shows a list of commands or help for one command
  
  GLOBAL OPTIONS:
     --port value, -p value  Port for server to listen incoming connections (default: "9003")
     --verbose      Enables debug messages print with SQL logging (default: false)
     --help, -h              show help
     --version, -v           print the version
```
```sh
$ ./manchester-server help
NAME:
   manchester-server - Manchester tire workshop API server

USAGE:
   manchester-server [global options] command [command options] [arguments...]

VERSION:
   v2.0.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value  Port for server to listen incoming connections (default: "9004")
   --verbose      Enables debug messages print with SQL logging (default: false)
   --help, -h              show help
   --version, -v           print the version
```

## API documentation
Documentation is provided for both applications by Swagger and can be accessed at ``http://localhost:{APPLICATION_PORT}/swagger/index.html`` 
