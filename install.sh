#!/bin/sh
set -e

URL="https://github.com/ArturC03/r2d2-cli/releases/download/v0.1.0/r2d2-installer_linux_amd64_v1/r2d2-installer"
TMPFILE=$(mktemp)

curl -L "$URL" -o "$TMPFILE"
chmod +x "$TMPFILE"
"$TMPFILE"
rm -f "$TMPFILE"
