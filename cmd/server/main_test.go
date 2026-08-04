package main

import (
	"errors"
	"testing"
)

func TestMain_InitializeSuccess_RunServerSuccess(t *testing.T) {

	oldInitialize := initialize
	oldRunServer := runServer
	oldInfo := infoLog
	oldFatal := fatalLog

	defer func() {
		initialize = oldInitialize
		runServer = oldRunServer
		infoLog = oldInfo
		fatalLog = oldFatal
	}()

	initCalled := false
	runCalled := false
	infoCount := 0
	fatalCalled := false

	initialize = func() error {
		initCalled = true
		return nil
	}

	runServer = func() error {
		runCalled = true
		return nil
	}

	infoLog = func(v ...any) {
		infoCount++
	}

	fatalLog = func(v ...any) {
		fatalCalled = true
	}

	main()

	if !initCalled {
		t.Fatal("initialize not called")
	}

	if !runCalled {
		t.Fatal("runServer not called")
	}

	if infoCount != 2 {
		t.Fatalf("expected 2 info logs got %d", infoCount)
	}

	if fatalCalled {
		t.Fatal("fatal should not be called")
	}
}

func TestMain_InitializeFailure(t *testing.T) {

	oldInitialize := initialize
	oldRunServer := runServer
	oldInfo := infoLog
	oldFatal := fatalLog

	defer func() {
		initialize = oldInitialize
		runServer = oldRunServer
		infoLog = oldInfo
		fatalLog = oldFatal
	}()

	runCalled := false
	fatalCalled := false

	initialize = func() error {
		return errors.New("init failed")
	}

	runServer = func() error {
		runCalled = true
		return nil
	}

	infoLog = func(v ...any) {}

	fatalLog = func(v ...any) {
		fatalCalled = true
	}

	main()

	if !fatalCalled {
		t.Fatal("expected fatal to be called")
	}

	if runCalled {
		t.Fatal("runServer should not be called")
	}
}

func TestMain_RunServerFailure(t *testing.T) {

	oldInitialize := initialize
	oldRunServer := runServer
	oldInfo := infoLog
	oldFatal := fatalLog

	defer func() {
		initialize = oldInitialize
		runServer = oldRunServer
		infoLog = oldInfo
		fatalLog = oldFatal
	}()

	fatalCalled := false

	initialize = func() error {
		return nil
	}

	runServer = func() error {
		return errors.New("server failed")
	}

	infoLog = func(v ...any) {}

	fatalLog = func(v ...any) {
		fatalCalled = true
	}

	main()

	if !fatalCalled {
		t.Fatal("expected fatal to be called")
	}
}
