# CLI

Just a couple of CLIs written to help me learn how golang works and its quirks.

To see the documentation (after building and moving into PATH):

```bash
serve --help
jwt --help
```

## Serve

Serves the current directory over localhost and the local network. Very similar to doing `python -m http.server`

### To run:

```bash
go run serve/main.go
```

### To build:

```bash
cd serve/
go build -o serve
sudo mv serve /usr/local/bin # requires elevated priviledges.
```

Once moved, you can just spin the server up via `serve`

## JWT Decoder

Decodes a JWT

### To run:

```bash
go run jwt_decoder/*.go eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c -h
```

### To build:

```bash
cd jwt_decoder/
go build -o jwt
sudo mv jwt /usr/local/bin/
```

Once moved, you can start decoding JWTs by simply running `jwt`
