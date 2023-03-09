package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func main() {
	checkDir()
	generateFile()
	time.Sleep(3 * time.Second)
	gitHandler()
}

func generateFile() {
	maintenanceQuantity()

	timestamp := fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String())))

	os.Create("./files/" + timestamp)

	log("create file: " + timestamp)
}

func maintenanceQuantity() {
	files, err := ioutil.ReadDir("./files")

	if err != nil {
		panic(err)
	}

	if len(files) > 10 {
		for _, file := range files {
			os.Remove("./files/" + file.Name())
			log("remove file: " + file.Name())
		}
	}
}

func checkDir() {
	if _, err := os.Stat("./files"); os.IsNotExist(err) {
		os.Mkdir("./files", os.ModePerm)
	}
}

func log(str string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + ": " + str)
}

func gitHandler() {
	execCommand("git add .")
	execCommand("git commit -m 'update'")
	execCommand("git push")
}

func execCommand(command string) {
	// windows
	cmd := exec.Command("cmd", "/C", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "./"
	cmd.Run()
}

// windows build command
// go build -ldflags "-H windowsgui" -o clockIn.exe main.go

// linux build command in Windows cmd : ubuntu x86_64
// set GOARCH=amd64 set GOOS=linux go build -o clockIn main.go
