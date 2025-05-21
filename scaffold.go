package main

import (
	"fmt"
	"os"
)

var services = []string{"auth", "course", "enrollment", "certification", "progress"}
var commonFolders = []string{"config", "handler", "model", "repository", "setting", "proto"}

func createStructure() {
	base := "services"

	for _, service := range services {
		for _, folder := range commonFolders {
			path := fmt.Sprintf("%s/%s/%s", base, service, folder)
			os.MkdirAll(path, os.ModePerm)
			// Create a placeholder Go file
			filePath := fmt.Sprintf("%s/%s/%s.go", base, service, folder, folder)
			os.WriteFile(filePath, []byte(fmt.Sprintf("package %s\n", folder)), 0644)
		}
		// Create main.go for the service
		mainFile := fmt.Sprintf("%s/%s/main.go", base, service)
		os.WriteFile(mainFile, []byte("package main\n\nfunc main() {}\n"), 0644)
	}

	// Common shared folder
	commonDirs := []string{"pkg/utils", "pkg/middleware", "api/proto"}
	for _, dir := range commonDirs {
		os.MkdirAll(dir, os.ModePerm)
	}

	// Root README
	os.WriteFile("README.md", []byte("# Nerd-station23: LMS built with Go (gRPC + MongoDB + Microservices)\n"), 0644)
}

func main() {
	createStructure()
}
