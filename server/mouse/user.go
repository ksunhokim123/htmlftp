package mouse

import "errors"

func e(str string) error {
	return errors.New(str)
}

type UserUpdater struct {
	AddUser       func(name string, password string) error
	RemoveUser    func(name string) error
	ChangeSetting func(name string, key string, value string) error //TODO
	StartService  func(users []*User) error
	StopService   func(users []*User) error
}

func (uu *UserUpdater) SetAddUser(adduser func(name string, password string) error) *UserUpdater {
	uu.AddUser = adduser
	return uu
}

func (uu *UserUpdater) SetRemoveUser(removeuser func(name string) error) *UserUpdater {
	uu.RemoveUser = removeuser
	return uu
}

func (uu *UserUpdater) SetChangeSetting(changesetting func(name string, key string, value string) error) *UserUpdater {
	uu.ChangeSetting = changesetting
	return uu
}

func (uu *UserUpdater) SetStartService(startservice func(users []*User) error) *UserUpdater {
	uu.StartService = startservice
	return uu
}

func (uu *UserUpdater) SetStopService(stopservice func(users []*User) error) *UserUpdater {
	uu.StopService = stopservice
	return uu
}

type UserContainer struct {
	Users    map[string]*User
	Updaters map[Type]*UserUpdater
}

func (uc *UserContainer) AddUser(name string, password string, typ Type) error {
	if _, ok := uc.Users[name]; ok {
		return e("Existing user")
	}
	updater, ok := uc.Updaters[typ]
	if !ok {
		return e("Undefined type")
	}

	if err := updater.AddUser(name, password); err == nil {
		uc.Users[name] = &User{
			Name: name,
			Type: typ,
		}
		return nil
	} else {
		return err
	}
}

func (uc *UserContainer) RemoveUser(name string, route string) error {
	user, ok := uc.Users[name]
	if !ok {
		return e("Nonexistent user")
	}

	updater := uc.Updaters[user.Type]
	if err := updater.RemoveUser(name); err == nil {
		delete(uc.Users, name)
		return nil
	} else {
		return err
	}
}

func (uc *UserContainer) ChangeSetting(name string, key string, value string) error {
	user, ok := uc.Users[name]
	if !ok {
		return e("Nonexistent user")
	}

	updater := uc.Updaters[user.Type]
	return updater.ChangeSetting(name, key, value)
}

type User struct {
	Name string
	Type Type
}
