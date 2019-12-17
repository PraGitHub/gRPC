log(){
    echo "`date ` ${1}"
}

log "Cleaning bin ..."
rm -rf ./bin
log "Done cleaning bin."

log "Building protocol ..."
protoc -I api/ \
    -I${GOPATH}/src \
    --go_out=plugins=grpc:api \
    api/api.proto
log "Done building protocol."

log "Formatting code ..."
gofmt -s -w -l .
log "Done formatting code."

log "Building server ..."
go build -i -v -o ./bin/server ./server/
log "Done building server."

log "Building client ..."
go build -i -v -o ./bin/client ./client/ 
log "Done building client."