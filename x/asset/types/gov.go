package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeAddPool = "AddPool"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeAddPool)
	govtypes.RegisterProposalTypeCodec(&AddPoolProposal{}, fmt.Sprintf("comdex/%s", ProposalTypeAddPool))
}

var (
	_ govtypes.Content = (*AddPoolProposal)(nil)
)

func (m *AddPoolProposal) GetTitle() string       { return m.Title }
func (m *AddPoolProposal) GetDescription() string { return m.Description }
func (m *AddPoolProposal) ProposalRoute() string  { return RouterKey }
func (m *AddPoolProposal) ProposalType() string   { return ProposalTypeAddPool }

func (m *AddPoolProposal) ValidateBasic() error {
	if err := govtypes.ValidateAbstract(m); err != nil {
		return err
	}
	if m.DenomIn == "" {
		return fmt.Errorf("denom_in cannot be empty")
	}
	if err := sdk.ValidateDenom(m.DenomIn); err != nil {
		return errors.Wrapf(err, "invalid denom_in %s", m.DenomIn)
	}
	if m.DenomOut == "" {
		return fmt.Errorf("denom_out cannot be empty")
	}
	if err := sdk.ValidateDenom(m.DenomOut); err != nil {
		return errors.Wrapf(err, "invalid denom_out %s", m.DenomOut)
	}
	if m.LiquidationRatio.IsNil() {
		return fmt.Errorf("liquidation_ratio cannot be nil")
	}
	if m.LiquidationRatio.IsNegative() {
		return fmt.Errorf("liquidation_ratio cannot be negative")
	}

	return nil
}
