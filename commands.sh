set -ex
export HASHCASH_BITS=20
export HASHCASH_SALT_LENGTH=10
server(){
  cd apps/server
  HASHCASH_BITS=${HASHCASH_BITS}\
  go run server.go
}

client(){
  cd apps/client
  SERVER_ADDR=localhost:9001 go run client.go
}

comp(){
  docker-compose build
  docker-compose up --force-recreate
}

"$@"