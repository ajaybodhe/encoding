package gencodeModels

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type GencodeUnsafeB struct {
	Location string
}

func (d *GencodeUnsafeB) Size() (s uint64) {

	{
		l := uint64(len(d.Location))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	return
}
func (d *GencodeUnsafeB) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Location))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Location)
		i += l
	}
	return buf[:i+0], nil
}

func (d *GencodeUnsafeB) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Location = string(buf[i+0 : i+0+l])
		i += l
	}
	return i + 0, nil
}

type GencodeUnsafeA struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int32
	Spouse   bool
	Money    float64
	Loc      *GencodeUnsafeB
}

func (d *GencodeUnsafeA) Size() (s uint64) {

	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Phone))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		if d.Loc != nil {

			{
				s += (*d.Loc).Size()
			}
			s += 0
		}
	}
	s += 22
	return
}
func (d *GencodeUnsafeA) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Name)
		i += l
	}
	{

		buf[i+0+0] = byte(d.BirthDay >> 0)

		buf[i+1+0] = byte(d.BirthDay >> 8)

		buf[i+2+0] = byte(d.BirthDay >> 16)

		buf[i+3+0] = byte(d.BirthDay >> 24)

		buf[i+4+0] = byte(d.BirthDay >> 32)

		buf[i+5+0] = byte(d.BirthDay >> 40)

		buf[i+6+0] = byte(d.BirthDay >> 48)

		buf[i+7+0] = byte(d.BirthDay >> 56)

	}
	{
		l := uint64(len(d.Phone))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.Phone)
		i += l
	}
	{

		buf[i+0+8] = byte(d.Siblings >> 0)

		buf[i+1+8] = byte(d.Siblings >> 8)

		buf[i+2+8] = byte(d.Siblings >> 16)

		buf[i+3+8] = byte(d.Siblings >> 24)

	}
	{
		if d.Spouse {
			buf[i+12] = 1
		} else {
			buf[i+12] = 0
		}
	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Money)))

		buf[i+0+13] = byte(v >> 0)

		buf[i+1+13] = byte(v >> 8)

		buf[i+2+13] = byte(v >> 16)

		buf[i+3+13] = byte(v >> 24)

		buf[i+4+13] = byte(v >> 32)

		buf[i+5+13] = byte(v >> 40)

		buf[i+6+13] = byte(v >> 48)

		buf[i+7+13] = byte(v >> 56)

	}
	{
		if d.Loc == nil {
			buf[i+21] = 0
		} else {
			buf[i+21] = 1

			{
				nbuf, err := (*d.Loc).Marshal(buf[i+22:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}
			i += 0
		}
	}
	return buf[:i+22], nil
}

func (d *GencodeUnsafeA) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i+0 : i+0+l])
		i += l
	}
	{

		d.BirthDay = 0 | (int64(buf[i+0+0]) << 0) | (int64(buf[i+1+0]) << 8) | (int64(buf[i+2+0]) << 16) | (int64(buf[i+3+0]) << 24) | (int64(buf[i+4+0]) << 32) | (int64(buf[i+5+0]) << 40) | (int64(buf[i+6+0]) << 48) | (int64(buf[i+7+0]) << 56)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Phone = string(buf[i+8 : i+8+l])
		i += l
	}
	{

		d.Siblings = 0 | (int32(buf[i+0+8]) << 0) | (int32(buf[i+1+8]) << 8) | (int32(buf[i+2+8]) << 16) | (int32(buf[i+3+8]) << 24)

	}
	{
		d.Spouse = buf[i+12] == 1
	}
	{

		v := 0 | (uint64(buf[i+0+13]) << 0) | (uint64(buf[i+1+13]) << 8) | (uint64(buf[i+2+13]) << 16) | (uint64(buf[i+3+13]) << 24) | (uint64(buf[i+4+13]) << 32) | (uint64(buf[i+5+13]) << 40) | (uint64(buf[i+6+13]) << 48) | (uint64(buf[i+7+13]) << 56)
		d.Money = *(*float64)(unsafe.Pointer(&v))

	}
	{
		if buf[i+21] == 1 {
			if d.Loc == nil {
				d.Loc = new(GencodeUnsafeB)
			}

			{
				ni, err := (*d.Loc).Unmarshal(buf[i+22:])
				if err != nil {
					return 0, err
				}
				i += ni
			}
			i += 0
		} else {
			d.Loc = nil
		}
	}
	return i + 22, nil
}
