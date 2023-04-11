// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package library

import "testing"

func TestToLowerWrapper(t *testing.T) {
	if ToLowerWrapper("DiscWorld") != "discworld" {
		t.Fail()
	}
}
