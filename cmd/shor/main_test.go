package main

import "os"

func Example() {
	os.Args = []string{"shor", "--N=21", "--a=5", "--seed=1"}
	main()

	// Output:
	// N=21, a=5, t=4

	// initial state
	// [0000 00001][  0   1]( 1.0000 0.0000i): 1.0000: ********************************

	// create superposition
	// [0000 00001][  0   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [0001 00001][  1   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [0010 00001][  2   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [0011 00001][  3   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [0100 00001][  4   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [0101 00001][  5   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [0110 00001][  6   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [0111 00001][  7   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [1000 00001][  8   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [1001 00001][  9   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [1010 00001][ 10   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [1011 00001][ 11   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [1100 00001][ 12   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [1101 00001][ 13   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [1110 00001][ 14   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [1111 00001][ 15   1]( 0.2500 0.0000i): 0.0625: ********************************

	// apply controlled-U
	// [0000 00001][  0   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [0001 10000][  1  16]( 0.2500 0.0000i): 0.0625: ********************************
	// [0010 00100][  2   4]( 0.2500 0.0000i): 0.0625: ********************************
	// [0011 00001][  3   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [0100 10000][  4  16]( 0.2500 0.0000i): 0.0625: ********************************
	// [0101 00100][  5   4]( 0.2500 0.0000i): 0.0625: ********************************
	// [0110 00001][  6   1]( 0.2500 0.0000i): 0.0625: ********************************
	// [0111 10000][  7  16]( 0.2500 0.0000i): 0.0625: ********************************
	// [1000 10001][  8  17]( 0.2500 0.0000i): 0.0625: ********************************
	// [1001 10100][  9  20]( 0.2500 0.0000i): 0.0625: ********************************
	// [1010 00101][ 10   5]( 0.2500 0.0000i): 0.0625: ********************************
	// [1011 10001][ 11  17]( 0.2500 0.0000i): 0.0625: ********************************
	// [1100 10100][ 12  20]( 0.2500 0.0000i): 0.0625: ********************************
	// [1101 00101][ 13   5]( 0.2500 0.0000i): 0.0625: ********************************
	// [1110 10001][ 14  17]( 0.2500 0.0000i): 0.0625: ********************************
	// [1111 10100][ 15  20]( 0.2500 0.0000i): 0.0625: ********************************

	// apply inverse QFT
	// [0000 00001][  0   1]( 0.1875 0.0000i): 0.0352: ********************************
	// [0000 00100][  0   4]( 0.1250 0.0000i): 0.0156: **************
	// [0000 00101][  0   5]( 0.1250 0.0000i): 0.0156: **************
	// [0000 10000][  0  16]( 0.1875 0.0000i): 0.0352: ********************************
	// [0000 10001][  0  17]( 0.1875 0.0000i): 0.0352: ********************************
	// [0000 10100][  0  20]( 0.1875 0.0000i): 0.0352: ********************************
	// [0001 00001][  1   1]( 0.0183-0.0183i): 0.0007:
	// [0001 00100][  1   4](-0.0442 0.0183i): 0.0023: **
	// [0001 00101][  1   5](-0.0478 0.0000i): 0.0023: **
	// [0001 10000][  1  16]( 0.0259 0.0000i): 0.0007:
	// [0001 10001][  1  17]( 0.0239-0.0099i): 0.0007:
	// [0001 10100][  1  20]( 0.0239 0.0099i): 0.0007:
	// [0010 00001][  2   1]( 0.0000-0.0625i): 0.0039: ***
	// [0010 00100][  2   4](-0.0625 0.0625i): 0.0078: *******
	// [0010 00101][  2   5](-0.0884 0.0000i): 0.0078: *******
	// [0010 10000][  2  16]( 0.0625 0.0000i): 0.0039: ***
	// [0010 10001][  2  17]( 0.0442-0.0442i): 0.0039: ***
	// [0010 10100][  2  20]( 0.0442 0.0442i): 0.0039: ***
	// [0011 00001][  3   1]( 0.1067 0.1067i): 0.0228: ********************
	// [0011 00100][  3   4]( 0.0442-0.1067i): 0.0133: ************
	// [0011 00101][  3   5]( 0.1155 0.0000i): 0.0133: ************
	// [0011 10000][  3  16](-0.1509 0.0000i): 0.0228: ********************
	// [0011 10001][  3  17](-0.0577 0.1394i): 0.0228: ********************
	// [0011 10100][  3  20](-0.0577-0.1394i): 0.0228: ********************
	// [0100 00001][  4   1]( 0.0625 0.0000i): 0.0039: ***
	// [0100 10000][  4  16](-0.0625 0.0000i): 0.0039: ***
	// [0100 10001][  4  17]( 0.0000 0.0625i): 0.0039: ***
	// [0100 10100][  4  20]( 0.0000-0.0625i): 0.0039: ***
	// [0101 00001][  5   1]( 0.1067-0.1067i): 0.0228: ********************
	// [0101 00100][  5   4]( 0.0442 0.1067i): 0.0133: ************
	// [0101 00101][  5   5](-0.1155 0.0000i): 0.0133: ************
	// [0101 10000][  5  16](-0.1509 0.0000i): 0.0228: ********************
	// [0101 10001][  5  17]( 0.0577 0.1394i): 0.0228: ********************
	// [0101 10100][  5  20]( 0.0577-0.1394i): 0.0228: ********************
	// [0110 00001][  6   1]( 0.0000 0.0625i): 0.0039: ***
	// [0110 00100][  6   4](-0.0625-0.0625i): 0.0078: *******
	// [0110 00101][  6   5]( 0.0884 0.0000i): 0.0078: *******
	// [0110 10000][  6  16]( 0.0625 0.0000i): 0.0039: ***
	// [0110 10001][  6  17](-0.0442-0.0442i): 0.0039: ***
	// [0110 10100][  6  20](-0.0442 0.0442i): 0.0039: ***
	// [0111 00001][  7   1]( 0.0183 0.0183i): 0.0007:
	// [0111 00100][  7   4](-0.0442-0.0183i): 0.0023: **
	// [0111 00101][  7   5]( 0.0478 0.0000i): 0.0023: **
	// [0111 10000][  7  16]( 0.0259 0.0000i): 0.0007:
	// [0111 10001][  7  17](-0.0239-0.0099i): 0.0007:
	// [0111 10100][  7  20](-0.0239 0.0099i): 0.0007:
	// [1000 00001][  8   1]( 0.1875 0.0000i): 0.0352: ********************************
	// [1000 00100][  8   4]( 0.1250 0.0000i): 0.0156: **************
	// [1000 00101][  8   5](-0.1250 0.0000i): 0.0156: **************
	// [1000 10000][  8  16]( 0.1875 0.0000i): 0.0352: ********************************
	// [1000 10001][  8  17](-0.1875 0.0000i): 0.0352: ********************************
	// [1000 10100][  8  20](-0.1875 0.0000i): 0.0352: ********************************
	// [1001 00001][  9   1]( 0.0183-0.0183i): 0.0007:
	// [1001 00100][  9   4](-0.0442 0.0183i): 0.0023: **
	// [1001 00101][  9   5]( 0.0478 0.0000i): 0.0023: **
	// [1001 10000][  9  16]( 0.0259 0.0000i): 0.0007:
	// [1001 10001][  9  17](-0.0239 0.0099i): 0.0007:
	// [1001 10100][  9  20](-0.0239-0.0099i): 0.0007:
	// [1010 00001][ 10   1]( 0.0000-0.0625i): 0.0039: ***
	// [1010 00100][ 10   4](-0.0625 0.0625i): 0.0078: *******
	// [1010 00101][ 10   5]( 0.0884 0.0000i): 0.0078: *******
	// [1010 10000][ 10  16]( 0.0625 0.0000i): 0.0039: ***
	// [1010 10001][ 10  17](-0.0442 0.0442i): 0.0039: ***
	// [1010 10100][ 10  20](-0.0442-0.0442i): 0.0039: ***
	// [1011 00001][ 11   1]( 0.1067 0.1067i): 0.0228: ********************
	// [1011 00100][ 11   4]( 0.0442-0.1067i): 0.0133: ************
	// [1011 00101][ 11   5](-0.1155 0.0000i): 0.0133: ************
	// [1011 10000][ 11  16](-0.1509 0.0000i): 0.0228: ********************
	// [1011 10001][ 11  17]( 0.0577-0.1394i): 0.0228: ********************
	// [1011 10100][ 11  20]( 0.0577 0.1394i): 0.0228: ********************
	// [1100 00001][ 12   1]( 0.0625 0.0000i): 0.0039: ***
	// [1100 10000][ 12  16](-0.0625 0.0000i): 0.0039: ***
	// [1100 10001][ 12  17]( 0.0000-0.0625i): 0.0039: ***
	// [1100 10100][ 12  20]( 0.0000 0.0625i): 0.0039: ***
	// [1101 00001][ 13   1]( 0.1067-0.1067i): 0.0228: ********************
	// [1101 00100][ 13   4]( 0.0442 0.1067i): 0.0133: ************
	// [1101 00101][ 13   5]( 0.1155 0.0000i): 0.0133: ************
	// [1101 10000][ 13  16](-0.1509 0.0000i): 0.0228: ********************
	// [1101 10001][ 13  17](-0.0577-0.1394i): 0.0228: ********************
	// [1101 10100][ 13  20](-0.0577 0.1394i): 0.0228: ********************
	// [1110 00001][ 14   1]( 0.0000 0.0625i): 0.0039: ***
	// [1110 00100][ 14   4](-0.0625-0.0625i): 0.0078: *******
	// [1110 00101][ 14   5](-0.0884 0.0000i): 0.0078: *******
	// [1110 10000][ 14  16]( 0.0625 0.0000i): 0.0039: ***
	// [1110 10001][ 14  17]( 0.0442 0.0442i): 0.0039: ***
	// [1110 10100][ 14  20]( 0.0442-0.0442i): 0.0039: ***
	// [1111 00001][ 15   1]( 0.0183 0.0183i): 0.0007:
	// [1111 00100][ 15   4](-0.0442-0.0183i): 0.0023: **
	// [1111 00101][ 15   5](-0.0478 0.0000i): 0.0023: **
	// [1111 10000][ 15  16]( 0.0259 0.0000i): 0.0007:
	// [1111 10001][ 15  17]( 0.0239 0.0099i): 0.0007:
	// [1111 10100][ 15  20]( 0.0239-0.0099i): 0.0007:

	// measure reg1
	// [0000 10001][  0  17]( 0.4330 0.0000i): 0.1875: ********************************
	// [0001 10001][  1  17]( 0.0552-0.0229i): 0.0036:
	// [0010 10001][  2  17]( 0.1021-0.1021i): 0.0208: ***
	// [0011 10001][  3  17](-0.1334 0.3219i): 0.1214: ********************
	// [0100 10001][  4  17]( 0.0000 0.1443i): 0.0208: ***
	// [0101 10001][  5  17]( 0.1334 0.3219i): 0.1214: ********************
	// [0110 10001][  6  17](-0.1021-0.1021i): 0.0208: ***
	// [0111 10001][  7  17](-0.0552-0.0229i): 0.0036:
	// [1000 10001][  8  17](-0.4330 0.0000i): 0.1875: ********************************
	// [1001 10001][  9  17](-0.0552 0.0229i): 0.0036:
	// [1010 10001][ 10  17](-0.1021 0.1021i): 0.0208: ***
	// [1011 10001][ 11  17]( 0.1334-0.3219i): 0.1214: ********************
	// [1100 10001][ 12  17]( 0.0000-0.1443i): 0.0208: ***
	// [1101 10001][ 13  17](-0.1334-0.3219i): 0.1214: ********************
	// [1110 10001][ 14  17]( 0.1021 0.1021i): 0.0208: ***
	// [1111 10001][ 15  17]( 0.0552 0.0229i): 0.0036:

	// * i= 0: N=21, a=5. s/r=11/16 ([1 0 1 1]=0.688). p=3, q=1.
	// * i= 1: N=21, a=5. s/r=11/16 ([1 0 1 1]=0.688). p=3, q=1.
	// * i= 2: N=21, a=5. s/r=11/16 ([1 0 1 1]=0.688). p=3, q=1.
	// * i= 3: N=21, a=5. s/r=11/16 ([1 0 1 1]=0.688). p=3, q=1.
	// * i= 4: N=21, a=5. s/r=11/16 ([1 0 1 1]=0.688). p=3, q=1.
	// * i= 5: N=21, a=5. s/r=11/16 ([1 0 1 1]=0.688). p=3, q=1.
	// * i= 6: N=21, a=5. s/r=11/16 ([1 0 1 1]=0.688). p=3, q=1.
	// * i= 7: N=21, a=5. s/r=11/16 ([1 0 1 1]=0.688). p=3, q=1.
	// * i= 8: N=21, a=5. s/r=11/16 ([1 0 1 1]=0.688). p=3, q=1.
	// * i= 9: N=21, a=5. s/r=11/16 ([1 0 1 1]=0.688). p=3, q=1.
}