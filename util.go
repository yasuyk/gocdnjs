package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	targetDirectory = "." + AppName
	cacheDirectory  = "cache"
)

func userAgent() string {
	return fmt.Sprintf("%s/%s (%s-%s-%s)",
		AppName, AppVersion, runtime.Version(),
		runtime.GOOS, runtime.GOARCH)
}

func CachePath(elem ...string) string {
	path := strings.Join(elem, string(filepath.Separator))
	return filepath.Join(homedir(), targetDirectory, cacheDirectory, path)
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func Exit(msg string) {
	fmt.Printf("%s\n", msg)
	os.Exit(0)
}

func GenereateLink(lib, ver, file string) string {
	return fmt.Sprintf(
		"http://cdnjs.cloudflare.com/ajax/libs/%s/%s/%s",
		lib, ver, file)
}

func getPackage(url string) (resp *http.Response, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent())

	return client.Do(req)
}

func HttpGetPackage(url string) []byte {
	response, err := getPackage(url)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return contents
}

func TrimNewLine(str string) string {
	return strings.Replace(str, "\n", "", -1)
}

func homedir() string {
	if runtime.GOOS == "linux" {
		//user.Current() is not implemented on linux
		return os.Getenv("HOME")
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
