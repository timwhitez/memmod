package main

import (
	"bufio"
	"fmt"
	"github.com/timwhitez/memmod"
	"os"
	"syscall"
)

func MemLD(moduleN string) (*memmod.Module, error) {
	f, e := os.ReadFile(moduleN)
	if e != nil {
		return nil, e
	}

	module, e := memmod.LoadLibrary(f)
	if e != nil {
		return nil, e
	}

	return module, nil
}

func main() {

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	calcf := "C:\\windows\\system32\\calc.exe"

	calc, e := MemLD(calcf)
	if e != nil {
		panic(e)
	}

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	base := calc.ModuleBase()

	entry := calc.EntryPoint()

	fmt.Printf("Base:0x%x\n", base)
	fmt.Printf("Entry:0x%x\n", entry)

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	syscall.SyscallN(entry)

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	calc.Free()

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
