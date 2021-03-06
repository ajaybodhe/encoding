package colferModels

// Code generated by colf(1); DO NOT EDIT.
// The compiler used schema file data.colf.

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

var intconv = binary.BigEndian

// Colfer configuration attributes
var (
	// ColferSizeMax is the upper limit for serial byte sizes.
	ColferSizeMax = 16 * 1024 * 1024
)

// ColferMax signals an upper limit breach.
type ColferMax string

// Error honors the error interface.
func (m ColferMax) Error() string { return string(m) }

// ColferError signals a data mismatch as as a byte index.
type ColferError int

// Error honors the error interface.
func (i ColferError) Error() string {
	return fmt.Sprintf("colfer: unknown header at byte %d", i)
}

// ColferTail signals data continuation as a byte index.
type ColferTail int

// Error honors the error interface.
func (i ColferTail) Error() string {
	return fmt.Sprintf("colfer: data continuation at byte %d", i)
}

type ColferA struct {
	Name string

	BirthDay int64

	Phone string

	Siblings int32

	Spouse bool

	Money float64

	Loc *ColferB
}

// MarshalTo encodes o as Colfer into buf and returns the number of bytes written.
// If the buffer is too small, MarshalTo will panic.
func (o *ColferA) MarshalTo(buf []byte) int {
	var i int

	if l := len(o.Name); l != 0 {
		buf[i] = 0
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		i += copy(buf[i:], o.Name)
	}

	if v := o.BirthDay; v != 0 {
		x := uint64(v)
		if v >= 0 {
			buf[i] = 1
		} else {
			x = ^x + 1
			buf[i] = 1 | 0x80
		}
		i++
		for n := 0; x >= 0x80 && n < 8; n++ {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if l := len(o.Phone); l != 0 {
		buf[i] = 2
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		i += copy(buf[i:], o.Phone)
	}

	if v := o.Siblings; v != 0 {
		x := uint32(v)
		if v >= 0 {
			buf[i] = 3
		} else {
			x = ^x + 1
			buf[i] = 3 | 0x80
		}
		i++
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
	}

	if o.Spouse {
		buf[i] = 4
		i++
	}

	if v := o.Money; v != 0 {
		buf[i] = 5
		intconv.PutUint64(buf[i+1:], math.Float64bits(v))
		i += 9
	}

	if v := o.Loc; v != nil {
		buf[i] = 6
		i++
		i += v.MarshalTo(buf[i:])
	}

	buf[i] = 0x7f
	i++
	return i
}

// MarshalLen returns the Colfer serial byte size.
// The error return option is colferModels.ColferMax.
func (o *ColferA) MarshalLen() (int, error) {
	l := 1

	if x := len(o.Name); x != 0 {
		if x > ColferSizeMax {
			return 0, ColferMax(fmt.Sprintf("colfer: field colferModels.ColferA.Name exceeds %d bytes", ColferSizeMax))
		}
		for l += x + 2; x >= 0x80; l++ {
			x >>= 7
		}
	}

	if v := o.BirthDay; v != 0 {
		l += 2
		x := uint64(v)
		if v < 0 {
			x = ^x + 1
		}
		for n := 0; x >= 0x80 && n < 8; n++ {
			x >>= 7
			l++
		}
	}

	if x := len(o.Phone); x != 0 {
		if x > ColferSizeMax {
			return 0, ColferMax(fmt.Sprintf("colfer: field colferModels.ColferA.Phone exceeds %d bytes", ColferSizeMax))
		}
		for l += x + 2; x >= 0x80; l++ {
			x >>= 7
		}
	}

	if v := o.Siblings; v != 0 {
		x := uint32(v)
		if v < 0 {
			x = ^x + 1
		}
		for l += 2; x >= 0x80; l++ {
			x >>= 7
		}
	}

	if o.Spouse {
		l++
	}

	if o.Money != 0 {
		l += 9
	}

	if v := o.Loc; v != nil {
		vl, err := v.MarshalLen()
		if err != nil {
			return 0, err
		}
		l += vl + 1
	}

	if l > ColferSizeMax {
		return l, ColferMax(fmt.Sprintf("colfer: struct colferModels.ColferA exceeds %d bytes", ColferSizeMax))
	}
	return l, nil
}

// MarshalBinary encodes o as Colfer conform encoding.BinaryMarshaler.
// The error return option is colferModels.ColferMax.
func (o *ColferA) MarshalBinary() (data []byte, err error) {
	l, err := o.MarshalLen()
	if err != nil {
		return nil, err
	}
	data = make([]byte, l)
	o.MarshalTo(data)
	return data, nil
}

// Unmarshal decodes data as Colfer and returns the number of bytes read.
// The error return options are io.EOF, colferModels.ColferError and colferModels.ColferMax.
func (o *ColferA) Unmarshal(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, io.EOF
	}
	header := data[0]
	i := 1

	if header == 0 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferSizeMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: colferModels.ColferA.Name size %d exceeds %d bytes", x, ColferSizeMax))
		}

		start := i
		i += int(x)
		if i >= len(data) {
			goto eof
		}
		o.Name = string(data[start:i])

		header = data[i]
		i++
	}

	if header == 1 {
		if i+1 >= len(data) {
			i++
			goto eof
		}
		x := uint64(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint64(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 || shift == 56 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.BirthDay = int64(x)

		header = data[i]
		i++
	} else if header == 1|0x80 {
		if i+1 >= len(data) {
			i++
			goto eof
		}
		x := uint64(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint64(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 || shift == 56 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.BirthDay = int64(^x + 1)

		header = data[i]
		i++
	}

	if header == 2 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferSizeMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: colferModels.ColferA.Phone size %d exceeds %d bytes", x, ColferSizeMax))
		}

		start := i
		i += int(x)
		if i >= len(data) {
			goto eof
		}
		o.Phone = string(data[start:i])

		header = data[i]
		i++
	}

	if header == 3 {
		if i+1 >= len(data) {
			i++
			goto eof
		}
		x := uint32(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint32(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.Siblings = int32(x)

		header = data[i]
		i++
	} else if header == 3|0x80 {
		if i+1 >= len(data) {
			i++
			goto eof
		}
		x := uint32(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				b := uint32(data[i])
				i++
				if i >= len(data) {
					goto eof
				}

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}
		o.Siblings = int32(^x + 1)

		header = data[i]
		i++
	}

	if header == 4 {
		if i >= len(data) {
			goto eof
		}
		o.Spouse = true
		header = data[i]
		i++
	}

	if header == 5 {
		start := i
		i += 8
		if i >= len(data) {
			goto eof
		}
		o.Money = math.Float64frombits(intconv.Uint64(data[start:]))
		header = data[i]
		i++
	}

	if header == 6 {
		o.Loc = new(ColferB)
		n, err := o.Loc.Unmarshal(data[i:])
		if err != nil {
			if err == io.EOF && len(data) >= ColferSizeMax {
				return 0, ColferMax(fmt.Sprintf("colfer: colferModels.ColferA size exceeds %d bytes", ColferSizeMax))
			}
			return 0, err
		}
		i += n

		if i >= len(data) {
			goto eof
		}
		header = data[i]
		i++
	}

	if header != 0x7f {
		return 0, ColferError(i - 1)
	}
	if i < ColferSizeMax {
		return i, nil
	}
eof:
	if i >= ColferSizeMax {
		return 0, ColferMax(fmt.Sprintf("colfer: struct colferModels.ColferA size exceeds %d bytes", ColferSizeMax))
	}
	return 0, io.EOF
}

// UnmarshalBinary decodes data as Colfer conform encoding.BinaryUnmarshaler.
// The error return options are io.EOF, colferModels.ColferError, colferModels.ColferTail and colferModels.ColferMax.
func (o *ColferA) UnmarshalBinary(data []byte) error {
	i, err := o.Unmarshal(data)
	if i < len(data) && err == nil {
		return ColferTail(i)
	}
	return err
}

type ColferB struct {
	Location string
}

// MarshalTo encodes o as Colfer into buf and returns the number of bytes written.
// If the buffer is too small, MarshalTo will panic.
func (o *ColferB) MarshalTo(buf []byte) int {
	var i int

	if l := len(o.Location); l != 0 {
		buf[i] = 0
		i++
		x := uint(l)
		for x >= 0x80 {
			buf[i] = byte(x | 0x80)
			x >>= 7
			i++
		}
		buf[i] = byte(x)
		i++
		i += copy(buf[i:], o.Location)
	}

	buf[i] = 0x7f
	i++
	return i
}

// MarshalLen returns the Colfer serial byte size.
// The error return option is colferModels.ColferMax.
func (o *ColferB) MarshalLen() (int, error) {
	l := 1

	if x := len(o.Location); x != 0 {
		if x > ColferSizeMax {
			return 0, ColferMax(fmt.Sprintf("colfer: field colferModels.ColferB.Location exceeds %d bytes", ColferSizeMax))
		}
		for l += x + 2; x >= 0x80; l++ {
			x >>= 7
		}
	}

	if l > ColferSizeMax {
		return l, ColferMax(fmt.Sprintf("colfer: struct colferModels.ColferB exceeds %d bytes", ColferSizeMax))
	}
	return l, nil
}

// MarshalBinary encodes o as Colfer conform encoding.BinaryMarshaler.
// The error return option is colferModels.ColferMax.
func (o *ColferB) MarshalBinary() (data []byte, err error) {
	l, err := o.MarshalLen()
	if err != nil {
		return nil, err
	}
	data = make([]byte, l)
	o.MarshalTo(data)
	return data, nil
}

// Unmarshal decodes data as Colfer and returns the number of bytes read.
// The error return options are io.EOF, colferModels.ColferError and colferModels.ColferMax.
func (o *ColferB) Unmarshal(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, io.EOF
	}
	header := data[0]
	i := 1

	if header == 0 {
		if i >= len(data) {
			goto eof
		}
		x := uint(data[i])
		i++

		if x >= 0x80 {
			x &= 0x7f
			for shift := uint(7); ; shift += 7 {
				if i >= len(data) {
					goto eof
				}
				b := uint(data[i])
				i++

				if b < 0x80 {
					x |= b << shift
					break
				}
				x |= (b & 0x7f) << shift
			}
		}

		if x > uint(ColferSizeMax) {
			return 0, ColferMax(fmt.Sprintf("colfer: colferModels.ColferB.Location size %d exceeds %d bytes", x, ColferSizeMax))
		}

		start := i
		i += int(x)
		if i >= len(data) {
			goto eof
		}
		o.Location = string(data[start:i])

		header = data[i]
		i++
	}

	if header != 0x7f {
		return 0, ColferError(i - 1)
	}
	if i < ColferSizeMax {
		return i, nil
	}
eof:
	if i >= ColferSizeMax {
		return 0, ColferMax(fmt.Sprintf("colfer: struct colferModels.ColferB size exceeds %d bytes", ColferSizeMax))
	}
	return 0, io.EOF
}

// UnmarshalBinary decodes data as Colfer conform encoding.BinaryUnmarshaler.
// The error return options are io.EOF, colferModels.ColferError, colferModels.ColferTail and colferModels.ColferMax.
func (o *ColferB) UnmarshalBinary(data []byte) error {
	i, err := o.Unmarshal(data)
	if i < len(data) && err == nil {
		return ColferTail(i)
	}
	return err
}
