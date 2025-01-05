package lnxk8s

import (
	"fmt"
	"log"
	"strings"

	"lnxk8x/util"
)

var {{ data.var }} string = `
{{ data.content }}
`

func {{ data.func }}(args ...string) {
	log.Println(util.GetCurrentFuncName())

	log.Printf("args: %+v\n", args)

	var script string
	script = strings.TrimSpace({{ data.var }})
	script = fmt.Sprintf("%s\n", script)

	util.ExecScript(script, args...)
}
