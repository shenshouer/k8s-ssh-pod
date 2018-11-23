package ssh

import (
	"fmt"
	"strconv"

	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func promptAndSelectNames(sess ssh.Session, term *terminal.Terminal, getNames func() ([]string, error)) (string, bool, error) {
	names, err := getNames()
	if err != nil {
		return "", false, err
	}

	// prompt
	for k, v := range names {
		_, _ = fmt.Fprintf(sess, "[%d] %s \r\n", k+1, v)
	}
	_, _ = fmt.Fprintln(sess, "Please select ID: ")
	sel, err := term.ReadLine()
	if err != nil {
		return "", false, err
	}

	if sel == "quit" || sel == "exit" {
		return "", true, nil
	}

	i, err := strconv.Atoi(sel)
	if err != nil {
		_, _ = fmt.Fprintf(sess, "Invalid input [%s], please again \r\n", sel)
		return promptAndSelectNames(sess, term, getNames)
	}

	if (i < 1) || (i > len(names)) {
		_, _ = fmt.Fprintf(sess, "ID[%d] out of range, please again\r\n", i)
		return promptAndSelectNames(sess, term, getNames)
	}

	return names[i-1], false, nil
}
