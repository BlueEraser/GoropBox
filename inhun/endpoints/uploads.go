package endpoints

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	ihttp "github.com/inhun/GoropBox/internal/http"

	"github.com/julienschmidt/httprouter"
)

func (e *Endpoints) Uploads(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var part *multipart.Part
	var n int

	mr, err := r.MultipartReader()
	if err != nil {
		ihttp.ResponseError(w, 500, err.Error())
		return
	}
	chunk := make([]byte, 4096)

	for {
		var filesize int
		var uploaded bool

		if part, err = mr.NextPart(); err != nil {
			if err != io.EOF {
				ihttp.ResponseError(w, 500, err.Error())
			} else {
				ihttp.ResponseOK(w, "success", "a")

			}
			return
		}

		buf := bytes.NewBuffer(nil)
		_, err = io.Copy(buf, part)
		if err != nil {
			ihttp.ResponseError(w, 500, err.Error())
			return
		}

		for !uploaded {
			if n, err = part.Read(chunk); err != nil {
				if err != io.EOF {
					ihttp.ResponseError(w, 500, err.Error())
					return
				}
				uploaded = true
			}

			filesize += n
		}

		output, err := e.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{Bucket: aws.String("goropbox"), Key: aws.String(part.FileName()), Body: buf, ContentLength: int64(filesize)})
		if err != nil {
			fmt.Println("SIBAL")
			ihttp.ResponseError(w, 500, err.Error())
			return
			// log.Fatal(err)
		}
		fmt.Println(output)

	}

}
