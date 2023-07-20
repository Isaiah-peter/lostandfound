package api

import (
	"context"

	"github.com/Isaiah-peter/lostandfound/util"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

const (
	cloud_name = "dieusg1qo"
    api_key = "874512347679454"
    api_secret = "VRynsM7TUYOjD9tS4NIZvDbKTBM"
)

var cld, _ = cloudinary.NewFromParams(cloud_name, api_key, api_secret)

func (server *Server) Upload(ctx context.Context, file interface{}) string {
	data, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: util.RandomWord(10)})
	if err != nil {
		panic("upload failed")
	}
	return data.SecureURL
}