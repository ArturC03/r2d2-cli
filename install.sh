#!/bin/sh
set -e

# Configuration
BINARY_URL="https://github.com/ArturC03/r2d2-cli/installer/build/r2d2-installer-linux-amd64"
TMP_BIN="/tmp/r2d2-installer"

# Root check
if [ "$(id -u)" -ne 0 ]; then
  echo "Este script precisa de ser executado com sudo. A reiniciar com sudo..."
  exec sudo "$0" "$@"
fi

echo "A transferir o instalador R2D2..."
curl -fsSL "$BINARY_URL" -o "$TMP_BIN"

chmod +x "$TMP_BIN"

echo "A executar o instalador..."
"$TMP_BIN"

echo "A limpar..."
rm -f "$TMP_BIN"
