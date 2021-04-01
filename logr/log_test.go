package logr

import (
	"testing"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
)

func TestBaseLogr(t *testing.T) {
	g := gomega.NewWithT(t)

	l := NewBase("base-logger")
	g.Expect(l).ToNot(BeNil())
	l.Info("something something")
	ls := l.Sugar()
	g.Expect(ls).ToNot(BeNil())
	ls.Infow("more", "key", "value")
}
