package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/sys/windows/registry"
)

const (
	appCompatKey = `Software\Microsoft\Windows NT\CurrentVersion\AppCompatFlags\Layers` //For compatibility shit
	waitSeconds  = 15
)

var searchPaths = []string{ //Tipical Path for Yahoo Messenger
	`C:\Program Files (x86)\Yahoo!\Messenger`,
	`C:\Program Files\Yahoo!\Messenger`,
}

type ExeFile struct {
	Index int
	Name  string
	Path  string
}

func main() {
	fmt.Println("================================================")
	fmt.Println("  Yahoo Web Search Fixer v0.1")
	fmt.Println("  For Yahoo Messenger")
	fmt.Println("================================================")
	fmt.Println()

	// Check administrator
	if !isAdmin() {
		fmt.Println("[ERROR] This program requires administrator privileges.")
		fmt.Println("Run with 'Run as administrator'.")
		fmt.Println()
		fmt.Print("Press Enter to exit...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}

	// Search for Yahoo path
	fmt.Println("[>>] Searching for Yahoo Messenger executable files...")
	fmt.Println()

	var exeFiles []ExeFile
	index := 1

	for _, searchPath := range searchPaths {
		if _, err := os.Stat(searchPath); err == nil {
			fmt.Printf("[OK] Directory found: %s\n", searchPath)
			
			files, err := filepath.Glob(filepath.Join(searchPath, "*.exe"))
			if err == nil {
				for _, file := range files {
					exeFiles = append(exeFiles, ExeFile{
						Index: index,
						Name:  filepath.Base(file),
						Path:  file,
					})
					fmt.Printf("  [%d] %s\n", index, filepath.Base(file))
					fmt.Printf("      Full path: %s\n", file)
					fmt.Println()
					index++
				}
			}
		} else {
			fmt.Printf("[--] The Director stated: %s\n", searchPath)
			fmt.Println()
		}
	}

	if len(exeFiles) == 0 {
		fmt.Println("[ERROR] No .exe file found")
		fmt.Println("Check if Yahoo Messenger is installed correctly.")
		fmt.Println()
		fmt.Print("Press Enter to exit...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}

	// select exe
	fmt.Println("================================================")
	fmt.Printf("  Select the file number (1-%d):\n", len(exeFiles))
	fmt.Println("================================================")
	fmt.Println()

	var choice int
	for {
		fmt.Print("Enter the number: ")
		_, err := fmt.Scanf("%d\n", &choice)
		if err == nil && choice >= 1 && choice <= len(exeFiles) {
			break
		}
		fmt.Printf("[ERROR] Invalid selection. Choose a number between 1 and %d.\n", len(exeFiles))
		// clear buffer
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}

	selectedFile := exeFiles[choice-1]
	fmt.Println()
	fmt.Printf("[OK] You have selected: %s\n", selectedFile.Path)

	// YahooMessenger.exe?
	if !strings.EqualFold(selectedFile.Name, "YahooMessenger.exe") {
		fmt.Printf("[ATTENTION] You have selected: %s\n", selectedFile.Name)
		fmt.Println("         This does not appear to be YahooMessenger.exe.")
		fmt.Print("Are you sure you want to continue? (Y/N): ")
		
		var confirm string
		fmt.Scanf("%s\n", &confirm)
		if strings.ToUpper(confirm) != "Y" {
			fmt.Println("[>>] Operation cancelled by the user.")
			fmt.Println()
			fmt.Print("Press Enter to exit...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		}
	}

	// Check if Yahoo Messenger is already running and close it
	if isProcessRunning(selectedFile.Name) {
		fmt.Println()
		fmt.Printf("[!!] %s is already running!\n", selectedFile.Name)
		fmt.Printf("[>>] Closing %s before starting the fix...\n", selectedFile.Name)
		closeApplication(selectedFile.Name)
		time.Sleep(2 * time.Second)
		
		// Check again if it's really closed
		if isProcessRunning(selectedFile.Name) {
			fmt.Printf("[ERROR] Could not close %s. Please close it manually and try again.\n", selectedFile.Name)
			fmt.Println()
			fmt.Print("Press Enter to exit...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			return
		}
		fmt.Printf("[OK] %s closed successfully!\n", selectedFile.Name)
	}

	// There's some weird bug (with no compatibility mode), After I noticed in IDA Pro that the program runs normally without any weird bugs after this steps ¯\_(ツ)_/¯. Well, it's a case of "it just works."
	// Set compatibility to Windows XP SP3
	Compatibility(selectedFile, "WINXPSP3", "Set compatibility to Windows XP SP3")

	// Set compatibility to Windows 7
	Compatibility(selectedFile, "WIN7RTM", "Set compatibility to Windows 7")

	// Remove compatibility 
	Compatibility(selectedFile, "", "Remove compatibility")

	// After completing those stages, it automatically fixes the bug, weird Windows shit.
	fmt.Println()
	fmt.Println("================================================")
	fmt.Println("  Now you'll have Yahoo Web Search on Yahoo Messenger again without using compatibility mode.")
	fmt.Println("================================================")
	fmt.Println()
	fmt.Printf("  Selected file: %s\n", selectedFile.Path)
	fmt.Println()
	fmt.Print("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func isAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

// Check if a process is running by its name
func isProcessRunning(processName string) bool {
	// Tasklist method
	cmd := exec.Command("tasklist", "/fi", fmt.Sprintf("IMAGENAME eq %s", processName))
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	
	// If process is found, the output will contain the process name
	return strings.Contains(string(output), processName)
}

func Compatibility(file ExeFile, mode string, description string) {
	fmt.Println()
	fmt.Println("================================================")
	fmt.Printf("  %s\n", description)
	fmt.Println("================================================")
	fmt.Println()

	// Clear the settings
	removeCompatibility(file.Path)

	// Setare compatibilitate (daca e cazul)
	if mode != "" {
		err := setCompatibility(file.Path, mode)
		if err != nil {
			fmt.Printf("[ERROR] Could not set compatibility %s: %v\n", mode, err)
			return
		}
		fmt.Printf("[OK] Compatibility %s set\n", mode)
	} else {
		fmt.Println("[OK] Compatibility DISABLED")
	}

	fmt.Printf("[>>] Start %s...\n", file.Name)
	cmd := exec.Command(file.Path)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("[ERROR] Could not start the application: %v\n", err)
		return
	}

	// 1..2...3... Let's GO BABYYY
	fmt.Printf("[>>] Waiting %d seconds...\n", waitSeconds)
	for i := waitSeconds; i > 0; i-- {
		fmt.Printf("     %d seconds remaining...\n", i)
		time.Sleep(1 * time.Second)
	}

	// Close yahoo messenger
	fmt.Printf("[>>] Closing %s...\n", file.Name)
	closeApplication(file.Name)

	// Wait the shit to close
	time.Sleep(2 * time.Second)

	fmt.Printf("[OK] %s complete!\n", description)
}

func setCompatibility(filePath string, mode string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, appCompatKey, registry.SET_VALUE)
	if err != nil {
		key, _, err = registry.CreateKey(registry.CURRENT_USER, appCompatKey, registry.SET_VALUE)
		if err != nil {
			return fmt.Errorf("Could not access the registry: %v", err)
		}
	}
	defer key.Close()

	return key.SetStringValue(filePath, mode)
}

func removeCompatibility(filePath string) {
	// Clear from HKCU
	key, err := registry.OpenKey(registry.CURRENT_USER, appCompatKey, registry.SET_VALUE)
	if err == nil {
		key.DeleteValue(filePath)
		key.Close()
	}

	// Clear from HKLM (need admin)
	key, err = registry.OpenKey(registry.LOCAL_MACHINE, appCompatKey, registry.SET_VALUE)
	if err == nil {
		key.DeleteValue(filePath)
		key.Close()
	}
}

func closeApplication(processName string) {
	// Taskkill 1/2
	exec.Command("taskkill", "/f", "/im", processName).Run()

	// FallBack Powershell
	exec.Command("powershell", "-Command", 
		fmt.Sprintf("Stop-Process -Name '%s' -Force -ErrorAction SilentlyContinue", 
			strings.TrimSuffix(processName, ".exe"))).Run()
}