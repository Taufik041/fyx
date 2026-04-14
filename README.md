# fyx

A smart CLI companion for your terminal. fyx does two things:

- **Fixes mistyped commands** — when you run a command that doesn't exist, fyx catches it, asks AI what you meant, and suggests the correct one before you run it.
- **Browses any tool interactively** — run `fyx <tool>` to get an arrow-key menu of all available subcommands. Select one, edit it if needed, then execute.

Supports OpenAI and Anthropic (Claude) as AI providers. Works on Windows (PowerShell), macOS, and Linux.

---

## Installation

Coming soon.

---

## Usage

### First time setup

```sh
fyx init
```

### Browse a tool's commands interactively

```sh
fyx kubectl
fyx terraform
fyx docker
```

### Refresh PATH after installing a new tool

```sh
fyx refresh
```

### Turn the command correction hook on/off

```sh
fyx activate
fyx deactivate
```

### Change your AI provider or API key

```sh
fyx config
```

---

## How it works

**Command correction** — fyx registers a hook in your shell profile. When you mistype a command and the shell can't find it, instead of just showing an error, fyx intercepts that moment, sends the failed command to your configured AI provider, and shows you the suggested correction before anything runs.

**Interactive browser** — fyx runs `<tool> --help` under the hood, parses the output, and renders it as a navigable menu in your terminal. Use arrow keys to move, Enter to select, then edit the command before executing.

---

## Configuration

fyx stores your config at `~/.fyx/config.json`. You can edit it directly or use `fyx config` to update it interactively.

---

## Supported Shells

| Shell | Platform | Status |
|---|---|---|
| PowerShell | Windows | ✅ Supported |
| zsh | macOS / Linux | ✅ Supported |
| bash | Linux | ✅ Supported |

---

## Contributing

This project is in early development. Feel free to open issues or PRs.

---

## License

MIT