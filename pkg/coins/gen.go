package coins

import (
	"bytes"
	"context"
	"fmt"
	"go/format"
	"os"
	"text/template"

	bc_client "github.com/NpoolPlatform/build-chain/pkg/client/v1"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/tokeninfo"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	proto "github.com/NpoolPlatform/message/npool/build-chain"
	"google.golang.org/protobuf/types/known/structpb"
)

const WriteFilePermisson = 0600

type GenTaskInfo struct {
	Host      string
	Package   string
	ListName  string
	TokenType string
	ChainType string
	Out       string
}

type tmplStruct struct {
	Package   string
	ListName  string
	TokenType string
	ChainType string

	Tokens []*proto.TokenInfo
}

func Gen(taskInfo *GenTaskInfo) {
	conn, err := bc_client.NewClientConn(context.Background(), taskInfo.Host)
	if err != nil {
		fmt.Printf("faild: host %v can not connect, %v\n", taskInfo.Host, err)
	}

	conds := cruder.NewFilterConds()
	conds.WithCond(tokeninfo.FieldChainType, cruder.EQ, structpb.NewStringValue(taskInfo.ChainType))
	conds.WithCond(tokeninfo.FieldTokenType, cruder.EQ, structpb.NewStringValue(taskInfo.TokenType))
	resp, err := conn.GetTokenInfos(context.Background(), &proto.GetTokenInfosRequest{Conds: conds})
	if err != nil {
		fmt.Printf("faild: get tokeninfos by %v %v, %v\n", taskInfo.ChainType, taskInfo.TokenType, err)
	}

	tmplData := tmplStruct{
		Package:   taskInfo.Package,
		ListName:  taskInfo.ListName,
		TokenType: taskInfo.TokenType,
		ChainType: taskInfo.ChainType,
		Tokens:    resp.Infos,
	}

	temp := template.Must(template.New("").Parse(tmplSourceGo))
	buff := new(bytes.Buffer)
	err = temp.Execute(buff, tmplData)
	if err != nil {
		fmt.Printf("faild: generate code, %v\n", err)
	}

	code, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("faild: format code, %v\n", err)
	}

	if err := os.WriteFile(taskInfo.Out, code, WriteFilePermisson); err != nil {
		fmt.Printf("faild: write file, %v\n", err)
	}
}

const tmplSourceGo = `package {{.Package}}

import (
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/coins"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/coins/register"
)

func init() {
	for i := range {{.ListName}} {
		{{.ListName}}[i].TokenType = "{{.TokenType}}"
		{{.ListName}}[i].Net = "main"
		{{.ListName}}[i].Waight = 1
		{{.ListName}}[i].Contract = {{.ListName}}[i].OfficialContract
		{{.ListName}}[i].CoinType = sphinxplugin.CoinType_CoinType{{.ChainType}}
		{{.ListName}}[i].Name = coins.GenerateName(&{{.ListName}}[i])
		register.RegisteTokenInfo(&{{.ListName}}[i])
	}
}

var {{.ListName}} = []coins.TokenInfo{
	{{- range .Tokens}}
	{OfficialName: "{{.Name}}", Decimal: {{.Decimal}}, Unit: "{{.Unit}}", OfficialContract: "{{.OfficialContract}}"},
	{{- end}}
}
`
