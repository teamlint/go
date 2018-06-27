// Copyright 2018 The Teamlint Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// You can obtain one at https://github.com/teamlint/go.

// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// Package convert 类型转化
package convert

import (
	"fmt"
	"strconv"
	"time"

	tbinary "github.com/teamlint/go/encoding/binary"
	ttime "github.com/teamlint/go/time"
)

// ToType 将变量i转换为字符串指定的类型t
func ToType(i interface{}, t string) interface{} {
	switch t {
	case "int":
		return ToInt(i)
	case "int8":
		return ToInt8(i)
	case "int16":
		return ToInt16(i)
	case "int32":
		return ToInt32(i)
	case "int64":
		return ToInt64(i)
	case "uint":
		return ToUint(i)
	case "uint8":
		return ToUint8(i)
	case "uint16":
		return ToUint16(i)
	case "uint32":
		return ToUint32(i)
	case "uint64":
		return ToUint64(i)
	case "float32":
		return ToFloat32(i)
	case "float64":
		return ToFloat64(i)
	case "bool":
		return ToBool(i)
	case "string":
		return ToString(i)
	case "[]byte":
		return ToBytes(i)
	case "time.Time":
		return ToTime(i)
	case "time.Duration":
		return ToTimeDuration(i)
	default:
		return i
	}
}

// ToTime 转化为time.Time
func ToTime(i interface{}, format ...string) time.Time {
	switch value := i.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		t := ToInt64(i)
		return time.Unix(t, 0)
	case string:
		var t time.Time
		var err error
		if len(format) > 0 {
			t, err = time.ParseInLocation(format[0], value, time.Local)
			if err == nil {
				return time.Time{}
			}
			return t
		}
		for _, layout := range ttime.TimeFormats {
			t, err = time.ParseInLocation(layout, value, time.Local)
			if err == nil {
				break
			}
		}
		if err != nil {
			return time.Time{}
		}
		return t
	default:
		return time.Time{}
	}
}

// ToTimeDuration 将变量i转换为time.Duration类型
func ToTimeDuration(i interface{}) time.Duration {
	return time.Duration(ToInt64(i))
}

// ToBytes 转化为[]byte
func ToBytes(i interface{}) []byte {
	if i == nil {
		return nil
	}
	if r, ok := i.([]byte); ok {
		return r
	}
	return tbinary.Encode(i)
}

// ToString 基础的字符串类型转换
func ToString(i interface{}) string {
	if i == nil {
		return ""
	}
	switch value := i.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.Itoa(int(value))
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(uint64(value), 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	default:
		return fmt.Sprintf("%v", value)
	}
}

// ToStringSlice 转化为字符串切片
func ToStringSlice(i interface{}) []string {
	if i == nil {
		return nil
	}
	if r, ok := i.([]string); ok {
		return r
	} else if r, ok := i.([]interface{}); ok {
		strs := make([]string, len(r))
		for k, v := range r {
			strs[k] = ToString(v)
		}
		return strs
	}
	return []string{fmt.Sprintf("%v", i)}
}

// ToBool false: "", 0, false, off
func ToBool(i interface{}) bool {
	if i == nil {
		return false
	}
	if v, ok := i.(bool); ok {
		return v
	}
	if s := ToString(i); s != "" && s != "0" && s != "false" && s != "off" {
		return true
	}
	return false
}

// ToInt 转化为int
func ToInt(i interface{}) int {
	if i == nil {
		return 0
	}
	switch value := i.(type) {
	case int:
		return value
	case int8:
		return int(value)
	case int16:
		return int(value)
	case int32:
		return int(value)
	case int64:
		return int(value)
	case uint:
		return int(value)
	case uint8:
		return int(value)
	case uint16:
		return int(value)
	case uint32:
		return int(value)
	case uint64:
		return int(value)
	case float32:
		return int(value)
	case float64:
		return int(value)
	case bool:
		if value {
			return 1
		}
		return 0
	default:
		v, _ := strconv.Atoi(ToString(value))
		return v
	}
}

// ToInt8 转化为int8
func ToInt8(i interface{}) int8 {
	if i == nil {
		return 0
	}
	if v, ok := i.(int8); ok {
		return v
	}
	return int8(ToInt(i))
}

// ToInt16 转化为int16
func ToInt16(i interface{}) int16 {
	if i == nil {
		return 0
	}
	if v, ok := i.(int16); ok {
		return v
	}
	return int16(ToInt(i))
}

// ToInt32 转化int32
func ToInt32(i interface{}) int32 {
	if i == nil {
		return 0
	}
	if v, ok := i.(int32); ok {
		return v
	}
	return int32(ToInt(i))
}

// ToInt64 转化为int64
func ToInt64(i interface{}) int64 {
	if i == nil {
		return 0
	}
	if v, ok := i.(int64); ok {
		return v
	}
	return int64(ToInt(i))
}

// ToUint 转化为unit
func ToUint(i interface{}) uint {
	if i == nil {
		return 0
	}
	switch value := i.(type) {
	case int:
		return uint(value)
	case int8:
		return uint(value)
	case int16:
		return uint(value)
	case int32:
		return uint(value)
	case int64:
		return uint(value)
	case uint:
		return value
	case uint8:
		return uint(value)
	case uint16:
		return uint(value)
	case uint32:
		return uint(value)
	case uint64:
		return uint(value)
	case float32:
		return uint(value)
	case float64:
		return uint(value)
	case bool:
		if value {
			return 1
		}
		return 0
	default:
		v, _ := strconv.ParseUint(ToString(value), 10, 64)
		return uint(v)
	}
}

// ToUint8 转化为uint8
func ToUint8(i interface{}) uint8 {
	if i == nil {
		return 0
	}
	if v, ok := i.(uint8); ok {
		return v
	}
	return uint8(ToUint(i))
}

// ToUint16 转化为uint16
func ToUint16(i interface{}) uint16 {
	if i == nil {
		return 0
	}
	if v, ok := i.(uint16); ok {
		return v
	}
	return uint16(ToUint(i))
}

// ToUint32 转化为uint32
func ToUint32(i interface{}) uint32 {
	if i == nil {
		return 0
	}
	if v, ok := i.(uint32); ok {
		return v
	}
	return uint32(ToUint(i))
}

// ToUint64 转化为uint64
func ToUint64(i interface{}) uint64 {
	if i == nil {
		return 0
	}
	if v, ok := i.(uint64); ok {
		return v
	}
	return uint64(ToUint(i))
}

// ToFloat32 转化为float32
func ToFloat32(i interface{}) float32 {
	if i == nil {
		return 0
	}
	if v, ok := i.(float32); ok {
		return v
	}
	v, _ := strconv.ParseFloat(ToString(i), 32)
	return float32(v)
}

// ToFloat64 转化为float64
func ToFloat64(i interface{}) float64 {
	if i == nil {
		return 0
	}
	if v, ok := i.(float64); ok {
		return v
	}
	v, _ := strconv.ParseFloat(ToString(i), 64)
	return v
}
