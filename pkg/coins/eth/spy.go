package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/NpoolPlatform/build-chain/pkg/coins"
	"github.com/NpoolPlatform/build-chain/pkg/coins/eth/erc20"
	coinsinfo_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/coinsinfo"
	deployedcoin_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/deployedcoin"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coinsinfo"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	proto "github.com/NpoolPlatform/message/npool/build-chain"
	"github.com/PuerkitoBio/goquery"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gocolly/colly"
	"github.com/google/uuid"
)

const (
	SimInitState    = 0    //it`s just crewl from explore
	SimOnlyContract = 10   // not with proxy
	SimDeployed     = 20   // can be deployed
	SimBaseERC20    = 1000 // simplest erc20 without complexy argments and proxy
)

func DeployBaseErc20Spy(ctx context.Context, client *rpc.Client) error {
	contract := &coins.Contract{}
	conds := cruder.NewConds()
	conds.WithCond(coinsinfo.FieldSimilarity, cruder.EQ, SimOnlyContract)

	infos, _, err := coinsinfo_crud.All(ctx, conds)
	if err != nil {
		return err
	}
	for _, info := range infos {
		if info.Data == nil || len(info.Data) == 0 {
			logger.Sugar().Error(err)
			continue
		}
		err = json.Unmarshal(info.Data, contract)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		contractAddr, err := DeployContract(client, contract.CreateCode)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		info.Similarity = 20
		_, err = coinsinfo_crud.Update(ctx, info)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		_, err = deployedcoin_crud.Create(ctx, &proto.DeployedCoin{
			CoinID:   info.ID,
			Contract: contractAddr.String(),
		})
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		fmt.Println("ctract deployed ", info.Name)
	}
	return nil
}

//nolint
func GetErc20TOP100() error {
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
			Similarity: SimInitState,
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
		infos, _, err := coinsinfo_crud.Rows(context.Background(), conds, 0, 1)
		if err != nil || len(infos) == 0 {
			fmt.Println(err)
			return
		}

		info := infos[0]
		contract := coins.Contract{
			Code:       e.ChildText("#editor"),
			ABI:        e.ChildText("pre#js-copytextarea2"),
			CreateCode: e.ChildText("div#verifiedbytecode2"),
		}

		e.DOM.Find("div#dividcode>div.mb-4").Each(func(i int, s *goquery.Selection) {
			text := strings.Trim(s.Find("pre.wordwrap").Text(), " ")
			target := strings.Trim(s.Find("h4.card-header-title").Text(), " ")
			switch {
			case strings.Contains(target, "Constructor Arguments"):
				contract.CreateArgs = text
			case strings.Contains(target, "Swarm Source"):
				contract.SwarmSource = text
			case strings.Contains(target, "Deployed ByteCode Sourcemap"):
				contract.SourceMap = text
			}
		})
		data, err := json.Marshal(contract)
		if err != nil {
			fmt.Println(err)
			return
		}

		info.Data = data
		info.Remark += "get contract data!"
		info.Similarity = SimOnlyContract
		_, err = coinsinfo_crud.Update(context.Background(), info)
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

func TransferSpy() {
	client, err := rpc.Dial("http://192.168.49.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	auth, err := GetAuth(client)
	if err != nil {
		log.Fatal(err)
	}

	infos, _, err := deployedcoin_crud.All(ctx, cruder.NewConds())
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		token, err := erc20.NewErc20(common.HexToAddress(info.Contract), ethclient.NewClient(client))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		_, err = token.Transfer(auth, common.HexToAddress("0xBcE9e4a7aa5eF6998439618771D4754596045b76"), big.NewInt(123456789))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
	}
}

func BalanceSpy() {
	client, err := rpc.Dial("http://192.168.49.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	infos, _, err := deployedcoin_crud.All(ctx, cruder.NewConds())
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		token, err := erc20.NewErc20(common.HexToAddress(info.Contract), ethclient.NewClient(client))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
		balance, err := token.BalanceOf(nil, common.HexToAddress("0xBcE9e4a7aa5eF6998439618771D4754596045b76"))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		if balance.Int64() == 0 {
			continue
		}

		coin, err := coinsinfo_crud.Row(ctx, uuid.MustParse(info.CoinID))
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}

		coin.Similarity = SimBaseERC20
		_, err = coinsinfo_crud.Update(ctx, coin)
		if err != nil {
			logger.Sugar().Error(err)
			continue
		}
	}
}
