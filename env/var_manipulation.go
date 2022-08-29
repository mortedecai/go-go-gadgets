package env

import (
	"os"
	"strconv"
)

func GetWithDefault(key, defVal string) (val string, found bool) {
	if val, found = os.LookupEnv(key); !found {
		val = defVal
	}
	return
}

func GetWithDefaultInt(key string, defVal int) (val int, found bool) {
	if str, f := os.LookupEnv(key); !f {
		val = defVal
		found = f
	} else {
		if v, err := strconv.Atoi(str); err != nil {
			val = defVal
			found = false
		} else {
			val = v
			found = f
		}
	}
	return
}
