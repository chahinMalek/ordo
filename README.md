# Ordo

**Ordo** is a fast, safe, and deterministic CLI tool written in **Go** for organizing files within a directory. It brings order to chaos by automatically categorizing files into folders based on their extensions or user-defined grouping rules.

## ✨ Features

- **Unified Rule System**: Everything is a rule. A rule maps extensions (like `jpg`, `png`) to a folder (like `images`).
- **Smart Fallback**: If a file extension isn't covered by a rule, Ordo automatically uses the extension as the folder name (e.g., `.pdf` → `pdf/`).
- **Predefined Groups**: Comes out-of-the-box with logical groupings for `images`, `documents`, `audio`, and more.
- **Native Performance**: Built with Go, it's a single binary with zero external dependencies and works across macOS, Linux, and Windows.
- **Safe by Design**: Ordo separates planning from execution. It detects name collisions and ensures no files are ever overwritten.
- **Dry Run Support**: Preview exactly what changes will be made before any file is moved.

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

Preview changes without moving any files:
```bash
ordo --dry-run
```

## 🛠 Rule Management

Manage your organization rules directly from the CLI:

```bash
# List current rules
ordo rules list

# Add a custom rule for images
ordo rules add images jpg png webp

# Delete a group
ordo rules delete documents
```

## 📝 Configuration

Ordo uses a standard TOML configuration file located at `~/.config/ordo/config.toml`. You can always reset to factory defaults using:

```bash
ordo config reset
```

---

*Ordo: Bringing order to your filesystem.*
