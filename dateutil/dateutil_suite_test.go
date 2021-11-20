package dateutil_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDateutil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dateutil Suite")
}
