# LLLMapi

A project for creating a local service that manages LLM interactions through browser automation.

## Setup

1. Install dependencies:

```bash
go mod tidy
```

1. Configure the service:

```bash
cp configs/config.yaml.example configs/config.yaml
```

1. Run the service:

```bash
go run cmd/proxy/main.go
```

## Development

### Prerequisites

- Go 1.21 or higher
- Python 3.8 or higher (for browser automation)
- Docker (optional)

### Project Structure

```plaintext
LLLMapi/
├── cmd/            # Main applications
├── internal/       # Private application code
├── pkg/           # Public library code
├── automation/    # Python browser automation
└── configs/       # Configuration files
```

### Building

```bash
make build
```

### Testing

```bash
make test
```

## License

[GNU License](LICENSE)
