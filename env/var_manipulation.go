package env

import (
	"fmt"
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
	fmt.Println("WTF?")
	if str, f := os.LookupEnv(key); !f {
		fmt.Printf("OS NOT FOUND: %s, %v\n", key, found)
		val = defVal
		found = f
	} else {
		fmt.Printf("Converting '%s' to int\n", str)
		if v, err := strconv.Atoi(str); err != nil {
			val = defVal
			found = false
			fmt.Printf("Error converting '%s' to int: %v\n", str, err)
		} else {
			val = v
			found = f
		}
	}
	return
}
