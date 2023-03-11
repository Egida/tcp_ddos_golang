set -ex

server(){
  cd apps/server
  go run server.go
}

client(){
  cd apps/client
  SERVER_ADDR=localhost:9001 go run client.go
}

comp(){
  stack=tcp_ddos
  docker-compose build
  set +e; docker stack rm $stack; set -e;
#  docker stack deploy --force-recreate -c docker-compose.yml $stack
  docker-compose up --force-recreate
  docker ps
}

"$@"