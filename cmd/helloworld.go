package main

import (
	"expvar"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const workdir string = "/opt/go_sample"

func main() {

	// Debug
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println("Executable path: ", exPath)

	fmt.Println("~~~~~~~~~~~ Files in working directory ~~~~~~~~~~~")
	tree(workdir)
	fmt.Println("~~~~~~~~~~~ Files in working directory ~~~~~~~~~~~")

	testExpvarPackage()
}

func testExpvarPackage() {
	tick := time.NewTicker(1 * time.Second)
	num_go := expvar.NewInt("runtime.goroutines")
	counters := expvar.NewMap("counters")
	counters.Set("cnt1", new(expvar.Int))
	counters.Set("cnt2", new(expvar.Float))

	go http.ListenAndServe(":8080", nil) // --> http://localhost:8081/debug/vars

	for {
		select {
		case <-tick.C:
			num_go.Set(int64(runtime.NumGoroutine()))
			counters.Add("cnt1", 1)
			counters.AddFloat("cnt2", 1.452)
		}
	}

}

func tree(root string) error {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(root, file.Name())

		if file.IsDir() {
			fmt.Println("Dir: " + filePath + "/")
			err := tree(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("File: " + filePath)
		}
	}
	return nil
}
