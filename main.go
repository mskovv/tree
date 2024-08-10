package main

import (
	_ "bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const mainPrefix = "├───"
const subPrefix = "└───"
const levelPrefix = "│\t"

var level = 1

func main() {
	out := os.Stdout
	//if !(len(os.Args) == 2 || len(os.Args) == 3) {
	//	panic("usage go run main.go . [-f]")
	//}
	//path := os.Args[1]
	//printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, "testdata", true)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, files bool) error {
	pathStat, err := os.Stat(path) // получаем информацию о полученном пути
	if err != nil {
		return err
	}

	if pathStat.IsDir() { // если это директория
		readDir, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		for idx, entry := range readDir {
			if entry.IsDir() {
				_, err := fmt.Fprintf(out, "%s\n", strings.Repeat(mainPrefix, level)+entry.Name())
				if err != nil {
					return err
				}

				level++
				err = dirTree(out, filepath.Join(path, entry.Name()), files)
				if err != nil {
					return err
				}

			} else if files {
				_, err := fmt.Fprintf(out, "%s\n", strings.Repeat(levelPrefix, level-1)+mainPrefix+entry.Name())
				if err != nil {
					return err
				}
			}

			if idx < len(readDir)-1 {
				continue
			}
		}
		level = 1

	} else if files {
		_, _ = fmt.Fprintf(out, "%s\n", strings.Repeat(mainPrefix, level)+path)
	}

	//for _, entry := range readDir {
	//	if entry.IsDir() {
	//		_, err := fmt.Fprintf(out, "%s\n", strings.Repeat(mainPrefix, level)+entry.Name())
	//		if err != nil {
	//			return err
	//		}
	//	} else if files {
	//		_, err := fmt.Fprintf(out, "%s\n", strings.Repeat(mainPrefix, level)+entry.Name())
	//		if err != nil {
	//			return err
	//		}
	//	}
	//
	//	level++
	//	err := dirTree(out, filepath.Join(path, entry.Name()), files)
	//	if err != nil {
	//		return err
	//	}
	//
	//	//if entry.IsDir() {
	//	//	level++
	//	//	dirPath := filepath.Join(path, entry.Name())
	//	//	entries, err := os.ReadDir(dirPath)
	//	//	if err != nil {
	//	//		continue
	//	//	}
	//	//	for _, dirEntry := range entries {
	//	//		if dirEntry.IsDir() {
	//	//			_, err := fmt.Fprintf(out, "%s\n", levelPrefix+mainPrefix+dirEntry.Name())
	//	//			if err != nil {
	//	//				return err
	//	//			}
	//	//		}
	//	//	}
	//	//
	//	//}
	//
	//	level = 1
	//}

	return fmt.Errorf("%v\n", out)
}
