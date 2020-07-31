package collectors

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"net/http"
)

type CurrentBlockCollector struct {
	apikey string
	desc *prometheus.Desc
}
type CurrentBlockJsonRPCResponse struct {
	JsonRpc string `json:jsonrpc`
	Id int `json:id`
	Result string `json:result`
}
func NewCurrentBlockCollector(apikey string) *CurrentBlockCollector {
	return &CurrentBlockCollector{
		apikey: apikey,
		desc: prometheus.NewDesc(
			"etherscan_current_block",
			"latest block number",
			nil,
			nil,
			),
	}
}

func (collector *CurrentBlockCollector) Describe(ch chan<- *prometheus.Desc) {
    ch <- collector.desc
}

func (collector *CurrentBlockCollector) Collect(ch chan<- prometheus.Metric) {
	response, err := http.Get(fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_blockNumber&apikey=%s", collector.apikey))
	if err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}
	var currentBlockJsonRpcResponse CurrentBlockJsonRPCResponse
	json.Unmarshal(content, &currentBlockJsonRpcResponse)
	currentBlock, err := hexutil.DecodeUint64(currentBlockJsonRpcResponse.Result)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}
	ch <- prometheus.MustNewConstMetric(collector.desc, prometheus.GaugeValue, float64(currentBlock))
}
