// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// 2.2练习题

// Package unitconv performs Meter and Foot conversions.
package unitconv

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }

// MToF converts a Meter to Foot.
func MToF(m Meter) Foot { return Foot(m / 0.3048) }

// FToC converts a Foot to Meter.
func FToM(f Foot) Meter { return Meter(f * 0.3048) }

//!-
//!-
