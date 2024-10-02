package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func isInitialized() bool {
	files, e := os.ReadDir("./")
	if e != nil {
		panic(e)
	}
	for _, file := range files {
		if file.Name() == ".go-git" {
			return true
		}
	}
	return false
}

func InitializeTracker() {
	if !isInitialized() {
		e := os.Mkdir(".go-git", 0755)
		if e != nil {
			panic(e)
		}
		e = os.Mkdir(".go-git/snapshots", 0755)
		if e != nil {
			panic(e)
		}
		e = os.WriteFile(".go-git/ignores.txt", []byte(""), 0755)
		if e != nil {
			panic(e)
		}
		fmt.Println("go-git init success")
		fmt.Println("[REMINDER] Go ahead and fill the ./.go-git/ignores.txt right away before you snap anything!")
	} else {
		fmt.Println("go-git already initialized in this directory!")
	}
}

func getFileContent(filePath string) string {
	data, e := os.ReadFile(filePath)
	if e != nil {
		panic(e)
	} else {

		return string(data)
	}
}

func WriteFileContent(filePath string, content string) {
	e := os.WriteFile(filePath, []byte(content), 0644)
	if e != nil {
		panic(e)
	}
}

func ReadJsonFile(filePath string) map[string]string {
	data, _ := os.ReadFile(filePath)
	var tracker map[string]string
	json.Unmarshal(data, &tracker)
	return tracker
}

func WriteJsonFile(snapShotName *string, content map[string]string) {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, ".go-git", "snapshots", *snapShotName+".json")
	fmt.Println(filePath)
	file, e := os.Create(filePath)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "	")
	encoder.Encode(content)
}

func EncodeToB64(content string) string {
	b64Content := b64.StdEncoding.EncodeToString([]byte(content))
	return b64Content
}

func DecodeFromB64(b64Content string) string {
	content, e := b64.StdEncoding.DecodeString(b64Content)
	if e != nil {
		panic(e)
	} else {
		return string(content)
	}
}

func getIgnores() []string {
	content := getFileContent(".go-git/ignores.txt")
	ignores := strings.Split(string(content), "\n")
	for i := range len(ignores) {
		ignores[i] = strings.TrimSpace(ignores[i])
	}
	return ignores
}

func isFileInIgnores(ignores *[]string, path string) bool {
	for _, ignore := range *ignores {
		if strings.Contains(path, ignore) || filepath.Ext(path) == ignore {
			return true
		}
	}
	return false
}

func traverseDirs(dirPath string) []string {
	var paths []string
	ignores := getIgnores()
	e := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, e error) error {
		if e != nil {
			return e
		}
		unixStyledPath := filepath.ToSlash(path)
		if !d.IsDir() && !strings.Contains(path, ".go-git") && !isFileInIgnores(&ignores, unixStyledPath) {
			paths = append(paths, unixStyledPath)
		}
		return nil
	})
	if e != nil {
		panic(e)
	}
	return paths
}

func Track() map[string]string {
	cwd, _ := os.Getwd()
	filePaths := traverseDirs(cwd)
	tracker := make(map[string]string)
	for _, filePath := range filePaths {
		data := getFileContent(filePath)
		encodedData := EncodeToB64(data)
		tracker[filePath] = encodedData
	}
	return tracker
}

func sanitizeSnapshotName(snapshotName *string) error {
	if strings.Contains(*snapshotName, "/") ||
		strings.Contains(*snapshotName, ".") ||
		strings.Contains(*snapshotName, "\\") ||
		*snapshotName == "" {
		return errors.New("invalid string")
	}
	// Remove leading/trailing spaces
	*snapshotName = strings.TrimSpace(*snapshotName)
	// Replace space and tabs with dash
	*snapshotName = strings.ReplaceAll(*snapshotName, " ", "-")
	*snapshotName = strings.ReplaceAll(*snapshotName, "\t", "-")
	// Replace \n and \r with ''
	*snapshotName = strings.ReplaceAll(*snapshotName, "\n", "")
	*snapshotName = strings.ReplaceAll(*snapshotName, "\r", "")
	return nil
}

func ProcessSnapshotName(snapshotName *string) {
	if e := sanitizeSnapshotName(snapshotName); e != nil {
		panic(e)
	}
	timestamp := time.Now().Format("20060102-1504")
	*snapshotName = fmt.Sprintf("%s-%s", timestamp, *snapshotName)
}
