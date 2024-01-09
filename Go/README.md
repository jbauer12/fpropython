# Checkers Game with Minimax Algorithm


This code implements the board game checkers with the minimax algorithm in Go

## Installation

After the Installation of Go (used 1.21.5) [Go](https://go.dev/dl/)

```bash
#Go to the right directory.
cd Go

#After that run the application
go run main.go

```
Usually it should download all dependencies automatically, but if it is not the case run
```bash
#Install those explicitly
go mod download


```
Then you can rerun the application. 


## Usage

```bash
#Go to the right directory.
cd Go

#After that run the application
go run main.go

```
## Usage Test
```bash
#run all tests at once
go test ./...

# run test from specific directory
# for example gameboard directory
cd gameboard 
go test
cd ..

```



go tool cover -func ./coverage.out

go test ./...  -coverpkg=./... -coverprofile ./coverage.out