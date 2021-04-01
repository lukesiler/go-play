package findsumpair

import "testing"

import "github.com/onsi/gomega"

var (
	ascOrderNums1 = []int{1, 2, 3, 4, 5, 10, 11, 12, 13, 14, 15, 105, 110, 113, 2000, 2001}
)

func TestHasSumPairImpls(t *testing.T) {
	g := gomega.NewWithT(t)

	funcImpls := []func(ascOrderNums []int, sum int) (found bool, index1 int, index2 int){HasSumPairOrder2, HasSumPairOrder1Ordered, HasSumPairOrder1Unordered}
	for _, funcImpl := range funcImpls {
		b, _, _ := funcImpl(ascOrderNums1, 10)
		g.Expect(b).To(gomega.BeFalse())
		//g.Expect(i1).To(gomega.BeNumerically("==", -1))
		//g.Expect(i2).To(gomega.BeNumerically("==", -1))

		b, _, _ = funcImpl(ascOrderNums1, 13)
		g.Expect(b).To(gomega.BeTrue())
		//g.Expect(i1).To(gomega.BeNumerically("==", 0))
		//g.Expect(i2).To(gomega.BeNumerically("==", 7))

		b, _, _ = funcImpl(ascOrderNums1, 2002)
		g.Expect(b).To(gomega.BeTrue())
		//g.Expect(i1).To(gomega.BeNumerically("==", 0))
		//g.Expect(i2).To(gomega.BeNumerically("==", 15))
	}
}
