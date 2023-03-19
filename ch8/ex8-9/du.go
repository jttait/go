package main

import (
	"time"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type rootDirectory struct {
	name string
	nbytes int64
	nfiles int64
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan rootDirectory)
	rootDirectories := make(map[string]*rootDirectory)
	for _, root := range roots {
		rootDirectories[root] = &rootDirectory{ root, 0, 0}
	}

	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes, root)
		}
		close(fileSizes)
	}()

	var tick <- chan time.Time
	tick = time.Tick(100 * time.Millisecond)

	loop:
	for {
		select {
		case rootDirectory, ok := <-fileSizes:
			if !ok {
				break loop
			}
			rootDirectories[rootDirectory.name].nfiles += 1
			rootDirectories[rootDirectory.name].nbytes += rootDirectory.nbytes
		case <- tick:
			printDiskUsage(rootDirectories)
		}
	}
	printDiskUsage(rootDirectories)
}

func printDiskUsage(l map[string]*rootDirectory) {
	for key, val := range l {
		fmt.Printf("%s: %d files %.1f Gb\n", key, val.nfiles, float64(val.nbytes)/1e9)
	}
}

func walkDir(dir string, fileSizes chan<- rootDirectory, root string) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, fileSizes, root)
		} else {
			fileSizes <- rootDirectory{ root, entry.Size(), 0}
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
