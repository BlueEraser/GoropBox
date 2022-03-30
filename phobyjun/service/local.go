package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
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
