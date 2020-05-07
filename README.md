# MARVEL

## HOW TO GEN PROTOC
- `protoc --go_out=plugins=grpc:. transport/grpc/proto/user/user.proto`

### Requirement
- Install golang
- Install docker

### Setup enviroment
- ```make init```

### Run server 
- ```make dev```

### Run test
-```make test```

# HOW TO GEN CERT 
- `cat gen_certs.sh.example > gen_certs.sh`
- `./gen_certs.sh`
