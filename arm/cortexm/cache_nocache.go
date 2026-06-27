//go:build !cortexm7 && !cortexm55 && !cortexm85

package cortexm

import "cache"

type DCache = cache.NoCache
