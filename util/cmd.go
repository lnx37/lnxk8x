package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func ExecCmd_(command string) (string, error) {
	var err error

	var cmd *exec.Cmd
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd = exec.Command("sh", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Start()

	err = cmd.Wait()

	var output string
	if err == nil {
		output = stdout.String()
	} else {
		output = stderr.String()
	}

	return output, err
}

func ExecCmd__(command string) (string, error) {
	var err error

	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", command)

	var combined_output []byte
	combined_output, err = cmd.CombinedOutput()

	var output string
	output = string(combined_output)

	return output, err
}

func ExecCmd(command string) (string, error) {
	var err error

	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", command)

	var stdout io.ReadCloser
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	cmd.Stderr = cmd.Stdout

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = cmd.Wait()
	if err != nil {
		panic(err)
	}

	var output string = "done"
	return output, err
}
