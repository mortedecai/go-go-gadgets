package env_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mortedecai/go-go-gadgets/env"
)

const (
	expKey          = "ENV_VAR_TEST_KEY"
	defValue        = "Some Value No One Would Ever Enter"
	assignedValue   = "42"
	assignedValInt  = 42
	defIntVal       = 316
	defBoolVal      = true
	assignedValBool = false
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
	Describe("GetWithDefaultBool", func() {
		It("should return the default value if the variable is not found", func() {
			val, found := env.GetWithDefaultBool(expKey, defBoolVal)
			Expect(val).To(Equal(defBoolVal))
			Expect(found).To(BeFalse())
		})
		It("should return the value if the variable is found", func() {
			origVal, valFound := os.LookupEnv(expKey)
			if err := os.Setenv(expKey, fmt.Sprintf("%v", assignedValBool)); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("No error setting variable")
			}
			val, found := env.GetWithDefaultBool(expKey, defBoolVal)
			Expect(val).To(Equal(assignedValBool))
			Expect(found).To(BeTrue())

			if valFound {
				os.Setenv(expKey, origVal)
			} else {
				os.Unsetenv(expKey)
			}
		})
		It("should return the default value if the variable is found but not parseable", func() {
			origVal, valFound := os.LookupEnv(expKey)
			os.Setenv(expKey, defValue)
			val, found := env.GetWithDefaultBool(expKey, defBoolVal)
			Expect(val).To(Equal(defBoolVal))
			Expect(found).To(BeFalse())

			if valFound {
				os.Setenv(expKey, origVal)
			} else {
				os.Unsetenv(expKey)
			}
		})
	})

	Describe("GetWithDefaultInt", func() {
		It("should return the default value if the variable is not found", func() {
			val, found := env.GetWithDefaultInt(expKey, defIntVal)
			Expect(val).To(Equal(defIntVal))
			Expect(found).To(BeFalse())
		})
		It("should return the value if the variable is found", func() {
			origVal, valFound := os.LookupEnv(expKey)
			if err := os.Setenv(expKey, assignedValue); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("No error setting variable")
			}
			val, found := env.GetWithDefaultInt(expKey, defIntVal)
			Expect(val).To(Equal(assignedValInt))
			Expect(found).To(BeTrue())

			if valFound {
				os.Setenv(expKey, origVal)
			} else {
				os.Unsetenv(expKey)
			}
		})
		It("should return the default value if the variable is found but not parseable", func() {
			origVal, valFound := os.LookupEnv(expKey)
			os.Setenv(expKey, defValue)
			val, found := env.GetWithDefaultInt(expKey, defIntVal)
			Expect(val).To(Equal(defIntVal))
			Expect(found).To(BeFalse())

			if valFound {
				os.Setenv(expKey, origVal)
			} else {
				os.Unsetenv(expKey)
			}
		})
	})
})
