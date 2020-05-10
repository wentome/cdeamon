package cdeamon

import (
	"bytes"
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
	if len(FindProcessPidByName(appName)) > 1 {
		return true
	} else {
		return false
	}
}

func Stop() error {
	appName := filepath.Base(os.Args[0])
	pids := FindProcessPidByName(appName)
	if len(pids) > 1 {
		for _, pid := range pids[:len(pids)-1] {
			err := KillProcess(pid)
			if err != nil {
				return err
			}
		}
		time.Sleep(time.Millisecond * 50)
	}
	return nil
}

func FindProcessPidByName(processName string) []int {
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

func KillProcess(pid int) error {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	err = proc.Kill()
	if err != nil {
		return err
	}
	return nil
}
