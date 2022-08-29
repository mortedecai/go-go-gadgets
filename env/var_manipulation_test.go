package env_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mortedecai/go-go-gadgets/env"
)

const (
	expKey        = "ENV_VAR_TEST_KEY"
	defValue      = "Some Value No One Would Ever Enter"
	assignedValue = "42"
)

var _ = Describe("Environment Variable Manipulation Tests", func() {
	Describe("GetWithDefault", func() {
		It("should return the default value if the variable is not found", func() {
			val, found := env.GetWithDefault(expKey, defValue)
			Expect(val).To(Equal(defValue))
			Expect(found).To(BeFalse())
		})
		It("should return the value if the variable is found", func() {
			origVal, valFound := os.LookupEnv(expKey)
			os.Setenv(expKey, assignedValue)
			val, found := env.GetWithDefault(expKey, defValue)
			Expect(val).To(Equal(assignedValue))
			Expect(found).To(BeTrue())

			if valFound {
				os.Setenv(expKey, origVal)
			} else {
				os.Unsetenv(expKey)
			}
		})
	})
})
