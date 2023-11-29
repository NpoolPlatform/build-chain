package eth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/NpoolPlatform/build-chain/pkg/coins"
	proto "github.com/NpoolPlatform/message/npool/build-chain/v1"
	"github.com/PuerkitoBio/goquery"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gocarina/gocsv"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/jedib0t/go-pretty/v6/table"
)

type erc20row struct {
	contract string
	num      int
}

type CrawlTaskInfo struct {
	Offset    int
	Limit     int
	Contract  string
	ChainType string
	TokenType string
}

var CrawlInterval = time.Second * 2

func CrawlERC20Rows(offset, limit int) ([]string, error) {
	if offset <= 0 || limit < 1 {
		return nil, errors.New("offset or limit is invalid, offset begin with 1, limit must GE 1")
	}
	var err error
	erc20rows := []erc20row{}
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	c.OnHTML("div#ContentPlaceHolder1_tblErc20Tokens table tbody tr", func(e *colly.HTMLElement) {
		_url := e.ChildAttr("td a", "href")
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
		fmt.Println(err)
	})

	itemsPrePage := 50
	startPage := (offset-1)/itemsPrePage + 1
	endPage := (offset-1+limit-1)/itemsPrePage + 1
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

func CrawlContractInfo(contractAddr string) (*proto.TokenInfo, error) {
	url := fmt.Sprintf("https://etherscan.io/token/%v#code", contractAddr)
	var err error
	var info *proto.TokenInfo
	var contract *coins.Contract
	// Instantiate default collector
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

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

		// remove redundant child node,then match tokenName
		e.DOM.Find("#content section.container-xxl span.fs-base.fw-medium span").Remove()
		tokenName := e.ChildText("#content section.container-xxl span.fs-base.fw-medium")
		if tokenName == "" {
			return
		}

		// match contract
		mainContractAddr := e.ChildAttr("#ContentPlaceHolder1_divSummary a.js-clipboard.link-secondary", "data-clipboard-text")
		if mainContractAddr == "" {
			return
		}

		// match decimal
		contractDecimal := e.ChildText("#ContentPlaceHolder1_divSummary h4 b")
		if contractDecimal == "" {
			return
		}

		// match unit
		contractUnit := e.ChildAttr("#ContentPlaceHolder1_hdnSymbol", "value")
		if contractUnit == "" {
			return
		}

		// filter erc20-proxy
		baseErc20NodeNum := 3
		if len(e.DOM.Find("#nav_subtabs").Children().Nodes) != baseErc20NodeNum {
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

		info = &proto.TokenInfo{
			Name:             tokenName,
			Unit:             contractUnit,
			Decimal:          contractDecimal,
			ChainType:        "ethereum",
			TokenType:        "erc20",
			OfficialContract: mainContractAddr,
			Data:             _contract,
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err)
	})

	err = c.Visit(url)
	if err != nil {
		return nil, err
	}

	if info == nil {
		return nil, errors.New("can not crawl contract info")
	}
	return info, nil
}

func Crawl(info *CrawlTaskInfo) {
	addresses, _ := CrawlERC20Rows(info.Offset, info.Limit)
	if info.Contract != "" {
		addresses = append(addresses, info.Contract)
	}

	successData := []*proto.TokenInfo{}
	var wg sync.WaitGroup
	var lock sync.Mutex
	for _, v := range addresses {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			info, err := CrawlContractInfo(addr)
			if err == nil {
				fmt.Printf("success get token name: %v, official contract :%v\n", info.Name, info.OfficialContract)
				lock.Lock()
				successData = append(successData, info)
				lock.Unlock()
				return
			}
			fmt.Printf("failed to get token info, official contract :%v, err :%v\n", v, err)
		}(v)
		// prevent to be baned
		time.Sleep(CrawlInterval)
	}
	wg.Wait()
	fmt.Println("success crawl table:")
	printTable(successData)
	err := printToCSV(successData, "./a.csv")
	if err != nil {
		fmt.Printf("failed to write token info to file, err :%v\n", err)
	}
	fmt.Printf("crawl end,found %v,success %v\n", len(addresses), len(successData))
}

func printToCSV(infos []*proto.TokenInfo, filepath string) error {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	return gocsv.MarshalFile(infos, file)

	// buf := bufio.NewWriter(file)
	// _, err = buf.WriteString("Name,Unit,Decimal,ChainType,TokenType,OfficialContract,PrivateContract,Remark,Data\n")
	// if err != nil {
	// 	return err
	// }
	// for _, v := range infos {
	// 	_, err = buf.WriteString(
	// 		fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v\n",
	// 			v.Name,
	// 			v.Unit,
	// 			v.Decimal,
	// 			v.ChainType,
	// 			v.TokenType,
	// 			v.OfficialContract,
	// 			v.PrivateContract,
	// 			v.Remark,
	// 			string(v.Data),
	// 		),
	// 	)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// return buf.Flush()
}

func printTable(infos []*proto.TokenInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Unit", "Decimal", "Message"})
	tableRows := []table.Row{}
	for _, v := range infos {
		tableRows = append(tableRows, []interface{}{v.Name, v.Unit, v.Decimal, v.Remark})
	}
	t.AppendRows(tableRows)
	t.Render()
}
