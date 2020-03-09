# libcgroups

This is a fork of github.com/uber-go/automaxprocs that exposes the existing functions
necessary to get the current process' cgroups, and the cpu quota allocation from that.


Importing the base package will automatically set `GOMAXPROCS` to match Linux container
CPU quota.

## Quick Start

To automatically set `GOMAXPROCS`:
```go
import _ "github.com/ovesh/automaxprocs"

func main() {
  // Your application logic here.
}
```

To get the current CPU quota
```go
allCGroups, err := cgroups.NewCGroupsForCurrentProcess()
if err != nil {
	// handle err
}
quota, defined, err := allCGroups.CPUQuota()
if !defined || err != nil {
	// handle err
}

fmt.Println("quota:", quota)
```

<hr>

Released under the [MIT License](LICENSE).
