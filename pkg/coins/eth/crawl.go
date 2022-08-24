package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	proto "github.com/NpoolPlatform/message/npool/build-chain"

	bc_client "github.com/NpoolPlatform/build-chain/pkg/client/v1"
	"github.com/NpoolPlatform/build-chain/pkg/coins"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/PuerkitoBio/goquery"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gocolly/colly"
)

type erc20row struct {
	contract string
	num      int
}

type CrawlTaskInfo struct {
	Host      string
	Offset    int
	Limit     int
	Force     bool
	Contract  string
	ChainType string
	TokenType string
}

var CrawlInterval = time.Millisecond * 300

func CrawlERC20Rows(offset, limit int) ([]string, error) {
	if offset <= 0 || limit < 1 {
		return nil, errors.New("offset or limit is invalid, offset begin with 1, limit must GE 1")
	}
	var err error
	erc20rows := []erc20row{}
	c := colly.NewCollector()

	c.OnHTML("div#ContentPlaceHolder1_divresult table tbody tr", func(e *colly.HTMLElement) {
		_url := e.ChildAttr("div.media-body a", "href")
		_items := strings.Split(_url, "/")
		contract := _items[len(_items)-1]
		if !common.IsHexAddress(contract) {
			return
		}

		_num := e.ChildText("td:nth-child(1)")
		num, err := strconv.Atoi(_num)
		if err != nil {
			return
		}

		erc20rows = append(erc20rows, erc20row{contract: contract, num: num})
	})

	c.OnError(func(r *colly.Response, err error) {
		logger.Sugar().Error(err)
	})

	startPage := (offset-1)/50 + 1
	endPage := (offset+limit-2)/50 + 1
	for i := startPage; i <= endPage; i++ {
		err = c.Visit(fmt.Sprintf("https://etherscan.io/tokens?p=%v", i))
		if err != nil {
			return nil, err
		}
	}

	contracts := make([]string, limit)

	for _, v := range erc20rows {
		index := v.num - offset
		if index >= 0 && index < len(contracts) {
			contracts[index] = v.contract
		}
	}

	return contracts, nil
}

func CrawlContractInfo(contractAddr string) (*proto.CoinInfo, error) {
	url := fmt.Sprintf("https://etherscan.io/token/%v#code", contractAddr)
	var err error
	var coininfo *proto.CoinInfo
	var contract *coins.Contract
	// Instantiate default collector
	c := colly.NewCollector()

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

		coinName := e.ChildText("#content > div.container.py-3 > div > div.mb-3.mb-lg-0 > h1 > div > span")
		if coinName == "" {
			return
		}

		contractAddr := e.ChildText("#ContentPlaceHolder1_divSummary div.row.align-items-center a.text-truncate.d-block.mr-2")
		if contractAddr == "" {
			return
		}

		// filter erc20-proxy
		if len(e.DOM.Find("#nav_subtabs").Children().Nodes) != 3 {
			return
		}

		contract = &coins.Contract{
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

		_contract, err := json.Marshal(contract)
		if err != nil {
			return
		}

		coininfo = &proto.CoinInfo{
			Name:             coinName,
			ChainType:        "ethereum",
			TokenType:        "erc20",
			OfficialContract: contractAddr,
			Data:             _contract,
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		logger.Sugar().Error(err)
	})

	err = c.Visit(url)
	if err != nil {
		return nil, err
	}

	if coininfo == nil {
		return nil, errors.New("can not crawl contract info")
	}
	return coininfo, nil
}

func Task(info *CrawlTaskInfo) {
	bcConn, err := bc_client.NewClientConn(info.Host)
	ctx := context.Background()
	if err != nil {
		log.Fatal(err)
	}

	addresses, _ := CrawlERC20Rows(info.Offset, info.Limit)
	if info.Contract != "" {
		addresses = append(addresses, info.Contract)
	}

	fmt.Println("start deploy contract,please wait for end !")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>")
	successNum := 0
	var wg sync.WaitGroup
	for _, v := range addresses {
		// prevent to be baned
		time.Sleep(CrawlInterval)

		wg.Add(1)
		go func(contractAddr string) {
			defer wg.Done()
			coininfo, err := CrawlContractInfo(contractAddr)
			if err != nil {
				fmt.Printf("faild: address %v not support, %v\n", contractAddr, err)
				return
			}
			resp, err := bcConn.CreateCoinInfo(ctx, &proto.CreateCoinInfoRequest{
				Force: info.Force,
				Info:  coininfo,
			})
			if err != nil {
				fmt.Printf("faild: token %v, %v\n", coininfo.Name, err)

				return
			}
			successNum++
			fmt.Printf("success: token %v, %v\n", coininfo.Name, resp.Msg)
		}(v)
	}
	wg.Wait()
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>")
	fmt.Printf("deploy end,found %v,success %v\n", len(addresses), successNum)
}
