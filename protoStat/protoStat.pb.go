// Code generated by protoc-gen-gogo.
// source: protoStat.proto
// DO NOT EDIT!

/*
	Package protoStat is a generated protocol buffer package.

	It is generated from these files:
		protoStat.proto

	It has these top-level messages:
		ProtoStat
		ProtoStats
*/
package protoStat

import proto "code.google.com/p/gogoprotobuf/proto"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import io "io"
import math1 "math"
import fmt "fmt"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"

import fmt1 "fmt"
import strings "strings"
import reflect "reflect"

import math2 "math"

import fmt2 "fmt"
import strings1 "strings"
import code_google_com_p_gogoprotobuf_proto1 "code.google.com/p/gogoprotobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect1 "reflect"

import fmt3 "fmt"
import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ProtoStat struct {
	Key              *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
	IndexKey         *string  `protobuf:"bytes,3,opt,name=indexKey" json:"indexKey,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ProtoStat) Reset()      { *m = ProtoStat{} }
func (*ProtoStat) ProtoMessage() {}

func (m *ProtoStat) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ProtoStat) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *ProtoStat) GetIndexKey() string {
	if m != nil && m.IndexKey != nil {
		return *m.IndexKey
	}
	return ""
}

type ProtoStats struct {
	Stats            []*ProtoStat `protobuf:"bytes,1,rep,name=stats" json:"stats,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ProtoStats) Reset()      { *m = ProtoStats{} }
func (*ProtoStats) ProtoMessage() {}

func (m *ProtoStats) GetStats() []*ProtoStat {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
}
func (m *ProtoStat) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.Key = &s
			index = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			i := index + 8
			if i > l {
				return io.ErrUnexpectedEOF
			}
			index = i
			v = uint64(data[i-8])
			v |= uint64(data[i-7]) << 8
			v |= uint64(data[i-6]) << 16
			v |= uint64(data[i-5]) << 24
			v |= uint64(data[i-4]) << 32
			v |= uint64(data[i-3]) << 40
			v |= uint64(data[i-2]) << 48
			v |= uint64(data[i-1]) << 56
			v2 := math1.Float64frombits(v)
			m.Value = &v2
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.IndexKey = &s
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *ProtoStats) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stats = append(m.Stats, &ProtoStat{})
			m.Stats[len(m.Stats)-1].Unmarshal(data[index:postIndex])
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (this *ProtoStat) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProtoStat{`,
		`Key:` + valueToStringProtoStat(this.Key) + `,`,
		`Value:` + valueToStringProtoStat(this.Value) + `,`,
		`IndexKey:` + valueToStringProtoStat(this.IndexKey) + `,`,
		`XXX_unrecognized:` + fmt1.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProtoStats) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProtoStats{`,
		`Stats:` + strings.Replace(fmt1.Sprintf("%v", this.Stats), "ProtoStat", "ProtoStat", 1) + `,`,
		`XXX_unrecognized:` + fmt1.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProtoStat(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt1.Sprintf("*%v", pv)
}
func (m *ProtoStat) Size() (n int) {
	var l int
	_ = l
	if m.Key != nil {
		l = len(*m.Key)
		n += 1 + l + sovProtoStat(uint64(l))
	}
	if m.Value != nil {
		n += 9
	}
	if m.IndexKey != nil {
		l = len(*m.IndexKey)
		n += 1 + l + sovProtoStat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *ProtoStats) Size() (n int) {
	var l int
	_ = l
	if len(m.Stats) > 0 {
		for _, e := range m.Stats {
			l = e.Size()
			n += 1 + l + sovProtoStat(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProtoStat(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtoStat(x uint64) (n int) {
	return sovProtoStat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedProtoStat(r randyProtoStat, easy bool) *ProtoStat {
	this := &ProtoStat{}
	v1 := randStringProtoStat(r)
	this.Key = &v1
	if r.Intn(10) != 0 {
		v2 := r.Float64()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		this.Value = &v2
	}
	if r.Intn(10) != 0 {
		v3 := randStringProtoStat(r)
		this.IndexKey = &v3
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedProtoStat(r, 4)
	}
	return this
}

func NewPopulatedProtoStats(r randyProtoStat, easy bool) *ProtoStats {
	this := &ProtoStats{}
	if r.Intn(10) != 0 {
		v4 := r.Intn(10)
		this.Stats = make([]*ProtoStat, v4)
		for i := 0; i < v4; i++ {
			this.Stats[i] = NewPopulatedProtoStat(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedProtoStat(r, 2)
	}
	return this
}

type randyProtoStat interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneProtoStat(r randyProtoStat) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringProtoStat(r randyProtoStat) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneProtoStat(r)
	}
	return string(tmps)
}
func randUnrecognizedProtoStat(r randyProtoStat, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldProtoStat(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldProtoStat(data []byte, r randyProtoStat, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateProtoStat(data, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		data = encodeVarintPopulateProtoStat(data, uint64(v6))
	case 1:
		data = encodeVarintPopulateProtoStat(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateProtoStat(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateProtoStat(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateProtoStat(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateProtoStat(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *ProtoStat) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ProtoStat) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key != nil {
		data[i] = 0xa
		i++
		i = encodeVarintProtoStat(data, i, uint64(len(*m.Key)))
		i += copy(data[i:], *m.Key)
	}
	if m.Value != nil {
		data[i] = 0x11
		i++
		i = encodeFixed64ProtoStat(data, i, uint64(math2.Float64bits(*m.Value)))
	}
	if m.IndexKey != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintProtoStat(data, i, uint64(len(*m.IndexKey)))
		i += copy(data[i:], *m.IndexKey)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *ProtoStats) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ProtoStats) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Stats) > 0 {
		for _, msg := range m.Stats {
			data[i] = 0xa
			i++
			i = encodeVarintProtoStat(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64ProtoStat(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ProtoStat(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintProtoStat(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *ProtoStat) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&protoStat.ProtoStat{` + `Key:` + valueToGoStringProtoStat(this.Key, "string"), `Value:` + valueToGoStringProtoStat(this.Value, "float64"), `IndexKey:` + valueToGoStringProtoStat(this.IndexKey, "string"), `XXX_unrecognized:` + fmt2.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *ProtoStats) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&protoStat.ProtoStats{` + `Stats:` + fmt2.Sprintf("%#v", this.Stats), `XXX_unrecognized:` + fmt2.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func valueToGoStringProtoStat(v interface{}, typ string) string {
	rv := reflect1.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect1.Indirect(rv).Interface()
	return fmt2.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringProtoStat(e map[int32]code_google_com_p_gogoprotobuf_proto1.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings1.Join(ss, ",") + "}"
	return s
}
func (this *ProtoStat) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt3.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ProtoStat)
	if !ok {
		return fmt3.Errorf("that is not of type *ProtoStat")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt3.Errorf("that is type *ProtoStat but is nil && this != nil")
	} else if this == nil {
		return fmt3.Errorf("that is type *ProtoStatbut is not nil && this == nil")
	}
	if this.Key != nil && that1.Key != nil {
		if *this.Key != *that1.Key {
			return fmt3.Errorf("Key this(%v) Not Equal that(%v)", *this.Key, *that1.Key)
		}
	} else if this.Key != nil {
		return fmt3.Errorf("this.Key == nil && that.Key != nil")
	} else if that1.Key != nil {
		return fmt3.Errorf("Key this(%v) Not Equal that(%v)", this.Key, that1.Key)
	}
	if this.Value != nil && that1.Value != nil {
		if *this.Value != *that1.Value {
			return fmt3.Errorf("Value this(%v) Not Equal that(%v)", *this.Value, *that1.Value)
		}
	} else if this.Value != nil {
		return fmt3.Errorf("this.Value == nil && that.Value != nil")
	} else if that1.Value != nil {
		return fmt3.Errorf("Value this(%v) Not Equal that(%v)", this.Value, that1.Value)
	}
	if this.IndexKey != nil && that1.IndexKey != nil {
		if *this.IndexKey != *that1.IndexKey {
			return fmt3.Errorf("IndexKey this(%v) Not Equal that(%v)", *this.IndexKey, *that1.IndexKey)
		}
	} else if this.IndexKey != nil {
		return fmt3.Errorf("this.IndexKey == nil && that.IndexKey != nil")
	} else if that1.IndexKey != nil {
		return fmt3.Errorf("IndexKey this(%v) Not Equal that(%v)", this.IndexKey, that1.IndexKey)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt3.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *ProtoStat) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ProtoStat)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Key != nil && that1.Key != nil {
		if *this.Key != *that1.Key {
			return false
		}
	} else if this.Key != nil {
		return false
	} else if that1.Key != nil {
		return false
	}
	if this.Value != nil && that1.Value != nil {
		if *this.Value != *that1.Value {
			return false
		}
	} else if this.Value != nil {
		return false
	} else if that1.Value != nil {
		return false
	}
	if this.IndexKey != nil && that1.IndexKey != nil {
		if *this.IndexKey != *that1.IndexKey {
			return false
		}
	} else if this.IndexKey != nil {
		return false
	} else if that1.IndexKey != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ProtoStats) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt3.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ProtoStats)
	if !ok {
		return fmt3.Errorf("that is not of type *ProtoStats")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt3.Errorf("that is type *ProtoStats but is nil && this != nil")
	} else if this == nil {
		return fmt3.Errorf("that is type *ProtoStatsbut is not nil && this == nil")
	}
	if len(this.Stats) != len(that1.Stats) {
		return fmt3.Errorf("Stats this(%v) Not Equal that(%v)", len(this.Stats), len(that1.Stats))
	}
	for i := range this.Stats {
		if !this.Stats[i].Equal(that1.Stats[i]) {
			return fmt3.Errorf("Stats this[%v](%v) Not Equal that[%v](%v)", i, this.Stats[i], i, that1.Stats[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt3.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *ProtoStats) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ProtoStats)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Stats) != len(that1.Stats) {
		return false
	}
	for i := range this.Stats {
		if !this.Stats[i].Equal(that1.Stats[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
