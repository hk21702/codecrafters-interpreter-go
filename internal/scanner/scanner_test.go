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
	expectedStdOut1 := `COMMA , null
DOT . null
LEFT_PAREN ( null
EOF  null
`
	expectedStdErr1 := `[line 1] Error: Unexpected character: $
[line 1] Error: Unexpected character: #
`

	input2 := []byte("={===}")
	expectedStdOut2 := `EQUAL = null
LEFT_BRACE { null
EQUAL_EQUAL == null
EQUAL = null
RIGHT_BRACE } null
EOF  null
`
	expectedStdErr2 := ""

	output1, error1 := captureScanOutput(input1)
	compareOutput(t, expectedStdOut1, output1.String())
	compareOutput(t, expectedStdErr1, error1.String())

	output2, error2 := captureScanOutput(input2)
	compareOutput(t, expectedStdOut2, output2.String())
	compareOutput(t, expectedStdErr2, error2.String())
}

func captureScanOutput(input []byte) (capturedStdout, capturedStderr bytes.Buffer) {
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()
	os.Stdout = wOut
	os.Stderr = wErr

	Scan(input)

	wErr.Close()
	wOut.Close()
	os.Stdout = oldStdout
	os.Stderr = oldStderr

	io.Copy(&capturedStdout, rOut)
	io.Copy(&capturedStderr, rErr)
	return capturedStdout, capturedStderr
}

func compareOutput(t *testing.T, expected, actual string) {
	expectedLines := strings.Split(expected, "\n")
	actualLines := strings.Split(actual, "\n")

	for i, expected_line := range expectedLines {
		if expected_line != actualLines[i] {
			t.Errorf("Mismatch at line %d:\nExpected: %s\nActual: %s", i+1, expected_line, actualLines[i])
		}
	}

	if len(expectedLines) != len(actualLines) {
		t.Errorf("Number of lines mismatch:\nExpected: %d lines\nActual: %d lines", len(expectedLines), len(actualLines))
	}
}
