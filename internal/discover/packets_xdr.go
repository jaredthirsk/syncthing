// ************************************************************
// This file is automatically generated by genxdr. Do not edit.
// ************************************************************

package discover

import (
	"bytes"
	"io"

	"github.com/calmh/xdr"
)

/*

Query Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Magic                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                      Length of Device ID                      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                  Device ID (variable length)                  \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Query {
	unsigned int Magic;
	opaque DeviceID<32>;
}

*/

func (o Query) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o Query) MarshalXDR() []byte {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o Query) AppendXDR(bs []byte) []byte {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	o.encodeXDR(xw)
	return []byte(aw)
}

func (o Query) encodeXDR(xw *xdr.Writer) (int, error) {
	xw.WriteUint32(o.Magic)
	if len(o.DeviceID) > 32 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteBytes(o.DeviceID)
	return xw.Tot(), xw.Error()
}

func (o *Query) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *Query) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *Query) decodeXDR(xr *xdr.Reader) error {
	o.Magic = xr.ReadUint32()
	o.DeviceID = xr.ReadBytesMax(32)
	return xr.Error()
}

/*

Announce Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Magic                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                            Device                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Number of Extra                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                Zero or more Device Structures                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Announce {
	unsigned int Magic;
	Device This;
	Device Extra<16>;
}

*/

func (o Announce) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o Announce) MarshalXDR() []byte {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o Announce) AppendXDR(bs []byte) []byte {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	o.encodeXDR(xw)
	return []byte(aw)
}

func (o Announce) encodeXDR(xw *xdr.Writer) (int, error) {
	xw.WriteUint32(o.Magic)
	_, err := o.This.encodeXDR(xw)
	if err != nil {
		return xw.Tot(), err
	}
	if len(o.Extra) > 16 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteUint32(uint32(len(o.Extra)))
	for i := range o.Extra {
		_, err := o.Extra[i].encodeXDR(xw)
		if err != nil {
			return xw.Tot(), err
		}
	}
	return xw.Tot(), xw.Error()
}

func (o *Announce) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *Announce) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *Announce) decodeXDR(xr *xdr.Reader) error {
	o.Magic = xr.ReadUint32()
	(&o.This).decodeXDR(xr)
	_ExtraSize := int(xr.ReadUint32())
	if _ExtraSize > 16 {
		return xdr.ErrElementSizeExceeded
	}
	o.Extra = make([]Device, _ExtraSize)
	for i := range o.Extra {
		(&o.Extra[i]).decodeXDR(xr)
	}
	return xr.Error()
}

/*

Device Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of ID                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     ID (variable length)                      \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                      Number of Addresses                      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                Zero or more Address Structures                \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Device {
	opaque ID<32>;
	Address Addresses<16>;
}

*/

func (o Device) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o Device) MarshalXDR() []byte {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o Device) AppendXDR(bs []byte) []byte {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	o.encodeXDR(xw)
	return []byte(aw)
}

func (o Device) encodeXDR(xw *xdr.Writer) (int, error) {
	if len(o.ID) > 32 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteBytes(o.ID)
	if len(o.Addresses) > 16 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteUint32(uint32(len(o.Addresses)))
	for i := range o.Addresses {
		_, err := o.Addresses[i].encodeXDR(xw)
		if err != nil {
			return xw.Tot(), err
		}
	}
	return xw.Tot(), xw.Error()
}

func (o *Device) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *Device) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *Device) decodeXDR(xr *xdr.Reader) error {
	o.ID = xr.ReadBytesMax(32)
	_AddressesSize := int(xr.ReadUint32())
	if _AddressesSize > 16 {
		return xdr.ErrElementSizeExceeded
	}
	o.Addresses = make([]Address, _AddressesSize)
	for i := range o.Addresses {
		(&o.Addresses[i]).decodeXDR(xr)
	}
	return xr.Error()
}

/*

Address Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Length of IP                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                     IP (variable length)                      \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|            0x0000             |             Port              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Address {
	opaque IP<16>;
	unsigned int Port;
}

*/

func (o Address) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.encodeXDR(xw)
}

func (o Address) MarshalXDR() []byte {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o Address) AppendXDR(bs []byte) []byte {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	o.encodeXDR(xw)
	return []byte(aw)
}

func (o Address) encodeXDR(xw *xdr.Writer) (int, error) {
	if len(o.IP) > 16 {
		return xw.Tot(), xdr.ErrElementSizeExceeded
	}
	xw.WriteBytes(o.IP)
	xw.WriteUint16(o.Port)
	return xw.Tot(), xw.Error()
}

func (o *Address) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.decodeXDR(xr)
}

func (o *Address) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.decodeXDR(xr)
}

func (o *Address) decodeXDR(xr *xdr.Reader) error {
	o.IP = xr.ReadBytesMax(16)
	o.Port = xr.ReadUint16()
	return xr.Error()
}
