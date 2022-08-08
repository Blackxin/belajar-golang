package constants

import (
	"os"
	"path/filepath"
)

const TotalFile = 3000
const ContentLength = 5000

var TempPath = filepath.Join(os.Getenv("XDG_CACHE_HOME"), "chapter-A.60-worker-pool")
