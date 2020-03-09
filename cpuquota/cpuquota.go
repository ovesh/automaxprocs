package cpuquota

import "go.uber.org/automaxprocs/internal/cgroups"

// NewCGroupsForCurrentProcess returns a new *CGroups instance for the current
// process.
func NewCGroupsForCurrentProcess() (cgroups.CGroups, error) {
	return cgroups.NewCGroupsForCurrentProcess()
}

// CPUQuota returns the CPU quota applied with the CPU cgroup controller.
// It is a result of `cpu.cfs_quota_us / cpu.cfs_period_us`. If the value of
// `cpu.cfs_quota_us` was not set (-1), the method returns `(-1, nil)`.
func CPUQuota(cg cgroups.CGroups) (float64, bool, error) {
	return cg.CPUQuota()
}
