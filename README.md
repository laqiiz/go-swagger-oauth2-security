# go-swagger-oauth2-security
SAMPLE APP: This repository is sample application intended to try oauth2 authentication using go-swagger with KeyCloak

## Abstract

This is POC application that confirm go-swagger server and keycloak authentication workflow. 

## Usage

Run KeyCloak Server.

> docker run -d -p 18080:8080 -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin --name keycloak jboss/keycloak go run main.go

Access localhost:18080 with your browser.

* xxx
* xxx
* xxx

Register Google callback URL

* follow bellow url
    * https://github.com/go-swagger/go-swagger/tree/master/examples/oauth2#register-the-callback-url

Run API server.
 
> set ClientID=<Your Client ID>
> set ClientSecret=<Your Client Secret>
> go run gen/cmd/hellouah-server/main.go --host 0.0.0.0 --port 8080

Send api request.

> curl -i -H 'Authorization: Bearer ya29.a0AfH6SMAijuEYH0qiV-HhyNXl3hHACKwrTbheWCJ2VzDEjVDokCRcJuZ832dC7KKkoxruIsMMDW1I-eb1x2t51G6G9DJWU_mZ2Ir16Uv_LyjFCJC9YvgZ2SK8xK3FVZQ_rrdqzt1uy6hxRCUIS0Xv9_IbKP07EK6gdl4' http://172.30.32.1:8080/v1/hello
  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Sun, 26 Jul 2020 07:34:28 GMT
  Content-Length: 20
  
  {"message":"hello"}

Send invalid token.

> curl -i -H 'Authorization: Bearer ya29.111111' http://172.30.32.1:8080/v1/hello
HTTP/1.1 401 Unauthorized
Content-Type: application/json
Date: Sun, 26 Jul 2020 07:35:20 GMT
Content-Length: 38

{"code":401,"message":"invalid token"}



## Development Memo

```sh
# Install go-swagger
go get -u github.com/go-swagger/go-swagger@0.25.0

# generate file
swagger generate server -a hellouah -A hellouah -P models.Principal --strict-additional-properties -t gen
```

