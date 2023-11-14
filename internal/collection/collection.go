package collection

import (
	"fmt"
	"os"

	postman "github.com/rbretecher/go-postman-collection"
)

func FromPaths(paths []string) []postman.Collection {
	var collections []postman.Collection

	possibleCollectionFiles, err := filesFromPaths(paths)
	if err != nil {
		return collections
	}

	for _, fileName := range possibleCollectionFiles {
		collection, err := parseCollectionFile(fileName)
		if err != nil {
			fmt.Println("Error parsing collection file", err)
			continue
		}
		if len(collection.Items) > 0 {
			collections = append(collections, collection)
		}
	}

	return collections
}

func CollectionNames(collections []postman.Collection) []string {
	var names []string
	for _, collection := range collections {
		names = append(names, collection.Info.Name)
	}
	return names
}

func IsZeroed(collection postman.Collection) bool {
	return collection.Info == (postman.Info{})
}

func parseCollectionFile(fileName string) (postman.Collection, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file", err)
		return postman.Collection{}, err
	}
	defer file.Close()

	collection, err := postman.ParseCollection(file)
	if err != nil {
		fmt.Println("Error parsing collection", err)
		return postman.Collection{}, err
	}

	return *collection, nil
}
