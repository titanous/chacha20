// The ChaCha20 core transform.
// An unrolled and inlined implementation in pure Go.

package chacha20

func core(input, output *[stateSize]uint32) {
	var (
		x00 = input[0]
		x01 = input[1]
		x02 = input[2]
		x03 = input[3]
		x04 = input[4]
		x05 = input[5]
		x06 = input[6]
		x07 = input[7]
		x08 = input[8]
		x09 = input[9]
		x10 = input[10]
		x11 = input[11]
		x12 = input[12]
		x13 = input[13]
		x14 = input[14]
		x15 = input[15]
	)

	var x uint32

	// Rounds 1 and 2
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 16) | (x >> 16)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 12) | (x >> 20)
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 8) | (x >> 24)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 7) | (x >> 25)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 16) | (x >> 16)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 12) | (x >> 20)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 8) | (x >> 24)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 7) | (x >> 25)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 16) | (x >> 16)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 12) | (x >> 20)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 8) | (x >> 24)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 7) | (x >> 25)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 16) | (x >> 16)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 12) | (x >> 20)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 8) | (x >> 24)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 7) | (x >> 25)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 16) | (x >> 16)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 12) | (x >> 20)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 8) | (x >> 24)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 7) | (x >> 25)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 16) | (x >> 16)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 12) | (x >> 20)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 8) | (x >> 24)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 7) | (x >> 25)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 16) | (x >> 16)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 12) | (x >> 20)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 8) | (x >> 24)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 7) | (x >> 25)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 16) | (x >> 16)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 12) | (x >> 20)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 8) | (x >> 24)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 7) | (x >> 25)

	// Rounds 3 and 4
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 16) | (x >> 16)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 12) | (x >> 20)
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 8) | (x >> 24)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 7) | (x >> 25)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 16) | (x >> 16)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 12) | (x >> 20)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 8) | (x >> 24)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 7) | (x >> 25)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 16) | (x >> 16)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 12) | (x >> 20)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 8) | (x >> 24)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 7) | (x >> 25)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 16) | (x >> 16)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 12) | (x >> 20)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 8) | (x >> 24)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 7) | (x >> 25)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 16) | (x >> 16)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 12) | (x >> 20)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 8) | (x >> 24)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 7) | (x >> 25)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 16) | (x >> 16)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 12) | (x >> 20)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 8) | (x >> 24)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 7) | (x >> 25)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 16) | (x >> 16)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 12) | (x >> 20)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 8) | (x >> 24)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 7) | (x >> 25)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 16) | (x >> 16)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 12) | (x >> 20)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 8) | (x >> 24)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 7) | (x >> 25)

	// Rounds 5 and 6
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 16) | (x >> 16)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 12) | (x >> 20)
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 8) | (x >> 24)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 7) | (x >> 25)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 16) | (x >> 16)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 12) | (x >> 20)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 8) | (x >> 24)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 7) | (x >> 25)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 16) | (x >> 16)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 12) | (x >> 20)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 8) | (x >> 24)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 7) | (x >> 25)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 16) | (x >> 16)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 12) | (x >> 20)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 8) | (x >> 24)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 7) | (x >> 25)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 16) | (x >> 16)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 12) | (x >> 20)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 8) | (x >> 24)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 7) | (x >> 25)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 16) | (x >> 16)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 12) | (x >> 20)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 8) | (x >> 24)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 7) | (x >> 25)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 16) | (x >> 16)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 12) | (x >> 20)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 8) | (x >> 24)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 7) | (x >> 25)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 16) | (x >> 16)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 12) | (x >> 20)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 8) | (x >> 24)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 7) | (x >> 25)

	// Rounds 7 and 8
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 16) | (x >> 16)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 12) | (x >> 20)
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 8) | (x >> 24)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 7) | (x >> 25)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 16) | (x >> 16)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 12) | (x >> 20)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 8) | (x >> 24)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 7) | (x >> 25)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 16) | (x >> 16)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 12) | (x >> 20)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 8) | (x >> 24)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 7) | (x >> 25)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 16) | (x >> 16)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 12) | (x >> 20)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 8) | (x >> 24)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 7) | (x >> 25)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 16) | (x >> 16)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 12) | (x >> 20)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 8) | (x >> 24)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 7) | (x >> 25)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 16) | (x >> 16)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 12) | (x >> 20)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 8) | (x >> 24)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 7) | (x >> 25)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 16) | (x >> 16)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 12) | (x >> 20)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 8) | (x >> 24)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 7) | (x >> 25)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 16) | (x >> 16)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 12) | (x >> 20)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 8) | (x >> 24)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 7) | (x >> 25)

	// Rounds 9 and 10
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 16) | (x >> 16)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 12) | (x >> 20)
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 8) | (x >> 24)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 7) | (x >> 25)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 16) | (x >> 16)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 12) | (x >> 20)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 8) | (x >> 24)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 7) | (x >> 25)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 16) | (x >> 16)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 12) | (x >> 20)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 8) | (x >> 24)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 7) | (x >> 25)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 16) | (x >> 16)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 12) | (x >> 20)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 8) | (x >> 24)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 7) | (x >> 25)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 16) | (x >> 16)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 12) | (x >> 20)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 8) | (x >> 24)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 7) | (x >> 25)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 16) | (x >> 16)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 12) | (x >> 20)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 8) | (x >> 24)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 7) | (x >> 25)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 16) | (x >> 16)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 12) | (x >> 20)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 8) | (x >> 24)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 7) | (x >> 25)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 16) | (x >> 16)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 12) | (x >> 20)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 8) | (x >> 24)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 7) | (x >> 25)

	// Rounds 11 and 12
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 16) | (x >> 16)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 12) | (x >> 20)
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 8) | (x >> 24)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 7) | (x >> 25)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 16) | (x >> 16)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 12) | (x >> 20)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 8) | (x >> 24)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 7) | (x >> 25)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 16) | (x >> 16)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 12) | (x >> 20)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 8) | (x >> 24)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 7) | (x >> 25)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 16) | (x >> 16)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 12) | (x >> 20)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 8) | (x >> 24)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 7) | (x >> 25)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 16) | (x >> 16)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 12) | (x >> 20)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 8) | (x >> 24)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 7) | (x >> 25)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 16) | (x >> 16)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 12) | (x >> 20)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 8) | (x >> 24)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 7) | (x >> 25)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 16) | (x >> 16)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 12) | (x >> 20)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 8) | (x >> 24)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 7) | (x >> 25)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 16) | (x >> 16)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 12) | (x >> 20)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 8) | (x >> 24)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 7) | (x >> 25)

	// Rounds 13 and 14
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 16) | (x >> 16)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 12) | (x >> 20)
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 8) | (x >> 24)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 7) | (x >> 25)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 16) | (x >> 16)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 12) | (x >> 20)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 8) | (x >> 24)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 7) | (x >> 25)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 16) | (x >> 16)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 12) | (x >> 20)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 8) | (x >> 24)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 7) | (x >> 25)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 16) | (x >> 16)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 12) | (x >> 20)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 8) | (x >> 24)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 7) | (x >> 25)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 16) | (x >> 16)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 12) | (x >> 20)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 8) | (x >> 24)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 7) | (x >> 25)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 16) | (x >> 16)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 12) | (x >> 20)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 8) | (x >> 24)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 7) | (x >> 25)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 16) | (x >> 16)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 12) | (x >> 20)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 8) | (x >> 24)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 7) | (x >> 25)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 16) | (x >> 16)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 12) | (x >> 20)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 8) | (x >> 24)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 7) | (x >> 25)

	// Rounds 15 and 16
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 16) | (x >> 16)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 12) | (x >> 20)
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 8) | (x >> 24)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 7) | (x >> 25)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 16) | (x >> 16)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 12) | (x >> 20)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 8) | (x >> 24)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 7) | (x >> 25)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 16) | (x >> 16)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 12) | (x >> 20)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 8) | (x >> 24)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 7) | (x >> 25)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 16) | (x >> 16)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 12) | (x >> 20)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 8) | (x >> 24)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 7) | (x >> 25)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 16) | (x >> 16)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 12) | (x >> 20)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 8) | (x >> 24)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 7) | (x >> 25)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 16) | (x >> 16)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 12) | (x >> 20)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 8) | (x >> 24)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 7) | (x >> 25)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 16) | (x >> 16)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 12) | (x >> 20)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 8) | (x >> 24)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 7) | (x >> 25)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 16) | (x >> 16)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 12) | (x >> 20)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 8) | (x >> 24)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 7) | (x >> 25)

	// Rounds 17 and 18
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 16) | (x >> 16)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 12) | (x >> 20)
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 8) | (x >> 24)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 7) | (x >> 25)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 16) | (x >> 16)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 12) | (x >> 20)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 8) | (x >> 24)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 7) | (x >> 25)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 16) | (x >> 16)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 12) | (x >> 20)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 8) | (x >> 24)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 7) | (x >> 25)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 16) | (x >> 16)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 12) | (x >> 20)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 8) | (x >> 24)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 7) | (x >> 25)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 16) | (x >> 16)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 12) | (x >> 20)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 8) | (x >> 24)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 7) | (x >> 25)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 16) | (x >> 16)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 12) | (x >> 20)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 8) | (x >> 24)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 7) | (x >> 25)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 16) | (x >> 16)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 12) | (x >> 20)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 8) | (x >> 24)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 7) | (x >> 25)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 16) | (x >> 16)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 12) | (x >> 20)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 8) | (x >> 24)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 7) | (x >> 25)

	// Rounds 19 and 20
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 16) | (x >> 16)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 12) | (x >> 20)
	x00 += x04
	x = x12 ^ x00
	x12 = (x << 8) | (x >> 24)
	x08 += x12
	x = x04 ^ x08
	x04 = (x << 7) | (x >> 25)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 16) | (x >> 16)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 12) | (x >> 20)
	x01 += x05
	x = x13 ^ x01
	x13 = (x << 8) | (x >> 24)
	x09 += x13
	x = x05 ^ x09
	x05 = (x << 7) | (x >> 25)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 16) | (x >> 16)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 12) | (x >> 20)
	x02 += x06
	x = x14 ^ x02
	x14 = (x << 8) | (x >> 24)
	x10 += x14
	x = x06 ^ x10
	x06 = (x << 7) | (x >> 25)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 16) | (x >> 16)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 12) | (x >> 20)
	x03 += x07
	x = x15 ^ x03
	x15 = (x << 8) | (x >> 24)
	x11 += x15
	x = x07 ^ x11
	x07 = (x << 7) | (x >> 25)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 16) | (x >> 16)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 12) | (x >> 20)
	x00 += x05
	x = x15 ^ x00
	x15 = (x << 8) | (x >> 24)
	x10 += x15
	x = x05 ^ x10
	x05 = (x << 7) | (x >> 25)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 16) | (x >> 16)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 12) | (x >> 20)
	x01 += x06
	x = x12 ^ x01
	x12 = (x << 8) | (x >> 24)
	x11 += x12
	x = x06 ^ x11
	x06 = (x << 7) | (x >> 25)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 16) | (x >> 16)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 12) | (x >> 20)
	x02 += x07
	x = x13 ^ x02
	x13 = (x << 8) | (x >> 24)
	x08 += x13
	x = x07 ^ x08
	x07 = (x << 7) | (x >> 25)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 16) | (x >> 16)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 12) | (x >> 20)
	x03 += x04
	x = x14 ^ x03
	x14 = (x << 8) | (x >> 24)
	x09 += x14
	x = x04 ^ x09
	x04 = (x << 7) | (x >> 25)

	output[0] = x00 + input[0]
	output[1] = x01 + input[1]
	output[2] = x02 + input[2]
	output[3] = x03 + input[3]
	output[4] = x04 + input[4]
	output[5] = x05 + input[5]
	output[6] = x06 + input[6]
	output[7] = x07 + input[7]
	output[8] = x08 + input[8]
	output[9] = x09 + input[9]
	output[10] = x10 + input[10]
	output[11] = x11 + input[11]
	output[12] = x12 + input[12]
	output[13] = x13 + input[13]
	output[14] = x14 + input[14]
	output[15] = x15 + input[15]
}
