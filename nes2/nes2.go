package nes2

type (
	//CHRRAM defines the amount of CHR RAM used by a NES ROM
	CHRRAM byte

	//PRGRAM defines the amount of PRG RAM used by a NES ROM
	PRGRAM byte

	//SUB defines the submapper number of NES ROM
	SUB byte

	//TV defines the region system of NES ROM
	//
	//NTSC = 0
	//
	//PAL = 1
	//
	//Both = 2
	TV byte

	//VS sets the NES ROM to use the Vs. Unisystem
	VS byte

	//BRAM defines the amount of battery backed PRG RAM in NES ROM
	BRAM byte

	//CHRBRAM defines the amount of battery backed CHR RAM in NES ROM
	CHRBRAM byte
)

//Default value
const (
	DefaultCHRRAM  = CHRRAM(0)
	DefaultPRGRAM  = PRGRAM(0)
	DefaultSUB     = SUB(0)
	DefaultTV      = TV(0)
	DefaultVS      = VS(0)
	DefaultBRAM    = BRAM(0)
	DefaultCHRBRAM = CHRBRAM(0)
)
