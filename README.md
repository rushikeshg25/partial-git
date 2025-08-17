# pgit

A fast, concurrent GitHub repository downloader that lets you download specific files, directories, or entire repositories without cloning the full git history.


https://github.com/user-attachments/assets/d8c7862a-3f07-42ad-af70-3891f4326e4f


## Features

- 🚀 **Fast concurrent downloads** using Go goroutines
- 📁 **Selective downloads** - files, directories, or entire repos
- 🔐 **GitHub token support** for private repos and higher rate limits
- ⏱️ **Smart timeouts** and cancellation support
- 🌳 **Branch/ref support** - download from specific branches or commits
- 💾 **Efficient** - only downloads what you need, no git history

## Installation

### Quick Install (Recommended)

**Using curl:**

```bash
curl -fsSL https://raw.githubusercontent.com/rushikeshg25/partial-git/main/scripts/install.sh | bash
```

**Using wget:**

```bash
wget -qO- https://raw.githubusercontent.com/rushikeshg25/partial-git/main/scripts/install.sh | bash
```

### Manual Installation

1. Download the latest binary for your platform from [Releases](https://github.com/rushikeshg25/partial-git/releases)
2. Make it executable: `chmod +x pgit`
3. Move to your PATH: `sudo mv pgit /usr/local/bin/`

### Supported Platforms

- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

## Usage

### Basic Usage

```bash
# Download entire repository
pgit https://github.com/user/repo

# Download specific directory
pgit https://github.com/user/repo/tree/main/src

# Download specific file
pgit https://github.com/user/repo/blob/main/README.md

# Download from specific branch
pgit https://github.com/user/repo/tree/develop
```

### GitHub Token Setup

For private repositories or higher rate limits:

```bash
# Set your GitHub Personal Access Token
pgit --set your_github_token_here

# Check token status
pgit --check

# View authenticated user info
pgit --auth

# Remove token
pgit --unset
```

### Examples

```bash
# Download VS Code's common utilities
pgit https://github.com/microsoft/vscode/tree/main/src/vs/base/common

# Download React's source code
pgit https://github.com/facebook/react/tree/main/packages/react/src

# Download a specific config file
pgit https://github.com/vercel/next.js/blob/canary/packages/next/package.json
```

## How It Works

1. **GitHub API Integration**: Uses GitHub's Contents API to fetch repository metadata
2. **Concurrent Downloads**: Downloads multiple files simultaneously using goroutines
3. **Smart Path Handling**: Automatically detects files vs directories and handles nested structures
4. **Rate Limiting**: Respects GitHub's API rate limits with optional authentication
5. **Timeout Protection**: 60-second timeout prevents hanging downloads

## Configuration

### Environment Variables

- `PGIT_GITHUB_TOKEN` - Your GitHub Personal Access Token (optional)

### Token Storage

Tokens are stored in your shell profile (`~/.zshrc` or `~/.bashrc`) for persistence across sessions.

## Development

### Building from Source

```bash
git clone https://github.com/rushikeshg25/partial-git.git
cd partial-git
go build -o pgit .
```

### Running Tests

```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
