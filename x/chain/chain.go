package chain

import (
	"context"
	"fmt"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"connectrpc.com/connect"
	"github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/rpc/client/local"
	"github.com/sonata-labs/sonata/common"
	"github.com/sonata-labs/sonata/config"
	v1 "github.com/sonata-labs/sonata/gen/api/v1"
	"github.com/sonata-labs/sonata/gen/api/v1/v1connect"
	"github.com/sonata-labs/sonata/types/module"
	"go.uber.org/zap"
)

var _ v1connect.ChainHandler = (*ChainService)(nil)

type ChainService struct {
	*module.BaseModule

	config *config.Config
	node   *node.Node
	rpc    *local.Local
}

func (c *ChainService) Name() string {
	return "chain"
}

func (c *ChainService) GetBlock(ctx context.Context, req *connect.Request[v1.GetBlockRequest]) (*connect.Response[v1.GetBlockResponse], error) {
	height := req.Msg.Height
	if height == 0 {
		height = c.node.BlockStore().Height()
	}

	block, blockMeta := c.node.BlockStore().LoadBlock(height)
	if block == nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("block not found"))
	}

	protoBlock, err := block.ToProto()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	protoBlockMeta := blockMeta.ToProto()

	return connect.NewResponse(&v1.GetBlockResponse{Block: protoBlock, BlockMeta: protoBlockMeta}), nil
}

func (c *ChainService) GetTransaction(ctx context.Context, req *connect.Request[v1.GetTransactionRequest]) (*connect.Response[v1.GetTransactionResponse], error) {
	if req.Msg.TxHash == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("tx hash is required"))
	}

	txHash := common.TxHashToBytes(req.Msg.TxHash)

	res, err := c.rpc.Tx(ctx, txHash, req.Msg.Prove)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	proof := res.Proof.ToProto()
	return connect.NewResponse(&v1.GetTransactionResponse{
		TxHash:   common.TxHashToString(res.Hash),
		Height:   res.Height,
		Index:    res.Index,
		TxResult: &res.TxResult,
		Tx:       res.Tx,
		Proof:    &proof,
	}), nil
}

func (c *ChainService) SendTransaction(context.Context, *connect.Request[v1.SendTransactionRequest]) (*connect.Response[v1.SendTransactionResponse], error) {
	panic("unimplemented")
}

func NewChainService(config *config.Config, logger *zap.Logger, node *node.Node) *ChainService {
	svc := &ChainService{config: config, node: node, rpc: local.New(node)}
	svc.BaseModule = module.NewBaseModule(logger.Named(svc.Name()))
	return svc
}

// ABCI++ Callbacks

func (c *ChainService) InitChain(ctx context.Context, req *abcitypes.InitChainRequest) (*abcitypes.InitChainResponse, error) {
	c.Logger.Info("initializing chain")
	return &abcitypes.InitChainResponse{}, nil
}

func (c *ChainService) CheckTx(ctx context.Context, req *abcitypes.CheckTxRequest) (*abcitypes.CheckTxResponse, error) {
	c.Logger.Info("checking tx")
	return &abcitypes.CheckTxResponse{}, nil
}

func (c *ChainService) PrepareProposal(ctx context.Context, req *abcitypes.PrepareProposalRequest) (*abcitypes.PrepareProposalResponse, error) {
	c.Logger.Info("preparing proposal")
	return &abcitypes.PrepareProposalResponse{Txs: req.Txs}, nil
}

func (c *ChainService) ProcessProposal(ctx context.Context, req *abcitypes.ProcessProposalRequest) (*abcitypes.ProcessProposalResponse, error) {
	c.Logger.Info("processing proposal")
	return &abcitypes.ProcessProposalResponse{Status: abcitypes.PROCESS_PROPOSAL_STATUS_ACCEPT}, nil
}

func (c *ChainService) FinalizeBlock(ctx context.Context, req *abcitypes.FinalizeBlockRequest) (*abcitypes.FinalizeBlockResponse, error) {
	c.Logger.Info("finalizing block")
	return &abcitypes.FinalizeBlockResponse{}, nil
}

func (c *ChainService) Commit(ctx context.Context, req *abcitypes.CommitRequest) (*abcitypes.CommitResponse, error) {
	c.Logger.Info("committing")
	return &abcitypes.CommitResponse{}, nil
}
