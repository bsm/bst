package bst

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("elements", func() {
	var subject elements

	BeforeEach(func() {
		subject = make(elements, 0, 5)
		subject, _ = subject.put(Int(2), false)
		subject, _ = subject.put(Int(4), false)
		subject, _ = subject.put(Int(6), false)
	})

	It("should have len", func() {
		Expect(subject.Len()).To(Equal(3))
	})

	It("should search", func() {
		Expect(subject.search(Int(2))).To(Equal(0))
		Expect(subject.search(Int(3))).To(Equal(1))
		Expect(subject.search(Int(4))).To(Equal(1))
		Expect(subject.search(Int(6))).To(Equal(2))
		Expect(subject.search(Int(7))).To(Equal(3))

		pos, ok := subject.searchFrom(Int(2), 0)
		Expect(ok).To(BeTrue())
		Expect(pos).To(Equal(0))
		pos, ok = subject.searchFrom(Int(2), 1)
		Expect(ok).To(BeFalse())
		Expect(pos).To(Equal(1))
		pos, ok = subject.searchFrom(Int(4), 1)
		Expect(ok).To(BeTrue())
		Expect(pos).To(Equal(1))
		pos, ok = subject.searchFrom(Int(5), 1)
		Expect(ok).To(BeFalse())
		Expect(pos).To(Equal(2))
		pos, ok = subject.searchFrom(Int(6), 1)
		Expect(ok).To(BeTrue())
		Expect(pos).To(Equal(2))
		pos, ok = subject.searchFrom(Int(6), 2)
		Expect(ok).To(BeTrue())
		Expect(pos).To(Equal(2))
		pos, ok = subject.searchFrom(Int(7), 2)
		Expect(ok).To(BeFalse())
		Expect(pos).To(Equal(3))
	})

	It("should remove data", func() {
		var ok bool
		subject, ok = subject.delete(Int(3))
		Expect(ok).To(BeFalse())
		Expect(subject.Len()).To(Equal(3))

		subject, ok = subject.delete(Int(2))
		Expect(ok).To(BeTrue())
		Expect(subject.Len()).To(Equal(2))
	})

	It("should check existence", func() {
		Expect(subject.Exists(Int(1))).To(BeFalse())
		Expect(subject.Exists(Int(2))).To(BeTrue())
		Expect(subject.Exists(Int(3))).To(BeFalse())
		Expect(subject.Exists(Int(4))).To(BeTrue())
	})

})
