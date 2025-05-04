package internal

const (
	Asm1 = "asm1"
	Asm2 = "asm2"
	Asm3 = "asm3"
)

var FundToUrlPostfix = map[string]string{
	Asm1: "ASM",
	Asm2: "ASW",
	Asm3: "AS1M",
}

const (
	Tngd  = "tngd"
	Boost = "boost"
	Fpx   = "fpx"
)

var AllPaymentMethods = []string{
	Tngd,
	Boost,
	Fpx,
}
