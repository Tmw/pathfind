package arena

type CellType uint8

const (
	CellTypeWalkable CellType = iota
	CellTypeNonWalkable
	CellTypeStart
	CellTypeFinish
	TyileTypePath
)

func (t CellType) String() string {
	switch t {
	case CellTypeNonWalkable:
		return SymbolNonWalkable

	case CellTypeStart:
		return SymbolStart

	case CellTypeFinish:
		return SymbolFinish

	case CellTypeWalkable:
		return SymbolWalkable

	default:
		return SymbolWalkable
	}
}

func symbolToType(i string) CellType {
	switch i {
	case SymbolNonWalkable:
		return CellTypeNonWalkable

	case SymbolStart:
		return CellTypeStart

	case SymbolFinish:
		return CellTypeFinish

	case SymbolWalkable:
		return CellTypeWalkable

	default:
		return CellTypeWalkable
	}
}
