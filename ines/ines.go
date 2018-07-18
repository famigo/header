package ines

type (
	//PRG defines the number of 16KB PRG ROM banks in a NES ROM
	PRG byte

	//CHR defines the number of 8KB CHR ROM banks in a NES ROM
	CHR byte

	//MAP defines the Mappers number of a NES ROM
	//
	//For a list of available mappers, see https://wiki.nesdev.com/w/index.php/Category:INES_Mappers
	MAP byte

	//MIR defines the mirroing mode of a NES ROM
	//
	//Horizontal = 0
	//
	//Vertical = 1
	MIR byte
)

//Default value
const (
	DefaultPRG = PRG(2)
	DefaultCHR = CHR(1)
	DefaultMAP = MAP(0)
	DefaultMIR = MIR(1)
)
