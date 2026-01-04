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
    token := "..."
    client, err := twentygo.NewClient("http://localhost:8000")
    if err != nil {
        return err
    }

    attachments, err := client.FindManyAttachments(context.Background(), token, &twentygo.FindManyAttachmentsParams{})
	if err != nil {
		return err
	}
	fmt.Println(attachments)
```

## Features

[TwentyIface](twenty_iface.go) holds all methods a client should fulfil.