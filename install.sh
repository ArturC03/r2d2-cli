#!/bin/sh

set -e

# Se root, continua com o resto
if [ "$(id -u)" -eq 0 ]; then
  INSTALLER_PATH="$1"

  if [ ! -f "$INSTALLER_PATH" ]; then
    echo "Installer não encontrado em: $INSTALLER_PATH"
    exit 1
  fi

  echo "Running installer as root..."
  chmod +x "$INSTALLER_PATH"
  "$INSTALLER_PATH"

  TMP_DIR="$(dirname "$INSTALLER_PATH")/../.."
  if [ -d "$TMP_DIR" ]; then
    echo "Cleaning up..."
    rm -rf "$TMP_DIR"
  fi
  exit 0
fi

# Parte não-root
echo "Cloning R2D2 CLI..."

TMP_DIR=$(mktemp -d)
git clone https://github.com/ArturC03/r2d2-cli.git "$TMP_DIR"
cd "$TMP_DIR"

echo "Building installer..."
go build -o installer/build/r2d2-installer ./installer

INSTALLER_PATH="$TMP_DIR/installer/build/r2d2-installer"

echo "Re-running as sudo..."
sudo bash "$0" "$INSTALLER_PATH"
