# cdeamon 注意程序名字最长15个字符
```
package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/wentome/cdeamon"
)

func main() {
	flag.Parse()
	flagArgs := flag.Args()
	fmt.Println(flagArgs)

	if len(flagArgs) > 0 {
		if flagArgs[0] == "stop" {
			err := cdeamon.Stop()
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		} else if flagArgs[0] == "restart" {
			err := cdeamon.Stop()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	if cdeamon.IsRunning() {
		fmt.Println("already run")
		return
	}
	if cdeamon.IsDeamon() {
		return
	}
	for i := 1; i < 10; i++ {
		time.Sleep(time.Second)
	}
}
```
