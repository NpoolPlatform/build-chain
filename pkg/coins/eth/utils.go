package eth

import (
	"context"

	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/NpoolPlatform/build-chain/pkg/config"

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

func UnlockCoinbase(client *rpc.Client) error {
	coinbase, err := Coinbase(client)
	if err != nil {
		return err
	}
	var unlock bool
	err = client.Call(&unlock, "personal_unlockAccount", coinbase, "")
	if err != nil {
		return err
	}
	if !unlock {
		return fmt.Errorf("coinbase unlock failed")
	}
	return nil
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
	if auth == nil {
		auth, err = getAuth(client)
		if err != nil {
			return nil, err
		}
	}
	return auth, nil
}

func getAuth(client *rpc.Client) (*bind.TransactOpts, error) {
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
	if chainID.Cmp(big.NewInt(1337)) != 0 {
		return nil, fmt.Errorf("only support testnet")
	}

	// get private key
	auth, err := bind.NewTransactorWithChainID(wallet, "", chainID)
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
			endpoint = config.GetENV().Endpoint
		}

		if endpoint == "" {
			endpoint = "http://192.168.49.1:8545"
		}
		client, err = rpc.Dial(endpoint)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}
