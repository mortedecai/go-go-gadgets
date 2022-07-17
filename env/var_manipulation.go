package env

import (
	"os"
)

func GetWithDefault(key, defVal string) (val string, found bool) {
	if val, found = os.LookupEnv(key); !found {
		val = defVal
	}
	return
}
