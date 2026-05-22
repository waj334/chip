//go:build !cortexm7 && !cortexm55 && !cortexm85

package runtime

import "cache"

type DCache = cache.NoCache
