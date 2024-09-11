#!/bin/bash

set -e

# SimpleArity Full Diagnostic Installation Script

# Function to print error messages
error() {
    echo "Error: $1" >&2
    exit 1
}

# Function to print debug messages
debug() {
    echo "Debug: $1" >&2
}

# Determine latest version
debug "Fetching latest version from GitHub API..."
GITHUB_API_RESPONSE=$(curl -sSi https://api.github.com/repos/kemalcanbora/simplearity/releases/latest)
debug "Full GitHub API Response Headers:"
echo "$GITHUB_API_RESPONSE" | sed 's/^/    /' >&2

GITHUB_API_BODY=$(echo "$GITHUB_API_RESPONSE" | sed -n '/^\r?$/,$p' | sed '1d')
debug "GitHub API Response Body:"
echo "$GITHUB_API_BODY" | sed 's/^/    /' >&2

LATEST_VERSION=$(echo "$GITHUB_API_BODY" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
debug "Latest version: $LATEST_VERSION"

if [ -z "$LATEST_VERSION" ]; then
    error "Failed to determine latest version. Please check if releases exist on GitHub."
fi

# Determine system architecture and OS
OS=$(uname -s)
ARCH=$(uname -m)
case $OS in
    Linux)
        OS="Linux"
        ;;
    Darwin)
        OS="Darwin"
        ;;
    *)
        error "Unsupported operating system: $OS"
        ;;
esac
case $ARCH in
    x86_64)
        ARCH="x86_64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    *)
        error "Unsupported architecture: $ARCH"
        ;;
esac
debug "Detected OS: $OS, Architecture: $ARCH"

# Construct download URL
DOWNLOAD_URL="https://github.com/kemalcanbora/simplearity/releases/download/${LATEST_VERSION}/simplearity_${OS}_${ARCH}.tar.gz"
debug "Download URL: $DOWNLOAD_URL"

# Create a directory for SimpleArity in the user's home
INSTALL_DIR="$HOME/.simplearity"
mkdir -p "$INSTALL_DIR"
debug "Installation directory: $INSTALL_DIR"

# Download SimpleArity
echo "Downloading SimpleArity ${LATEST_VERSION} for ${OS} ${ARCH}..."
CURL_RESPONSE=$(curl -sSL -w "%{http_code}" -o "$INSTALL_DIR/simplearity.tar.gz" $DOWNLOAD_URL)
HTTP_STATUS=${CURL_RESPONSE: -3}
debug "HTTP Status Code: $HTTP_STATUS"

if [ $HTTP_STATUS -ne 200 ]; then
    error "Failed to download SimpleArity. HTTP Status: $HTTP_STATUS"
fi

# Check file type
FILE_TYPE=$(file -b "$INSTALL_DIR/simplearity.tar.gz")
debug "Downloaded file type: $FILE_TYPE"

# Check file size
FILE_SIZE=$(du -h "$INSTALL_DIR/simplearity.tar.gz" | cut -f1)
debug "Downloaded file size: $FILE_SIZE"

# Print contents if file is small
if [ "$(wc -c < "$INSTALL_DIR/simplearity.tar.gz")" -lt 1000 ]; then
    debug "Contents of downloaded file:"
    cat "$INSTALL_DIR/simplearity.tar.gz" | sed 's/^/    /' >&2
fi

# Extract the binary
echo "Extracting SimpleArity..."
if ! tar -xzf "$INSTALL_DIR/simplearity.tar.gz" -C "$INSTALL_DIR"; then
    error "Failed to extract SimpleArity"
fi

# Check if the binary was extracted successfully
if [ ! -f "$INSTALL_DIR/simplearity" ]; then
    error "SimpleArity binary not found in the extracted archive"
fi

# Make the binary executable
chmod +x "$INSTALL_DIR/simplearity"

# Clean up
rm "$INSTALL_DIR/simplearity.tar.gz"

echo "SimpleArity ${LATEST_VERSION} has been installed to $INSTALL_DIR/simplearity"

# Add the installation directory to PATH in shell configuration file
SHELL_CONFIG="$HOME/.bashrc"
if [[ $SHELL == *"zsh"* ]]; then
    SHELL_CONFIG="$HOME/.zshrc"
fi

if ! grep -q "$INSTALL_DIR" "$SHELL_CONFIG"; then
    echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$SHELL_CONFIG"
    echo "Added $INSTALL_DIR to your PATH in $SHELL_CONFIG"
    echo "Please restart your terminal or run 'source $SHELL_CONFIG' to update your PATH"
else
    echo "$INSTALL_DIR is already in your PATH"
fi

echo "You can now use SimpleArity by running 'simplearity' in your terminal after updating your PATH"

# Verify installation
if "$INSTALL_DIR/simplearity" --version >/dev/null 2>&1; then
    echo "Installation verified. SimpleArity is ready to use."
else
    error "Installation seems to have failed. Unable to run SimpleArity."
fi