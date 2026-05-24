# 栞 shiori

> *Shiori (栞) is the Japanese word for bookmark — a mark that shows you the way back.*

A CLI tool for managing and quickly opening configuration files. Instead of manually navigating to `~/.config` every time, you register your config files once and open them with a single command.

```bash
shiori open nvim
```

---

## Motivation

Config files are scattered all over the system — in `~/.config/`, `~/.local/`, `~/`, or `/etc/`. Shiori creates a central, named registry and gives you direct access by name, without having to remember paths.

---

## Planned Features

### V1 – Core

| Command | Description |
|---|---|
| `shiori add <name> <path>` | Register a config file or directory |
| `shiori open <name>` | Open the entry in your configured editor |
| `shiori list` | List all registered entries |
| `shiori remove <name>` | Remove an entry from the registry |

- Support for **both files and directories** (e.g. an entire Neovim config directory)
- Falls back to `$EDITOR`, with an optional per-entry editor override
- Registry stored at `~/.config/shiori/registry.json`
- Helpful error messages (file not found, name already taken, etc.)

### V2 – Planned Extensions

- `shiori edit <name>` — update the path or editor of an existing entry
- `shiori scan` — auto-discover known config files on the system
- `--editor` flag on `open` to override the editor for a single invocation
- Tags and groups (`shiori list --tag work`)
- `shiori which <name>` — print only the path, useful for scripting
- Shell completions (Bash, Zsh, Fish)

---

## Technology

Written in **Go** using [Cobra](https://github.com/spf13/cobra) for CLI structure.

---

## Status

🚧 Actively in development — V1 not yet released.

---

## Installation

*Coming with the first stable release.*

---

## License

MIT
