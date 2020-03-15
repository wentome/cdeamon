package cdeamon

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

func IsDeamon() bool {
	if os.Getppid() != 1 {
		filePath, _ := filepath.Abs(os.Args[0])
		args := append([]string{filePath}, os.Args[1:]...)
		os.StartProcess(filePath, args, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
		return true
	} else {
		return false
	}
}
func IsRunning() bool {
	appName := filepath.Base(os.Args[0])
	if len(findProcessPidByName(appName)) > 1 {
		return true
	} else {
		return false
	}
}

func Stop() {
	appName := filepath.Base(os.Args[0])
	pids := findProcessPidByName(appName)
	fmt.Println(pids)
	if len(pids) > 1 {
		for _, pid := range pids[:len(pids)-1] {
			fmt.Println(pid)
			killProcess(pid)
		}
		time.Sleep(time.Millisecond * 50)
	}
}

func findProcessPidByName(processName string) []int {
	var pids []int
	fd, _ := ioutil.ReadDir("/proc")
	for _, fi := range fd {
		fiName := fi.Name()
		pid, err := strconv.Atoi(fiName)
		if err == nil {
			statusFile := path.Join("/proc", fiName, "status")
			f, err := ioutil.ReadFile(statusFile)
			if err != nil {
				continue
			}

			name := string(f[6:bytes.IndexByte(f, '\n')])
			if name == processName {
				pids = append(pids, pid)
			}

		} else {
			continue
		}
	}
	return pids
}

func killProcess(pid int) {
	proc, _ := os.FindProcess(pid)
	proc.Kill()
}
