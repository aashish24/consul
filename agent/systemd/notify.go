package systemd

import (
	"errors"
	"net"
	"os"
)

var NotifyNoSocket = errors.New("No socket")

// Notifier provides a method to send a message to systemd.
type Notifier struct{}

// Notify sends a message to the init daemon. It is common to ignore the error.
func (n *Notifier) Notify(state string) error {
	addr := &net.UnixAddr{
		Name: os.Getenv("NOTIFY_SOCKET"),
		Net:  "unixgram",
	}

	if addr.Name == "" {
		return NotifyNoSocket
	}

	conn, err := net.DialUnix(addr.Net, nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(state))
	return err
}