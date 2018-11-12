# go-starter
GoLang example for starter

```go build``` --> generate *.exe (manually run *.exe)


```go run``` --> generate *exe & run 

```go install``` -> move *.exe to GOBIN


## Syntax Note
- Go use compiler for build *.go into platform(OS)'s binary (executable).
- unused import, variable will throw compile error.
- ***panic*** is name of runtime error.
- multiple return value allowed (most found for *error).
- Go use copy value to assign variable (also copy all in struct).
- Must use pointer(```&``` reference, ```*``` dereference) if need to reference to address.
- ```defer``` use for close resource(s) [eachwordcountfromfile.go](./workshop/eachwordcountfromfile.go).
- Go not have Inheritance but Polymophism available.
- Method Receiver use same meaning of ```this``` in Java [custom_type.go](./custom_type.go).
- ```nil``` is no initiate variable.
- goroutine is very small concurrency work.
- ```go``` keyword use for create new goroutine [Concurrency Example](./concurrency/main.go).
- ```var ch chan int``` or ```ch := make(chan int)``` is creation of channel variable with type int.
- channel use for communication of goroutine [channel.go](./channel.go) 
- ```make``` can use with channel, slice, map.
- Go can do cross-compile (build other OS' binary).
- Go binary larger than C, cause it bundle runtime in it (sometimes C need sharelib eg. *.dll).