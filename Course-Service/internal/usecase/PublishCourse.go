package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/User"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/client"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/common"
	"net/http"
)

func (uc *coursesUseCase) PublishCourse(ctx context.Context, rq *Course.PublishCourseRequest) ([]string, error) {
	courseIdDecoded, err := uc.h.Decode(rq.CourseId)
	userIdDecoded, err := uc.h.Decode(rq.UserId)
	if err != nil {
		return nil, err
	}
	course, err := uc.repo.GetCourse(courseIdDecoded)
	if err != nil {
		return nil, err
	}
	if course.InstructorID != userIdDecoded {
		return nil, common.NewCustomError(err, http.StatusNotFound, "Your are not an instructor of this course!")
	}
	userClient, err := client.InitUserServiceClient(uc.cf)

	if err != nil {
		return nil, err
	}
	if course.Price.Value != "free" {
		res, err := userClient.GetProfileInstructor(ctx, &User.GetUserInformationRequest{UserId: rq.UserId})
		if err != nil {
			return nil, err
		}
		fmt.Println(res)
		if res.AccountPaypal == "" {
			return nil, common.NewCustomError(errors.New("a paying account is required to publish a paid course"), 403, "a paying account is required to publish a paid course")
		}
	}

	if _, err = govalidator.ValidateStruct(course); err != nil {
		var errors []string
		for _, e := range err.(govalidator.Errors) {
			errors = append(errors, e.Error())
		}
		return errors, nil
	}

	if err = uc.repo.PublishCourse(&course); err != nil {
		return nil, err
	}
	return nil, nil
}
