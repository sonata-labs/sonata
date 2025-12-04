package media

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
)

const (
	MinFFmpegMajor = 7
)

type MediaEncoder struct {
	sem chan struct{} // worker pool
}

func NewMediaEncoder(maxWorkers int) (*MediaEncoder, error) {
	if maxWorkers <= 0 {
		maxWorkers = 1
	}
	if err := ensureFFmpegVersion(); err != nil {
		return nil, err
	}
	return &MediaEncoder{
		sem: make(chan struct{}, maxWorkers),
	}, nil
}

func (m *MediaEncoder) withWorker(fn func() error) error {
	m.sem <- struct{}{}
	defer func() { <-m.sem }()
	return fn()
}

// Transcodes any audio input into untouched FLAC.
func (m *MediaEncoder) EncodeAudio(ctx context.Context, in io.Reader, out io.Writer) error {
	return m.withWorker(func() error {
		return runFFmpegStream(ctx, in, out,
			"-i", "pipe:0",
			"-c:a", "flac",
			"-compression_level", "5",
			"-f", "flac",
			"pipe:1",
		)
	})
}

// Transcodes any image input into untouched PNG.
func (m *MediaEncoder) EncodeImage(ctx context.Context, in io.Reader, out io.Writer) error {
	return m.withWorker(func() error {
		return runFFmpegStream(ctx, in, out,
			"-i", "pipe:0",
			"-c:v", "png",
			"-f", "png",
			"pipe:1",
		)
	})
}

// runFFmpegStream:
// - streams caller input → ffmpeg stdin
// - streams ffmpeg stdout → caller output
// - stderr captured for error reporting
func runFFmpegStream(ctx context.Context, in io.Reader, out io.Writer, args ...string) error {
	cmd := exec.CommandContext(ctx, "ffmpeg",
		"-hide_banner",
		"-loglevel", "error", // keep stderr clean
	)
	cmd.Args = append(cmd.Args, args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// streaming pipes
	cmd.Stdin = in
	cmd.Stdout = out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg: %w, stderr: %s", err, stderr.String())
	}

	return nil
}

func ensureFFmpegVersion() error {
	out, err := exec.Command("ffmpeg", "-version").Output()
	if err != nil {
		return fmt.Errorf("ffmpeg not found: %w", err)
	}

	// Example first line: "ffmpeg version 7.0.1 ..."
	var major int
	_, scanErr := fmt.Sscanf(string(out), "ffmpeg version %d", &major)
	if scanErr != nil {
		return fmt.Errorf("unexpected ffmpeg version parse: %w", scanErr)
	}

	if major < MinFFmpegMajor {
		return fmt.Errorf("ffmpeg major version %d < required %d", major, MinFFmpegMajor)
	}

	return nil
}
