module github.com/netcloth/netcloth-chain

go 1.12

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd v0.0.0-20190115013929-ed77733ec07d
	github.com/cosmos/go-bip39 v0.0.0-20180618194314-52158e4697b8
	github.com/cosmos/ledger-cosmos-go v0.10.3
	github.com/ethereum/go-ethereum v1.9.9
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.3.1-0.20190508161146-9fa652df1129
	github.com/gorilla/mux v1.7.3
	github.com/mattn/go-isatty v0.0.11
	github.com/pelletier/go-toml v1.8.0
	github.com/pkg/errors v0.9.0
	github.com/rakyll/statik v0.1.6
	github.com/spf13/afero v1.2.1 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.1
	github.com/stretchr/testify v1.4.0
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/iavl v0.12.4
	github.com/tendermint/tendermint v0.32.12
	github.com/tendermint/tm-db v0.2.0
	golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/tendermint/tendermint => github.com/netcloth/tendermint v0.32.15
