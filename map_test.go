package bst

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map", func() {
	var subject *Map
	var collect = func() map[int]string {
		res := make(map[int]string, subject.Len())
		for iter := subject.Iterator(); iter.Next(); {
			res[int(iter.Key().(Int))] = iter.Value().(string)
		}
		return res
	}

	BeforeEach(func() {
		subject = NewMap(5)
		Expect(subject.Add(Int(2), "B")).To(BeTrue())
		Expect(subject.Add(Int(4), "D")).To(BeTrue())
		Expect(subject.Add(Int(6), "F")).To(BeTrue())
	})

	It("should have len", func() {
		Expect(subject.Len()).To(Equal(3))
	})

	It("should get", func() {
		val, ok := subject.Get(Int(1))
		Expect(ok).To(BeFalse())
		Expect(val).To(BeNil())

		val, ok = subject.Get(Int(2))
		Expect(ok).To(BeTrue())
		Expect(val).To(Equal("B"))
	})

	It("should set data", func() {
		Expect(subject.Set(Int(3), "C")).To(BeFalse())
		Expect(subject.Set(Int(1), "A")).To(BeFalse())
		Expect(subject.Len()).To(Equal(5))

		Expect(subject.Set(Int(2), "NEWB")).To(BeTrue())
		Expect(subject.Set(Int(3), "NEWC")).To(BeTrue())
		Expect(subject.Set(Int(4), "NEWD")).To(BeTrue())
		Expect(subject.Len()).To(Equal(5))

		Expect(collect()).To(Equal(map[int]string{1: "A", 2: "NEWB", 3: "NEWC", 4: "NEWD", 6: "F"}))
	})

	It("should add data", func() {
		Expect(subject.Add(Int(3), "C")).To(BeTrue())
		Expect(subject.Add(Int(1), "A")).To(BeTrue())
		Expect(subject.Len()).To(Equal(5))

		Expect(subject.Add(Int(2), "NEWB")).To(BeFalse())
		Expect(subject.Add(Int(3), "NEWC")).To(BeFalse())
		Expect(subject.Add(Int(4), "NEWD")).To(BeFalse())
		Expect(subject.Len()).To(Equal(5))

		Expect(collect()).To(Equal(map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 6: "F"}))
	})

	It("should delete data", func() {
		Expect(subject.Delete(Int(3))).To(BeFalse())
		Expect(subject.Len()).To(Equal(3))
		Expect(subject.Delete(Int(2))).To(BeTrue())
		Expect(subject.Len()).To(Equal(2))
		Expect(subject.Delete(Int(2))).To(BeFalse())
		Expect(subject.Len()).To(Equal(2))
	})

	It("should check if exists", func() {
		Expect(subject.Exists(Int(1))).To(BeFalse())
		Expect(subject.Exists(Int(2))).To(BeTrue())
		Expect(subject.Exists(Int(3))).To(BeFalse())
		Expect(subject.Exists(Int(4))).To(BeTrue())
	})

	It("should iterate", func() {
		iter := subject.Iterator()

		Expect(iter.Next()).To(BeTrue())
		Expect(iter.Key()).To(Equal(Int(2)))
		Expect(iter.Value()).To(Equal("B"))

		Expect(iter.Next()).To(BeTrue())
		Expect(iter.Key()).To(Equal(Int(4)))
		Expect(iter.Value()).To(Equal("D"))

		Expect(iter.Next()).To(BeTrue())
		Expect(iter.Key()).To(Equal(Int(6)))
		Expect(iter.Value()).To(Equal("F"))

		Expect(iter.Next()).To(BeFalse())
		Expect(iter.Previous()).To(BeTrue())
		Expect(iter.Key()).To(Equal(Int(4)))
		Expect(iter.Value()).To(Equal("D"))
	})

	It("should seek", func() {
		iter := subject.Iterator()
		Expect(iter.Seek(Int(4))).To(BeTrue())
		Expect(iter.Key()).To(Equal(Int(4)))
		Expect(iter.Next()).To(BeTrue())
		Expect(iter.Key()).To(Equal(Int(6)))
		Expect(iter.Next()).To(BeFalse())

		Expect(iter.Seek(Int(6))).To(BeTrue())
		Expect(iter.Key()).To(Equal(Int(6)))
		Expect(iter.Next()).To(BeFalse())

		Expect(iter.Seek(Int(7))).To(BeFalse())
	})

})
