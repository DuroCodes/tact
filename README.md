# Tact

<div align="center">
  <img src="./logo.png" width="400" />
</div>

Tact is a CLI tool that measures your typing speed and accuracy, similar to [monkeytype](https://monkeytype.com/). It's an easy way to practice your typing skills without needing to open your browser.

## Usage

> [!TIP]
> You can use `go install github.com/durocodes/tact@latest` to install the CLI tool on your system (if I set it up correctly and you have Go installed)

1. Clone the repository
2. Run `go build` to build the binary
3. Run `./tact` (or `./tact.exe` on Windows) to start the CLI

### Flags

- `n` - number of words to type (min: 1, max: 100, default: 25)
- `w` - wordset to use (default: "english", available wordsets in `/wordsets`, thanks to [monkeytype](https://github.com/monkeytypegame/monkeytype/tree/master/frontend/static/languages))

> [!TIP]
> You can use a custom wordset by creating a new file in `/wordsets` and using the filename as the flag value
