package main

import (
	"bufio"
	"fmt"
	"github.com/moloch--/memmod"
	"os"
	"syscall"
)

func main() {
	dllfile := "MsgBox.x86.dll"
	funcs := "_MessageBoxThread@4"

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

	p, _ := mod.ProcAddressByName(funcs)
	fmt.Printf("%s: 0x%x\n", funcs, p)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	syscall.SyscallN(p)

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	mod.Free()
}
