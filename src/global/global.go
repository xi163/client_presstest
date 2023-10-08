package global

import (
	"os"
	"path/filepath"
)

var (
	path, _  = os.Executable()
	Dir, Exe = filepath.Split(path)
)
var (
	Name string
)
