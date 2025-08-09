#!/usr/bin/env bash

# Simple pgit installer
# Usage: curl -fsSL https://raw.githubusercontent.com/rushikeshg25/partial-git/main/scripts/install-simple.sh | bash

set -e

REPO="rushikeshg25/partial-git" 
BINARY_NAME="pgit"
INSTALL_DIR="/usr/local/bin"

# Detect OS and arch
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64|amd64) ARCH="amd64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    armv7l) ARCH="arm" ;;
    i386|i686) ARCH="386" ;;
    *) echo "Unsupported architecture: $ARCH" && exit 1 ;;
esac

case $OS in
    linux) OS="linux" ;;
    darwin) OS="darwin" ;;
    *) echo "Unsupported OS: $OS" && exit 1 ;;
esac

PLATFORM="${OS}_${ARCH}"

# Get latest version
VERSION=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$VERSION" ]; then
    echo "Error: Could not determine latest version"
    exit 1
fi

echo "Installing ${BINARY_NAME} ${VERSION} for ${PLATFORM}..."

# Download and install
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${BINARY_NAME}_${PLATFORM}"

if command -v sudo >/dev/null 2>&1 && [ ! -w "$INSTALL_DIR" ]; then
    sudo curl -fsSL "$DOWNLOAD_URL" -o "$INSTALL_DIR/$BINARY_NAME"
    sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
else
    curl -fsSL "$DOWNLOAD_URL" -o "$INSTALL_DIR/$BINARY_NAME"
    chmod +x "$INSTALL_DIR/$BINARY_NAME"
fi

echo "✅ ${BINARY_NAME} installed successfully!"
echo "Run: ${BINARY_NAME} --help"
