# R2D2 CLI Project Makefile

INSTALLER_SRC = ./installer
INSTALLER_OUT = installer/build/r2d2-installer

.PHONY: installer installer-all clean

installer:
	@mkdir -p installer/build
	GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) go build -o $(INSTALLER_OUT) $(INSTALLER_SRC)
	@echo "Built: $(INSTALLER_OUT) for $(shell go env GOOS)/$(shell go env GOARCH)"

installer-all:
	@mkdir -p installer/build
	GOOS=linux   GOARCH=amd64 go build -o installer/build/r2d2-installer-linux-amd64   $(INSTALLER_SRC)
	GOOS=linux   GOARCH=arm64 go build -o installer/build/r2d2-installer-linux-arm64   $(INSTALLER_SRC)
	GOOS=windows GOARCH=amd64 go build -o installer/build/r2d2-installer-windows-amd64.exe $(INSTALLER_SRC)
	GOOS=windows GOARCH=arm64 go build -o installer/build/r2d2-installer-windows-arm64.exe $(INSTALLER_SRC)
	GOOS=darwin  GOARCH=amd64 go build -o installer/build/r2d2-installer-darwin-amd64   $(INSTALLER_SRC)
	GOOS=darwin  GOARCH=arm64 go build -o installer/build/r2d2-installer-darwin-arm64   $(INSTALLER_SRC)
	@echo "Built all platform installers in installer/build/"

clean:
	rm -rf installer/build/*
	@echo "Cleaned installer/build directory" 
