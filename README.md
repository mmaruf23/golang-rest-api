# Golang REST API

## Overview

A RESTful API built with Go, providing efficient and scalable endpoints for data management.

## Features

- Fast and lightweight HTTP server
- RESTful endpoint design
- Easy to deploy and maintain

## Prerequisites

- Go 1.16 or higher
- Git

## Installation

```bash
git clone <repository-url>
cd golang-rest-api
go mod download
```

## Usage

SETUP database and env.
and then :

```bash
go run cmd/main.go
```

The API will be available at `http://localhost:8080`

## Project Structure

```
.
├── main.go
├── handlers/
├── models/
└── go.mod
```

## API Endpoints

- `GET /api/health` - Health check

## Contributing

Pull requests are welcome. Please open an issue first to discuss changes.

## License

MIT
