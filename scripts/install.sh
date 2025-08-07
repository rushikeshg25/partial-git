#!/usr/bin/env bash

set -e

VERSION="v0.0.2"
REPO="rushikeshg25/partial-git"

# Detect OS
OS="$(uname | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

# Normalize architecture
if [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
elif [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
  ARCH="arm64"
else
  echo "❌ Unsupported architecture: $ARCH"
  exit 1
fi

# Set file extension based on OS
if [[ "$OS" == "windows" ]]; then
  FILE_EXT=".zip"
else
  FILE_EXT=".tar.gz"
fi

FILENAME="pgit-${VERSION}-${OS}-${ARCH}${FILE_EXT}"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${FILENAME}"

echo "📥 Downloading pgit ${VERSION} for ${OS}-${ARCH}..."
curl -L -o "$FILENAME" "$URL"

echo "📦 Extracting..."
if [[ "$FILE_EXT" == ".zip" ]]; then
  unzip -o "$FILENAME"
else
  tar -xzf "$FILENAME"
fi

echo "🚚 Installing..."
chmod +x pgit
sudo mv pgit /usr/local/bin/pgit

echo "✅ Installed pgit to /usr/local/bin/pgit"
echo "🔍 Run 'pgit --help' to get started."

# Cleanup
rm "$FILENAME"
