package account

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"connectrpc.com/connect"
	"github.com/sonata-labs/sonata/config"
	v1 "github.com/sonata-labs/sonata/gen/api/v1"
	"github.com/sonata-labs/sonata/gen/api/v1/v1connect"
	"github.com/sonata-labs/sonata/types/module"
	"go.uber.org/zap"
)

type AccountService struct {
	*module.BaseModule
	config *config.Config
}

func (a *AccountService) Name() string {
	return "account"
}

var _ module.Module = (*AccountService)(nil)

func (a *AccountService) GetAccount(context.Context, *connect.Request[v1.GetAccountRequest]) (*connect.Response[v1.GetAccountResponse], error) {
	panic("unimplemented")
}

var _ v1connect.AccountHandler = (*AccountService)(nil)

func NewAccountService(config *config.Config, logger *zap.Logger) *AccountService {
	svc := &AccountService{config: config}
	svc.BaseModule = module.NewBaseModule(logger.Named(svc.Name()))
	return svc
}

// ABCI++ Callbacks

func (a *AccountService) CheckTx(ctx context.Context, req *abcitypes.CheckTxRequest) (*abcitypes.CheckTxResponse, error) {
	a.Logger.Info("checking tx")
	return &abcitypes.CheckTxResponse{}, nil
}

func (a *AccountService) PrepareProposal(ctx context.Context, req *abcitypes.PrepareProposalRequest) (*abcitypes.PrepareProposalResponse, error) {
	a.Logger.Info("preparing proposal")
	return &abcitypes.PrepareProposalResponse{}, nil
}

func (a *AccountService) ProcessProposal(ctx context.Context, req *abcitypes.ProcessProposalRequest) (*abcitypes.ProcessProposalResponse, error) {
	a.Logger.Info("processing proposal")
	return &abcitypes.ProcessProposalResponse{}, nil
}

func (a *AccountService) FinalizeBlock(ctx context.Context, req *abcitypes.FinalizeBlockRequest) (*abcitypes.FinalizeBlockResponse, error) {
	a.Logger.Info("finalizing block")
	return &abcitypes.FinalizeBlockResponse{}, nil
}
