package info

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/Anandhu3301/osStatChecker/models"
	"github.com/Anandhu3301/osStatChecker/pkg"
)

var (
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	globalMemorystatusEx = kernel32.NewProc("GlobalMemoryStatusEx")
)

func MemoryValues() models.SystemGenerator {
	return func(sv *models.SystemValues) {
		var memoryStatus memoryStatusEx

		memoryStatus.Length = uint32(unsafe.Sizeof(memoryStatus))
		ret, _, err := globalMemorystatusEx.Call(uintptr(unsafe.Pointer(&memoryStatus)))
		if ret == 0 {
			sv.Error = fmt.Sprintf("failed in GlobalMemoryStatusEx: %s", err)
			return
		}
		memoryArray := [7]uint64{
			memoryStatus.AvailPhys,
			memoryStatus.TotalPhys,
			memoryStatus.TotalPhys - memoryStatus.AvailPhys,
			memoryStatus.TotalPageFile,
			memoryStatus.AvailPageFile,
			memoryStatus.TotalVirtual,
			memoryStatus.AvailVirtual,
		}
		// sv.Memory.Free = memoryStatus.AvailPhys
		// sv.Memory.Total = memoryStatus.TotalPhys
		// sv.Memory.Used = sv.Memory.Total - sv.Memory.Free
		// sv.Memory.PageFileTotal = memoryStatus.TotalPageFile
		// sv.Memory.PageFileFree = memoryStatus.AvailPageFile
		// sv.Memory.VirtualTotal = memoryStatus.TotalVirtual
		// sv.Memory.VirtualFree = memoryStatus.AvailVirtual

		  revisedMemArray := pkg.ByteToGiga(memoryArray[:])

		// statusArray := [7]any {
		// 	sv.Memory.Free,
		// 	sv.Memory.Total,
		// 	sv.Memory.Used,
		// 	sv.Memory.PageFileTotal,
		// 	sv.Memory.PageFileFree,
		// 	sv.Memory.VirtualTotal,
		// 	sv.Memory.VirtualFree,
		// }

		// if len(revisedMemArray) == len(statusArray){
		// 	for i := 0; i < len(memoryArray); i++ {
		// 		statusArray[i] = revisedMemArray[i]
		// 	}
		// }

		sv.Memory.Free = revisedMemArray[0]
		sv.Memory.Total = revisedMemArray[1]
		sv.Memory.Used = revisedMemArray[2]
		sv.Memory.PageFileTotal = revisedMemArray[3]
		sv.Memory.PageFileFree = revisedMemArray[4]
		sv.Memory.VirtualTotal = revisedMemArray[5]
		sv.Memory.VirtualFree =revisedMemArray[6]
	}
}

type memoryStatusEx struct {
	Length               uint32
	MemoryLoad           uint32
	TotalPhys            uint64 
	AvailPhys            uint64
	TotalPageFile        uint64
	AvailPageFile        uint64
	TotalVirtual         uint64
	AvailVirtual         uint64
	AvailExtendedVirtual uint64
}
