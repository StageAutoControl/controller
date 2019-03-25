package process

import (
	"context"
	"fmt"
	"time"

	"github.com/StageAutoControl/controller/pkg/internal/logging"
)

type processInfo struct {
	process Process
	status  Status
}

type manager struct {
	ctx       context.Context
	logger    logging.Logger
	processes map[string]*processInfo
}

// NewManager returns a new process manager instance
func NewManager(ctx context.Context, logger logging.Logger) Manager {
	m := &manager{
		ctx:       ctx,
		logger:    logger,
		processes: make(map[string]*processInfo),
	}

	go m.listenExit()

	return m
}

func (m *manager) listenExit() {
	<-m.ctx.Done()
	for name := range m.processes {
		if p, _, err := m.GetProcess(name); err != nil {
			m.logger.Errorf("failed to find process %q while shutting down: %v", name, err)

		} else if err := p.Stop(); err != nil {
			m.logger.Errorf("failed to stop process %q: %v", name, err)

		}
	}
}

func (m *manager) AddProcess(name string, process Process, verbose bool) error {
	if _, ok := m.processes[name]; ok {
		return errProcessAlreadyExists
	}

	m.processes[name] = &processInfo{
		process: process,
		status: Status{
			Name:    name,
			Running: false,
			Logs:    make([]Log, 0),
			Verbose: verbose,
		},
	}
	return nil
}

func (m *manager) GetProcess(name string) (Process, *Status, error) {
	info, ok := m.processes[name]
	if !ok {
		return nil, nil, errProcessNotFound
	}

	return info.process, &info.status, nil
}

func (m *manager) Start(name string) (*Status, error) {
	info, ok := m.processes[name]
	if !ok {
		return nil, errProcessNotFound
	}

	if info.status.Running {
		return nil, errProcessAlreadyRunning
	}

	info.status.Running = true
	info.status.Error = nil
	info.status.StartedAt = &JSONTime{Time: time.Now()}
	info.status.StoppedAt = nil
	info.status.Logs = make([]Log, 0)

	logger := NewBufferedLogger(&info.status.Logs, info.status.Verbose)
	info.process.SetLogger(logger)

	go func() {
		if err := info.process.Start(m.ctx); err != nil {
			info.status.Error = err
			info.status.Running = false
			m.logger.Errorf("failed to start process %s: %v", name, err)
		}
	}()

	return &info.status, nil
}

// Stop a given process ID
func (m *manager) Stop(name string) (*Status, error) {
	p, ok := m.processes[name]
	if !ok {
		return nil, errProcessNotFound
	}

	if !p.status.Running {
		return nil, errProcessNotRunning
	}

	if err := p.process.Stop(); err != nil {
		return nil, fmt.Errorf("failed to stop process %q: %v", name, err)
	}

	p.status.Running = false
	p.status.StoppedAt = &JSONTime{Time: time.Now()}

	return &p.status, nil
}
