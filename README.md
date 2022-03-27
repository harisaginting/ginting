# ginting
[GIN](https://github.com/gin-gonic/gin) Framework Boilerplate
untuk kamu yang dikejar waktu


#### Development Work Flow with Dependecty Injection Google Wire
- ```go get github.com/google/wire/cmd/wire```
- ```$GOPATH/bin/wire pkg/wire/wire.go```

#### Enable Keycloak
- edit .env file and set:
- KEYCLOAK=1
- KEYCLOAK_CERTS={{keycloak_host}}/auth/realms/{{realms_name}}/protocol/openid-connect/certs
- KEYCLOAK_ISSUER={{keycloak_host}}/auth/realms/{{realms_name}}

### RUN Sample GRPC SERVER
- ```go run ext/sample-grpc-server/main.go```