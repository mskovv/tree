package main

import (
	_ "bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

const mainPrefix = "├───"
const subPrefix = "└───"
const levelPrefix = "│"
const tab = "\t"

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, files bool) error {
	err := printTree(out, path, files, "")
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}

	return nil
}

func printTree(out io.Writer, path string, files bool, prefix string) error {
	readDir, _ := os.ReadDir(path)
	var filteredDir []fs.DirEntry

	if !files {
		for _, f := range readDir {
			if f.IsDir() {
				filteredDir = append(filteredDir, f)
			}
		}
	} else {
		filteredDir = readDir
	}

	for i, entry := range filteredDir {
		curPrefix := prefix
		var nextPrefix string

		if i == len(filteredDir)-1 {
			nextPrefix = curPrefix + tab
			curPrefix = prefix + subPrefix
		} else {
			nextPrefix = prefix + levelPrefix + tab
			curPrefix = prefix + mainPrefix
		}

		output := curPrefix + entry.Name()

		if !entry.IsDir() {
			sizeInfo := getFileSize(filepath.Join(path, entry.Name()))
			output += " " + sizeInfo
		}
		_, err := fmt.Fprintf(out, "%s\n", output)
		if err != nil {
			return err
		}
		err = printTree(out, filepath.Join(path, entry.Name()), files, nextPrefix)
		if err != nil {
			return err
		}
	}

	return nil
}

func getFileSize(fileName string) string {
	fInfo, _ := os.Stat(fileName)
	var sizeInfo string
	if fInfo.Size() == 0 {
		sizeInfo = "(empty)"
	} else {
		sizeInfo = "(" + strconv.FormatInt(fInfo.Size(), 10) + "b)"
	}

	return sizeInfo
}
