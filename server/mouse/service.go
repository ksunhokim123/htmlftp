package mouse

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Address struct {
	Ip   string
	Port string
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
		updater.StartService(sv.Config)
		for _, user := range sv.UserContainer.Users {
			updater.AddUser(user)
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
		log.Fatalf(err.Error())
	}
	config := Config{}
	yaml.Unmarshal(file, &config)

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
