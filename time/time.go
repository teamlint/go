// Copyright 2018 The Teamlint Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// you can obtain one at https://github.com/teamlint/go.

// Package time 实现time常用操作
package time

import "time"

const (
	DefaultDatetimeFormat = "2006-01-02 15:04:05"
	DefaultDateFormat     = "2006-01-02"
	DefaultTimeFormat     = "15:04:05"
)

var (
	TimeFormats = []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}
)
