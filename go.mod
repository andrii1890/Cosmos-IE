module github.com/andrii1890/Cosmos-IE

go 1.15

require (
	github.com/bandprotocol/chain v1.0.0
	github.com/cosmos/cosmos-sdk v0.41.3
	github.com/e-money/em-ledger v0.9.4
	github.com/irisnet/irishub v1.0.1
	github.com/kava-labs/kava v0.12.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/prometheus/client_golang v1.8.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/tendermint/tendermint v0.34.7
	github.com/terra-project/core v0.4.2
	github.com/palomachain/paloma v0.11.6
	go.uber.org/zap v1.15.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
