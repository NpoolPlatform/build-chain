package main

import (
	"database/sql"
	"time"

	// proto "github.com/NpoolPlatform/message/npool/build-chain"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger.Init(logger.DebugLevel, "")
	// ---------------test api
	// ret, err := sqliteQuery()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// conn, err := client.NewClientConn("192.168.49.1:12315")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// conn.CreateCoinInfo(context.Background(), &proto.CreateCoinInfoRequest{
	// 	Info: &proto.CoinInfo{
	// 		Name:             ret.Name,
	// 		ChainType:        ret.ChainType,
	// 		TokenType:        ret.TokenType,
	// 		OfficialContract: ret.Contract,
	// 		Remark:           "ret.Remark",
	// 		Data:             ret.Data,
	// 	},
	// 	Force: true,
	// })

	// ------------------test transfer
	// client, err := eth.Client()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = eth.TransferSpy(context.Background(), client, common.HexToAddress("0x698e90B0bB6Ec413B7036443da977d92246f0350"))
	// log.Fatal(err)

	// --------------------test crawl contract
	// ret, err := eth.CrawlERC20Rows(0, 2)
	// fmt.Println(err, ret)

	cc := make(chan struct{}, 10)
	ccc := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			cc <- struct{}{}
			time.Sleep(time.Second)
		}
		close(ccc)
	}()

	select {
	case <-ccc:
	}
}

type DCoinInfo struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ChainType  string `json:"chain_type"`
	TokenType  string `json:"token_type"`
	Contract   string `json:"contract"`
	Remark     string `json:"remark"`
	Data       []byte `json:"data"`
	CreatedAt  int    `json:"created_at"`
	UpdatedAt  int    `json:"updated_at"`
	DeletedAt  int    `json:"deleted_at"`
	Similarity int    `json:"similarity"`
}

func sqliteQuery() (*DCoinInfo, error) {
	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./test.sqlite.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	// query
	rows, err := db.Query("SELECT * FROM coins_infos WHERE id = " + "'5012faa9-18d3-4f6a-bfcb-4e55546b2811'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dcoin := DCoinInfo{}
	for rows.Next() {

		err = rows.Scan(&dcoin.ID, &dcoin.CreatedAt, &dcoin.UpdatedAt, &dcoin.DeletedAt, &dcoin.Name, &dcoin.ChainType, &dcoin.TokenType, &dcoin.Contract, &dcoin.Similarity, &dcoin.Remark, &dcoin.Data)
		if err != nil {
			return nil, err
		}
		return &dcoin, nil
	}
	return nil, nil
}
