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

BIN_FILE="${BIN_NAME}_${OS}_${ARCH}"  # Sem o _v1 se não existir
BIN_URL="${BASE_URL}/${BIN_FILE}"

echo "A transferir de: $BIN_URL"

HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" "$BIN_URL")
if [ "$HTTP_STATUS" != "200" ]; then
  echo "Erro: ficheiro não encontrado na URL: $BIN_URL"
  exit 1
fi

TMP_DIR=$(mktemp -d)
BIN_PATH="$TMP_DIR/$BIN_NAME"

curl -L -o "$BIN_PATH" "$BIN_URL"
chmod +x "$BIN_PATH"

echo "A executar o binário..."
sudo "$BIN_PATH"

rm -rf "$TMP_DIR"
