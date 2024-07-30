package scanner

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	input1 := []byte(",.$(#")
	expectedOutput1 := `[line 1] Error: Unexpected character: $
[line 1] Error: Unexpected character: #
COMMA , null
DOT . null
LEFT_PAREN ( null
EOF  null`

	input2 := []byte("={===}")
	expectedOutput2 := `EQUAL = null
LEFT_BRACE { null
EQUAL_EQUAL == null
EQUAL = null\n
RIGHT_BRACE } null
EOF  null`

	output1 := captureScanOutput(input1)
	compareOutput(t, expectedOutput1, output1.String())

	output2 := captureScanOutput(input2)
	compareOutput(t, expectedOutput2, output2.String())
}

func captureScanOutput(input []byte) (capturedOutput bytes.Buffer) {
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w

	Scan(input)

	w.Close()
	os.Stdout = oldStdout
	os.Stderr = oldStderr

	io.Copy(&capturedOutput, r)
	return capturedOutput
}

func compareOutput(t *testing.T, expected, actual string) {
	expectedLines := strings.Split(expected, "\n")
	actualLines := strings.Split(expected, "\n")

	for i, expected_line := range expectedLines {
		if expected_line != actualLines[i] {
			t.Errorf("Mismatch at line %d:\nExpected: %s\nActual: %s", i+1, expected_line, actualLines[i])
		}
	}

	if len(expectedLines) != len(actualLines) {
		t.Errorf("Number of lines mismatch:\nExpected: %d lines\nActual: %d lines", len(expectedLines), len(actualLines))
	}
}
