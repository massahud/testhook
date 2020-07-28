// Package testhook is a simple helper to add
// default beforeEach (setup) and afterEach (teardown)
// hooks to subtests.
//
// To use it just wrap *testing.T into a TestHook,
// configure the hooks and call Run on the TestHook.
//
// Example:
// func TestSomething(t *testing.T) {
// 	th := testhook.Wrap(t)
//
// 	th.BeforeEach(func(t *testing.T) {
// 		fmt.Println(t.Name(), "started")
// 	})
//
// 	th.AfterEach(func(t *testing.T) {
// 		fmt.Println(t.Name(), "passed?", !t.Failed())
// 	})
//
// 	th.Run("a subtest", func(t *testing.T) {
// 		t.Log("foo")
// 	})
// }
//
// The afterEach hook runs even if the test fails. If you do not want it to run
// check t.Failed in your hook.
//
package testhook
