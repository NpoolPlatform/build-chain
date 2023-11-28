package eth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrawlContract(t *testing.T) {
	contractInfo, err := CrawlContractInfo("0x111111111117dc0aa78b770fa6a738034120c302")
	assert.Nil(t, err)
	assert.Equal(t, "1INCH Token", contractInfo.Name)
	assert.Equal(t, "0x111111111117dC0aa78b770fA6A738034120C302", contractInfo.OfficialContract)
	assert.Equal(t, "18", contractInfo.Decimal)
	assert.Equal(t, "1INCH", contractInfo.Unit)

	contractInfo, err = CrawlContractInfo("0x111")
	assert.NotNil(t, err)
	assert.Nil(t, contractInfo)
}

func TestCrawlRows(t *testing.T) {
	contracts, err := CrawlERC20Rows(0, 1)
	assert.NotNil(t, err)
	assert.Empty(t, contracts)

	contracts, err = CrawlERC20Rows(2, 1)
	assert.NotNil(t, err)
	assert.Equal(t, 1, len(contracts))

	contracts, err = CrawlERC20Rows(25, 50)
	assert.NotNil(t, err)
	assert.Equal(t, 50, len(contracts))
}
