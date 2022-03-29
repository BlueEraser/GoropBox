package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	ihttp "github.com/inhun/GoropBox/internal/http"

	"github.com/julienschmidt/httprouter"
)

func (e *Endpoints) Downloads(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	presignclient := s3.NewPresignClient(e.S3Client)
	getobjectoutput, err := presignclient.PresignGetObject(context.TODO(), &s3.GetObjectInput{Bucket: aws.String("goropbox"), Key: aws.String("demisoda apple.jpg")})
	if err != nil {
		ihttp.ResponseError(w, 500, err.Error())
		return
	}
	fmt.Println(getobjectoutput)
	ihttp.ResponseOK(w, "success", getobjectoutput)
	return
}
