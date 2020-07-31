# Etherscan Prometheus Exporter

Export gas price and latest block number as prometheus metrics

## How to build

Just run ```make build```, etherscan-prometheus will be built using the official golang docker container.


## How to run

Etherscan Exporter requires a listening address and a valid Etherscan API Key

```
etherscan-exporter -bind 0.0.0.0:9638 -apikey <KEY>
```

## Exposed metrics

This exporter exposes the following metrics:

- **etherscan_gas_price**: current gas price
- **etherscan_current_block**: latest block
