//go:build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Check if rsrc is available
	_, err := exec.LookPath("rsrc")
	if err != nil {
		fmt.Println("Note: rsrc not found. To embed manifest, install:")
		fmt.Println("  go install github.com/akavel/rsrc@latest")
		fmt.Println("Then run: go generate")
		return
	}

	// Create .syso file with embedded manifest
	manifest := "hellogang.manifest"
	output := "hellogang.syso"

	if _, err := os.Stat(manifest); err != nil {
		fmt.Printf("Manifest file %s not found\n", manifest)
		return
	}

	cmd := exec.Command("rsrc", "-manifest", manifest, "-o", output)
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("Failed to create syso: %v\n%s\n", err, output)
		os.Exit(1)
	}

	fmt.Printf("Created %s from %s\n", output, manifest)
}
