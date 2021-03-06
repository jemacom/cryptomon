# CryptoMon

CryptoMon gives you an insight about the vibrant world of cryptocurrencies by listing their prices.

## Description

The project is composed of three microservices:
* `pricingsvc` Pricing Service that keeps the up-to-date pricing information
* `rankingsvc` Ranking Service that keeps the up-to-date ranking information
* `api` HTTP-API Service that exposes a HTTP endpoint that returns the up-to-date list of top coins prices.

Those microservices use HTTP to communicate with each others.

## Technical choices

* `pricingsvc` uses a CoinMarketCap client `github.com/coincircle/go-coinmarketcap` to get the prices of all cryptocurrencies.

* `rankingsvc` uses a CryptoCompare client `github.com/lucazulian/cryptocomparego` to get the ranking.

* We use `multi-stage builds` feature of Docker to build small images.

* `Docker compose` help us with the orchestration discovery of services.

* We use `Golang` because we love it !


## Run

To see CryptoMon in action run docker compose
```
docker-compose up -d
```

The HTTP API service will bind to localhost on port 8080 and you can send requests to the API
```
curl 'http://localhost:8080/v1?limit=4'
```

The output is in a json format
```json
[
    {
        "Rank": 1,
        "Symbol": "BTC",
        "PriceUSD": 7927.73
    },
    {
        "Rank": 2,
        "Symbol": "ETH",
        "PriceUSD": 625.752
    },
    {
        "Rank": 3,
        "Symbol": "LTC",
        "PriceUSD": 125.566
    },
    {
        "Rank": 4,
        "Symbol": "DASH",
        "PriceUSD": 351.144
    }
]
```

# Next steps

* Use `gRPC` for microservices communication
* Write `kubernetes` config files
* Use `Opentracing` for debugging and trancing requests
* Use `swagger` for API definition.
