package k8s

import (
	"io"

	"github.com/gliderlabs/ssh"
	"k8s.io/client-go/tools/remotecommand"
)

const END_OF_TRANSMISSION = "\u0004"

// PtyHandler is what remotecommand expects from a pty
type PtyHandler interface {
	io.Reader
	io.Writer
	remotecommand.TerminalSizeQueue
}

// TerminalSession implements PtyHandler (using a SHH Session connection)
type TerminalSession struct {
	SSHSession ssh.Session
}

// TerminalSize handles pty->process resize events
// Called in a loop from remotecommand as long as the process is running
func (t TerminalSession) Next() *remotecommand.TerminalSize {
	_, winChan, _ := t.SSHSession.Pty()
	select {
	case win := <-winChan:
		return &remotecommand.TerminalSize{Width: uint16(win.Width), Height: uint16(win.Height)}
	}
}

// Read handles pty->process messages (stdin, resize)
func (t TerminalSession) Read(p []byte) (int, error) {
	// log.Info("=Read: ==", string(p), "==byte:", p)
	// n, err := t.SSHSession.Read(p)
	// if err != nil {
	// 	return copy(p, END_OF_TRANSMISSION), err
	// }
	// return n, nil
	return t.SSHSession.Read(p)
}

// Write handles process->pty stdout
// Called from remotecommand whenever there is any output
func (t TerminalSession) Write(p []byte) (int, error) {
	// log.Info("=Write: ==", string(p), "==byte:", p)
	return t.SSHSession.Write(p)
}
