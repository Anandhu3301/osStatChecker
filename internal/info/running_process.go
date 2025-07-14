package info

import (
	"fmt"
	"strconv"

	"github.com/Anandhu3301/osStatChecker/pkg"
	"golang.org/x/sys/windows"
)

const processEntrySize = 568

func Get_Process_List() {
	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if e != nil {
		panic(e)
	}
	p := windows.ProcessEntry32{Size: processEntrySize}
	var process []string
	var uniqueProcess []string;

	// s := windows.UTF16ToString(p.ExeFile[:])
	for {
		process = append(process, windows.UTF16ToString(p.ExeFile[:]) + "--->"+ strconv.FormatUint(uint64(p.ProcessID),10))
		e := windows.Process32Next(h, &p)
		if e != nil {
			break
		}
	}

	uniqueProcess = pkg.RemoveDuplicateStr(process);
	fmt.Println(uniqueProcess);
}
