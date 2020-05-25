# run app
make deps
make run

# test client side
open several terminal
`telnet localhost 8080` connect to local server, then can send messages
```
$ go test -cover -coverprofile=testcover.out
$ go tool cover -html=testcover.out
```
