package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ExecScript(script string, args ...string) {
	var err error

	var ip string
	ip = ""
	if len(args) == 1 {
		ip = args[0]
	}

	var file string
	file = ".tmp.sh"

	{
		var script2 []byte
		script2 = []byte(script)

		err = ioutil.WriteFile(file, script2, 0644)
		if err != nil {
			panic(err)
		}
	}

	defer os.Remove(file)

	{
		var bash string
		bash = "bash"
		if SETTINGS.DEBUG {
			bash = "bash -x"
		}

		var cmd string
		if ip == "" {
			cmd = fmt.Sprintf("%s %s", bash, file)
		} else {
			cmd = fmt.Sprintf("%s %s %s", bash, file, ip)
		}

		var output string
		output, err = ExecCmd(cmd)

		log.Println(output)

		if err != nil {
			panic(err)
		}
	}
}
