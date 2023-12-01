package mylib

type PanicStruct struct {
	anField string
}

func (p *PanicStruct) OhnoIMayPanic() *string {
	//nolint:nilaway
	return &p.anField
}

func MakeANilPanicStruct() *PanicStruct {
	//nolint:nilaway
    return nil
}
