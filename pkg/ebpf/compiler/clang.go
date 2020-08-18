// +build linux_bpf

package compiler

/*
#cgo LDFLAGS: -lclangTooling -lclangCodeGen -lclangFrontend -lclangFrontendTool -lclangDriver -lclangSerialization -lclangParse -lclangSema -lclangStaticAnalyzerFrontend -lclangStaticAnalyzerCheckers -lclangStaticAnalyzerCore -lclangAnalysis -lclangARCMigrate -lclangRewrite -lclangRewriteFrontend -lclangEdit -lclangAST -lclangLex -lclangBasic -lclang -lclangASTMatchers
#cgo CPPFLAGS: -I/usr/include -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS

#include <stdlib.h>
#include "shim.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"unsafe"

	"golang.org/x/sys/unix"
)

type EBPFCompiler struct {
	compiler *C.struct_bpf_compiler

	verbose       bool
	defaultCflags []string
}

func (e *EBPFCompiler) CompileToObjectFile(inputFile, outputFile string, cflags ...string) error {
	inputC := C.CString(inputFile)
	defer C.free(unsafe.Pointer(inputC))

	outputC := C.CString(outputFile)
	defer C.free(unsafe.Pointer(outputC))

	cflagsC := make([]*C.char, len(e.defaultCflags)+len(cflags)+1)
	for i, cflag := range e.defaultCflags {
		cflagsC[i] = C.CString(cflag)
	}
	for i, cflag := range cflags {
		cflagsC[i] = C.CString(cflag)
	}
	cflagsC[len(cflagsC)-1] = nil

	defer func() {
		for _, cflag := range cflagsC {
			if cflag != nil {
				C.free(unsafe.Pointer(cflag))
			}
		}
	}()

	verboseC := C.char(0)
	if e.verbose {
		verboseC = 1
	}

	// TODO I don't think there is a guarantee cflagsC will be laid out in a continuous memory segment
	if err := C.bpf_compile_to_object_file(e.compiler, inputC, outputC, (**C.char)(&cflagsC[0]), verboseC); err != 0 {
		errs := C.GoString(C.bpf_compiler_get_errors(e.compiler))
		return errors.New(errs)
	}

	return nil
}

func (e *EBPFCompiler) Close() {
	runtime.SetFinalizer(e, nil)
	C.delete_bpf_compiler(e.compiler)
	e.compiler = nil
}

func kernelTarget() (string, error) {
	var uname unix.Utsname
	err := unix.Uname(&uname)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(uname.Release[:]), "\x00"), nil
}

func kernelArch() string {
	switch runtime.GOARCH {
	case "386", "amd64":
		return "x86"
	case "arm", "armbe":
		return "arm"
	case "arm64", "arm64be":
		return "arm64"
	case "mips", "mipsle", "mips64", "mips64le":
		return "mips"
	case "ppc", "ppc64", "ppc64le":
		return "powerpc"
	case "riscv", "riscv64":
		return "riscv"
	case "s390", "s390x":
		return "s390"
	case "sparc", "sparc64":
		return "sparc64"
	default:
		return ""
	}
}

func NewEBPFCompiler(verbose bool) (*EBPFCompiler, error) {
	releaseName, err := kernelTarget()
	if err != nil {
		return nil, err
	}
	arch := kernelArch()

	var defaultFlags []string
	subdirs := []string{
		"include",
		"include/uapi",
		"include/generated/uapi",
	}
	cDir, err := filepath.Abs("../c")
	if err != nil {
		return nil, err
	}
	defaultFlags = append(defaultFlags, "-include", path.Join(cDir, "asm_goto_workaround.h"))

	base := path.Join("/usr/src", fmt.Sprintf("linux-headers-%s", releaseName))
	for _, d := range subdirs {
		defaultFlags = append(defaultFlags, "-isystem", path.Join(base, d))
		defaultFlags = append(defaultFlags, "-isystem", path.Join(base, "arch", arch, d))
	}

	ebpfCompiler := &EBPFCompiler{
		compiler:      C.new_bpf_compiler(),
		defaultCflags: defaultFlags,
		verbose:       verbose,
	}

	runtime.SetFinalizer(ebpfCompiler, func(e *EBPFCompiler) {
		e.Close()
	})

	return ebpfCompiler, nil
}
