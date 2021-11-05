package keys

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	sdkerrors "github.com/irisnet/core-sdk-go/types/errors"
)

type keysClient struct {
	sdk.KeyManager
}

func NewClient(keyManager sdk.KeyManager) Client {
	return keysClient{keyManager}
}

func (k keysClient) Add(name, password string) (string, string, error) {
	address, mnemonic, err := k.Insert(name, password)
	if err != nil {
		return "", "", sdkerrors.Wrapf(ErrInsert, err.Error())
	}
	return address, mnemonic, nil
}

func (k keysClient) Recover(name, password, mnemonic string) (string, error) {
	address, err := k.KeyManager.Recover(name, password, mnemonic, "")
	if err != nil {
		return "", sdkerrors.Wrapf(ErrRecover, err.Error())
	}
	return address, nil
}

func (k keysClient) RecoverWithHDPath(name, password, mnemonic, hdPath string) (string, error) {
	address, err := k.KeyManager.Recover(name, password, mnemonic, hdPath)
	if err != nil {
		return "", sdkerrors.Wrapf(ErrRecover, err.Error())
	}
	return address, nil
}

func (k keysClient) Import(name, password, privKeyArmor string) (string, error) {
	address, err := k.KeyManager.Import(name, password, privKeyArmor)
	return address, err
}

func (k keysClient) Export(name, password string) (string, error) {
	keystore, err := k.KeyManager.Export(name, password)
	if err != nil {
		return "", sdkerrors.Wrapf(ErrExport, err.Error())
	}
	return keystore, nil
}

func (k keysClient) Delete(name, password string) error {
	err := k.KeyManager.Delete(name, password)
	if err != nil {
		return sdkerrors.Wrapf(ErrDelete, err.Error())
	}
	return nil
}

func (k keysClient) Show(name, password string) (string, error) {
	_, address, err := k.KeyManager.Find(name, password)
	if err != nil {
		return "", sdkerrors.Wrapf(ErrShow, err.Error())
	}
	return address.String(), nil
}
