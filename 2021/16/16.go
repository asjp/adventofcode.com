package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

type TypeId byte

type PacketHeader struct {
	version byte
	typeID  TypeId
}

func (ph PacketHeader) Version() byte {
	return ph.version
}

type Packet interface {
	Version() byte
}

type LiteralPacket struct {
	PacketHeader
	literal int
}

type OperatorPacket struct {
	PacketHeader
	lengthType byte
	length     int
	sub        []Packet
}

const (
	HEADER_SIZE           = 6    // bits
	HEADER_BITMASK        = 0x3F // 00111111
	VERSION_SIZE          = 3    // bits
	TYPE_BITMASK          = 0x7
	LITERAL_SIZE          = 5 // bits
	LITERAL_GROUP_CONT    = 1
	LITERAL_GROUP_END     = 0
	LITERAL_GROUP_BITMASK = 0x10
	LITERAL_VALUE_BITMASK = 0xF
	LENGTH_TYPE_BITS      = 0
	LENGTH_TYPE_PACKETS   = 1
)

const (
	TYPE_OP_SUM TypeId = iota
	TYPE_OP_PRODUCT
	TYPE_OP_MINIMUM
	TYPE_OP_MAXIMUM
	TYPE_LITERAL
	TYPE_OP_GREATER
	TYPE_OP_LESS
	TYPE_OP_EQUAL
)

func ReadHeader(p *PacketHeader, offset int, buf []byte) int {
	if offset == -1 {
		return -1
	}
	nextByte := offset / 8
	if nextByte >= len(buf)-1 {
		// not enough data left for any more packets
		return -1
	}
	b := buf[nextByte]
	bitsRemain := 8 - (offset % 8)
	if bitsRemain < HEADER_SIZE {
		b = (b << (HEADER_SIZE - bitsRemain)) & HEADER_BITMASK
		b += buf[offset/8+1] >> (8 - (HEADER_SIZE - bitsRemain))
		bitsRemain = HEADER_SIZE
	}
	p.version = (b >> (bitsRemain - VERSION_SIZE)) & TYPE_BITMASK
	p.typeID = TypeId((b >> (bitsRemain - HEADER_SIZE)) & TYPE_BITMASK)
	//fmt.Printf("HEADER %+v\n", p)
	return offset + HEADER_SIZE
}

func ReadLiteral(p *LiteralPacket, offset int, buf []byte) int {
	p.literal = 0
	for {
		b := buf[offset/8]
		bitsRemain := 8 - (offset % 8)
		if bitsRemain < LITERAL_SIZE {
			b = (b << (LITERAL_SIZE - bitsRemain)) & (LITERAL_GROUP_BITMASK | LITERAL_VALUE_BITMASK)
			b += buf[offset/8+1] >> (8 - (LITERAL_SIZE - bitsRemain))
		} else if bitsRemain > LITERAL_SIZE {
			b >>= (bitsRemain - LITERAL_SIZE)
		}
		if b&LITERAL_GROUP_BITMASK == LITERAL_GROUP_BITMASK {
			p.literal += int(b & LITERAL_VALUE_BITMASK)
			p.literal <<= 4
		} else {
			p.literal += int(b & LITERAL_VALUE_BITMASK)
		}
		//fmt.Printf("literal %b %d\n", p.literal, p.literal)
		offset += 5
		if b&LITERAL_GROUP_BITMASK == LITERAL_GROUP_END {
			break
		}
	}

	//fmt.Println("literal", p.literal, "(", offset, ")")
	return offset
}

func ReadOperator(p *OperatorPacket, offset int, buf []byte, depth int) int {
	b := buf[offset/8]
	bitsRemain := 8 - (offset % 8)
	p.lengthType = (b >> (bitsRemain - 1)) & 0x1
	offset++
	bitsRemain = 8 - (offset % 8)
	if p.lengthType == LENGTH_TYPE_BITS {
		// next 15 bits are the length in bits
		data := append([]byte{0}, buf[offset/8:(offset+15)/8+1]...)
		if len(data) < 4 {
			data = append([]byte{0}, data...)
		}
		p.length = int(binary.BigEndian.Uint32(data))
		p.length = p.length >> (8 - ((offset + 15) % 8))
		p.length = p.length & 0x7FFF
		offset += 15
	} else {
		// next 11 bits are the length in packets
		data := append([]byte{0}, buf[offset/8:(offset+11)/8+1]...)
		for len(data) < 4 {
			data = append([]byte{0}, data...)
		}
		p.length = int(binary.BigEndian.Uint32(data))
		p.length = p.length >> (8 - ((offset + 11) % 8))
		p.length = p.length & 0x7FF
		offset += 11
	}
	//fmt.Printf("OPERATOR depth=%d T(%d) L(%d) offset %d\n", depth, p.lengthType, p.length, offset)

	// read sub-packets, if any
	from := offset
	n := 1
	for {
		var subp Packet
		if p.lengthType == LENGTH_TYPE_BITS && offset >= from+p.length {
			break
		}
		if p.lengthType == LENGTH_TYPE_PACKETS && n > p.length {
			break
		}
		subp, offset = ReadPacket(buf, offset, depth+1)
		p.sub = append(p.sub, subp)
		n++
	}

	return offset
}

func ReadPacket(buf []byte, offset int, depth int) (Packet, int) {
	if offset > (len(buf)-1)*8 {
		return nil, -1
	}
	hdr := &PacketHeader{}
	offset = ReadHeader(hdr, offset, buf)
	if offset == -1 {
		return nil, -1
	}
	if hdr.typeID == TYPE_LITERAL {
		ltrl := LiteralPacket{*hdr, 0}
		offset = ReadLiteral(&ltrl, offset, buf)
		return &ltrl, offset
	}
	// else operator packet
	op := OperatorPacket{*hdr, 0, 0, nil}
	offset = ReadOperator(&op, offset, buf, depth)
	return &op, offset
}

func ReadInput(r io.ReadSeeker) Packet {
	_, _ = r.Seek(0, io.SeekStart)
	rdr := bufio.NewReader(r)
	s, _ := rdr.ReadString('\n')
	hexbytes, _ := hex.DecodeString(s)
	p, _ := ReadPacket(hexbytes, 0, 0)
	return p
}

func VersionSum(p Packet) int {
	sum := int(p.Version())
	if op, ok := p.(*OperatorPacket); ok {
		for _, sp := range op.sub {
			sum += VersionSum(sp)
		}
	}
	return sum
}

func Operations(p Packet) int {
	if op, ok := p.(*OperatorPacket); ok {
		switch op.typeID {
		case TYPE_OP_SUM:
			{
				sum := 0
				for _, sp := range op.sub {
					sum += Operations(sp)
				}
				return sum
			}
		case TYPE_OP_PRODUCT:
			{
				prod := 1
				for _, sp := range op.sub {
					prod *= Operations(sp)
				}
				return prod
			}
		case TYPE_OP_MINIMUM:
			{
				min := math.MaxInt32
				for _, sp := range op.sub {
					t := Operations(sp)
					if t < min {
						min = t
					}
				}
				return min
			}
		case TYPE_OP_MAXIMUM:
			{
				max := 0
				for _, sp := range op.sub {
					t := Operations(sp)
					if t > max {
						max = t
					}
				}
				return max
			}
		case TYPE_OP_GREATER:
			{
				if Operations(op.sub[0]) > Operations(op.sub[1]) {
					return 1
				} else {
					return 0
				}
			}
		case TYPE_OP_LESS:
			{
				if Operations(op.sub[0]) < Operations(op.sub[1]) {
					return 1
				} else {
					return 0
				}
			}
		case TYPE_OP_EQUAL:
			{
				if Operations(op.sub[0]) == Operations(op.sub[1]) {
					return 1
				} else {
					return 0
				}
			}
		default:
			{
				panic("unknown type ID")
			}
		}
	}
	if lp, ok := p.(*LiteralPacket); ok {
		return lp.literal
	}
	panic("unknown packet type")
}

func Part1(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	packets := ReadInput(r)
	return VersionSum(packets)
}

func Part2(r io.ReadSeeker) int {
	defer timeTrack(time.Now())
	packets := ReadInput(r)
	//PrintPackets(packets, 0)
	return Operations(packets)
}

func PrintPackets(p Packet, depth int) {
	for i := 0; i < depth+1; i++ {
		fmt.Printf("   ")
	}
	fmt.Println(p)
	if op, ok := p.(*OperatorPacket); ok {
		for _, sp := range op.sub {
			PrintPackets(sp, depth+1)
		}
	}
}

func timeTrack(start time.Time) {
	fmt.Printf("(%10s) ", time.Since(start))
}

func expect(expected, actual int, msg string) {
	if expected == actual {
		fmt.Printf("\033[1;32mOK\033[0m")
	} else {
		fmt.Printf("\033[1;31mFAIL\033[0m (expected %d, actual %d)", expected, actual)
	}
	fmt.Println(" ", msg)
}

func main() {
	expect(6, Part1(strings.NewReader("D2FE28")), "Part1 - D2FE28")
	expect(9, Part1(strings.NewReader("38006F45291200")), "Part1 - 38006F45291200")
	expect(14, Part1(strings.NewReader("EE00D40C823060")), "Part1 - EE00D40C823060")
	expect(16, Part1(strings.NewReader("8A004A801A8002F478")), "Part1 - 8A004A801A8002F478")
	expect(12, Part1(strings.NewReader("620080001611562C8802118E34")), "Part1 - 620080001611562C8802118E34")
	expect(23, Part1(strings.NewReader("C0015000016115A2E0802F182340")), "Part1 - C0015000016115A2E0802F182340")
	expect(31, Part1(strings.NewReader("A0016C880162017C3686B18A3D4780")), "Part1 - A0016C880162017C3686B18A3D4780")

	input, err := os.Open("2021/16/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part1 - puzzle", Part1(input))

	expect(3, Part2(strings.NewReader("C200B40A82")), "Part2 - C200B40A82")
	expect(54, Part2(strings.NewReader("04005AC33890")), "Part2 - 04005AC33890")
	expect(7, Part2(strings.NewReader("880086C3E88112")), "Part2 - 880086C3E88112")
	expect(9, Part2(strings.NewReader("CE00C43D881120")), "Part2 - CE00C43D881120")
	expect(1, Part2(strings.NewReader("D8005AC2A8F0")), "Part2 - D8005AC2A8F0")
	expect(1, Part2(strings.NewReader("9C0141080250320F1802104A08")), "Part2 - 9C0141080250320F1802104A08")

	fmt.Println("Part2 - puzzle", Part2(input))
}
