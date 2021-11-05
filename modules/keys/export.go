package keys

type Client interface {
	Add(name, password string) (address string, mnemonic string, err error)
	Recover(name, password, mnemonic string) (address string, err error)
	RecoverWithHDPath(name, password, mnemonic, hdPath string) (address string, err error)
	Import(name, password, privKeyArmor string) (address string, err error)
	Export(name, password string) (privKeyArmor string, err error)
	Delete(name, password string) error
	Show(name, password string) (string, error)
}
