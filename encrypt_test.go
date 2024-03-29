package dare_test

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/da-moon/dare-cli/pkg/dare/encryptor"
	assert "github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	testString := []byte("this is a test")
	expectedEncryptionResult := []byte{193, 169, 253, 213, 94, 44, 240, 65, 32, 194, 251, 86, 65, 75, 104, 6, 205, 132, 60, 67, 10, 68, 228, 244, 232, 29, 54, 141, 97, 141}
	nonce := [24]byte{58, 230, 79, 44, 187, 45, 107, 226, 245, 53, 169, 118, 218, 116, 235, 95, 132, 127, 166, 200, 203, 141, 251, 51}
	sharedKey := &[32]byte{199, 156, 103, 110, 157, 5, 107, 139, 94, 138, 53, 214, 74, 100, 211, 97, 106, 48, 11, 179, 200, 19, 244, 108, 138, 167, 49, 163, 156, 176, 66, 64}
	buf := bytes.NewBuffer(testString)
	reader := bufio.NewReader(buf)
	t.Run("NewReader", func(t *testing.T) {
		enc := encryptor.NewReader(
			reader,
			nonce,
			sharedKey,
		)
		result := new(bytes.Buffer)
		n, err := io.Copy(result, enc)
		if err != nil {
			if err != io.EOF {
				assert.NoError(t, err)
			}
		}
		assert.Equal(t, len(expectedEncryptionResult), int(n))
		assert.True(t, bytes.Equal(expectedEncryptionResult, result.Bytes()), fmt.Sprintf("want: %v \n got: %v", expectedEncryptionResult, result.Bytes()))

	})
}
