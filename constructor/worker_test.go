package constructor

import (
	"context"
	"errors"
	"testing"

	mocks "github.com/coinbase/rosetta-sdk-go/mocks/constructor"
	"github.com/coinbase/rosetta-sdk-go/types"

	"github.com/stretchr/testify/assert"
)

func TestWaitMessage(t *testing.T) {
	tests := map[string]struct {
		input *FindBalanceInput

		message string
	}{
		"simple message": {
			input: &FindBalanceInput{
				MinimumBalance: &types.Amount{
					Value: "100",
					Currency: &types.Currency{
						Symbol:   "BTC",
						Decimals: 8,
					},
				},
			},
			message: `Waiting for balance {"value":"100","currency":{"symbol":"BTC","decimals":8}}`,
		},
		"message with address": {
			input: &FindBalanceInput{
				Address: "hello",
				MinimumBalance: &types.Amount{
					Value: "100",
					Currency: &types.Currency{
						Symbol:   "BTC",
						Decimals: 8,
					},
				},
			},
			message: `Waiting for balance {"value":"100","currency":{"symbol":"BTC","decimals":8}} on account {"address":"hello"}`,
		},
		"message with address and subaccount": {
			input: &FindBalanceInput{
				Address: "hello",
				SubAccount: &types.SubAccountIdentifier{
					Address: "sub hello",
				},
				MinimumBalance: &types.Amount{
					Value: "100",
					Currency: &types.Currency{
						Symbol:   "BTC",
						Decimals: 8,
					},
				},
			},
			message: `Waiting for balance {"value":"100","currency":{"symbol":"BTC","decimals":8}} on account {"address":"hello","sub_account":{"address":"sub hello"}}`,
		},
		"message with address and not address": {
			input: &FindBalanceInput{
				Address: "hello",
				NotAddress: []string{
					"good",
					"bye",
				},
				MinimumBalance: &types.Amount{
					Value: "100",
					Currency: &types.Currency{
						Symbol:   "BTC",
						Decimals: 8,
					},
				},
			},
			message: `Waiting for balance {"value":"100","currency":{"symbol":"BTC","decimals":8}} on account {"address":"hello"} != to addresses ["good","bye"]`,
		},
		"message with address and not coins": {
			input: &FindBalanceInput{
				Address: "hello",
				NotCoins: []*types.CoinIdentifier{
					{
						Identifier: "coin1",
					},
				},
				MinimumBalance: &types.Amount{
					Value: "100",
					Currency: &types.Currency{
						Symbol:   "BTC",
						Decimals: 8,
					},
				},
			},
			message: `Waiting for balance {"value":"100","currency":{"symbol":"BTC","decimals":8}} on account {"address":"hello"} != to coins [{"identifier":"coin1"}]`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.message, waitMessage(test.input))
		})
	}
}

func TestFindBalanceWorker(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		input *FindBalanceInput

		mockHelper *mocks.WorkerHelper

		output *FindBalanceOutput
		err    error
	}{
		"simple find balance with wait": {
			input: &FindBalanceInput{
				MinimumBalance: &types.Amount{
					Value: "100",
					Currency: &types.Currency{
						Symbol:   "BTC",
						Decimals: 8,
					},
				},
				Wait:       true,
				NotAddress: []string{"addr4"},
			},
			mockHelper: func() *mocks.WorkerHelper {
				helper := &mocks.WorkerHelper{}
				helper.On("AllAddresses", ctx).Return([]string{"addr2", "addr1", "addr3", "addr4"}, nil).Twice()
				helper.On("LockedAddresses", ctx).Return([]string{"addr2"}, nil).Twice()
				helper.On("Balance", ctx, &types.AccountIdentifier{
					Address:    "addr1",
					SubAccount: (*types.SubAccountIdentifier)(nil),
				}).Return([]*types.Amount{
					{
						Value: "99",
						Currency: &types.Currency{
							Symbol:   "BTC",
							Decimals: 8,
						},
					},
				}, nil).Once()
				helper.On("Balance", ctx, &types.AccountIdentifier{
					Address:    "addr3",
					SubAccount: (*types.SubAccountIdentifier)(nil),
				}).Return([]*types.Amount{
					{
						Value: "101",
						Currency: &types.Currency{
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
				}, nil).Once()
				helper.On("Balance", ctx, &types.AccountIdentifier{
					Address:    "addr1",
					SubAccount: (*types.SubAccountIdentifier)(nil),
				}).Return([]*types.Amount{
					{
						Value: "100",
						Currency: &types.Currency{
							Symbol:   "BTC",
							Decimals: 8,
						},
					},
				}, nil).Once()

				return helper
			}(),
			output: &FindBalanceOutput{
				Account: &types.AccountIdentifier{
					Address: "addr1",
				},
				Balance: &types.Amount{
					Value: "100",
					Currency: &types.Currency{
						Symbol:   "BTC",
						Decimals: 8,
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			worker := NewWorker(test.mockHelper)
			output, err := worker.FindBalanceWorker(ctx, types.PrintStruct(test.input))
			if test.err != nil {
				assert.Equal(t, "", output)
				assert.True(t, errors.Is(err, test.err))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, types.PrintStruct(test.output), output)
			}

			test.mockHelper.AssertExpectations(t)
		})
	}
}
