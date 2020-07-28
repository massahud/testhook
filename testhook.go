package testhook

import "testing"

// TestHook is a function executed during specific times on a test cycle
type TestHook func(t *testing.T)

// TestSpawner is a interface with the testing.T#Run function to help test
// the package.
type TestSpawner interface {
	Run(name string, f func(t *testing.T)) bool
}

// Wrapper is a wrapper to run subtests
type Wrapper struct {
	t        TestSpawner
	setup    TestHook
	tearDown TestHook
}

// Wrap wraps a TestSpawner into a Wrapper.
func Wrap(t TestSpawner) *Wrapper {
	return &Wrapper{t: t}
}

// BeforeEach sets the specified TestHook to execute before every subtest.
// If setup fails the given testing.T, the subtest function is not executed,
// but teardown hook will still execute.
func (w *Wrapper) BeforeEach(setup TestHook) {
	w.setup = setup
}

// AfterEach sets the specified TestHook to execut after every subtest.
// The hook is executed even if the setup TestHook or the subtest failed.
func (w *Wrapper) AfterEach(teardown TestHook) {
	w.tearDown = teardown
}

// Run wraps testing.T#Run, executing BeforeEach and AfterEach TestHooks
func (w *Wrapper) Run(name string, f func(t *testing.T)) bool {
	fWrapped := func(tt *testing.T) {
		if w.setup != nil {
			w.setup(tt)
		}
		if !tt.Failed() {
			f(tt)
		}
		if w.tearDown != nil {
			w.tearDown(tt)
		}
	}
	return w.t.Run(name, fWrapped)
}
