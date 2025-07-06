#!/bin/sh

set -e

# Detecta se está a correr como root
if [ "$(id -u)" -ne 0 ]; then
  echo "Cloning R2D2 CLI..."

  TMP_DIR=$(mktemp -d)
  git clone https://github.com/ArturC03/r2d2-cli.git "$TMP_DIR"
  cd "$TMP_DIR"

  echo "Building installer..."
  go build -o installer/build/r2d2-installer ./installer

  INSTALLER_PATH="$TMP_DIR/installer/build/r2d2-installer"

  echo "Re-running as sudo..."
  exec sudo bash -c "INSTALLER_PATH='$INSTALLER_PATH' bash -s" < "$0"
fi

# Parte que corre como root
echo "Running installer as root..."
chmod +x "$INSTALLER_PATH"
"$INSTALLER_PATH"

# Limpeza só se o tmp ainda existir
if [ -d "$(dirname "$INSTALLER_PATH")/../.." ]; then
  echo "Cleaning up..."
  rm -rf "$(dirname "$INSTALLER_PATH")/../.."
fi
