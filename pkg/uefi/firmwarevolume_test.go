// Copyright 2018 the LinuxBoot Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uefi

import (
	"fmt"
	"os"
	"testing"

	"github.com/linuxboot/fiano/pkg/log"
)

var (
	// FV examples
	emptyFV  = []byte{} // Empty file
	sampleFV []byte     // Sample FV from OVMF
)

func init() {
	var err error
	sampleFV, err = os.ReadFile("../../integration/roms/ovmfSECFV.fv")
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func TestNewFirmwareVolume(t *testing.T) {
	var tests = []struct {
		name string
		buf  []byte
		msg  string
	}{
		{"emptyFV", emptyFV, fmt.Sprintf("Firmware Volume size too small: expected %d bytes, got %d",
			FirmwareVolumeMinSize, len(emptyFV))},
		{"sampleFV", sampleFV, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewFirmwareVolume(test.buf, 0, false)
			if err == nil && test.msg != "" {
				t.Errorf("Error was not returned, expected %v", test.msg)
			} else if err != nil && err.Error() != test.msg {
				t.Errorf("Mismatched Error returned, expected \n%v\n got \n%v\n", test.msg, err.Error())
			}
		})
	}
}
