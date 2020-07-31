package collectors

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"net/http"
)

type GasPriceCollector struct {
	apikey string
	desc *prometheus.Desc
}
type GasPriceJsonRPCResponse struct {
	JsonRpc string `json:jsonrpc`
	Id int `json:id`
	Result string `json:result`
}
func NewGasPriceCollector(apikey string) *GasPriceCollector {
	return &GasPriceCollector{
		apikey: apikey,
		desc: prometheus.NewDesc(
			"etherscan_gas_price",
			"the current price per gas in wei",
			nil,
			nil,
			),
	}
}

func (collector *GasPriceCollector) Describe(ch chan<- *prometheus.Desc) {
    ch <- collector.desc
}

func (collector *GasPriceCollector) Collect(ch chan<- prometheus.Metric) {
	response, err := http.Get(fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_gasPrice&apikey=%s", collector.apikey))
	if err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}
	var gasPriceJsonRpcResponse GasPriceJsonRPCResponse
	json.Unmarshal(content, &gasPriceJsonRpcResponse)
	gasPrice, err := hexutil.DecodeUint64(gasPriceJsonRpcResponse.Result)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}
	ch <- prometheus.MustNewConstMetric(collector.desc, prometheus.GaugeValue, float64(gasPrice))
}
