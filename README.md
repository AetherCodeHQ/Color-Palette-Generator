# 💻 Color Palette Generator

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-00ADD8?style=for-the-badge)
![CI](https://img.shields.io/badge/CI-GitHub%20Actions-2088FF?style=for-the-badge&logo=githubactions&logoColor=white)
![CodeQL](https://img.shields.io/badge/CodeQL-Security-00ADD8?style=for-the-badge)
![Lint](https://img.shields.io/badge/GolangCI--Lint-Passing-00ADD8?style=for-the-badge)
![Dependabot](https://img.shields.io/badge/Dependabot-Enabled-025E9C?style=for-the-badge&logo=dependabot&logoColor=white)

> Color-Palette-Generator - Open source tool by AetherCodeHQ

`cli` `color` `design` `golang` `utilities`

## What is it?

**Color Palette Generator** is a cli tool designed for developers who need fast, reliable, and offline-capable tools. Built with Go for maximum performance and minimal resource usage.

## Why should you care?

- 🚀 **Fast** — Compiled Go binary, no runtime dependencies
- 🔒 **Secure** — CodeQL security analysis + Dependabot
- 🌐 **Offline-first** — Works without internet connection
- 📦 **Lightweight** — Single binary, minimal footprint
- 🛠️ **Developer-friendly** — Clean CLI with helpful documentation

## Quick Start

### Prerequisites
- Go 1.21 or higher

### Install from source
```bash
git clone https://github.com/AetherCodeHQ/Color-Palette-Generator.git
cd Color-Palette-Generator
go build -o Color-Palette-Generator .
```

### Run
```bash
./Color-Palette-Generator --help
```

## Usage

```bash
# Basic usage
./Color-Palette-Generator --path ./target

# With options
./Color-Palette-Generator --path ./target --format json --output report.json

# Verbose mode
./Color-Palette-Generator --path ./target --verbose
```

## Features

- ✅ High-performance Go implementation
- ✅ Cross-platform support (Windows, Linux, macOS)
- ✅ JSON export for CI/CD integration
- ✅ Colored terminal output
- ✅ Configurable via YAML/JSON
- ✅ Comprehensive documentation

## CLI Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--path` | Target directory | `.` |
| `--format` | Output format (json, text, badge) | `text` |
| `--output` | Output filename | `stdout` |
| `--verbose` | Enable verbose output | `false` |
| `--config` | Config file path | - |

## Examples

### Basic scan
```bash
./Color-Palette-Generator --path ./my-project
```

### JSON report
```bash
./Color-Palette-Generator --path ./my-project --format json --output report.json
```

### CI/CD integration
```yaml
# .github/workflows/scan.yml
- name: Run Color Palette Generator
  run: ./Color-Palette-Generator --path . --format json --output report.json
```

## Development

```bash
# Clone the repo
git clone https://github.com/AetherCodeHQ/Color-Palette-Generator.git
cd Color-Palette-Generator

# Build
go build -o Color-Palette-Generator .

# Run tests
go test ./...

# Lint
golangci-lint run
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Security

If you discover a security vulnerability, please report it responsibly. See [SECURITY.md](SECURITY.md) for details.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  Built with ❤️ by <a href="https://github.com/AetherCodeHQ">AetherCode</a> • <a href="https://github.com/AetherCode-Core">AetherCode-Core</a>
</p>
