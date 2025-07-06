#!/bin/sh
set -e

ARCH=$(uname -m)
OS=$(uname | tr '[:upper:]' '[:lower:]')

BIN_NAME="r2d2-installer"
VERSION="v0.1.0"  # ajusta para a tag correta
BASE_URL="https://github.com/ArturC03/r2d2-cli/releases/download/${VERSION}"

case $ARCH in
  x86_64) ARCH="amd64" ;;
  aarch64 | arm64) ARCH="arm64" ;;
  *) echo "Arquitetura $ARCH não suportada"; exit 1 ;;
esac

BIN_URL="${BASE_URL}/${BIN_NAME}_${OS}_${ARCH}_v1"

TMP_DIR=$(mktemp -d)
BIN_PATH="$TMP_DIR/$BIN_NAME"

echo "A transferir $BIN_URL..."
curl -L -o "$BIN_PATH" "$BIN_URL"

chmod +x "$BIN_PATH"

echo "A executar o binário..."
"$BIN_PATH"

rm -rf "$TMP_DIR"
