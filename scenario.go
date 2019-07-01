package whatev

import (
	"testing"

	"github.com/franela/goblin"
	"github.com/onsi/gomega"
)

// G is an alias to *goblin.G
type G = *goblin.G

// Feature does goblin plumbing setup and executes passed test closure
func Feature(name string, t *testing.T, body func(G)) {
	g := goblin.Goblin(t)
	gomega.RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	g.Describe(name, func() {
		body(g)
	})
}
