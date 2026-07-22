package hashs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper to compute CRC8 using the exported type (single pass)
func computeCRC8(data []byte) Crc8 {
	var c Crc8
	written, err := c.Write(data)
	// sanity: no error and all bytes consumed
	if err != nil || written != len(data) {
		panic("unexpected write failure in helper")
	}
	return c
}

func TestCrc8Empty(t *testing.T) {
	var c Crc8

	assert.Equal(t, uint8(0x00), uint8(c))
}

func TestCrc8WriteEmpty(t *testing.T) {
	var c Crc8
	before := c
	written, err := c.Write([]byte{})

	assert.NoError(t, err)
	assert.Equal(t, 0, written)
	assert.Equal(t, before, c)
}

func TestCrc8SingleZero(t *testing.T) {
	c := computeCRC8([]byte{0x00})

	assert.Equal(t, uint8(0x00), uint8(c))
}

func TestCrc8SingleFF(t *testing.T) {
	c := computeCRC8([]byte{0xFF})

	assert.Equal(t, uint8(0xF3), uint8(c))
}

func TestCrc8Sequence123456789(t *testing.T) {
	c := computeCRC8([]byte("123456789"))

	assert.Equal(t, uint8(0xF4), uint8(c))
}

// TODO
//func TestCrc8Sequence1To5(t *testing.T) {
//	c := computeCRC8([]byte{1, 2, 3, 4, 5})
//	assert.Equal(t, uint8(0x16), uint8(c))
//}

func TestCrc8MultiPartWriteConsistency(t *testing.T) {
	data := []byte("123456789")
	// one shot
	c1 := computeCRC8(data)
	// multi part
	var c2 Crc8
	w1, err1 := c2.Write([]byte("1234"))
	w2, err2 := c2.Write([]byte("56789"))

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Equal(t, 4, w1)
	assert.Equal(t, 5, w2)
	assert.Equal(t, c1, c2)
}

func TestCrc8SumAppends(t *testing.T) {
	var c Crc8
	_, _ = c.Write([]byte("123456789"))
	b := []byte{0xAA, 0xBB}
	orig := append([]byte{}, b...)       // keep original
	res := c.Sum(append([]byte{}, b...)) // copy input slice to ensure no aliasing assumptions

	assert.Equal(t, orig, b) // ensure original slice untouched
	assert.Equal(t, len(b)+1, len(res))
	assert.Equal(t, byte(0xF4), res[len(res)-1])
	assert.Equal(t, []byte{0xAA, 0xBB, 0xF4}, res)
}

func TestCrc8Reset(t *testing.T) {
	var c Crc8
	_, _ = c.Write([]byte{0xFF})

	assert.NotEqual(t, uint8(0x00), uint8(c))
	c.Reset()
	assert.Equal(t, uint8(0x00), uint8(c))
}

func TestCrc8SizeAndBlockSize(t *testing.T) {
	var c Crc8

	assert.Equal(t, 1, c.Size())
	assert.Equal(t, 1, c.BlockSize())
}
