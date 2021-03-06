// this file was generated by gomacro command: import _b "debug/pe"
// DO NOT EDIT! Any change will be lost when the file is re-generated

// +build go1.11

package go1_11

import (
	. "reflect"
	pe "debug/pe"
)

// reflection: allow interpreted code to import "debug/pe"
func init() {
	Packages["debug/pe"] = Package{
	Binds: map[string]Value{
		"COFFSymbolSize":	ValueOf(pe.COFFSymbolSize),
		"IMAGE_DIRECTORY_ENTRY_ARCHITECTURE":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_ARCHITECTURE),
		"IMAGE_DIRECTORY_ENTRY_BASERELOC":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_BASERELOC),
		"IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT),
		"IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR),
		"IMAGE_DIRECTORY_ENTRY_DEBUG":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_DEBUG),
		"IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT),
		"IMAGE_DIRECTORY_ENTRY_EXCEPTION":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_EXCEPTION),
		"IMAGE_DIRECTORY_ENTRY_EXPORT":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_EXPORT),
		"IMAGE_DIRECTORY_ENTRY_GLOBALPTR":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_GLOBALPTR),
		"IMAGE_DIRECTORY_ENTRY_IAT":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_IAT),
		"IMAGE_DIRECTORY_ENTRY_IMPORT":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_IMPORT),
		"IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG),
		"IMAGE_DIRECTORY_ENTRY_RESOURCE":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_RESOURCE),
		"IMAGE_DIRECTORY_ENTRY_SECURITY":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_SECURITY),
		"IMAGE_DIRECTORY_ENTRY_TLS":	ValueOf(pe.IMAGE_DIRECTORY_ENTRY_TLS),
		"IMAGE_FILE_MACHINE_AM33":	ValueOf(pe.IMAGE_FILE_MACHINE_AM33),
		"IMAGE_FILE_MACHINE_AMD64":	ValueOf(pe.IMAGE_FILE_MACHINE_AMD64),
		"IMAGE_FILE_MACHINE_ARM":	ValueOf(pe.IMAGE_FILE_MACHINE_ARM),
		"IMAGE_FILE_MACHINE_ARM64":	ValueOf(pe.IMAGE_FILE_MACHINE_ARM64),
		"IMAGE_FILE_MACHINE_EBC":	ValueOf(pe.IMAGE_FILE_MACHINE_EBC),
		"IMAGE_FILE_MACHINE_I386":	ValueOf(pe.IMAGE_FILE_MACHINE_I386),
		"IMAGE_FILE_MACHINE_IA64":	ValueOf(pe.IMAGE_FILE_MACHINE_IA64),
		"IMAGE_FILE_MACHINE_M32R":	ValueOf(pe.IMAGE_FILE_MACHINE_M32R),
		"IMAGE_FILE_MACHINE_MIPS16":	ValueOf(pe.IMAGE_FILE_MACHINE_MIPS16),
		"IMAGE_FILE_MACHINE_MIPSFPU":	ValueOf(pe.IMAGE_FILE_MACHINE_MIPSFPU),
		"IMAGE_FILE_MACHINE_MIPSFPU16":	ValueOf(pe.IMAGE_FILE_MACHINE_MIPSFPU16),
		"IMAGE_FILE_MACHINE_POWERPC":	ValueOf(pe.IMAGE_FILE_MACHINE_POWERPC),
		"IMAGE_FILE_MACHINE_POWERPCFP":	ValueOf(pe.IMAGE_FILE_MACHINE_POWERPCFP),
		"IMAGE_FILE_MACHINE_R4000":	ValueOf(pe.IMAGE_FILE_MACHINE_R4000),
		"IMAGE_FILE_MACHINE_SH3":	ValueOf(pe.IMAGE_FILE_MACHINE_SH3),
		"IMAGE_FILE_MACHINE_SH3DSP":	ValueOf(pe.IMAGE_FILE_MACHINE_SH3DSP),
		"IMAGE_FILE_MACHINE_SH4":	ValueOf(pe.IMAGE_FILE_MACHINE_SH4),
		"IMAGE_FILE_MACHINE_SH5":	ValueOf(pe.IMAGE_FILE_MACHINE_SH5),
		"IMAGE_FILE_MACHINE_THUMB":	ValueOf(pe.IMAGE_FILE_MACHINE_THUMB),
		"IMAGE_FILE_MACHINE_UNKNOWN":	ValueOf(pe.IMAGE_FILE_MACHINE_UNKNOWN),
		"IMAGE_FILE_MACHINE_WCEMIPSV2":	ValueOf(pe.IMAGE_FILE_MACHINE_WCEMIPSV2),
		"NewFile":	ValueOf(pe.NewFile),
		"Open":	ValueOf(pe.Open),
	}, Types: map[string]Type{
		"COFFSymbol":	TypeOf((*pe.COFFSymbol)(nil)).Elem(),
		"DataDirectory":	TypeOf((*pe.DataDirectory)(nil)).Elem(),
		"File":	TypeOf((*pe.File)(nil)).Elem(),
		"FileHeader":	TypeOf((*pe.FileHeader)(nil)).Elem(),
		"FormatError":	TypeOf((*pe.FormatError)(nil)).Elem(),
		"ImportDirectory":	TypeOf((*pe.ImportDirectory)(nil)).Elem(),
		"OptionalHeader32":	TypeOf((*pe.OptionalHeader32)(nil)).Elem(),
		"OptionalHeader64":	TypeOf((*pe.OptionalHeader64)(nil)).Elem(),
		"Reloc":	TypeOf((*pe.Reloc)(nil)).Elem(),
		"Section":	TypeOf((*pe.Section)(nil)).Elem(),
		"SectionHeader":	TypeOf((*pe.SectionHeader)(nil)).Elem(),
		"SectionHeader32":	TypeOf((*pe.SectionHeader32)(nil)).Elem(),
		"StringTable":	TypeOf((*pe.StringTable)(nil)).Elem(),
		"Symbol":	TypeOf((*pe.Symbol)(nil)).Elem(),
	}, Untypeds: map[string]string{
		"COFFSymbolSize":	"int:18",
		"IMAGE_DIRECTORY_ENTRY_ARCHITECTURE":	"int:7",
		"IMAGE_DIRECTORY_ENTRY_BASERELOC":	"int:5",
		"IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT":	"int:11",
		"IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR":	"int:14",
		"IMAGE_DIRECTORY_ENTRY_DEBUG":	"int:6",
		"IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT":	"int:13",
		"IMAGE_DIRECTORY_ENTRY_EXCEPTION":	"int:3",
		"IMAGE_DIRECTORY_ENTRY_EXPORT":	"int:0",
		"IMAGE_DIRECTORY_ENTRY_GLOBALPTR":	"int:8",
		"IMAGE_DIRECTORY_ENTRY_IAT":	"int:12",
		"IMAGE_DIRECTORY_ENTRY_IMPORT":	"int:1",
		"IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG":	"int:10",
		"IMAGE_DIRECTORY_ENTRY_RESOURCE":	"int:2",
		"IMAGE_DIRECTORY_ENTRY_SECURITY":	"int:4",
		"IMAGE_DIRECTORY_ENTRY_TLS":	"int:9",
		"IMAGE_FILE_MACHINE_AM33":	"int:467",
		"IMAGE_FILE_MACHINE_AMD64":	"int:34404",
		"IMAGE_FILE_MACHINE_ARM":	"int:448",
		"IMAGE_FILE_MACHINE_ARM64":	"int:43620",
		"IMAGE_FILE_MACHINE_EBC":	"int:3772",
		"IMAGE_FILE_MACHINE_I386":	"int:332",
		"IMAGE_FILE_MACHINE_IA64":	"int:512",
		"IMAGE_FILE_MACHINE_M32R":	"int:36929",
		"IMAGE_FILE_MACHINE_MIPS16":	"int:614",
		"IMAGE_FILE_MACHINE_MIPSFPU":	"int:870",
		"IMAGE_FILE_MACHINE_MIPSFPU16":	"int:1126",
		"IMAGE_FILE_MACHINE_POWERPC":	"int:496",
		"IMAGE_FILE_MACHINE_POWERPCFP":	"int:497",
		"IMAGE_FILE_MACHINE_R4000":	"int:358",
		"IMAGE_FILE_MACHINE_SH3":	"int:418",
		"IMAGE_FILE_MACHINE_SH3DSP":	"int:419",
		"IMAGE_FILE_MACHINE_SH4":	"int:422",
		"IMAGE_FILE_MACHINE_SH5":	"int:424",
		"IMAGE_FILE_MACHINE_THUMB":	"int:450",
		"IMAGE_FILE_MACHINE_UNKNOWN":	"int:0",
		"IMAGE_FILE_MACHINE_WCEMIPSV2":	"int:361",
	}, 
	}
}
