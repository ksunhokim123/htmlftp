package mouse

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func e(str string) error {
	return errors.New(str)
}

type Address struct {
	Ip   string
	Port string
}

func (ad Address) IsValid() bool {
	if len(ad.Port) == 0 {
		return false
	}
	return true
}

func (ad Address) String() string {
	return ad.Ip + ":" + ad.Port
}

type Config struct {
	Username   string
	Password   string
	Userfile   string
	Keyfile    string
	FtpAddress struct {
		API Address
		FTP Address
	} `yaml:"ftp_address"`
	Address Address
}

type Service struct {
	UserContainer *UserContainer
	KeyContainer  *KeyContainer
	Config        *Config
}

func (sv *Service) Start() {
	for _, updater := range sv.UserContainer.updaters {
		err := updater.StartService(sv.Config)
		if err != nil {
			log.Fatalf(`error during serveice.start %v`, err)
		}
		for _, user := range sv.UserContainer.Users {
			err := updater.AddUser(user)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (sv *Service) Stop() {
	for _, updater := range sv.UserContainer.updaters {
		updater.StopService()
	}
}

func (sv *Service) AddDefaultUpdaters() {
	sv.UserContainer.AddUpdater(&FTPUpdater{})
}

func NewService(configfile string) *Service {
	file, err := ioutil.ReadFile(configfile)
	if err != nil {
		log.Fatalf(`config file open error %v`, err)
	}

	config := Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf(`config file parse error %v`, err)
	}

	if !config.FtpAddress.API.IsValid() {
		log.Fatal("FTP API address is not valid")
	}
	if !config.FtpAddress.FTP.IsValid() {
		log.Fatal("FTP FTP address is not valid")
	}
	if !config.Address.IsValid() {
		log.Fatal("address is not valid")
	}

	sv := &Service{
		UserContainer: &UserContainer{
			Users:    make(map[string]*User),
			updaters: []UserUpdater{},
		},
		KeyContainer: new(KeyContainer),
		Config:       &config,
	}

	sv.AddDefaultUpdaters()
	return sv
}
