package transcoder

import (
	"time"

	"github.com/sonata-labs/sonata/media"
	"github.com/sonata-labs/sonata/store/localstore"
	"github.com/sonata-labs/sonata/types/module"
	"go.uber.org/zap"
)

const (
	MaxWorkers = 10
)

// Transcoder walks through the local store for media files and transcodes them to the desired format.
// It uses a worker pool to transcode the media files in parallel.
type Transcoder struct {
	*module.BaseModule

	localStore *localstore.LocalStore
	encoder    *media.MediaEncoder
}

var _ module.Module = (*Transcoder)(nil)

func (t *Transcoder) Name() string {
	return "transcoder"
}

func NewTranscoder(logger *zap.Logger, localStore *localstore.LocalStore) (*Transcoder, error) {
	encoder, err := media.NewMediaEncoder(MaxWorkers)
	if err != nil {
		return nil, err
	}

	transcoder := &Transcoder{
		encoder:    encoder,
		localStore: localStore,
	}

	transcoder.BaseModule = module.NewBaseModule(logger.Named(transcoder.Name()))
	return transcoder, nil
}

func (t *Transcoder) Start() error {
	t.AwaitStartupDeps()
	t.Logger.Info("starting")

	for {
		// TODO: walk through the local store for media files and transcode them to the desired format
		select {
		case <-t.Stopped():
			return nil
		default:
		}
		time.Sleep(1 * time.Second)
	}
}

func (t *Transcoder) Stop() error {
	t.AwaitShutdownDeps()
	t.Logger.Info("stopping")
	return nil
}
