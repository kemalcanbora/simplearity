#!/bin/bash

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to log debug messages
debug() {
    echo "Debug: $1"
}

# Function to log errors
error() {
    echo -e "${RED}Error: $1${NC}" >&2
}

# Function to log success messages
success() {
    echo -e "${GREEN}$1${NC}"
}

# Function to log warnings
warn() {
    echo -e "${YELLOW}Warning: $1${NC}"
}

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$ARCH" == "x86_64" ]]; then
    ARCH="amd64"
elif [[ "$ARCH" == "aarch64" || "$ARCH" == "arm64" ]]; then
    ARCH="arm64"
else
    error "Unsupported architecture: $ARCH"
    exit 1
fi

# Fetch latest version from GitHub API
debug "Fetching latest version from GitHub API..."
GITHUB_API_RESPONSE=$(curl -sSL https://api.github.com/repos/kemalcanbora/simplearity/releases/latest)
debug "Full GitHub API Response: $GITHUB_API_RESPONSE"

# Extract version using jq
if command -v jq >/dev/null 2>&1; then
    VERSION=$(echo "$GITHUB_API_RESPONSE" | jq -r .tag_name)
else
    VERSION=$(echo "$GITHUB_API_RESPONSE" | grep -oP '"tag_name": "\K[^"]+')
fi

if [[ -z "$VERSION" ]]; then
    error "Failed to determine latest version. Please check if releases exist on GitHub."
    exit 1
fi

debug "Latest version: $VERSION"

# Construct download URL
DOWNLOAD_URL="https://github.com/kemalcanbora/simplearity/releases/download/${VERSION}/simplearity_${OS}_${ARCH}.tar.gz"
debug "Download URL: $DOWNLOAD_URL"

# Download the tarball
TARBALL="simplearity_${OS}_${ARCH}.tar.gz"
curl -sSL "$DOWNLOAD_URL" -o "$TARBALL"

# Extract the tarball
tar -xzf "$TARBALL"

# Get the full path of the current directory
INSTALL_DIR=$(pwd)

# Update PATH in .zshrc
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo "export PATH=\"$INSTALL_DIR:\$PATH\"" >> ~/.zshrc
    export PATH="$INSTALL_DIR:$PATH"
    success "Updated PATH in ~/.zshrc and current session."
    echo "You can now run 'simplearity' from any directory."
    echo "Please restart your terminal or run 'source ~/.zshrc' to apply changes in new terminal windows."
else
    warn "The installation directory is already in your PATH."
fi

# Clean up
rm "$TARBALL"

# Verify simplearity is accessible
if command -v simplearity &> /dev/null; then
    success "Simplearity $VERSION has been installed and is accessible from anywhere."
else
    error "Error: simplearity is not accessible. Please check the file permissions and try again."
fi

echo "Installation complete. You may need to restart your terminal or run 'source ~/.zshrc' to use the 'simplearity' command in this session."