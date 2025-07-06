![R2D2 Logo](https://r2d2-lang.kinsta.app/assets/images/r2d2-dark.svg)

# R2D2 - The programming Language

**R2D2** is a programming language designed for building modular JavaScript applications using a clear, explicit, and structured syntax.

It is **written in Go** and **compiles to JavaScript**, embracing a module-based architecture where each module can contain variables and functions. These modules are compiled into native JavaScript objects, making it easy to integrate with existing frontend or backend projects.

R2D2 introduces the concept of **pseudo-functions** — a special type of function that can only contain calls to other functions. This enforces composability and encourages building code through reusable, isolated behaviors.

The language is built with simplicity and clarity in mind. It avoids hidden behavior and favors explicit patterns, making it approachable for those who want to write structured, maintainable code that compiles directly to JavaScript.

You can start using **R2D2** by executing the following command on your system.

## Installation

### Quick Install (Recommended)

**Linux/macOS:**
```bash
curl -sSL https://raw.githubusercontent.com/ArturC03/r2d2-cli/main/install.sh | bash
```

**Windows (PowerShell):**
```powershell
iwr -useb https://raw.githubusercontent.com/ArturC03/r2d2-cli/main/install.ps1 | iex
```

This will:
- 🚀 Download the appropriate installer for your platform
- 🎨 Launch a beautiful TUI installer with R2D2 branding
- 📦 Install R2D2 CLI to your system automatically
- 🧹 Clean up temporary files when done

**Simple and fast** - just one command!

### Manual Installation

If you prefer to download manually:

1. Go to [Releases](https://github.com/ArturC03/r2d2-cli/releases/latest)
2. Download `r2d2-installer-{your-os}-{your-arch}.tar.gz`
3. Extract: `tar -xzf r2d2-installer-*.tar.gz`
4. Run: `./r2d2-installer`

### Requirements

- **Linux/macOS**: `curl` and `tar`
- **Windows**: PowerShell 3.0+ and Windows 10 1903+ (for tar)

### For Windows

```bash
Just use WSL, believe me. I'm doing you a favor.
```

### Verification

To verify the installation, try executing `r2d2`, if it returns something you're good to go!

## Hello World

Let's try to make a hello world program!

1. Create a `.r2d2` file such as `helloworld.r2d2`
2. Put the following content in the file

```js
module HelloWorld {
  export fn main() {
    console.log("Hello World!");
  }
}
```
3. Now you can choose to compile, run or even convert it into a js file by using on of the following commands

```bash
r2d2 build helloworld.r2d2
``` 

```bash
r2d2 run helloworld.r2d2
``` 

```bash
r2d2 js helloworld.r2d2
```

For more information go to the [site](https://r2d2-lang.kinsta.app/docs). Any doughts just make an issue.

## Installation Details

The installation process is designed to be **simple and automatic**:

1. **Platform Detection**: Automatically detects your OS and architecture
2. **Download**: Gets the right installer binary from GitHub releases
3. **TUI Installer**: Launches a beautiful terminal interface
4. **Installation**: Installs R2D2 CLI and dependencies automatically
5. **Cleanup**: Removes temporary files

### TUI Installer Features

- 🎨 **R2D2 Branding**: Purple R, Green D, adaptive colors
- 📱 **Responsive**: Adapts to any terminal size
- 🔄 **Progress Tracking**: Real-time installation progress
- ✨ **Error Handling**: Clear messages and troubleshooting

### Supported Platforms

- **Linux**: x86_64, ARM64, i386
- **macOS**: Intel and Apple Silicon  
- **Windows**: x86_64, i386
- **FreeBSD**: x86_64

### What Gets Installed

The installer will:
- Install Git, Go, and Deno if needed
- Build and install the R2D2 CLI
- Add R2D2 to your system PATH
- Verify the installation works
