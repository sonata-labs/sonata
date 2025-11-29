package storage

import (
	"context"

	"connectrpc.com/connect"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/sonata-labs/sonata/config"
	v1 "github.com/sonata-labs/sonata/gen/api/v1"
	"github.com/sonata-labs/sonata/gen/api/v1/v1connect"
	"github.com/sonata-labs/sonata/types/module"
	"go.uber.org/zap"
)

type StorageService struct {
	*module.BaseModule
	config *config.Config
}

func (s *StorageService) Name() string {
	return "storage"
}

func (s *StorageService) DownloadFile(context.Context, *connect.Request[v1.DownloadFileRequest]) (*connect.Response[v1.DownloadFileResponse], error) {
	panic("unimplemented")
}

func (s *StorageService) DownloadFileChunk(context.Context, *connect.Request[v1.DownloadFileChunkRequest], *connect.ServerStream[v1.DownloadFileChunkResponse]) error {
	panic("unimplemented")
}

func (s *StorageService) Upload(context.Context, *connect.Request[v1.UploadRequest]) (*connect.Response[v1.UploadResponse], error) {
	panic("unimplemented")
}

func (s *StorageService) UploadChunk(context.Context, *connect.Request[v1.UploadChunkRequest]) (*connect.Response[v1.UploadChunkResponse], error) {
	panic("unimplemented")
}

var _ v1connect.StorageHandler = (*StorageService)(nil)

func NewStorageService(config *config.Config, logger *zap.Logger) *StorageService {
	svc := &StorageService{config: config}
	svc.BaseModule = module.NewBaseModule(logger.Named(svc.Name()))
	return svc
}

// ABCI++ Callbacks

func (s *StorageService) CheckTx(ctx context.Context, req *abcitypes.CheckTxRequest) (*abcitypes.CheckTxResponse, error) {
	s.Logger.Info("checking tx")
	return &abcitypes.CheckTxResponse{}, nil
}

func (s *StorageService) PrepareProposal(ctx context.Context, req *abcitypes.PrepareProposalRequest) (*abcitypes.PrepareProposalResponse, error) {
	s.Logger.Info("preparing proposal")
	return &abcitypes.PrepareProposalResponse{}, nil
}

func (s *StorageService) ProcessProposal(ctx context.Context, req *abcitypes.ProcessProposalRequest) (*abcitypes.ProcessProposalResponse, error) {
	s.Logger.Info("processing proposal")
	return &abcitypes.ProcessProposalResponse{}, nil
}

func (s *StorageService) FinalizeBlock(ctx context.Context, req *abcitypes.FinalizeBlockRequest) (*abcitypes.FinalizeBlockResponse, error) {
	s.Logger.Info("finalizing block")
	return &abcitypes.FinalizeBlockResponse{}, nil
}
