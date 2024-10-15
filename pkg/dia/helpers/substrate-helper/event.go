package bifrosthelper

import (
	"fmt"

	gsrpc "github.com/didaunesp/no-signature-go-substrate-rpc-client-v4"
	"github.com/didaunesp/no-signature-go-substrate-rpc-client-v4/registry/parser"
	"github.com/didaunesp/no-signature-go-substrate-rpc-client-v4/registry/retriever"
	"github.com/didaunesp/no-signature-go-substrate-rpc-client-v4/registry/state"

	"github.com/didaunesp/no-signature-go-substrate-rpc-client-v4/types"

	"github.com/sirupsen/logrus"
)

type SubstrateEventHelper struct {
	logger *logrus.Entry
	API    *gsrpc.SubstrateAPI
}

func NewSubstrateEventHelper(nodeURL string, logger *logrus.Entry) (*SubstrateEventHelper, error) {
	api, err := gsrpc.NewSubstrateAPI(nodeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Substrate node: %v", err)
	}
	return &SubstrateEventHelper{API: api, logger: logger}, nil
}

func (s *SubstrateEventHelper) ListenForSpecificBlock(blockNumber uint64, callback func([]*parser.Event)) error {
	blockHash, err := s.API.RPC.Chain.GetBlockHash(blockNumber)
	if err != nil {
		message := fmt.Sprintf("Failed to fetch block hash: %v", err)
		s.logger.Errorf(message, err)
		return fmt.Errorf(message)
	}

	events, err := s.DecodeEvents(blockHash)
	if err != nil {
		message := fmt.Sprintf("Failed to decode events: %v", err)
		s.logger.Errorf(message, err)
		return err
	}

	callback(events)

	return nil
}

// TODO: MOVE THIS TO COMMON PARACHAIN HELPER
// DecodeEvents fetches and decodes events for a specific block hash using CustomEventRecords
// func (s *SubstrateEventHelper) DecodeEvents(blockHash types.Hash) (*CustomEventRecords, error) {

type EventStableAssetTokenSwapped struct {
	Phase           types.Phase     // The phase of the event (applies to all events)
	Swapper         types.AccountID `json:"swapper"`           // Account ID of the swapper
	PoolID          types.U32       `json:"pool_id"`           // Pool ID
	A               types.U128      `json:"a"`                 // Arbitrary value 'a'
	InputAsset      AssetID         `json:"input_asset"`       // Input asset (custom type bifrost_primitives:currency:CurrencyId)
	OutputAsset     AssetID         `json:"output_asset"`      // Output asset (custom type bifrost_primitives:currency:CurrencyId)
	InputAmount     types.U128      `json:"input_amount"`      // Input amount
	MinOutputAmount types.U128      `json:"min_output_amount"` // Minimum output amount
	Balances        []types.U128    `json:"balances"`          // Balances (Vec<Balance>)
	TotalSupply     types.U128      `json:"total_supply"`      // Total supply
	OutputAmount    types.U128      `json:"output_amount"`     // Output amount
	Topics          []types.Hash    // Event topics (applies to all events)
}

// AssetID represents the custom bifrost_primitives:currency:CurrencyId type.
type AssetID struct {
	VToken2 types.U32 `json:"VToken2,omitempty"` // Token variant 2
	Token2  types.U32 `json:"Token2,omitempty"`  // Another token variant
}

type CustomEventRecords struct {
	StableAsset_TokenSwapped []EventStableAssetTokenSwapped
}

func (s *SubstrateEventHelper) DecodeEvents(blockHash types.Hash) ([]*parser.Event, error) {
	r, err := retriever.NewDefaultEventRetriever(state.NewEventProvider(s.API.RPC.State), s.API.RPC.State)

	if err != nil {
		return nil, fmt.Errorf("Couldn't create event retriever: %s", err)
	}

	events, err := r.GetEvents(blockHash)

	if err != nil {
		return nil, fmt.Errorf("Couldn't retrieve events for block hash %d: %s\n", blockHash, err)
	}

	s.logger.Infof("Found %d events\n", len(events))

	return events, nil
}

// ListenForNewBlocks listens for new blocks and continuously decodes events.
func (s *SubstrateEventHelper) ListenForNewBlocks(callback func([]*parser.Event)) error {
	sub, err := s.API.RPC.Chain.SubscribeNewHeads()
	if err != nil {
		return fmt.Errorf("failed to subscribe to new heads: %v", err)
	}

	for {
		head := <-sub.Chan()

		blockHash, err := s.API.RPC.Chain.GetBlockHash(uint64(head.Number))
		if err != nil {
			s.logger.Errorf("Failed to fetch block hash: %v\n", err)
			continue
		}
		fmt.Printf("\nNew block detected! Block number: %v, Block hash: %v\n", head.Number, blockHash.Hex())

		events, err := s.DecodeEvents(blockHash)
		if err != nil {
			s.logger.Errorf("Failed to decode events: %v\n", err)
			continue
		}

		callback(events)
	}
}
