package htmlftp

type KeyContainer []string

func (kc *KeyContainer) Generate() string {
	return "unimplemented"
}

func (kc *KeyContainer) Exist(key string) int {
	return 0
}

func (kc *KeyContainer) Remove(key string) {

}
