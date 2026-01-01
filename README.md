# Ordo

**Ordo** is a fast, safe, and deterministic CLI tool written in **Go** for organizing files within a directory. It brings order to chaos by automatically categorizing files into folders based on their extensions or user-defined grouping rules.

## ✨ Features

- **Customizable Rules**: Define your own rules to group specific extensions (like `jpg` and `png`) together into target folders (like `images`).
- **Smart Fallback**: If a file extension isn't covered by a rule, Ordo automatically uses the file extension as the folder name (e.g., `.pdf` → `pdf/`).
- **Predefined Groups**: Comes out-of-the-box with logical groupings for `images`, `documents`, `audio`, `videos`, and `archives`.
- **Undo Capability**: Regret a move? Ordo tracks its actions in a local `.ordo_history` file, allowing you to instantly revert the last operation.
- **Safe by Design**: Ordo separates planning from execution. It detects path blockages and ensures no files are ever overwritten.
- **Dry Run Support**: Preview exactly what changes will be made before any file is moved.
- **Native Performance**: Built with Go, it's a single binary with zero external dependencies and works across macOS, Linux, and Windows.

## 🚀 Quick Start

### Installation

```bash
go install github.com/chahinMalek/ordo@latest
```

### Basic Usage

Organize the current directory using active rules (with extension fallback):
```bash
ordo
```

Organize a specific directory:
```bash
ordo --path ~/Downloads
```

Preview changes without moving any files:
```bash
ordo --dry-run
```

Undo the most recent organization operation:
```bash
ordo revert
```

## 🛠 Rule Management

Manage your organization rules directly from the CLI:

```bash
# List current rules
ordo rules list

# Add extensions to a group (creates group if it doesn't exist)
ordo rules add images jpg png webp

# Remove a specific extension from a group
ordo rules remove images webp

# Delete an entire rule group
ordo rules delete documents
```

## 📝 Configuration

Ordo uses a standard TOML configuration file located at `~/.config/ordo/config.toml`. You can always reset to factory defaults using:

```bash
ordo config reset
```

---

*Ordo: Bringing order to your filesystem.*
