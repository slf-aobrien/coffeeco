# Coffee Company
## Description
Test project from the book Domain Driven Design with GoLang

## Instructions
### Setup
I worked through this chapter differently.  I wrote the code before initializing and running the project here are the notes to make it runnable:
```
 go mod init github.com/slf-aobrien/coffeeco
 go mod tidy
```
Based on the examples of the book all of the local import statements had to be changed from something like this:
```
"coffeeco/internal/loyalty"
```
to
```
"github.com/slf-aobrien/coffeeco/internal/loyalty"
```
### Building
#### Requirements
* GoLang version: go version go1.20.2 windows/amd64  
#### Notes

### Testing
### Running
step 1) 
Start the Mongo Container
```
docker compose up
```
step 2)
Execute the code:
```
go build cmd/main.go
```

## Chapter Notes
* Chapter 5 - Implement a charge service other than Strip (maybe square)
* [Book Code Repo](https://github.com/PacktPublishing/Domain-Driven-Design-with-GoLang/tree/main/chapter5)