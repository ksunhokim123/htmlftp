package htmlftp

type Service struct {
	UserContainer *UserContainer
	KeyContainer  *KeyContainer
}

func (sv *Service) Start() {
	for _, val := range sv.UserContainer.Updaters {
		val.StartService([]*User{}) // TODO
	}
}

func NewService(userfile string, keyfile string) *Service {
	sv := &Service{
		UserContainer: &UserContainer{
			Users:    make(map[string]*User),
			Updaters: make(map[Type]*UserUpdater),
		},
		KeyContainer: new(KeyContainer),
	}
	sv.UserContainer.Updaters[FTP] = NewFTPUpdater()
	return sv
}
