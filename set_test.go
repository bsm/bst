package bst

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewSet", func() {
	var subject *Set
	var collect = func() []int {
		res := make([]int, 0, subject.Len())
		for iter := subject.Iterator(); iter.Next(); {
			res = append(res, int(iter.Value().(Int)))
		}
		return res
	}

	BeforeEach(func() {
		subject = NewSet(5)
		Expect(subject.Add(Int(2))).To(BeTrue())
		Expect(subject.Add(Int(4))).To(BeTrue())
		Expect(subject.Add(Int(6))).To(BeTrue())
	})

	It("should have len", func() {
		Expect(subject.Len()).To(Equal(3))
	})

	It("should add data", func() {
		Expect(subject.Add(Int(3))).To(BeTrue())
		Expect(subject.Add(Int(1))).To(BeTrue())
		Expect(subject.Len()).To(Equal(5))

		Expect(subject.Add(Int(2))).To(BeFalse())
		Expect(subject.Add(Int(3))).To(BeFalse())
		Expect(subject.Add(Int(4))).To(BeFalse())
		Expect(subject.Len()).To(Equal(5))

		Expect(collect()).To(Equal([]int{1, 2, 3, 4, 6}))
	})

	It("should delete data", func() {
		Expect(subject.Delete(Int(3))).To(BeFalse())
		Expect(subject.Len()).To(Equal(3))
		Expect(subject.Delete(Int(2))).To(BeTrue())
		Expect(subject.Len()).To(Equal(2))
		Expect(subject.Delete(Int(2))).To(BeFalse())
		Expect(subject.Len()).To(Equal(2))

		Expect(collect()).To(Equal([]int{4, 6}))
	})

	It("should check if exists", func() {
		Expect(subject.Exists(Int(1))).To(BeFalse())
		Expect(subject.Exists(Int(2))).To(BeTrue())
		Expect(subject.Exists(Int(3))).To(BeFalse())
		Expect(subject.Exists(Int(4))).To(BeTrue())
	})

	It("should check for intersections", func() {
		oth := NewSet(3)
		Expect(subject.Intersects(oth)).To(BeFalse())

		oth.Add(Int(3))
		oth.Add(Int(5))
		Expect(subject.Intersects(oth)).To(BeFalse())

		oth.Add(Int(7))
		oth.Add(Int(4))
		Expect(subject.Intersects(oth)).To(BeTrue())
	})

	It("should iterate", func() {
		iter := subject.Iterator()

		Expect(iter.Next()).To(BeTrue())
		Expect(iter.Value()).To(Equal(Int(2)))

		Expect(iter.Next()).To(BeTrue())
		Expect(iter.Value()).To(Equal(Int(4)))

		Expect(iter.Next()).To(BeTrue())
		Expect(iter.Value()).To(Equal(Int(6)))

		Expect(iter.Next()).To(BeFalse())
		Expect(iter.Previous()).To(BeTrue())
		Expect(iter.Value()).To(Equal(Int(4)))
	})

	It("should seek", func() {
		iter := subject.Iterator()
		Expect(iter.Seek(Int(4))).To(BeTrue())
		Expect(iter.Value()).To(Equal(Int(4)))
		Expect(iter.Next()).To(BeTrue())
		Expect(iter.Value()).To(Equal(Int(6)))
		Expect(iter.Next()).To(BeFalse())

		Expect(iter.Seek(Int(6))).To(BeTrue())
		Expect(iter.Value()).To(Equal(Int(6)))
		Expect(iter.Next()).To(BeFalse())

		Expect(iter.Seek(Int(7))).To(BeFalse())
	})
})

// --------------------------------------------------------------------

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "set")
}
