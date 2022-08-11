package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	// npool "github.com/NpoolPlatform/message/npool/build-chain"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/NpoolPlatform/build-chain/pkg/coins/eth"
	res "github.com/NpoolPlatform/build-chain/resource"
)

func init() {
	rand.Seed(time.Now().Unix())
}
func _ethT() {
	client, err := rpc.Dial("http://192.168.49.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(eth.UnlockCoinbase(client))
	fmt.Println(eth.DeployContract(client, strings.TrimSpace(res.GetContract().ERC20Coins[0].ConstuctData)))
}

func main() {
	// _ethT()
	eth.GetTOP100()

	// fmt.Println(res.GetContract().ERC20Coins[0].ConstuctData)
	// okContract := []string{}
	// for _, v := range AllContract {
	// 	time.Sleep(time.Second)
	// 	_, err := eth.ERC20Faucet(common.HexToAddress(v), common.HexToAddress("0xBcE9e4a7aa5eF6998439618771D4754596045b76"), big.NewFloat(0.123))
	// 	fmt.Println(err)
	// 	if err == nil {
	// 		okContract = append(okContract, v)
	// 	}
	// }

	// for _, v := range okContract {
	// 	fmt.Println(v)
	// 	fmt.Println(eth.ERC20Balance(common.HexToAddress(v), common.HexToAddress("0xBcE9e4a7aa5eF6998439618771D4754596045b76")))
	// }

	// var err error
	// // Instantiate default collector
	// c := colly.NewCollector(
	// 	colly.MaxDepth(2),
	// )

	// // get erc20 abi and bin
	// c.OnHTML("#conntent", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.ChildText("editor"))
	// })

	// c.OnError(func(r *colly.Response, err error) {
	// 	fmt.Println(err)
	// })

	// c.OnResponse(func(r *colly.Response) {
	// 	if !strings.Contains(strings.ToLower(r.Headers.Get("Content-Type")), "html") {
	// 		return
	// 	}
	// 	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(r.Body))
	// 	if err != nil {
	// 		return
	// 	}
	// 	ele := colly.NewHTMLElementFromSelectionNode(r, doc.Selection, doc.Nodes[0], 0)
	// 	fmt.Println(ele.DOM.Find("#content").Nodes[0].Type)
	// })

	// // erc20 top50
	// err = c.Visit("https://etherscan.io/token/0xB8c77482e45F1F44dE1745F52C74426C631bDD52#code")
	// if err != nil {
	// 	fmt.Println(err)
	// }

}

var AllContract = []string{
	"0xE892C9f6899d1971d9E03f9B948bABbE8d7b0372",
	"0x45D8C413991797EcE0E069187eB36382E07F20A6",
	"0x19c876505A8d8f1F1b30C21a8740dE1Cecbe99B3",
	"0x5c8c9cE95c49Ad3BF3f27306Aa49CB60B0bCda00",
	"0x9661A8ff47Ff095EAac4a2929B36F5B8b87FD897",
	"0x44EB573F8BB31C854A717002d6B40c7E8AAd3F62",
	"0x15C2C2df4f5D7c5c0f15AE2E4dc63a52E1dBd5E9",
	"0x170D87ab0fDE0df72473d43614473dCe09Dd9170",
	"0x84Ab1Cbffd6b7B8F8d83319b291469b70A5E4103",
	"0x7A5A308A78E1C69cE720b46247a683C82cA407b7",
	"0xDa0624159f13E9B553dF394E9E3880bbC5D057CC",
	"0x60405849662E42b081505b4BC4958C571e081567",
	"0x867B5fdFDb3C9eC66300F9135ba603C242953695",
	"0x84E0836514a5D04E56E3474ad47d362F4B430C0d",
	"0xb7aa97251E059F3Ac3B39D8415FdcBD5Dba5F14F",
	"0xc5ea14f121d501f98427b55deaB85075488438Ad",
	"0x490B9E1fbE7b0173172EC583f6002c6BD38364fC",
	"0x175Ec2f6B8718C36BD78ACD41fBE375607f88f01",
	"0xE1355B888B700295c756151377C1bD44EEcD78F9",
	"0x5fd6b8153348B5BeACbA535920acC6638580b63E",
	"0xF4CF2D175F21474B6EBf651c06016527bC9eF750",
	"0x94F04597aD7f6f9C456a9B4890b9F1fbd7545f96",
	"0x638dC48a4c3f3D330fF5C0b1DEb570fD38236a46",
	"0x76203d6Ff017F978066A1cE35a30923E79d4301b",
	"0x1cc65A86444b12cFb4B09AA50A2e30274dd75889",
	"0x2022bc038fb82758c397BF20050CE7AD748ccD8E",
	"0x5668079a22E83416404BdcE35b046827878E6377",
	"0xf49Bf1c011Ea85856830852C012C0456335B0188",
	"0x50810Bf6213d10B1A8A0046fD1D7bAE430b0cfd0",
	"0x19f853aCf6880435fFf699874064299d53761319",
	"0x22A22977Da3eF7065cda296e01817d62f1Fb306d",
	"0x56c77035B32a3ab3e1D2A124946d33ED96338D6d",
	"0xEe3D2D7d6ABdeb6FEF4fCacCa920391E788809AB",
	"0xACCFe058C4F69FE37BbB7c98D3331Fe53d7487DB",
	"0xf1012005C7d59547FFE08c6e49CDA125A21270Eb",
	"0x9B139697bEAB0e153CCff2c6fd657e7bF5476315",
	"0x53f92a552ea35DE9E664Bcfc22a4b79EF3fD1B1b",
	"0x76fFed500c66f37032ACdaC4C209b4beEa3B8899",
	"0x54963eab135833A948b81Ebec9FACcF13559d911",
	"0xAc4D7d9F7593dD19777d21F13721e4e8B440C5d2",
	"0x742c46C7708EF95a99e4DAc54577979E8B0f2B44",
	"0xaa64f539A049aeaf7709A8baFbd472fD5D375e31",
	"0xBb49f33C85DA8402FEBA4a7A62088Cab0F004407",
	"0x2FDaf02543F7fFe7f04c5a332e6009e38dbd0D8F",
	"0xf40CE9B0F4281Fd2a94741E0B217455507b0A57D",
	"0xeBF9e2122630aB1859a111ED01018De0D865018e",
	"0x8fEb68E08C701a93b6dc8a319D1433F946850BcB",
	"0x7AB3ae57BaBAd73253a08e4811ce8779BDa2e92F",
	"0x0C8578Aeb6a552fE7Fc3dbF4d9E0B908A5Fd3f82",
	"0x8017DE8B9302FC5FE92a8c14e8bb33C75321E4A5",
	"0x47bD3A9c96c19d1a1043189f1Ccc8bC2DFfA6a1E",
	"0x41E2c72EbC589eFeB82ca5D6083739930e5b84aB",
	"0xb2b25F4DAa98bf5C7D84bC9Ee388099BA42A9606",
	"0xE7237FcF3291155304Bdc6f89A6cB03f01085891",
	"0x3ee9A363c0AAE2e29FbA48eE1E024a9B478fB164",
	"0x93a8C0C896A0933dB8D1b5E23f5E9deC983D6bB8",
	"0x8F7Dd26f664418fEB452594e494f67A3e5713c87",
	"0x9b8F7dC62faDe8778555b5EB582F8b55429C4E93",
	"0x4aD41010fa8eDC20d8700D519BeBfa4CcBB16b7C",
	"0x8e5b0C68c3aE37f206a24Db14725b75E0C28D76f",
	"0x77792e6E29F39AacCcfF27F3bD25689AB725710E",
	"0x119c5eBbD37c02445b0F369725834CceA9454b30",
	"0x1165eAA0fD3DbB0fC81d8a63b41C45C41221b148",
	"0x27dC805268b19C3FAadb063eB17f2bC1e7558fEe",
	"0x79dE4C1bB7C91aEB2DDF1D9c64FC80796f379490",
	"0x200699A76849d73dDcD2E5c1EE3327d30515ADc7",
	"0x35ca8d01DCbd8917ea5035C1CaE8590426c628Bc",
	"0x871D7D214FfEB464c0Cb4CF35F94Dc839CBFa9FD",
	"0xCF2F09777D8F1B54d17861a81679Db8eD1203FDd",
	"0x8b4F9eaB4340E33c41bcd7915b91B1628b012538",
	"0xcC169dDb11D488ee414c249fc89F67c805AeA151",
	"0xA3D65Cb311cd9eE7AC24D63827f171cEb5DF99dF",
	"0xe6C334bb5AB18943c0851DCac1A215d46C252151",
	"0x4bb7fBE81490EA3C751aEdB0BB1b3Df08194888c",
	"0x8D4FaDaF33d420bDD6A83A56531cdAf5E5A19DCa",
	"0xa305a73386cC4336346e6e059D74f79b2b12db93",
	"0xBCCE3De8eba3FbbBC812D6e5C1bA7bC8Dd71bD44",
}
