# api-server
API backend server for ShootingCoin project

## Install
```
go mod tidy
go run server.go
```

## Configuration
```
./abigen --abi=./abi/abi.json --pkg=chain --type=ShootingCoinManager --out=./abi/contract.go
```