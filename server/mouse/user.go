package mouse

import (
	"errors"
	"log"
)

func e(str string) error {
	return errors.New(str)
}

type User struct {
	Name     string
	Password string
	Settings map[string]string
}

type UserUpdater interface {
	AddUser(user *User) error
	RemoveUser(name string) error
	StartService(config *Config) error
	StopService()
}

type UserContainer struct {
	Users    map[string]*User
	updaters []UserUpdater
}

func (uc *UserContainer) AddUser(name string, password string) error {
	if _, ok := uc.Users[name]; ok {
		return e("Existing user")
	}

	uc.Users[name] = &User{
		Name:     name,
		Password: password,
		Settings: make(map[string]string),
	}

	for _, updater := range uc.updaters {
		if err := updater.AddUser(uc.Users[name]); err != nil {
			log.Fatalf("Updater error")
		}
	}

	return nil
}

func (uc *UserContainer) RemoveUser(name string) error {
	if _, ok := uc.Users[name]; ok {
		return e("Nonexistent user")
	}

	for _, updater := range uc.updaters {
		if err := updater.RemoveUser(name); err != nil {
			panic("Updater error")
		}
	}

	delete(uc.Users, name)

	return nil
}

func (uc *UserContainer) AddUpdater(updater UserUpdater) {
	uc.updaters = append(uc.updaters, updater)
}
