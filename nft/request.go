package nft

import (
	"github.com/cosmos/cosmos-sdk/types"
	"strings"
)

type CreateClassRequest struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Schema           string `json:"schema"`
	Symbol           string `json:"symbol"`
	Description      string `json:"description"`
	Uri              string `json:"uri"`
	UriHash          string `json:"uri_hash"`
	Data             string `json:"data"`
	MintRestricted   bool   `json:"mint_restricted"`
	UpdateRestricted bool   `json:"update_restricted"`
}

func NewCreateClassRequest(
	id string,
	name string,
	schema string,
	symbol string,
	description string,
	uri string,
	uriHash string,
	data string,
	mintRestricted bool,
	updateRestricted bool,
) CreateClassRequest {
	req := CreateClassRequest{}
	req.ID = strings.TrimSpace(id)
	req.Name = strings.TrimSpace(name)
	req.Schema = strings.TrimSpace(schema)
	req.Symbol = strings.TrimSpace(symbol)
	req.Description = strings.TrimSpace(description)
	req.Uri = strings.TrimSpace(uri)
	req.UriHash = strings.TrimSpace(uriHash)
	req.Data = strings.TrimSpace(data)
	req.MintRestricted = mintRestricted
	req.UpdateRestricted = updateRestricted
	return req
}

func (req *CreateClassRequest) Validate() error {
	if err := ValidateDenomID(req.ID); err != nil {
		return err
	}
	return ValidateKeywords(req.ID)
}

func (req *CreateClassRequest) ToMsg() (*MsgIssueDenom, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &MsgIssueDenom{
		Id:          req.ID,
		Name:        req.Name,
		Schema:      req.Schema,
		Symbol:      req.Symbol,
		Description: req.Description,
		Uri:         req.Uri,
		UriHash:     req.UriHash,
		Data:        req.Data,
	}, nil
}

type TransferClassRequest struct {
	ID        string `json:"id"`
	Recipient string `json:"recipient"`
}

func NewTransferClassRequest(
	id string,
	recipient string,
) TransferClassRequest {
	req := TransferClassRequest{}
	req.ID = strings.TrimSpace(id)
	req.Recipient = strings.TrimSpace(recipient)
	return req
}

func (req *TransferClassRequest) Validate() error {
	if err := ValidateDenomID(req.ID); err != nil {
		return err
	}
	if _, err := types.AccAddressFromBech32(req.Recipient); err != nil {
		return err
	}
	return ValidateKeywords(req.ID)
}

func (req *TransferClassRequest) ToMsg() (*MsgTransferDenom, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &MsgTransferDenom{
		Id:        req.ID,
		Recipient: req.Recipient,
	}, nil
}

type MintNFTRequest struct {
	ClassId   string `json:"class_id"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	URI       string `json:"uri"`
	Data      string `json:"data"`
	Recipient string `json:"recipient"`
	URIHash   string `json:"uri_hash"`
}

func NewMintNFTRequest(
	classId string,
	id string,
	name string,
	uri string,
	data string,
	recipient string,
	uriHash string,
) MintNFTRequest {
	req := MintNFTRequest{}
	req.ClassId = strings.TrimSpace(classId)
	req.ID = strings.TrimSpace(id)
	req.Name = strings.TrimSpace(name)
	req.URI = strings.TrimSpace(uri)
	req.Data = strings.TrimSpace(data)
	req.URIHash = strings.TrimSpace(uriHash)
	req.Recipient = strings.TrimSpace(recipient)
	return req
}

func (req *MintNFTRequest) Validate() error {
	if err := ValidateDenomID(req.ClassId); err != nil {
		return err
	}
	if err := ValidateTokenID(req.ID); err != nil {
		return err
	}
	if len(req.Recipient) > 0 {
		if _, err := types.ValAddressFromBech32(req.Recipient); err != nil {
			return err
		}
	}

	return ValidateKeywords(req.ID)
}

func (req *MintNFTRequest) ToMsg() (*MsgMintNFT, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &MsgMintNFT{
		DenomId:   req.ClassId,
		Id:        req.ID,
		Name:      req.Name,
		URI:       req.URI,
		Data:      req.Data,
		Recipient: req.Recipient,
		UriHash:   req.URIHash,
	}, nil
}

type EditNFTRequest struct {
	ClassId string `json:"class_id"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	URI     string `json:"uri"`
	Data    string `json:"data"`
	URIHash string `json:"uri_hash"`
}

func NewEditNFTRequest(
	classId string,
	id string,
	name string,
	uri string,
	data string,
	uriHash string,
) EditNFTRequest {
	req := EditNFTRequest{}
	req.ClassId = strings.TrimSpace(classId)
	req.ID = strings.TrimSpace(id)
	req.Name = strings.TrimSpace(name)
	req.URI = strings.TrimSpace(uri)
	req.Data = strings.TrimSpace(data)
	req.URIHash = strings.TrimSpace(uriHash)
	return req
}

func (req *EditNFTRequest) Validate() error {
	if err := ValidateDenomID(req.ClassId); err != nil {
		return err
	}
	if err := ValidateTokenID(req.ID); err != nil {
		return err
	}
	return ValidateKeywords(req.ID)
}

func (req *EditNFTRequest) ToMsg() (*MsgEditNFT, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &MsgEditNFT{
		DenomId: req.ClassId,
		Id:      req.ID,
		Name:    req.Name,
		URI:     req.URI,
		Data:    req.Data,
		UriHash: req.URIHash,
	}, nil
}

type TransferNFTRequest struct {
	ClassId   string `json:"class_id"`
	ID        string `json:"id"`
	URI       string `json:"uri"`
	Data      string `json:"data"`
	Name      string `json:"name"`
	Recipient string `json:"recipient"`
	URIHash   string `json:"uri_hash"`
}

func NewTransferNFTRequest(
	classId string,
	id string,
	uri string,
	data string,
	name string,
	recipient string,
	uriHash string,
) TransferNFTRequest {
	req := TransferNFTRequest{}
	req.ClassId = strings.TrimSpace(classId)
	req.ID = strings.TrimSpace(id)
	req.URI = strings.TrimSpace(uri)
	req.Data = strings.TrimSpace(data)
	req.Name = strings.TrimSpace(name)
	req.Recipient = strings.TrimSpace(recipient)
	req.URIHash = strings.TrimSpace(uriHash)
	return req
}

func (req *TransferNFTRequest) Validate() error {
	if err := ValidateDenomID(req.ClassId); err != nil {
		return err
	}
	if err := ValidateTokenID(req.ID); err != nil {
		return err
	}
	if _, err := types.ValAddressFromBech32(req.Recipient); err != nil {
		return err
	}
	return ValidateKeywords(req.ID)
}

func (req *TransferNFTRequest) ToMsg() (*MsgTransferNFT, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &MsgTransferNFT{
		DenomId:   req.ClassId,
		Id:        req.ID,
		URI:       req.URI,
		Data:      req.Data,
		Name:      req.Name,
		Recipient: req.Recipient,
		UriHash:   req.URIHash,
	}, nil
}

type BurnNFTRequest struct {
	ClassId string `json:"class_id"`
	ID      string `json:"id"`
}

func NewBurnNFTRequest(classId string, id string) BurnNFTRequest {
	req := BurnNFTRequest{}
	req.ClassId = strings.TrimSpace(classId)
	req.ID = strings.TrimSpace(id)
	return req
}

func (req *BurnNFTRequest) Validate() error {
	if err := ValidateDenomID(req.ClassId); err != nil {
		return err
	}
	if err := ValidateTokenID(req.ID); err != nil {
		return err
	}
	return ValidateKeywords(req.ID)
}

func (req *BurnNFTRequest) ToMsg() (*MsgBurnNFT, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &MsgBurnNFT{
		DenomId: req.ClassId,
		Id:      req.ID,
	}, nil
}

type PaginationRequest struct {
	NextKey string `json:"next_key"`
	Offset  uint64 `json:"offset"`
	Limit   uint64 `json:"limit"`
}
