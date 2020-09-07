package twobit

import (
	"bytes"
	"log"
	"testing"
)

// func (service *serviceResults) GenomicInterval(chr string, start, end int) (string, error) {
func TestGenomicInterval(t *testing.T) {
	// mock2bit := fmt.Sprintf("%s/%s", config.DataFolder, "out.2bit")
	mockService := new(twobitResults)

	var err error

	// Insulin 2-bit fasta in bytes - using chr4 as header
	rdr := []byte{67, 39, 65, 26, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 4, 99, 104, 114, 52, 25, 0, 0, 0, 151, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 181, 69, 190, 111, 77, 134, 235, 189, 97, 173, 188, 76, 22, 191, 80, 55, 27, 207, 209, 190, 5, 191, 61, 62, 85, 189, 86, 209, 54, 219, 251, 231, 61, 63, 71, 58, 216, 207, 252, 237, 91, 253, 86, 189, 191, 101, 61, 65, 181, 53, 27, 84, 212, 196, 86, 225, 147, 20, 19, 88, 245, 76, 248, 221, 69, 53, 83, 79, 116, 211, 213, 19, 254, 83, 149, 181, 219, 80, 51, 165, 166, 83, 55, 209, 153, 79, 62, 180, 68, 148, 179, 55, 254, 158, 244, 16, 73, 153, 90, 229, 117, 254, 246, 239, 148, 219, 243, 181, 164, 213, 131, 77, 84, 245, 213, 91, 89, 85, 52, 83, 221, 21, 149, 182, 63, 110, 191, 246, 251, 211, 89, 91, 111, 252, 111, 54, 64, 2, 170, 235, 4, 67, 198, 113, 74, 172, 229, 180, 84, 207, 86, 198, 232, 70, 212, 239, 159, 48, 244, 31, 109, 87, 184, 152, 110, 252, 253, 157, 20, 84, 89, 29, 84, 106, 106, 53, 93, 181, 96, 17, 101, 70, 3, 142, 93, 184, 26, 204, 3, 10, 202, 177, 79, 206, 83, 252, 102, 252, 213, 103, 77, 77, 68, 253, 233, 149, 97, 157, 95, 190, 253, 207, 77, 77, 78, 207, 214, 229, 83, 29, 111, 81, 159, 109, 22, 44, 111, 184, 255, 174, 52, 255, 155, 213, 63, 238, 178, 79, 225, 148, 193, 189, 21, 147, 57, 211, 85, 255, 127, 250, 251, 207, 230, 51, 247, 15, 245, 50, 241, 102, 86, 204, 252, 229, 69, 68, 165, 63, 22, 213, 244, 251, 143, 207, 236, 222, 82, 253, 61, 253, 189, 253, 147, 49, 21, 57, 51, 20, 83, 49, 81, 53, 29, 116, 193, 126, 148, 209, 55, 125, 156, 83, 219, 63, 219, 207, 180, 253, 255, 213, 60, 219, 219, 83, 109, 80, 245, 79, 191, 197, 77, 186, 220, 246, 12, 250, 104, 211, 37, 182, 19, 69, 68, 150, 211, 238, 146, 77, 164, 185, 219, 87, 111, 109, 86, 101, 117, 212, 83, 101, 238, 238, 62, 138, 181, 67, 165, 180}

	referenceFile := bytes.NewReader(rdr)

	mockService.tb, err = NewReader(referenceFile)
	if err != nil {
		log.Fatal(err)
	}

	tt := []struct {
		chr      string
		start    int
		end      int
		expected string
	}{
		{"4", 0, 5, "AGCCC"},   // Near start
		{"4", -10, 5, "AGCCC"}, // Overruns start is clipped
		{"4", 5, 5, ""},        // No bases
		{"4", 1400, 1431, "GAGAGAGATGGAATAAAGCCCTTGAACCAGC"},  // Near end
		{"4", 1400, 10000, "GAGAGAGATGGAATAAAGCCCTTGAACCAGC"}, // Overruns end is clipped
	}
	for i, test := range tt {
		seq, err := mockService.GenomicInterval(test.chr, test.start, test.end)
		if seq != test.expected {
			t.Errorf("expected %v got %v on test number %d", test.expected, seq, i)
		}
		if err != nil {
			t.Errorf("expected nil got %v on test number %d", err, i)
		}
	}
}
