# TraceGraph CLI

Command-line tool for uploading case bundles to TraceGraph backend.

## Install

```bash
go build -o tracegraph-cli.exe .
```

## Configuration

Backend URL priority:
1. `--backend` flag
2. `TRACEGRAPH_BACKEND` env var
3. Default: `http://localhost:8000`

## Commands

### upload

Upload a case bundle from a folder.

```bash
tracegraph-cli upload <path> [--backend <url>]
```

The folder must contain:
- `alert.json`
- `dns.log`
- `conn.log`
- `ground_truth.json`

Example:
```bash
tracegraph-cli upload ./testdata/case_001
# Case created: id=abc-123, status=processing
```

### status

Show case status.

```bash
tracegraph-cli status <id> [--backend <url>]
```

Example:
```bash
tracegraph-cli status abc-123
# Case abc-123: status=done
```

### get

Show full case.

```bash
tracegraph-cli get <id> [--backend <url>]
```

Example:
```bash
tracegraph-cli get abc-123
# { ... JSON ... }
```

## Flags

| Flag | Description |
|------|-------------|
| `--backend <url>` | Backend URL. Overrides env var and default. |
| `--help, -h` | Show usage. |

**Note:** Flags must come **before** positional arguments:
```bash
tracegraph-cli upload --backend http://x:9000 ./case_001
```

## Development

Mock server for local testing:
```bash
go run ./tools/mockserver
```

Run tests:
```bash
go test ./...
```
