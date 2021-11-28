package typeconfig_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTypeconfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Typeconfig Suite")
}
