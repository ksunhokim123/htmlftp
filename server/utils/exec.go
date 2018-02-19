package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func copyAndCapture(w io.Writer, r io.Reader) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				fmt.Println(out, " err:", err)
				return
			}
		}
		if err != nil {
			fmt.Print(out)
			if err == io.EOF {
				fmt.Println(" err:", err)
			}
			return
		}
	}
}

func ExecCommand(name string, args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(name, args...)
	stdoutIn, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderrIn, _ := cmd.StderrPipe()
	// :thinking:

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	go copyAndCapture(os.Stdout, stdoutIn)
	go copyAndCapture(os.Stderr, stderrIn)

	// IDEA to figure out does the errStdout matter
	return cmd, nil
}
