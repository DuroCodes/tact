# Tact

<div align="center">
  <img src="./logo.png" width="400" />
</div>

Tact is a CLI tool that measures your typing speed and accuracy, similar to [MonkeyType](https://monkeytype.com/). It's an easy way to practice your typing skills without needing to open your browser.

## Usage

> [!TIP]
> You can use `go install github.com/durocodes/tact@latest` to install the CLI tool on your system (if I set it up correctly and you have Go installed)

1. Clone the repository
2. Run `go build` to build the binary
3. Run `./tact` (or `./tact.exe` on Windows) to start the CLI

### Flags

- `n` - number of words to type (min: 1, max: 100, default: 25)
