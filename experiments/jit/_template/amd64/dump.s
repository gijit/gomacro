	.file	"dump.s"
	.text
	.p2align 4,,15
	.globl	Dump
	.type	Dump, @function
Dump:
	.cfi_startproc
	.byte 0x48, 0x8b, 0x7c, 0x24, 0x8, 0x48, 0x8b, 0x87, 0x0, 0x0, 0x0, 0x0, 0x48, 0x69, 0xc0, 0x2, 0x0, 0x0, 0x0, 0x48, 0x81, 0xc0, 0x3, 0x0, 0x0, 0x0, 0x48, 0x81, 0xc8, 0x4, 0x0, 0x0, 0x0, 0x48, 0x81, 0xe0, 0xfa, 0xff, 0xff, 0xff, 0x48, 0x81, 0xf0, 0x6, 0x0, 0x0, 0x0, 0x48, 0x8b, 0x8f, 0x0, 0x0, 0x0, 0x0, 0x81, 0xe1, 0x2, 0x0, 0x0, 0x0, 0x48, 0x81, 0xc9, 0x1, 0x0, 0x0, 0x0, 0x48, 0xf7, 0xf9, 0x48, 0x89, 0x87, 0x8, 0x0, 0x0, 0x0, 0xc3
	.cfi_endproc

