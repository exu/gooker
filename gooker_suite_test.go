package gooker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGooker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gooker Suite")
}
