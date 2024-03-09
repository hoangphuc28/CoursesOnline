package rpc

import (
	"context"
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Error"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/User"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/config"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"
	client2 "github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/client"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/common"
)

type CoursesUseCase interface {
	CreateCourse(data *model.Course) error
	GetCourses(limit int, page int) (coursesResponse Course.GetCoursesResponse, err error)
	GetCourse(fakeId string) (*Course.GetCourseResponse, error)
	GetAllCategories() (*Course.GetAllCategoriesResponse, error)
	NewEnrollment(userId, courseId string) error
	GetCourseContent(userId, courseId string) (*Course.GetCourseContentResponse, error)
	GetEnrollments(userId string) (*Course.GetEnrollmentsResponse, error)
	NewCourse(rq *Course.CreateCourseRequest, instructorId string) (string, error)
	GetPrices() (*Course.GetPricesResponse, error)
	UpdateCourse(rq *Course.UpdateCourseRequest) error
	PublishCourse(ctx context.Context, rq *Course.PublishCourseRequest) ([]string, error)
	GetCoursesWithInstructor(rq *Course.GetCourseWithInstructorRequest) (*Course.GetCourseWithInstructorResponse, error)
	DeleteCourse(req *Course.DeleteCourseRequest) error

	UpdateSection(req *Course.UpdateSectionRequest) error
	DeleteSection(req *Course.DeleteSectionRequest) error
	CreateSection(req *Course.CreateSectionRequest) (*Course.CreateSectionResponse, error)
	GetSections(request *Course.GetSectionsRequest) (*Course.GetSectionsResponse, error)

	CreateLecture(req *Course.CreateLectureRequest) (*Course.CreateLectureResponse, error)
	UpdateLecture(req *Course.UpdateLectureRequest) error
	DeleteLecture(req *Course.DeleteLectureRequest) error

	CheckOwner(rq *Course.CheckOwnerRequest) error
}
type coursesHandler struct {
	uc CoursesUseCase
	cf *config.Config
	Course.UnimplementedCourseServiceServer
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

func NewCoursesHandler(uc CoursesUseCase, cf *config.Config) *coursesHandler {
	return &coursesHandler{uc: uc, cf: cf}
}
func (hdl *coursesHandler) GetCourses(ctx context.Context, rq *Course.GetCoursesRequest) (*Course.GetCoursesResponse, error) {
	res, err := hdl.uc.GetCourses(int(rq.PageSize), int(rq.Page))
	client, err := client2.InitUserServiceClient(hdl.cf)
	for index, item := range res.Courses {
		instructor, _ := client.GetInstructor(ctx, &User.GetInstructorInformationRequest{
			Id:  item.Instructor.Id,
			Key: "instructor",
		})
		res.Courses[index].Instructor = instructor.Information
	}

	if err != nil {
		return &Course.GetCoursesResponse{
			Error: HandleError(err),
		}, nil
	}
	return &res, nil
}
func (hdl *coursesHandler) GetCourse(ctx context.Context, rq *Course.GetCourseRequest) (*Course.GetCourseResponse, error) {
	res, err := hdl.uc.GetCourse(rq.Id)
	if err != nil {
		return &Course.GetCourseResponse{
			Error: HandleError(err),
		}, nil
	}
	client, err := client2.InitUserServiceClient(hdl.cf)
	if err != nil {
		return &Course.GetCourseResponse{
			Error: HandleError(err),
		}, nil
	}
	instructor, err := client.GetInstructor(ctx, &User.GetInstructorInformationRequest{
		Id:  res.Course.Instructor.Id,
		Key: "instructor",
	})

	if err != nil {
		fmt.Println(err)
		return &Course.GetCourseResponse{
			Error: HandleError(err),
		}, nil
	}
	res.Course.Instructor = instructor.Information

	return res, nil
}
func (hdl *coursesHandler) GetAllCategories(ctx context.Context, rq *Course.GetAllCategoriesRequest) (*Course.GetAllCategoriesResponse, error) {
	res, err := hdl.uc.GetAllCategories()
	if err != nil {
		return &Course.GetAllCategoriesResponse{
			Error: HandleError(err),
		}, nil
	}
	return res, nil
}
