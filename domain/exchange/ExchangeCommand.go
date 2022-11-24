package exchange

type ExportCommand struct {
	ShopeCode   string
	ProductCode string
	Quantity    int32
	ExportTo    string
}
type ImportCommand struct {
	ShopeCode   string
	ProductCode string
	Quantity    int32
	ImportFrom  string
}
