package eth

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"io"

	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/NpoolPlatform/build-chain/pkg/coins"
	"github.com/NpoolPlatform/build-chain/pkg/config"
	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/rpc"
)

func Coinbase(client *rpc.Client) (string, error) {
	var coinbase string
	err := client.Call(&coinbase, "eth_coinbase")
	return coinbase, err
}

func DeployContract(client *rpc.Client, bytecode string) (common.Address, error) {
	ctx := context.Background()
	backend := ethclient.NewClient(client)
	var contractAddr common.Address
	auth, err := GetAuth(client)
	if err != nil {
		return contractAddr, err
	}
	// TODO: ethscan 的contractcode中是有构造参数的 直接发送就可以了
	nonce, err := backend.PendingNonceAt(ctx, auth.From)
	if err != nil {
		return contractAddr, err
	}
	gasPrice, err := backend.SuggestGasPrice(ctx)
	if err != nil {
		return contractAddr, err
	}
	gas, err := backend.EstimateGas(ctx, ethereum.CallMsg{
		From:     auth.From,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     common.Hex2Bytes(bytecode),
	})
	if err != nil {
		return contractAddr, err
	}
	legacyTx := &types.LegacyTx{
		Nonce:    nonce,
		Gas:      gas,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     common.Hex2Bytes(bytecode),
	}
	tx := types.NewTx(legacyTx)
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return contractAddr, err
	}

	err = backend.SendTransaction(ctx, signedTx)

	if err != nil {
		return contractAddr, err
	}
	contractAddr = crypto.CreateAddress(auth.From, signedTx.Nonce())

	return contractAddr, nil
}

var auth *bind.TransactOpts

func GetAuth(client *rpc.Client) (*bind.TransactOpts, error) {
	var err error
	if auth != nil {
		return auth, nil
	}
	iKey := ""
	if config.GetENV() != nil {
		iKey = config.GetENV().InverstorKey
	}

	if iKey != "" {
		auth, err = getAuthFromPri(client, iKey)
	} else {
		auth, err = getAuthFromCoinbase(client)
	}

	if err != nil {
		return nil, err
	}

	return auth, nil
}

func getAuthFromPri(client *rpc.Client, priKey string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		return nil, fmt.Errorf("parse key err: %v", err)
	}

	chainID, err := ethclient.NewClient(client).NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get eth chainID err: %v", err)
	}

	// build auth
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("conn eth client err: %v", err)
	}
	return auth, nil
}

func getAuthFromCoinbase(client *rpc.Client) (*bind.TransactOpts, error) {
	coinbase, err := Coinbase(client)
	if err != nil {
		return nil, err
	}
	// TODO default wallet(maybe no wallet)
	type EthAccount struct {
		Address string `json:"address"`
		URL     string `json:"url"`
	}
	type EthWallet struct {
		Accounts []EthAccount `json:"accounts"`
		Status   string       `json:"status"`
		URL      string       `json:"url"`
	}
	wallets := []EthWallet{}
	err = client.Call(&wallets, "personal_listWallets")
	if err != nil {
		return nil, err
	}

	defaultWallet := wallets[0].Accounts[0]
	if coinbase != defaultWallet.Address {
		return nil, fmt.Errorf("compare geth wallet can only use coinbase wallet")
	}

	walletPath := strings.Split(defaultWallet.URL, "//")
	wallet, err := os.Open(walletPath[1])
	if err != nil {
		return nil, fmt.Errorf("open wallet err: %v", err)
	}
	defer wallet.Close()

	chainID, err := ethclient.NewClient(client).NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get eth chainID err: %v", err)
	}

	// only support testnet
	if chainID.Cmp(big.NewInt(coins.TestNetChainID)) != 0 {
		return nil, fmt.Errorf("only support testnet")
	}

	// getPrivateKey
	json, err := io.ReadAll(wallet)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(json, "")
	if err != nil {
		return nil, err
	}

	// parse privateKey to hex string
	// privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
	// privateKeyBytesHex := make([]byte, len(privateKeyBytes)*2)
	// hex.Encode(privateKeyBytesHex, privateKeyBytes)
	// fmt.Println("privateKey:", string(privateKeyBytesHex))

	// build auth
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("conn eth client err: %v", err)
	}
	return auth, nil
}

var client *rpc.Client

func Client() (*rpc.Client, error) {
	var err error
	if client == nil {
		endpoint := ""
		if config.GetENV() != nil {
			endpoint = config.GetENV().EthEndpoint
		}

		if endpoint == "" {
			endpoint = "http://127.0.0.1:8545"
		}
		client, err = rpc.Dial(endpoint)

		if err != nil {
			return nil, err
		}
	}
	return client, nil
}

func GenPriAndPubKey() (pri *ecdsa.PrivateKey, pub common.Address, err error) {
	pri, err = crypto.GenerateKey()
	if err != nil {
		return
	}

	privateKeyBytes := crypto.FromECDSA(pri)
	lenRate := 2
	priKeyBytesHexLen := len(privateKeyBytes) * lenRate
	privateKeyBytesHex := make([]byte, priKeyBytesHexLen)
	hex.Encode(privateKeyBytesHex, privateKeyBytes)

	// privateKey.PublicKey
	publicKey := pri.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("create account error casting public key to ECDSA")
		return
	}

	pub = crypto.PubkeyToAddress(*publicKeyECDSA)
	return pri, pub, nil
}
