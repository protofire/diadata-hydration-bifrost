// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package whirlpool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// OpenPositionWithMetadata is the `openPositionWithMetadata` instruction.
type OpenPositionWithMetadata struct {
	Bumps          *OpenPositionWithMetadataBumps
	TickLowerIndex *int32
	TickUpperIndex *int32

	// [0] = [WRITE, SIGNER] funder
	//
	// [1] = [] owner
	//
	// [2] = [WRITE] position
	//
	// [3] = [WRITE, SIGNER] positionMint
	//
	// [4] = [WRITE] positionMetadataAccount
	//
	// [5] = [WRITE] positionTokenAccount
	//
	// [6] = [] whirlpool
	//
	// [7] = [] tokenProgram
	//
	// [8] = [] systemProgram
	//
	// [9] = [] rent
	//
	// [10] = [] associatedTokenProgram
	//
	// [11] = [] metadataProgram
	//
	// [12] = [] metadataUpdateAuth
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewOpenPositionWithMetadataInstructionBuilder creates a new `OpenPositionWithMetadata` instruction builder.
func NewOpenPositionWithMetadataInstructionBuilder() *OpenPositionWithMetadata {
	nd := &OpenPositionWithMetadata{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetBumps sets the "bumps" parameter.
func (inst *OpenPositionWithMetadata) SetBumps(bumps OpenPositionWithMetadataBumps) *OpenPositionWithMetadata {
	inst.Bumps = &bumps
	return inst
}

// SetTickLowerIndex sets the "tickLowerIndex" parameter.
func (inst *OpenPositionWithMetadata) SetTickLowerIndex(tickLowerIndex int32) *OpenPositionWithMetadata {
	inst.TickLowerIndex = &tickLowerIndex
	return inst
}

// SetTickUpperIndex sets the "tickUpperIndex" parameter.
func (inst *OpenPositionWithMetadata) SetTickUpperIndex(tickUpperIndex int32) *OpenPositionWithMetadata {
	inst.TickUpperIndex = &tickUpperIndex
	return inst
}

// SetFunderAccount sets the "funder" account.
func (inst *OpenPositionWithMetadata) SetFunderAccount(funder ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(funder).WRITE().SIGNER()
	return inst
}

// GetFunderAccount gets the "funder" account.
func (inst *OpenPositionWithMetadata) GetFunderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *OpenPositionWithMetadata) SetOwnerAccount(owner ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner)
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *OpenPositionWithMetadata) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPositionAccount sets the "position" account.
func (inst *OpenPositionWithMetadata) SetPositionAccount(position ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(position).WRITE()
	return inst
}

// GetPositionAccount gets the "position" account.
func (inst *OpenPositionWithMetadata) GetPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPositionMintAccount sets the "positionMint" account.
func (inst *OpenPositionWithMetadata) SetPositionMintAccount(positionMint ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(positionMint).WRITE().SIGNER()
	return inst
}

// GetPositionMintAccount gets the "positionMint" account.
func (inst *OpenPositionWithMetadata) GetPositionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPositionMetadataAccountAccount sets the "positionMetadataAccount" account.
func (inst *OpenPositionWithMetadata) SetPositionMetadataAccountAccount(positionMetadataAccount ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(positionMetadataAccount).WRITE()
	return inst
}

// GetPositionMetadataAccountAccount gets the "positionMetadataAccount" account.
func (inst *OpenPositionWithMetadata) GetPositionMetadataAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPositionTokenAccountAccount sets the "positionTokenAccount" account.
func (inst *OpenPositionWithMetadata) SetPositionTokenAccountAccount(positionTokenAccount ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(positionTokenAccount).WRITE()
	return inst
}

// GetPositionTokenAccountAccount gets the "positionTokenAccount" account.
func (inst *OpenPositionWithMetadata) GetPositionTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetWhirlpoolAccount sets the "whirlpool" account.
func (inst *OpenPositionWithMetadata) SetWhirlpoolAccount(whirlpool ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(whirlpool)
	return inst
}

// GetWhirlpoolAccount gets the "whirlpool" account.
func (inst *OpenPositionWithMetadata) GetWhirlpoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *OpenPositionWithMetadata) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *OpenPositionWithMetadata) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *OpenPositionWithMetadata) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *OpenPositionWithMetadata) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetRentAccount sets the "rent" account.
func (inst *OpenPositionWithMetadata) SetRentAccount(rent ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *OpenPositionWithMetadata) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *OpenPositionWithMetadata) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *OpenPositionWithMetadata) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetMetadataProgramAccount sets the "metadataProgram" account.
func (inst *OpenPositionWithMetadata) SetMetadataProgramAccount(metadataProgram ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(metadataProgram)
	return inst
}

// GetMetadataProgramAccount gets the "metadataProgram" account.
func (inst *OpenPositionWithMetadata) GetMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetMetadataUpdateAuthAccount sets the "metadataUpdateAuth" account.
func (inst *OpenPositionWithMetadata) SetMetadataUpdateAuthAccount(metadataUpdateAuth ag_solanago.PublicKey) *OpenPositionWithMetadata {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(metadataUpdateAuth)
	return inst
}

// GetMetadataUpdateAuthAccount gets the "metadataUpdateAuth" account.
func (inst *OpenPositionWithMetadata) GetMetadataUpdateAuthAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

func (inst OpenPositionWithMetadata) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_OpenPositionWithMetadata,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst OpenPositionWithMetadata) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *OpenPositionWithMetadata) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Bumps == nil {
			return errors.New("Bumps parameter is not set")
		}
		if inst.TickLowerIndex == nil {
			return errors.New("TickLowerIndex parameter is not set")
		}
		if inst.TickUpperIndex == nil {
			return errors.New("TickUpperIndex parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Funder is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Position is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PositionMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.PositionMetadataAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.PositionTokenAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Whirlpool is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.MetadataProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.MetadataUpdateAuth is not set")
		}
	}
	return nil
}

func (inst *OpenPositionWithMetadata) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("OpenPositionWithMetadata")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("         Bumps", *inst.Bumps))
						paramsBranch.Child(ag_format.Param("TickLowerIndex", *inst.TickLowerIndex))
						paramsBranch.Child(ag_format.Param("TickUpperIndex", *inst.TickUpperIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                funder", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                 owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("              position", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          positionMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      positionMetadata", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         positionToken", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("             whirlpool", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("       metadataProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("    metadataUpdateAuth", inst.AccountMetaSlice.Get(12)))
					})
				})
		})
}

func (obj OpenPositionWithMetadata) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bumps` param:
	err = encoder.Encode(obj.Bumps)
	if err != nil {
		return err
	}
	// Serialize `TickLowerIndex` param:
	err = encoder.Encode(obj.TickLowerIndex)
	if err != nil {
		return err
	}
	// Serialize `TickUpperIndex` param:
	err = encoder.Encode(obj.TickUpperIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *OpenPositionWithMetadata) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bumps`:
	err = decoder.Decode(&obj.Bumps)
	if err != nil {
		return err
	}
	// Deserialize `TickLowerIndex`:
	err = decoder.Decode(&obj.TickLowerIndex)
	if err != nil {
		return err
	}
	// Deserialize `TickUpperIndex`:
	err = decoder.Decode(&obj.TickUpperIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewOpenPositionWithMetadataInstruction declares a new OpenPositionWithMetadata instruction with the provided parameters and accounts.
func NewOpenPositionWithMetadataInstruction(
	// Parameters:
	bumps OpenPositionWithMetadataBumps,
	tickLowerIndex int32,
	tickUpperIndex int32,
	// Accounts:
	funder ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	position ag_solanago.PublicKey,
	positionMint ag_solanago.PublicKey,
	positionMetadataAccount ag_solanago.PublicKey,
	positionTokenAccount ag_solanago.PublicKey,
	whirlpool ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	metadataProgram ag_solanago.PublicKey,
	metadataUpdateAuth ag_solanago.PublicKey) *OpenPositionWithMetadata {
	return NewOpenPositionWithMetadataInstructionBuilder().
		SetBumps(bumps).
		SetTickLowerIndex(tickLowerIndex).
		SetTickUpperIndex(tickUpperIndex).
		SetFunderAccount(funder).
		SetOwnerAccount(owner).
		SetPositionAccount(position).
		SetPositionMintAccount(positionMint).
		SetPositionMetadataAccountAccount(positionMetadataAccount).
		SetPositionTokenAccountAccount(positionTokenAccount).
		SetWhirlpoolAccount(whirlpool).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetMetadataProgramAccount(metadataProgram).
		SetMetadataUpdateAuthAccount(metadataUpdateAuth)
}