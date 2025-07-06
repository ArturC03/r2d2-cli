# Minimal R2D2 CLI Installer (PowerShell)

$installerPath = "installer/build/r2d2-installer.exe"
if (-Not (Test-Path $installerPath)) {
    $installerPath = "installer/build/r2d2-installer"
    if (-Not (Test-Path $installerPath)) {
        Write-Host "Installer binary not found at installer/build/r2d2-installer(.exe). Please build it first."
        exit 1
    }
}

Copy-Item $installerPath -Destination .\r2d2-installer.exe -Force
& .\r2d2-installer.exe
Remove-Item .\r2d2-installer.exe -Force 