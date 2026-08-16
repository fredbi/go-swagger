// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build race

package generator

// raceEnabled tells whether the tests run under the race detector.
//
// A plugin has to be built the same way as the program loading it, so the build of the fixture
// needs to know.
const raceEnabled = true
