# Phi

> *Find your configs. Instantly.*

A CLI tool for managing and quickly opening configuration files. Instead of manually navigating to `~/.config/`, `/etc/`, or hidden home directories every time, you register your config files once and open them with a single command.

```bash
phi nvim
```

---

## Motivation

Config files are scattered all over the system. Phi creates a central, named registry (stored in a clean, readable TOML format) and gives you direct access by simple names, eliminating the need to remember complex paths.

---

## Planned Features

### V1 – Core

| Command | Description |
|---|---|
| `phi add <name> <path>` | Register a config file or directory |
| `phi open <name>` | Open the entry in your configured editor |
| `phi list` | List all registered entries |
| `phi remove <name>` | Remove an entry from the registry |

- Support for **both files and directories** (e.g., an entire Neovim config directory)
- Uses your `$EDITOR` environment variable by default, with support for per-entry overrides
- Registry is safely stored at `~/.config/phi/config.toml`
- Clean, user-friendly error handling (file not found, name already taken, etc.)

### V2 – Planned Extensions

- `phi edit <name>` — Update the path or editor of an existing entry
- `phi scan` — Auto-discover known config files on the system
- `--editor` flag on `open` to override the editor for a single invocation
- Tags and groups (`phi list --tag work`)
- `phi which <name>` — Print only the absolute path (useful for shell scripting)
- Shell completions (Bash, Zsh, Fish)

---

## Technology

Written in **Go** using [Cobra](https://github.com/spf13/cobra) for the CLI architecture and `go-toml/v2` for fast, allocation-free TOML parsing.

---

## Status

🚧 Actively in development — V1 not yet released.

---

## Installation

*Coming with the first stable release.*

---

## License

MIT
