package input

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/benthosdev/benthos/v4/public/service"
)

var fsnotifyConfigSpec = service.NewConfigSpec().
	Summary("Input that waits for file system events from a directory.").
	Field(service.NewStringField("path").Default("./"))

func newFSNotifyInput(conf *service.ParsedConfig) (service.Input, error) {
	path, err := conf.FieldString("path")
	if err != nil {
		return nil, err
	}
	if path == "" {
		return nil, errors.New("path cannot be empty")
	}
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("cannot access path %s", path)
	}
	return service.AutoRetryNacks(&fsnotifyInput{path}), nil
}

func init() {
	err := service.RegisterInput(
		"fsnotify", fsnotifyConfigSpec,
		func(conf *service.ParsedConfig, mgr *service.Resources) (service.Input, error) {
			return newFSNotifyInput(conf)
		})
	if err != nil {
		panic(err)
	}
}

//------------------------------------------------------------------------------

type fsnotifyInput struct {
	path string
}

func (f *fsnotifyInput) Connect(ctx context.Context) error {
	return nil
}

func (f *fsnotifyInput) Read(ctx context.Context) (*service.Message, service.AckFunc, error) {
	b := []byte("test")
	return service.NewMessage(b), func(ctx context.Context, err error) error {
		// Nacks are retried automatically when we use service.AutoRetryNacks
		return nil
	}, nil
}

func (f *fsnotifyInput) Close(ctx context.Context) error {
	return nil
}

// func try() {
// 	// Create new watcher.
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer watcher.Close()

// 	// Start listening for events.
// 	go func() {
// 		for {
// 			select {
// 			case event, ok := <-watcher.Events:
// 				if !ok {
// 					return
// 				}
// 				log.Println("event:", event)
// 				if event.Has(fsnotify.Write) {
// 					log.Println("modified file:", event.Name)
// 				}
// 			case err, ok := <-watcher.Errors:
// 				if !ok {
// 					return
// 				}
// 				log.Println("error:", err)
// 			}
// 		}
// 	}()

// 	// Add a path.
// 	err = watcher.Add("/tmp")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Block main goroutine forever.
// 	<-make(chan struct{})
// }
