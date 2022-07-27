package handlers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handlers Test", func() {
	Describe("/books/", func() {
		When("records is available", func() {
			It("should return list of books", func() {
				Expect(1 + 1).To(Equal(4))
			})
		})
	})
})
