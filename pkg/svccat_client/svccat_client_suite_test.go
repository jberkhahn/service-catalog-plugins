package svccat_client_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSvcCatClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PluginClient Suite")
}

var _ = BeforeEach(func() {
	os.Setenv("KUBECTL_PLUGINS_GLOBAL_FLAG_AS_GROUP", "[]")
})
