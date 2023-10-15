# golang-grpc
Golang gRPC skeleton with SQLC

## Initiate
`bash cmd.sh --cmd=init`
will install code requirement & running sample database at port 6543

## Generate SQLC
`bash cmd.sh --cmd=sqlc --path=pathfolder`
will generate the sqlc generated file for example:
`bash cmd.sh --cmd=sqlc --path=app/master-data/user` will place the generated file to `app/master-data/user/sqlc`
## Generate Proto
`bash cmd.sh --cmd=proto --path=pathfolder`
will generate the protobuf generated file for example:
`bash cmd.sh --cmd=sqlc --path=app/master-data/user` will place the generated file to `app/master-data/user/proto`