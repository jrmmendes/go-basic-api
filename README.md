# Go Sample API

This is a basic API made with go.

## Setup

> This project requires `go` 1.23+, `node` and `make`. 

To run the development server, install `air` and `pino-pretty` for code reload and better log visualization:

```bash
go install github.com/air-verse/air@latest; npm i -g pino-pretty
```

```bash
npm i -g pino-pretty
```

Then, you can start the development server with the command `make dev`.

## Useful commands
- Run integration tests: 
```bash
make test-integration
```

- Clear temp files: 
```bash
make clear
```

- Build the application:
```bash
make build
```

- Start the production server (you must generate a valid build before):
```bash
make start
```
