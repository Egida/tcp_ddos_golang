# TCP golang server with POW DOS protection
Hascash was chosen as simplest one among other POW algorithms 
because here in my opinion any other more complex POW algorithm is not required here.

The client sends hashcash stamp to server. 
The server checks the stamp and in case it's valid, it 
responds with a random joke from the list of Dad jokes.
# How to run:
```
docker-compose build
docker-compose up --force-recreate
```
output:
```
client_1  | 2023/03/12 05:03:44 common.go:34: bits: 20, saltLength: 10, extra:
server_1  | 2023/03/12 05:03:41 common.go:34: bits: 20, saltLength: 10, extra:
server_1  | 2023/03/12 05:03:41 server.go:39: loaded 250 jokes
server_1  | 2023/03/12 05:03:41 server.go:106: listening tcp :9001 ....
client_1  | 2023/03/12 05:03:46 client.go:40: time spent on generating the stamp: 1.623595s
server_1  | 2023/03/12 05:03:46 server.go:64: received from client: 1:20:230312:client_id::HSgqjgFj1B:136984
client_1  | 2023/03/12 05:03:46 client.go:51: reply from server= What do you call a man with a rubber toe?<>Roberto
client_1  | 2023/03/12 05:03:46 client.go:53: sleeping for 10s
client_1  | 2023/03/12 05:03:56 client.go:40: time spent on generating the stamp: 578.2302ms
server_1  | 2023/03/12 05:03:56 server.go:64: received from client: 1:20:230312:client_id::RJTm+6VBzD:68c7e
client_1  | 2023/03/12 05:03:56 client.go:51: reply from server= Just watched a documentary about beavers.<>It was the best damn program I’ve ever seen.
client_1  | 2023/03/12 05:03:56 client.go:53: sleeping for 10s
client_1  | 2023/03/12 05:04:08 client.go:40: time spent on generating the stamp: 1.6741424s
server_1  | 2023/03/12 05:04:08 server.go:64: received from client: 1:20:230312:client_id::G31i3pZ6Xs:120a41
client_1  | 2023/03/12 05:04:08 client.go:51: reply from server= I could never be a plumber<>it’s too hard watching your life’s work go down the drain.
client_1  | 2023/03/12 05:04:08 client.go:53: sleeping for 10s
client_1  | 2023/03/12 05:04:22 client.go:40: time spent on generating the stamp: 3.6743512s
server_1  | 2023/03/12 05:04:22 server.go:64: received from client: 1:20:230312:client_id::8QhA5qVsYl:2c2afc
client_1  | 2023/03/12 05:04:22 client.go:51: reply from server= I used to be addicted to the hokey pokey<> but I turned myself around.
client_1  | 2023/03/12 05:04:22 client.go:53: sleeping for 10s
```
# Run tests
```go test -v ./...```
outputs
```
?       github.com/ypapax/tcp_ddos_golang/apps/client   [no test files]
=== RUN   TestTcpServe
2023/03/12 19:42:36 common.go:36: bits: 20, saltLength: 8, extra:
2023/03/12 19:42:36 common.go:36: bits: 20, saltLength: 8, extra:
2023/03/12 19:42:36 server.go:45: loaded 250 jokes
2023/03/12 19:42:36 server.go:112: listening tcp :9002 ....
=== RUN   TestTcpServe/0
2023/03/12 19:42:37 server.go:70: received from client: 1:20:230312:test::LmiI7Upm:db6f6
    server_test.go:30: a: What do you call someone with no body and no nose? <>Nobody knows.
=== RUN   TestTcpServe/1
2023/03/12 19:42:37 server.go:70: received from client: 1:20:230312:test::LmiI7Upm:db6f6
2023/03/12 19:42:37 server.go:80: this token was already detected in the past 1:20:230312:test::LmiI7Upm:db6f6 : 2023-03-12 19:42:37.849357 +0600 +06 m=+1.148311542
    server_test.go:30: a: the request is not verified by proof of work hashcash
=== RUN   TestTcpServe/2
2023/03/12 19:42:37 server.go:70: received from client: some random stamp
    server_test.go:30: a: the request is not verified by proof of work hashcash
--- PASS: TestTcpServe (1.15s)
    --- PASS: TestTcpServe/0 (0.01s)
    --- PASS: TestTcpServe/1 (0.00s)
    --- PASS: TestTcpServe/2 (0.00s)
PASS
ok      github.com/ypapax/tcp_ddos_golang/apps/server   1.632s
?       github.com/ypapax/tcp_ddos_golang/common    [no test files]
```
# Links
https://en.wikipedia.org/wiki/Proof_of_work

https://github.com/catalinc/hashcash

https://raw.githubusercontent.com/yesinteractive/dadjokes/master/controllers/jokes.txt