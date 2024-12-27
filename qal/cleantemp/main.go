package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func cleanTempFiles() {
	tempPath := os.Getenv("TEMP")

	if tempPath == "" {
		fmt.Println("TEMP path is not valid.")
		return
	}

	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil
		}

		if !info.IsDir() {
			err = os.Remove(path)
			if err != nil {
				fmt.Printf("Failed to delete file: %s. Reason: %v\n", path, err)
			} else {
				fmt.Printf("Deleted file: %s\n", path)
			}
		} else if path != tempPath {
			err = os.RemoveAll(path)
			if err != nil {
				fmt.Printf("Failed to delete directory: %s. Reason: %v\n", path, err)
			} else {
				fmt.Printf("Deleted directory: %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through TEMP directory: %v\n", err)
	}
}

func cleanPercentTempFiles() {
	systemRoot := os.Getenv("SYSTEMROOT")
	if systemRoot == "" {
		fmt.Println("SYSTEMROOT path is not valid.")
		return
	}

	percentTempPath := filepath.Join(systemRoot, "Temp")

	err := filepath.Walk(percentTempPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil
		}

		if !info.IsDir() {
			err = os.Remove(path)
			if err != nil {
				fmt.Printf("Failed to delete file: %s. Reason: %v\n", path, err)
			} else {
				fmt.Printf("Deleted file: %s\n", path)
			}
		} else if path != percentTempPath {
			err = os.RemoveAll(path)
			if err != nil {
				fmt.Printf("Failed to delete directory: %s. Reason: %v\n", path, err)
			} else {
				fmt.Printf("Deleted directory: %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through SYSTEMROOT\\Temp directory: %v\n", err)
	}
}

func runCleanmgr() {
	cmd := exec.Command("cleanmgr", "/sagerun:1")
	err := cmd.Run()

	if err != nil {
		fmt.Printf("Failed to run Disk Cleanup. Reason: %v\n", err)
	} else {
		fmt.Println("Disk Cleanup completed successfully.")
	}
}

func main() {
	fmt.Println("Cleaning TEMP files...")
	cleanTempFiles()

	fmt.Println("Cleaning %TEMP% files...")
	cleanPercentTempFiles()

	fmt.Println("Running Disk Cleanup...")
	runCleanmgr()

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
