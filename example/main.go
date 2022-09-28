package main

import (
	"bufio"
	"fmt"
	"github.com/timwhitez/memmod"
	"os"
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

	if len(os.Args) == 3 {
		p, _ := mod.ProcAddressByName(funcs)
		fmt.Printf("%s: 0x%x\n", funcs, p)
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		p, _ = mod1.ProcAddressByName(funcs)
		fmt.Printf("Syscall_%s: 0x%x\n", funcs, p)
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	mod.Free()
	mod1.Free()
}
