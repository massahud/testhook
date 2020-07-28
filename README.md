# TestHook

![Master](https://github.com/massahud/testhook/workflows/Master/badge.svg)

## About The Project

Package testhook is a tiny module for adding beforeEach (setup) and afterEach (teardown) hooks to Go subtests.

All of the documentation can be found on the [go.dev](https://pkg.go.dev/github.com/massahud/testhook?tab=doc) website.

### Very quick start

```go
package some_test

import (
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
        t.Error("foo")
    })
}
```

## Licensing

```text
Copyright 2020, Geraldo Augusto Massahud Rodrigues dos Santos

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
