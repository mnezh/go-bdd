// The test below produces spec-like output
// lutra-aquae :: ~/Gits/go-bdd ‹goblin*› » go test                                                                                                                                       1 ↵
//   duckduckgo search
//     Looking for something
//       ✓ http status OK
// 	  	 ✓ abstractURL contains something

package whatev

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	Feature("duckduckgo search", t, func(g G) {
		g.Describe("Looking for something", func() {
			var responseBody *map[string]interface{}
			var response *http.Response
			g.Before(func() {
				client, err := NewClient("http://api.duckduckgo.com")
				Expect(err).To(BeNil())
				responseBody, response, err = client.GetJSON("/?q=something%20&format=json")
				Expect(err).To(BeNil())
			})
			g.It("http status OK", func() {
				Expect(response.StatusCode).To(Equal(200))
			})
			g.It("abstractURL contains something", func() {
				abstractURL := strings.ToLower(fmt.Sprintf("%v", (*responseBody)["AbstractURL"]))
				Expect(strings.Contains(abstractURL, "something")).To(BeTrue())
			})
		})
	})
}
