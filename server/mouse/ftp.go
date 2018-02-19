package mouse

import (
	"os/exec"

	"github.com/sunho/mouse-hosting/server/utils"
)

type FTPUpdater struct {
	cmd *exec.Cmd
}

func (fu *FTPUpdater) AddUser(user *User) error {
	return nil
}

func (fu *FTPUpdater) RemoveUser(name string) error {
	return nil
}

func (fu *FTPUpdater) StartService(config *Config) error {
	address := config.FtpAddress
	fu.cmd = utils.ExecCommand("mouseftp", "-fi", address.FTP.Ip, address.FTP.Port, "-ai", address.API.Ip, address.API.Port)
	return nil
}

func (fu *FTPUpdater) StopService() {
	fu.cmd.Process.Kill()
}
