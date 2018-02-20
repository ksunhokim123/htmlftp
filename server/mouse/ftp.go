package mouse

import (
	"net/http"
	"net/url"
	"os/exec"

	"github.com/sunho/mouse-hosting/server/utils"
)

type FTPUpdater struct {
	cmd    *exec.Cmd
	ftpapi string
}

func (fu *FTPUpdater) AddUser(user *User) error {
	_, err := http.PostForm(fu.ftpapi+"/users",
		url.Values{"username": {user.Name}, "password": {user.Password}})
	if err != nil {
		return err
	}
	//TODO error resp
	return nil
}

func (fu *FTPUpdater) RemoveUser(name string) error {
	return nil
}

func (fu *FTPUpdater) StartService(config *Config) error {
	address := config.FtpAddress
	cmd, err := utils.ExecCommand("mouseftp", "-fi", address.FTP.Ip, address.FTP.Port, "-ai", address.API.Ip, address.API.Port)
	if err != nil {
		return err
	}
	fu.cmd = cmd
	fu.ftpapi = "http://" + config.FtpAddress.API.String()
	return nil
}

func (fu *FTPUpdater) StopService() {
	fu.cmd.Process.Kill()
}
