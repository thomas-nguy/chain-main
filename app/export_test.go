package app

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store/rootmulti"
	"encoding/hex"
	dbm "github.com/cosmos/cosmos-db"
	cronos "github.com/crypto-org-chain/cronos/store/rootmulti"
	"github.com/stretchr/testify/require"
	"strings"
)

func TestMemIAVL(t *testing.T) {
	store := cronos.NewStore("/Users/thomasnguy/Documents/localnet/level-chainmain/new_node/data/memiavl.db", log.NewNopLogger(), true, false)
	err := store.LoadVersion(500)
	require.NoError(t, err)
	hash := store.LastCommitID().Hash
	//version := store.LatestVersion()

	//require.Equal(t, 10, version)
	require.Equal(t, "a", strings.ToUpper(hex.EncodeToString(hash)))
}

func TestVersionCommitID(t *testing.T) {
	db, _ := dbm.NewGoLevelDB("application", "/Users/thomasnguy/Documents/localnet/level-chainmain/data/chainmaind/node0/data", nil)
	store := rootmulti.NewStore(db, log.NewNopLogger(), nil)
	err := store.LoadVersion(500)
	require.NoError(t, err)
	hash := store.LastCommitID().Hash
	//version := store.LatestVersion()

	//require.Equal(t, 10, version)
	require.Equal(t, "a", strings.ToUpper(hex.EncodeToString(hash)))
}

func TestVersionCommitIDNewNode(t *testing.T) {
	db, _ := dbm.NewGoLevelDB("application", "/Users/thomasnguy/Documents/localnet/level-chainmain/new_node_version/data", nil)
	store := rootmulti.NewStore(db, log.NewNopLogger(), nil)
	err := store.LoadVersion(500)
	require.NoError(t, err)
	hash := store.LastCommitID().Hash
	//version := store.LatestVersion()

	//require.Equal(t, 10, version)
	require.Equal(t, "a", strings.ToUpper(hex.EncodeToString(hash)))
}
