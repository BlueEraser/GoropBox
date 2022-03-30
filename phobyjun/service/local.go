package service

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"phobyjun/model"
	"strings"
)

const (
	baseDir    = "uploaded"
	bufferSize = 16 * 1024
	ivSize     = 16
	V1         = 0x1
	hmacSize   = sha512.Size
)

func UploadFileToLocal(fileDto *model.File, file *multipart.FileHeader, aesKey, hmacKey []byte) (*model.File, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(src)

	encryptedName := encryptFileNameDir(fileDto)
	fileDto.EncryptedName = encryptedName

	dstDir := strings.Join([]string{
		baseDir,
		encryptedName,
	}, "/")

	dst, err := os.Create(dstDir)
	if err != nil {
		return nil, err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dst)

	if err := encryptFile(src, dst, aesKey, hmacKey); err != nil {
		return nil, err
	}

	return fileDto, nil
}

func DownloadFileFromLocal(fileDto *model.File, aesKey, hmacKey []byte) ([]byte, error) {
	encryptedName := fileDto.EncryptedName
	srcDir := strings.Join([]string{
		baseDir,
		encryptedName,
	}, "/")

	src, err := os.Open(srcDir)
	if err != nil {
		return nil, err
	}
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(src)

	dstDir := strings.Join([]string{
		baseDir,
		base64.StdEncoding.EncodeToString([]byte(encryptedName)),
	}, "/")

	dst, err := os.Create(dstDir)
	if err != nil {
		return nil, err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dst)

	if err := decryptFile(src, dst, aesKey, hmacKey); err != nil {
		return nil, err
	}

	return []byte(dst.Name()), nil
}

func encryptFileNameDir(fileDto *model.File) string {
	fileNameDir := fileDto.FileNameDir

	return base64.StdEncoding.EncodeToString([]byte(fileNameDir))
}

func encryptFile(src io.Reader, dst io.Writer, aesKey, hmacKey []byte) error {
	iv := make([]byte, ivSize)
	_, err := rand.Read(iv)
	if err != nil {
		return err
	}

	AES, err := aes.NewCipher(aesKey)
	if err != nil {
		return err
	}

	ctr := cipher.NewCTR(AES, iv)
	HMAC := hmac.New(sha512.New, hmacKey)

	_, err = dst.Write([]byte{V1})
	if err != nil {
		return err
	}

	writer := io.MultiWriter(dst, HMAC)

	_, err = writer.Write(iv)
	if err != nil {
		return err
	}

	buffer := make([]byte, bufferSize)
	for {
		n, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}

		if n != 0 {
			outBuffer := make([]byte, n)
			ctr.XORKeyStream(outBuffer, buffer[:n])
			_, err := writer.Write(outBuffer)
			if err != nil {
				return err
			}
		}

		if err == io.EOF {
			break
		}
	}

	_, err = dst.Write(HMAC.Sum(nil))
	return err
}

func decryptFile(src io.Reader, dst io.Writer, aesKey, hmacKey []byte) error {
	var version int8

	err := binary.Read(src, binary.LittleEndian, &version)
	if err != nil {
		return err
	}

	iv := make([]byte, ivSize)
	_, err = io.ReadFull(src, iv)
	if err != nil {
		return err
	}

	AES, err := aes.NewCipher(aesKey)
	if err != nil {
		return err
	}

	ctr := cipher.NewCTR(AES, iv)
	h := hmac.New(sha512.New, hmacKey)
	h.Write(iv)
	mac := make([]byte, hmacSize)

	w := dst

	buf := bufio.NewReaderSize(src, bufferSize)
	var limit int
	var b []byte
	for {
		b, err = buf.Peek(bufferSize)
		if err != nil && err != io.EOF {
			return err
		}

		limit = len(b) - hmacSize

		if err == io.EOF {

			left := buf.Buffered()
			if left < hmacSize {
				return errors.New("not enough left")
			}

			copy(mac, b[left-hmacSize:left])

			if left == hmacSize {
				break
			}
		}

		h.Write(b[:limit])

		outBuf := make([]byte, int64(limit))
		_, err = buf.Read(b[:limit])
		if err != nil {
			return err
		}
		ctr.XORKeyStream(outBuf, b[:limit])
		_, err = w.Write(outBuf)
		if err != nil {
			return err
		}

		if err == io.EOF {
			break
		}
	}

	return nil
}
