package main

import (
	"reflect"
	"testing"
)

func Test_execBashPipedCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := execBashPipedCommand(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("execBashPipedCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("execBashPipedCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_installClamAV(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			installClamAV()
		})
	}
}

func Test_refreshClamAVVirusDatabase(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			refreshClamAVVirusDatabase()
		})
	}
}

func Test_configureClamAVDaemon(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configureClamAVDaemon()
		})
	}
}

func Test_configureScanService(t *testing.T) {
	type args struct {
		scanDir string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configureScanService(tt.args.scanDir)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
