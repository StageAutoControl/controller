package playback

import (
	"context"
	"fmt"

	"github.com/StageAutoControl/controller/pkg/artnet"
	"github.com/StageAutoControl/controller/pkg/cntl/transport"
	"github.com/StageAutoControl/controller/pkg/cntl/waiter"
	"github.com/StageAutoControl/controller/pkg/internal/logging"
	"github.com/StageAutoControl/controller/pkg/visualizer"
)

// Process handles the playback of a single song
type Process struct {
	logger     logging.Logger
	loader     loader
	storage    storage
	params     Params
	controller artnet.Controller
	player     *Player
	cancel     context.CancelFunc
	visualizer *visualizer.Server
}

// NewProcess returns a new playback process instance
func NewProcess(loader loader, storage storage, controller artnet.Controller, visualizer *visualizer.Server) *Process {
	return &Process{
		loader:     loader,
		storage:    storage,
		controller: controller,
		visualizer: visualizer,
	}
}

// SetParams tells the playback process whether to playback a song or set list and the corresponding ID
func (p *Process) SetParams(params Params) {
	p.params = params
}

// GetParams returns the params the process is currently running with
func (p *Process) GetParams() Params {
	return p.params
}

// SetLogger sets the logger for the process
func (p *Process) SetLogger(logger logging.Logger) {
	p.logger = logger
}

// Start the process, i.e. start the player with all the collected information
func (p *Process) Start(ctx context.Context) error {
	ds, err := p.loader.Load()
	if err != nil {
		return fmt.Errorf("failed to load data from disk: %v", err)
	}

	config := &Config{}
	if err := p.storage.Read(paramsStorageKey, config); err != nil {
		return fmt.Errorf("failed to find playback config: %v", err)
	}

	cfg, err := p.parseConfig(config)
	if err != nil {
		return err
	}
	p.player = NewPlayer(p.logger, ds, cfg.writers, cfg.waiters)
	ctx, p.cancel = context.WithCancel(ctx)

	if p.params.SetList.ID != "" {
		if err := p.player.PlaySetList(ctx, p.params.SetList.ID); err != nil && err != ErrCancelled {
			return fmt.Errorf("failed to start setlist playbaack: %v", err)
		}
	} else if p.params.Song.ID != "" {
		if err := p.player.PlaySong(ctx, p.params.Song.ID); err != nil && err != ErrCancelled {
			return fmt.Errorf("failed to start song playback: %v", err)
		}
	} else {
		return ErrNoSongIDOrSetListIDGiven
	}

	// return p.Stop()
	// we don't need to explicitly stop the process when it's done as it's marked as blocking
	return nil
}

func (p *Process) parseConfig(config *Config) (*parsedConfig, error) {
	cfg := &parsedConfig{
		waiters: []Waiter{},
		writers: []TransportWriter{},
	}

	if config.TransportWriters.ArtNet.Enabled {
		aw, err := transport.NewArtNet(p.controller)
		if err != nil {
			return nil, fmt.Errorf("failed to create artnet transport writer: %v", err)
		}

		cfg.writers = append(cfg.writers, aw)
	}

	if config.TransportWriters.MIDI.Enabled {
		mw, err := transport.NewMIDI(p.logger, config.TransportWriters.MIDI.OutputDeviceID)
		if err != nil {
			return nil, fmt.Errorf("failed to create midi transport writer: %v", err)
		}

		cfg.writers = append(cfg.writers, mw)
	}

	if config.TransportWriters.Visualizer.Enabled {
		cfg.writers = append(cfg.writers, p.visualizer)
	}

	if config.Waiters.Audio.Enabled {
		cfg.waiters = append(cfg.waiters, waiter.NewAudio(p.logger, config.Waiters.Audio.Threshold))
	}

	return cfg, nil
}

// Stop the process, i.e. cancel the playback context
func (p *Process) Stop() error {
	if p.cancel != nil {
		p.cancel()
	}
	p.player = nil

	return nil
}

// Blocking returns true if calling Start() is a blocking operation and the process is stopped after start returned
func (p *Process) Blocking() bool {
	return true
}
