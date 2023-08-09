package integration_test

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/irisnet/core-sdk-go/common/chain/irishub"

	"github.com/irisnet/core-sdk-go/common/log"
	"github.com/irisnet/core-sdk-go/crypto/keyring"
	"github.com/irisnet/core-sdk-go/store"
	sdktypes "github.com/irisnet/core-sdk-go/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irismod-sdk-go/client"
	"github.com/stretchr/testify/suite"
)

const (
	nodeURI  = "tcp://localhost:26657"
	grpcAddr = "localhost:9090"
	chainID  = "testnet-1"
	charset  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	addr     = "iaa1w9lvhwlvkwqvg08q84n2k4nn896u9pqx93velx"
	password = "12345678"
	node0    = "node0"
)

type ClientTestSuite struct {
	suite.Suite

	client.Client
	r            *rand.Rand
	rootAccount  MockAccount
	randAccounts []MockAccount
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (s *ClientTestSuite) SetupSuite() {
	sdkAddressCfg := irishub.NewConfig()
	options := []sdktypes.Option{
		sdktypes.KeyDAOOption(store.NewMemory(nil)),
		sdktypes.TimeoutOption(10),
		sdktypes.TokenManagerOption(irishub.TokenManager{}),
		sdktypes.KeyManagerOption(keyring.NewKeyManager()),
		sdktypes.Bech32AddressPrefixOption(sdkAddressCfg),
		sdktypes.BIP44PathOption(""),
		sdktypes.FeeOption(cosmostypes.NewDecCoins(
			cosmostypes.NewDecCoin("uiris", cosmostypes.NewInt(400000)))),
	}
	cfg, err := sdktypes.NewClientConfig(nodeURI, grpcAddr, chainID, options...)
	if err != nil {
		panic(err)
	}

	cli, err := client.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	s.Client = cli
	s.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	s.rootAccount = MockAccount{
		Name:     node0,
		Password: password,
		Address:  cosmostypes.MustAccAddressFromBech32(addr),
	}
	s.SetLogger(log.NewLogger(log.Config{
		Format: log.FormatJSON,
		Level:  log.DebugLevel,
	}))
	s.initAccount()
}

// Account return a test account
func (s *ClientTestSuite) Account() MockAccount {
	return s.rootAccount
}

// RandStringOfLength return a random string
func (s *ClientTestSuite) RandStringOfLength(l int) string {
	var result []byte
	bytes := []byte(charset)
	for i := 0; i < l; i++ {
		result = append(result, bytes[s.r.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandAccount return a random test account
func (s *ClientTestSuite) GetRandAccount() MockAccount {
	return s.randAccounts[s.r.Intn(len(s.randAccounts))]
}

func (s *ClientTestSuite) initAccount() {
	_, err := s.Import(
		s.Account().Name,
		s.Account().Password,
		string(getPrivKeyArmor()),
	)
	if err != nil {
		panic(err)
	}

	//var receipts bank.Receipts
	for i := 0; i < 5; i++ {
		name := s.RandStringOfLength(10)
		pwd := s.RandStringOfLength(16)
		address, _, err := s.Add(name, password)
		if err != nil {
			panic("generate test account failed")
		}

		s.randAccounts = append(s.randAccounts, MockAccount{
			Name:     name,
			Password: pwd,
			Address:  cosmostypes.MustAccAddressFromBech32(address),
		})
	}
}

func getPrivKeyArmor() []byte {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Dir(path)
	path = filepath.Join(path, "integration_test/scripts/priv.key")
	bz, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bz
}

// MockAccount define mock account for test
type MockAccount struct {
	Name, Password string
	Address        cosmostypes.AccAddress
}
