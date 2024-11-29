package api

import (
	"github.com/cryft-labs/coreth/plugin/evm"
	"github.com/cryft-labs/cryftgo/api/admin"
	"github.com/cryft-labs/cryftgo/api/health"
	"github.com/cryft-labs/cryftgo/api/info"
	"github.com/cryft-labs/cryftgo/api/keystore"
	"github.com/cryft-labs/cryftgo/indexer"
	"github.com/cryft-labs/cryftgo/vms/avm"
	"github.com/cryft-labs/cryftgo/vms/platformvm"
)

// Issues API calls to a node
// TODO: byzantine api. check if appropriate. improve implementation.
type Client interface {
	PChainAPI() platformvm.Client
	XChainAPI() avm.Client
	XChainWalletAPI() avm.WalletClient
	CChainAPI() evm.Client
	CChainEthAPI() EthClient // ethclient websocket wrapper that adds mutexed calls, and lazy conn init (on first call)
	InfoAPI() info.Client
	HealthAPI() health.Client
	KeystoreAPI() keystore.Client
	AdminAPI() admin.Client
	PChainIndexAPI() indexer.Client
	CChainIndexAPI() indexer.Client
	// TODO add methods
}
