//go:build !windows

package install

import "fmt"

// installCMD is not supported on non-Windows platforms
func installCMD(opts InstallOptions) error {
	return fmt.Errorf("CMD AutoRun is only supported on Windows. Use --shell bash instead")
}

// uninstallCMD is not supported on non-Windows platforms
func uninstallCMD(opts InstallOptions) error {
	return fmt.Errorf("CMD AutoRun is only supported on Windows")
}

// InstallStartupApp is not supported on non-Windows platforms
func InstallStartupApp(opts InstallOptions) error {
	return fmt.Errorf("Windows Startup App is only supported on Windows")
}

// UninstallStartupApp is not supported on non-Windows platforms
func UninstallStartupApp(opts InstallOptions) error {
	return fmt.Errorf("Windows Startup App is only supported on Windows")
}
