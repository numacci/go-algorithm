package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	STLDir   = "stl/function/"
	MainFile = "main.go"
	OutFile  = "atcoder/" + MainFile
	Template = "atcoder/templates/template.go"

	FuncFileSuffix = ".go"
	TestFileSuffix = "_test.go"

	FuncPrefix = "func"
	FuncSuffix = "}"
	NewLine    = "\n"
)

type Function struct {
	Lines []string
}

// gen.go extracts functions from stl directory and generates main.go
func main() {
	fmt.Println("sync operation started")

	fmt.Printf("initialize output file: %s\n", OutFile)
	if err := initialize(); err != nil {
		fmt.Printf("error occured when initializing: %+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("collect functions from %s\n", STLDir)
	functions, err := collectFunctions()
	if err != nil {
		fmt.Printf("error occured when collecting functions: %+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("write functions to %s\n", OutFile)
	if err = writeFunctions(functions); err != nil {
		fmt.Printf("error occured when writing functions: %+v\n", err)
		os.Exit(1)
	}

	fmt.Println("copy generated file to atcoder-cli template directory")
	if err = copyToAtcoderCliDir(); err != nil {
		fmt.Printf("error occured when copying to atcoder-cli directory: %+v", err)
		os.Exit(1)
	}

	fmt.Println("sync operation completed")
}

func initialize() error {
	err := os.Remove(OutFile)
	if err != nil {
		return err
	}

	src, err := os.Open(Template)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(OutFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}

func collectFunctions() ([]Function, error) {
	files, err := getSTLFiles()
	if err != nil {
		return nil, err
	}

	functions := make([]Function, 0)
	for _, file := range files {
		tmp, err := doCollect(file)
		if err != nil {
			return nil, err
		}
		functions = append(functions, tmp...)
	}
	return functions, nil
}

func getSTLFiles() ([]string, error) {
	var files []string
	err := filepath.Walk(STLDir, func(path string, info fs.FileInfo, err error) error {
		paths := strings.Split(path, "/")
		filename := paths[len(paths)-1]
		if strings.HasSuffix(filename, TestFileSuffix) || !strings.HasSuffix(filename, FuncFileSuffix) {
			return nil
		}
		files = append(files, filename)
		return nil
	})
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fmt.Printf("functions will be collected from %s\n", file)
	}
	return files, nil
}

func doCollect(filename string) ([]Function, error) {
	f, err := os.Open(STLDir + filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReaderSize(f, 64)
	inFunc := false
	functions := make([]Function, 0)
	var ln []byte
	for {
		line, isPrefix, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		ln = append(ln, line...)
		if !isPrefix {
			if strings.HasPrefix(string(ln), FuncPrefix) {
				functions = append(functions, Function{Lines: make([]string, 0)})
				inFunc = true
			}
			if inFunc {
				size := len(functions)
				functions[size-1].Lines = append(functions[size-1].Lines, NewLine+string(ln))
			}
			if string(ln) == FuncSuffix {
				size := len(functions)
				functions[size-1].Lines = append(functions[size-1].Lines, NewLine)
				inFunc = false
			}
			ln = nil
		}
	}
	return functions, nil
}

func writeFunctions(functions []Function) error {
	f, err := os.OpenFile(OutFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, function := range functions {
		for _, line := range function.Lines {
			_, err = f.WriteString(line)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func copyToAtcoderCliDir() error {
	out, err := exec.Command("acc", "config-dir").Output()
	if err != nil {
		return err
	}
	accTemplateFile := fmt.Sprintf("%s/go/%s", strings.TrimRight(string(out), "\n"), MainFile)

	if err = os.Remove(accTemplateFile); err != nil {
		return err
	}

	src, err := os.Open(OutFile)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(accTemplateFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}
