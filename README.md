# Etherscan Exporter

Export gas price and latest block number as prometheus metrics

## How to build

just run ```make build```, caddy will be built using the offical golang docker container.


## How to run

Etherscan Exporter requires a listening address and a valid Etherscan API Key

```
etherscan-exporter -bind 0.0.0.0:9638 -apikey <KEY>
```

