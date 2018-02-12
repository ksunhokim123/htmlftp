package mouse

//TODO seperate
func NewFTPUpdater() *UserUpdater {
	updater := &UserUpdater{}
	updater = updater.SetAddUser(addUser)
	updater = updater.SetRemoveUser(removeUser)
	updater = updater.SetStartService(startService)
	return updater
}

func addUser(name string, password string) error {
	return nil
}

func removeUser(name string) error {
	return nil
}

func startService(users []*User) error {
	go FTPRun()
	return nil
}
