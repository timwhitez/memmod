package main

import (
	"bufio"
	"fmt"
	"github.com/timwhitez/memmod"
	"os"
	"syscall"
)

func main() {
	dllfile := "MsgBox.dll"
	funcs := "MessageBoxThread"

	f, e := os.ReadFile(dllfile)
	if e != nil {
		panic(e)
	}

	mod, e := memmod.LoadLibrary(f)
	if e != nil {
		panic(e)
	}

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	mod1, e := memmod.LoadLibrarySyscall(f)
	if e != nil {
		panic(e)
	}
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	p, _ := mod.ProcAddressByName(funcs)
	fmt.Printf("%s: 0x%x\n", funcs, p)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	syscall.SyscallN(p)

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	mod.Free()

	p, _ = mod1.ProcAddressByName(funcs)
	fmt.Printf("Syscall_%s: 0x%x\n", funcs, p)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	syscall.SyscallN(p)

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	mod1.Free()
}
