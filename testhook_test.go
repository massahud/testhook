package testhook

import (
	"testing"
)

type mockSpawner struct {
	shouldFail bool
	executed   bool
}

func (ts *mockSpawner) Run(name string, f func(t *testing.T)) bool {
	var t testing.T
	f(&t)
	ts.executed = true
	return !ts.shouldFail
}

func TestNoHook(t *testing.T) {
	ts := mockSpawner{}
	th := Wrap(&ts)
	var testRun bool
	th.Run("test", func(t *testing.T) {
		testRun = true
	})
	if !ts.executed {
		t.Error("Spanwer not executed")
	}
	if !testRun {
		t.Error("test function did not run")
	}
}

func TestBeforeEach(t *testing.T) {
	t.Run("SetupPass", func(t *testing.T) {
		ts := mockSpawner{}
		th := Wrap(&ts)
		var setupRun bool
		th.BeforeEach(func(t *testing.T) {
			setupRun = true
		})
		var testRun bool
		th.Run("test", func(t *testing.T) {
			testRun = true
		})
		if !ts.executed {
			t.Error("Spanwer not executed")
		}
		if !setupRun {
			t.Error("setup did not run")
		}
		if !testRun {
			t.Error("test function did not run")
		}
	})
	t.Run("SetupFail", func(t *testing.T) {
		ts := mockSpawner{}
		th := Wrap(&ts)
		var setupRun bool
		th.BeforeEach(func(t *testing.T) {
			setupRun = true
			t.Fail()
		})
		var testRun bool
		th.Run("test", func(t *testing.T) {
			testRun = true
		})
		if !ts.executed {
			t.Error("Spanwer not executed")
		}
		if !setupRun {
			t.Error("setup did not run")
		}
		if testRun {
			t.Error("test function should not execute when setup fails")
		}
	})
}

func TestAfterEach(t *testing.T) {
	t.Run("TestPass", func(t *testing.T) {
		t.Log("should execute")
		ts := mockSpawner{}
		th := Wrap(&ts)
		var teardownRun bool
		th.AfterEach(func(t *testing.T) {
			teardownRun = true
		})
		var testRun bool
		th.Run("test", func(t *testing.T) {
			testRun = true
		})
		if !ts.executed {
			t.Error("Spanwer not executed")
		}
		if !testRun {
			t.Error("test function did not run")
		}
		if !teardownRun {
			t.Error("teardown did not run")
		}
	})
	t.Run("TestFail", func(t *testing.T) {
		t.Log("should execute if test function fails")
		ts := mockSpawner{}
		th := Wrap(&ts)
		var teardownRun bool
		th.AfterEach(func(t *testing.T) {
			teardownRun = true
		})
		var testRun bool
		th.Run("test", func(t *testing.T) {
			testRun = true
			t.Fail()
		})
		if !ts.executed {
			t.Error("Spanwer not executed")
		}
		if !testRun {
			t.Error("test function did not run")
		}
		if !teardownRun {
			t.Error("teardown did not run")
		}
	})

	t.Run("SetupFail", func(t *testing.T) {
		t.Log("should execute if setup fails")
		ts := mockSpawner{}
		th := Wrap(&ts)
		var setupRun bool
		th.BeforeEach(func(t *testing.T) {
			setupRun = true
			t.Fail()
		})
		var teardownRun bool
		th.AfterEach(func(t *testing.T) {
			teardownRun = true
		})
		var testRun bool
		th.Run("test", func(t *testing.T) {
			testRun = true
		})
		if !ts.executed {
			t.Error("Spanwer not executed")
		}
		if !setupRun {
			t.Error("setup did not run")
		}
		if testRun {
			t.Error("test function should not run")
		}
		if !teardownRun {
			t.Error("teardown did not run")
		}
	})
}
