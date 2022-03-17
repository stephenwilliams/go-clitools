package oshelpers

import "os"

func IsExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}
