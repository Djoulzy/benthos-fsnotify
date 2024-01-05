package main

import (
	"github.com/benthosdev/benthos/v4/public/service"

	// Import all standard Benthos components
	_ "github.com/benthosdev/benthos/v4/public/components/all"

	// Add your plugin packages here
	"github.com/Djoulzy/benthos-fsnotify/input"
)

func main() {
	service.RunCLI(context.Background())
}