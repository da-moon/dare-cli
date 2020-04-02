package dare

import (
	encryptor "github.com/da-moon/coe817-dare/internal/encryptor"
	"io"
)

// EncryptReader returns an io.reader that
// encrypts data with a passed key as it is reading it
// from an io stream (eg socket , file).
func EncryptReader(reader io.Reader, key []byte) (io.Reader, error) {
	return encryptor.New(reader, nil, key)
}

// Encrypt a convenience method that uses encrypt reader
// to encrypt data as it is reading it from io stream
// and would write the encrypted data to another sink
// through an io.writer
func Encrypt(writer io.Writer, reader io.Reader, key []byte) (n int64, err error) {
	encryptReader, err := EncryptReader(reader, key)
	if err != nil {
		return 0, err
	}
	return io.Copy(writer, encryptReader)

}
