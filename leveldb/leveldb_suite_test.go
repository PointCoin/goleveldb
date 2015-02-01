package leveldb

import (
	"testing"

	"github.com/PointCoin/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
