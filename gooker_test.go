package gooker_test

import (
	. "github.com/exu/gooker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
	It("gets some books", func() {
		book := GetSomeBook()

		Expect(book.Title).To(Equal("Rainbow Six"))
		Expect(book.Author).To(Equal("Tom Clancy"))
		Expect(book.Pages).To(Equal(1354))
	})

	It("can scream LOUD!", func() {
		book := GetSomeBook()
		scream := book.ScreamLoud()

		Expect(scream).To(Equal("LOUD!"))
	})
})
