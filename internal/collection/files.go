package collection

import (
	"fmt"
	"os"
	"strings"
)

func filesFromPaths(paths []string) ([]string, error) {
	var collectionFiles []string
	for _, path := range paths {
		file, err := os.Stat(path)
		if err != nil {
			continue
		}

		if file.IsDir() {
			collectionDir, err := os.Open(path)
			if err != nil {
				fmt.Println("Error opening directory", err)
				return nil, err
			}
			defer collectionDir.Close()

			entries, err := collectionDir.ReadDir(0)
			if err != nil {
				fmt.Println("Error reading directory", err)
				return nil, err
			}
			collectionFiles = append(collectionFiles, collectionFilesFromDir(collectionDir, entries)...)
		} else {
			if isCollectionFile(file.Name()) {
				collectionFiles = append(collectionFiles, path)
			}
		}
	}

	return collectionFiles, nil
}

func collectionFilesFromDir(dir *os.File, dirEntries []os.DirEntry) []string {
	var collectionFiles []string
	for _, entry := range dirEntries {
		if !entry.IsDir() && isCollectionFile(entry.Name()) {
			collectionFiles = append(collectionFiles, dir.Name()+"/"+entry.Name())
		}
	}
	return collectionFiles
}

func isCollectionFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".json")
}
