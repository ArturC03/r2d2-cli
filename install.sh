#!/bin/sh
set -e

ARCH=$(uname -m)
OS=$(uname | tr '[:upper:]' '[:lower:]')

BIN_NAME="r2d2-installer"
VERSION="v0.1.0"
BASE_URL="https://github.com/ArturC03/r2d2-cli/releases/download/${VERSION}"

case $ARCH in
  x86_64) ARCH="amd64" ;;
  aarch64 | arm64) ARCH="arm64" ;;
  *) echo "Arquitetura $ARCH não suportada"; exit 1 ;;
esac

BIN_FILE="${BIN_NAME}_${OS}_${ARCH}"
BIN_URL="${BASE_URL}/${BIN_FILE}"

TMP_DIR=$(mktemp -d)
BIN_PATH="$TMP_DIR/$BIN_NAME"

download_and_run() {
  echo "A transferir $BIN_URL..."
  curl -L -o "$BIN_PATH" "$BIN_URL"

  chmod +x "$BIN_PATH"

  echo "A executar o binário..."
  "$BIN_PATH"
}

if [ "$(id -u)" -ne 0 ]; then
  echo "Este script precisa de ser executado com sudo. A reiniciar com sudo..."
  exec sudo sh "$0" "$@"
else
  download_and_run
  rm -rf "$TMP_DIR"
fi
