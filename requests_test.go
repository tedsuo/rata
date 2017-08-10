package rata_test

import (
	. "github.com/tedsuo/rata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Requests", func() {
	var (
		host             string
		requestGenerator *RequestGenerator
	)
	const (
		PathWithSlash    = "WithSlash"
		PathWithoutSlash = "WithoutSlash"
	)

	JustBeforeEach(func() {
		routes := Routes{
			{Name: PathWithSlash, Method: "GET", Path: "/some-route"},
			{Name: PathWithoutSlash, Method: "GET", Path: "some-route"},
		}
		requestGenerator = NewRequestGenerator(
			host,
			routes,
		)
	})

	Describe("CreateRequest", func() {
		Context("when the host does not have a trailing slash", func() {
			BeforeEach(func() {
				host = "http://example.com"
			})

			Context("when the path starts with a slash", func() {
				It("generates a URL with one slash between the host and the path", func() {
					request, err := requestGenerator.CreateRequest(PathWithSlash, Params{}, nil)
					Expect(err).NotTo(HaveOccurred())

					Expect(request.URL.String()).To(Equal("http://example.com/some-route"))
				})
			})

			Context("when the path does not start with a slash", func() {
				It("generates a URL with one slash between the host and the path", func() {
					request, err := requestGenerator.CreateRequest(PathWithoutSlash, Params{}, nil)
					Expect(err).NotTo(HaveOccurred())

					Expect(request.URL.String()).To(Equal("http://example.com/some-route"))
				})
			})
		})

		Context("when host has a trailing slash", func() {
			BeforeEach(func() {
				host = "example.com/"
			})

			Context("when the path starts with a slash", func() {
				It("generates a URL with one slash between the host and the path", func() {
					request, err := requestGenerator.CreateRequest(PathWithSlash, Params{}, nil)
					Expect(err).NotTo(HaveOccurred())

					Expect(request.URL.String()).To(Equal("example.com/some-route"))
				})
			})

			Context("when the path does not start with a slash", func() {
				It("generates a URL with one slash between the host and the path", func() {
					request, err := requestGenerator.CreateRequest(PathWithoutSlash, Params{}, nil)
					Expect(err).NotTo(HaveOccurred())

					Expect(request.URL.String()).To(Equal("example.com/some-route"))
				})
			})
		})
	})
})
