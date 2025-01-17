package types

func NewGenesisState(assets []Asset, pairs []Pair, params Params) *GenesisState {
	return &GenesisState{
		Assets:  assets,
		Pairs:   pairs,
		Params:  params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		nil,
		nil,
		DefaultParams(),
	)
}

func ValidateGenesis(_ *GenesisState) error {
	return nil
}
