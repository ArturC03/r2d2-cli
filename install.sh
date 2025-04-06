#!/bin/bash

sudo pacman -S deno
sudo pacman -S golang
git clone https://github.com/ArturC03/r2d2-cli
go build .
mv r2d2-cli /usr/bin/r2d2

