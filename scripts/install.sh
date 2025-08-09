#!/usr/bin/env bash

set -e

# pget installer script
# Usage: curl -fsSL https://raw.githubusercontent.com/rushikeshg25/partial-git/main/scripts/install.sh | bash
# or:    wget -qO- https://raw.githubusercontent.com/rushikeshg25/partial-git/main/scripts/install.sh | bash

REPO="rushikeshg25/partial-git"  # Replace with your GitHub username/repo
BINARY_NAME="pget"
INSTALL_DIR="$HOME/.local/bin"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Logging functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Detect OS and architecture
detect_platform() {
    local os
    local arch
    
    # Detect OS
    case "$(uname -s)" in
        Linux*)     os="linux";;
        Darwin*)    os="darwin";;
        CYGWIN*)    os="windows";;
        MINGW*)     os="windows";;
        MSYS*)      os="windows";;
        *)          
            log_error "Unsupported operating system: $(uname -s)"
            exit 1
            ;;
    esac
    
    # Detect architecture
    case "$(uname -m)" in
        x86_64|amd64)   arch="amd64";;
        arm64|aarch64)  arch="arm64";;
        armv7l)         arch="arm";;
        i386|i686)      arch="386";;
        *)              
            log_error "Unsupported architecture: $(uname -m)"
            exit 1
            ;;
    esac
    
    echo "${os}_${arch}"
}

# Get latest release version from GitHub API
get_latest_version() {
    local latest_url="https://api.github.com/repos/${REPO}/releases/latest"
    
    if command -v curl >/dev/null 2>&1; then
        curl -s "$latest_url" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'
    elif command -v wget >/dev/null 2>&1; then
        wget -qO- "$latest_url" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'
    else
        log_error "Neither curl nor wget is available. Please install one of them."
        exit 1
    fi
}

# Download and install binary
install_binary() {
    local version="$1"
    local platform="$2"
    local binary_url="https://github.com/${REPO}/releases/download/${version}/${BINARY_NAME}_${platform}"
    
    log_info "Downloading ${BINARY_NAME} ${version} for ${platform}..."
    
    # Create install directory if it doesn't exist
    mkdir -p "$INSTALL_DIR"
    
    # Download binary
    local temp_file
    temp_file=$(mktemp)
    
    if command -v curl >/dev/null 2>&1; then
        if ! curl -fsSL "$binary_url" -o "$temp_file"; then
            log_error "Failed to download binary from $binary_url"
            rm -f "$temp_file"
            exit 1
        fi
    elif command -v wget >/dev/null 2>&1; then
        if ! wget -qO "$temp_file" "$binary_url"; then
            log_error "Failed to download binary from $binary_url"
            rm -f "$temp_file"
            exit 1
        fi
    fi
    
    # Move to install directory and make executable
    mv "$temp_file" "$INSTALL_DIR/$BINARY_NAME"
    chmod +x "$INSTALL_DIR/$BINARY_NAME"
    
    log_success "Installed ${BINARY_NAME} to $INSTALL_DIR/$BINARY_NAME"
}

# Add to PATH if not already there
update_path() {
    local shell_profile
    
    # Detect shell and profile file
    case "$SHELL" in
        */zsh)
            shell_profile="$HOME/.zshrc"
            ;;
        */bash)
            if [[ -f "$HOME/.bash_profile" ]]; then
                shell_profile="$HOME/.bash_profile"
            else
                shell_profile="$HOME/.bashrc"
            fi
            ;;
        */fish)
            shell_profile="$HOME/.config/fish/config.fish"
            ;;
        *)
            shell_profile="$HOME/.profile"
            ;;
    esac
    
    # Check if install directory is in PATH
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        log_info "Adding $INSTALL_DIR to PATH in $shell_profile"
        
        # Add to shell profile
        echo "" >> "$shell_profile"
        echo "# Added by pget installer" >> "$shell_profile"
        echo "export PATH=\"$INSTALL_DIR:\$PATH\"" >> "$shell_profile"
        
        log_warning "Please restart your terminal or run: source $shell_profile"
        log_warning "Or add $INSTALL_DIR to your PATH manually"
    else
        log_info "$INSTALL_DIR is already in PATH"
    fi
}

# Verify installation
verify_installation() {
    if [[ ":$PATH:" == *":$INSTALL_DIR:"* ]] && command -v "$BINARY_NAME" >/dev/null 2>&1; then
        local version
        version=$("$BINARY_NAME" version 2>/dev/null || echo "unknown")
        log_success "Installation verified! ${BINARY_NAME} ${version} is ready to use."
    elif [[ -x "$INSTALL_DIR/$BINARY_NAME" ]]; then
        local version
        version=$("$INSTALL_DIR/$BINARY_NAME" version 2>/dev/null || echo "unknown")
        log_success "Installation complete! ${BINARY_NAME} ${version} installed to $INSTALL_DIR/$BINARY_NAME"
        log_info "Run: export PATH=\"$INSTALL_DIR:\$PATH\" to use $BINARY_NAME from anywhere"
    else
        log_error "Installation verification failed"
        exit 1
    fi
}

# Main installation function
main() {
    log_info "Installing ${BINARY_NAME}..."
    
    # Detect platform
    local platform
    platform=$(detect_platform)
    log_info "Detected platform: $platform"
    
    # Get latest version
    local version
    version=$(get_latest_version)
    if [[ -z "$version" ]]; then
        log_error "Could not determine latest version"
        exit 1
    fi
    log_info "Latest version: $version"
    
    # Install binary
    install_binary "$version" "$platform"
    
    # Update PATH
    update_path
    
    # Verify installation
    verify_installation
    
    log_success "🎉 ${BINARY_NAME} installation complete!"
    log_info "Usage: ${BINARY_NAME} <github-url>"
    log_info "Example: ${BINARY_NAME} https://github.com/user/repo"
    log_info "Run '${BINARY_NAME} --help' for more options"
}

# Run main function
main "$@"