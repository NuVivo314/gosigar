// Copyright (c) 2012 VMware, Inc.

package sigar

type LoadAverage struct {
	One, Five, Fifteen float64
}

type Uptime struct {
	Length float64
}

type Mem struct {
	Total      uint64
	Used       uint64
	Free       uint64
	ActualFree uint64
	ActualUsed uint64
	Active     uint64
	Inactive   uint64
	Buffers    uint64
	Cached     uint64
}

type Swap struct {
	Total uint64
	Used  uint64
	Free  uint64
}

type Cpu struct {
	User    uint64
	Nice    uint64
	Sys     uint64
	Idle    uint64
	Wait    uint64
	Irq     uint64
	SoftIrq uint64
	Stolen  uint64
}

func (self *Cpu) Total() uint64 {
	return self.User + self.Nice + self.Sys + self.Idle +
		self.Wait + self.Irq + self.SoftIrq + self.Stolen
}

type CpuList struct {
	List []Cpu
}

type FileSystem struct {
	DirName     string
	DevName     string
	TypeName    string
	SysTypeName string
	Options     string
	Flags       uint32
}

type DiskIO struct {
	Name        string
	ReadIssued  uint64
	ReadMerged  uint64
	ReadSectors uint64
	ReadSpentMs uint64

	WriteComplet uint64
	WriteMerged  uint64
	WriteSectors uint64
	WriteSpentMs uint64

	WaitIo                uint64
	WaitIoSpentMs         uint64
	WaitIoSpentMsWeighted uint64
}

type DiskIOList struct {
	List []DiskIO
}

type FileSystemList struct {
	List []FileSystem
}

type FileSystemUsage struct {
	Total     uint64
	Used      uint64
	Free      uint64
	Avail     uint64
	Files     uint64
	FreeFiles uint64
}

type ProcList struct {
	List []int
}

type RunState byte

const (
	RunStateSleep   = 'S'
	RunStateRun     = 'R'
	RunStateStop    = 'T'
	RunStateZombie  = 'Z'
	RunStateIdle    = 'D'
	RunStateUnknown = '?'
)

type ProcState struct {
	Name      string
	State     RunState
	Ppid      int
	Tty       int
	Priority  int
	Nice      int
	Processor int
}

type ProcMem struct {
	Size        uint64
	Resident    uint64
	Share       uint64
	MinorFaults uint64
	MajorFaults uint64
	PageFaults  uint64
}

type ProcTime struct {
	StartTime uint64
	User      uint64
	Sys       uint64
	Total     uint64
}

type ProcArgs struct {
	List []string
}

type ProcExe struct {
	Name string
	Cwd  string
	Root string
}

// From: https://code.google.com/p/psutil/source/browse/psutil/_pslinux.py#535
type ProcIO struct {
	ReadCount  uint64
	WriteCount uint64
	ReadBytes  uint64
	WriteBytes uint64
}

type ProcUser struct {
	UidReal      string
	UidEffective string
	UidSaveSet   string
	UidFs        string
	GidReal      string
	GidEffective string
	GidSaveSet   string
	GidFs        string
	OtherGroups  []string
}

type networkData struct {
	Bytes      uint64
	Packets    uint64
	Errs       uint64
	Drop       uint64
	Fifo       uint64
	Frame      uint64
	Compressed uint64
	Multicast  uint64
}

type Network struct {
	Name string
	Tx   networkData
	Rx   networkData
}

type NetworkList struct {
	List []Network
}
