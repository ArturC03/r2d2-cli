#!/bin/sh

set -e

# Criar diretório temporário
TMP_DIR=$(mktemp -d)
echo "Cloning R2D2 CLI to $TMP_DIR..."
git clone https://github.com/ArturC03/r2d2-cli.git "$TMP_DIR"

# Ir para o diretório
cd "$TMP_DIR"

# Construir o instalador
echo "Building installer..."
go build -o installer/build/r2d2-installer ./installer

# Caminho para o binário
INSTALLER_PATH="$TMP_DIR/installer/build/r2d2-installer"

# Verifica se é root
if [ "$(id -u)" -ne 0 ]; then
  echo "Este script precisa de sudo. A reiniciar com sudo..."
  exec sudo "$0" "$@"
fi

# Executar o instalador
chmod +x "$INSTALLER_PATH"
"$INSTALLER_PATH"

# Limpar
echo "Cleaning up..."
rm -rf "$TMP_DIR"
