package net_test

import (
	binet "github.com/cloudfoundry/bosh-cli/common/net"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net"
	"strconv"
)

var _ = Describe("common.net freeport", func() {
	Describe("GetFreePort", func() {
		It("Return an available free port that is ready to use", func() {
			port, err := binet.GetFreePort()
			Expect(err).ToNot(HaveOccurred())

			Expect(port).ShouldNot(Equal(0))

			// Try to listen on the port
			l, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(port))
			Expect(err).ToNot(HaveOccurred())
			defer l.Close()
		})
	})
})
