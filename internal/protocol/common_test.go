// Copyright (C) 2014 Jakob Borg and Contributors (see the CONTRIBUTORS file).
// All rights reserved. Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package protocol

import (
	"io"
	"time"
)

type TestModel struct {
	data     []byte
	folder   string
	name     string
	offset   int64
	size     int
	closedCh chan bool
}

func newTestModel() *TestModel {
	return &TestModel{
		closedCh: make(chan bool),
	}
}

func (t *TestModel) Index(deviceID DeviceID, folder string, files []FileInfo) {
}

func (t *TestModel) IndexUpdate(deviceID DeviceID, folder string, files []FileInfo) {
}

func (t *TestModel) Request(deviceID DeviceID, folder, name string, offset int64, size int) ([]byte, error) {
	t.folder = folder
	t.name = name
	t.offset = offset
	t.size = size
	return t.data, nil
}

func (t *TestModel) Close(deviceID DeviceID, err error) {
	close(t.closedCh)
}

func (t *TestModel) ClusterConfig(deviceID DeviceID, config ClusterConfigMessage) {
}

func (t *TestModel) isClosed() bool {
	select {
	case <-t.closedCh:
		return true
	case <-time.After(1 * time.Second):
		return false // Timeout
	}
}

type ErrPipe struct {
	io.PipeWriter
	written int
	max     int
	err     error
	closed  bool
}

func (e *ErrPipe) Write(data []byte) (int, error) {
	if e.closed {
		return 0, e.err
	}
	if e.written+len(data) > e.max {
		n, _ := e.PipeWriter.Write(data[:e.max-e.written])
		e.PipeWriter.CloseWithError(e.err)
		e.closed = true
		return n, e.err
	}
	return e.PipeWriter.Write(data)
}
