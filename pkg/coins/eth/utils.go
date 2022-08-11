package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/NpoolPlatform/build-chain/pkg/coins"
	"github.com/NpoolPlatform/build-chain/pkg/config"
	coinsinfo_crud "github.com/NpoolPlatform/build-chain/pkg/crud/coinsinfo"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coinsinfo"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	proto "github.com/NpoolPlatform/message/npool/build-chain"
	"github.com/PuerkitoBio/goquery"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gocolly/colly"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

type DeployTx struct {
	// From     *common.Address `json:"to"`
	GasPrice *big.Int `json:"gasPrice"`
	Gas      uint64   `json:"gas"`
	Value    *big.Int `json:"value"`
	Data     []byte   `json:"data"`
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

func GetTOP100() error {
	endpoint := "https://etherscan.io"
	var err error
	// Instantiate default collector
	c := colly.NewCollector(
		colly.MaxDepth(2),
	)

	// get erc20 Tokens url and insert it`s info to database
	c.OnHTML("div#ContentPlaceHolder1_divresult table tbody tr", func(e *colly.HTMLElement) {
		_url := e.ChildAttr("div.media-body a", "href")
		_items := strings.Split(_url, "/")

		name := e.ChildText("div.media-body a")
		contract := _items[len(_items)-1]
		if !common.IsHexAddress(contract) {
			fmt.Printf("%v is invalid url\n", _url)
		}

		_, err := coinsinfo_crud.CheckContractAndCreate(context.Background(), &proto.CoinsInfo{
			Name:       name,
			ChainType:  coins.EthereumChain,
			TokenType:  coins.ERC20TOKEN,
			Contract:   contract,
			Similarity: 0,
			Remark:     "scrach it!",
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		err = c.Visit(fmt.Sprintf("https://etherscan.io/%v#code", _url))
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Second)
	})

	// get contract detail
	c.OnResponse(func(r *colly.Response) {
		// construct html element
		if !strings.Contains(strings.ToLower(r.Headers.Get("Content-Type")), "html") {
			return
		}
		doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(r.Body))
		if err != nil {
			return
		}
		e := colly.NewHTMLElementFromSelectionNode(r, doc.Selection, doc.Nodes[0], 0)

		// match cointype
		coinType := e.ChildText("div.media-body span.text-secondary")
		if coinType == "" {
			return
		}

		// filter erc20-proxy
		if len(e.DOM.Find("#nav_subtabs").Children().Nodes) != 3 {
			fmt.Printf("%v is erc20-proxy,skip!,%v\n", coinType, len(e.DOM.Find("#nav_subtabs").Children().Nodes))
			return
		}

		// get contract address
		contractAddr := e.ChildText("div a.text-truncate.d-block.mr-2")
		conds := cruder.NewConds()
		conds.WithCond(coinsinfo.FieldContract, cruder.EQ, contractAddr)
		infos, num, err := coinsinfo_crud.Rows(context.Background(), conds, 0, 1)
		if err != nil || num == 0 {
			fmt.Println(err)
			return
		}

		info := infos[0]
		contract := coins.Contract{
			Code:       e.ChildText("#editor"),
			ABI:        e.ChildText("pre#js-copytextarea2"),
			CreateCode: e.ChildText("div#verifiedbytecode2"),
			// CreateArgs: e.ChildText(), // TODO
		}
		data, err := json.Marshal(contract)
		if err != nil {
			fmt.Println(err)
			return
		}

		info.Data = data
		info.Remark += "get contract data!"
		info.Similarity = 10
		err = coinsinfo_crud.Update(context.Background(), info)
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err)
	})

	// erc20 top50
	err = c.Visit(endpoint + "/tokens?p=1")
	if err != nil {
		return err
	}
	// erc20 top51-100
	err = c.Visit(endpoint + "/tokens?p=2")
	if err != nil {
		return err
	}

	return nil
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
