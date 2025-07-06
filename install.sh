#!/bin/sh

# Minimal R2D2 CLI Installer

SCRIPT_DIR="$(dirname "$(realpath "$0")")"
INSTALLER_PATH="$SCRIPT_DIR/installer/build/r2d2-installer"
echo "Installer path: $INSTALLER_PATH"

# Check if the script is run as root
if [ "$(id -u)" -ne 0 ]; then
  echo "This script must be run with sudo. Re-running with sudo..."
  exec sudo "$0" "$@"
fi

if [ ! -f "$INSTALLER_PATH" ]; then
  echo "Installer binary not found at $INSTALLER_PATH Please build it first."
  exit 1
fi

cp "$INSTALLER_PATH" ./r2d2-installer
chmod +x ./r2d2-installer
./r2d2-installer
rm -f ./r2d2-installer 
