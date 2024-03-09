package rpc

import (
	"context"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Error"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/File"
	"github.com/hoangphuc28/CoursesOnline/File-Service/config"
	"github.com/hoangphuc28/CoursesOnline/File-Service/pkg/common"
	"github.com/hoangphuc28/CoursesOnline/File-Service/pkg/upload"
	"io"
)

type uploadHandler struct {
	provider upload.UploadProvider
	cf       *config.Config
	File.UnimplementedFileServiceServer
	FirebaseProvider upload.FireBaseProvider
}

func HandleError(err error) *Error.ErrorResponse {
	if errors, ok := err.(*common.AppError); ok {
		return &Error.ErrorResponse{
			Code:    int64(errors.StatusCode),
			Message: errors.Message,
		}
	}
	appErr := common.ErrInternal(err.(error))
	return &Error.ErrorResponse{
		Code:    int64(appErr.StatusCode),
		Message: appErr.Message,
	}
}
func NewUploadHandler(provider upload.UploadProvider, cf *config.Config) *uploadHandler {
	return &uploadHandler{provider: provider, cf: cf}
}
func (hdl *uploadHandler) UploadAvatar(ctx context.Context, rq *File.UploadAvatarRequest) (*File.UploadAvatarResponse, error) {
	//if err := hdl.s3Provider.DeleteFile(rq.OldUrl); err != nil {
	//	fmt.Println(err)
	//	return &pb.UploadAvatarResponse{
	//		Error: HandleError(err),
	//	}, nil
	//}
	//url, err := hdl.provider.UploadFile(
	//	ctx,
	//	rq.File.Content,
	//	fmt.Sprintf("%s/%s", rq.File.Folder, rq.File.FileName),
	//)
	url, err := hdl.provider.UploadFileFireBase(
		rq.File.Content,
		"user/"+rq.File.FileName,
	)
	if err != nil {
		fmt.Println(err)
		return &File.UploadAvatarResponse{
			Error: HandleError(err),
		}, nil
	}

	return &File.UploadAvatarResponse{
		Url: url,
	}, nil
}

func (hdl *uploadHandler) UploadAsset(stream File.FileService_UploadAssetServer) error {
	var file File.File
	rq, err := stream.Recv()
	file.FileName = rq.File.FileName
	file.Size = rq.File.Size
	file.Folder = rq.File.Folder
	if err != nil {
		return err
	}
	for {
		rq, err = stream.Recv()
		if err == io.EOF {
			url, err := hdl.provider.UploadFileFireBase(
				file.Content,
				fmt.Sprintf("%s/%s", file.Folder, file.FileName),
			)
			if err != nil {
				return err
			}
			var video File.Video
			//if file.Folder == "video" {
			//	v, err := utils.GetDurationVideo(hdl.cf, url)
			//	if err != nil {
			//		return err
			//	}
			//	video.Duration = int64(math.Ceil(v.Duration))
			//	video.Unit = "seconds"
			//}
			if err = stream.Send(&File.UploadAssetResponse{
				Url:   url,
				Video: &video,
			}); err != nil {
				return err
			}
			return nil
		}
		if err != nil {
			return err
		}
		file.Content = append(file.Content, rq.File.Content...)
	}

}
