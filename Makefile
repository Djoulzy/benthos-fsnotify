all:
	make main fstest

fstest:
	go build -o fstest cmd/fsnotify/*

main:
	go build -o benthos-fsnotify