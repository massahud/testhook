// Runnable example test
package example_test

import (
	"fmt"
	"testing"

	"github.com/massahud/testhook"
)

func TestSomething(t *testing.T) {
	th := testhook.Wrap(t)

	th.BeforeEach(func(t *testing.T) {
		fmt.Println(t.Name(), "started")
	})

	th.AfterEach(func(t *testing.T) {
		fmt.Println(t.Name(), "passed?", !t.Failed())
	})

	th.Run("a subtest", func(t *testing.T) {
		t.Log("foo")
	})
}
