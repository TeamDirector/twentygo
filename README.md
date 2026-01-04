# twentygo

Golang Twenty CRM API Package


## Contribution

PRs are welcome

## Changelog

For release notes please consult the specific releases [here](https://github.com/Nerzal/gocloak/releases)

## Usage

### Installation

```shell
go get github.com/TeamDirector/twentygo
```

### Importing

```go
 import "github.com/TeamDirector/twentygo"
```

### Create New User

```go
 client := gocloak.NewClient("https://mycool.keycloak.instance")
 ctx := context.Background()
 token, err := client.LoginAdmin(ctx, "user", "password", "realmName")
 if err != nil {
  panic("Something wrong with the credentials or url")
 }

 user := gocloak.User{
  FirstName: gocloak.StringP("Bob"),
  LastName:  gocloak.StringP("Uncle"),
  Email:     gocloak.StringP("something@really.wrong"),
  Enabled:   gocloak.BoolP(true),
  Username:  gocloak.StringP("CoolGuy"),
 }

 _, err = client.CreateUser(ctx, token.AccessToken, "realm", user)
 if err != nil {
  panic("Oh no!, failed to create user :(")
 }
```

## Features

[TwentyIface](twenty_iface.go) holds all methods a client should fulfil.


## Configure gocloak to skip TLS Insecure Verification

```go
    client := gocloak.NewClient(serverURL)
    restyClient := client.RestyClient()
    restyClient.SetDebug(true)
    restyClient.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })
```