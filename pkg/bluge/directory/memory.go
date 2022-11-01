package directory

import (
	"github.com/blugelabs/bluge"
	"github.com/blugelabs/bluge/index"
)

func GetMemoryConfig(bucket string, indexName string, timeRange ...int64) bluge.Config {
	config := index.InMemoryOnlyConfig()
	return bluge.DefaultConfigWithIndexConfig(config)
}