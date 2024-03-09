package usecase

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/User"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/config"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Course-Service/pkg/utils"
	"strconv"
)

type CoursesRepository interface {
	Create(course *model.Course) error
	GetCourses(limit int, page int) (res []model.Course, total int64, err error)
	GetCourse(id int) (res model.Course, err error)
	GetAllCategories() ([]model.Category, error)
	NewEnrollment(enrollment *model.Enrollment) error
	GetCourseContent(e *model.Enrollment) (*model.Course, error)
	GetEnrollments(u *model.User) ([]model.Course, error)
	NewCourse(course *model.Course) (*int, error)
	GetPrices() (res []model.Price, err error)
	UpdateCourse(course *model.Course) error
	PublishCourse(course *model.Course) error
	FindDataWithCondition(conditions map[string]any) ([]model.Course, error)
	DeleteCourse(course model.Course) error

	NewLecture(lecture model.Lecture) (*model.Lecture, error)
	UpdateLecture(lecture *model.Lecture, condition map[string]any) error
	DeleteLecture(s *model.Lecture) error
	FindLectureWithCondition(condition map[string]any) (*model.Lecture, error)

	FindSectionWithCondition(condition map[string]any) (*model.Section, error)
	NewSection(section model.Section) (*model.Section, error)
	UpdateSection(section *model.Section, condition map[string]any) error
	DeleteSection(s *model.Section) error
	FindSectionsWithCondition(condition map[string]any) ([]model.Section, error)
}
type coursesUseCase struct {
	repo CoursesRepository
	cf   *config.Config
	h    *utils.Hasher
}

func NewCoursesUseCase(repo CoursesRepository, h *utils.Hasher, cf *config.Config) *coursesUseCase {
	return &coursesUseCase{repo: repo, h: h, cf: cf}
}
func (uc *coursesUseCase) CreateCourse(data *model.Course) error {

	if err := uc.repo.Create(data); err != nil {
		return err
	}

	return nil
}
func (uc *coursesUseCase) GetCourses(limit int, page int) (coursesResponse Course.GetCoursesResponse, err error) {
	res, total, err := uc.repo.GetCourses(limit, page)
	var courses []*Course.Course
	for _, course := range res {
		course.FakeId = uc.h.Encode(course.Id)
		price := &Course.Price{
			Value:    course.Price.Value,
			Currency: course.Price.Currency,
		}
		//if !res[i].IsPaid {
		//	price.Value = "free"
		//}
		var img Course.Image
		if course.Thumbnail != nil {
			img.Url = course.Thumbnail.Url
			img.Width = course.Thumbnail.Width
			img.Height = course.Thumbnail.Height
		}
		courses = append(courses, &Course.Course{
			Id:          course.FakeId,
			Title:       course.Title,
			Description: course.Description,
			Level:       course.Level,
			Language:    course.Language,
			Price:       price,
			IsPublish:   course.IsPublish,
			NumReviews:  strconv.Itoa(course.NumReviews),
			AvgRating:   strconv.Itoa(int(course.Rating)),
			Thumbnail:   &img,
			SubCategory: uc.h.Encode(course.SubcategoryId),
			Instructor: &User.Instructor{
				Id: uc.h.Encode(course.InstructorID),
			},
			Requirement: course.Requirement,
		})
	}
	prePage := &Course.Link{
		Method: "POST",
		Href:   uc.cf.Service.Host + "/courses?" + "pageSize=" + strconv.Itoa(limit) + "page=" + strconv.Itoa(page-1),
		Rel:    "pre_page",
	}

	nextPage := &Course.Link{
		Method: "POST",
		Href:   uc.cf.Service.Host + "/courses?" + "pageSize=" + strconv.Itoa(limit) + "page=" + strconv.Itoa(page+1),
		Rel:    "next_page",
	}

	if int(total) <= (page * limit) {
		nextPage.Href = uc.cf.Service.Host + "/courses?" + "pageSize=" + strconv.Itoa(limit) + "page=" + strconv.Itoa(1)

	}
	if page <= 1 {
		prePage.Href = uc.cf.Service.Host + "/courses?" + "pageSize=" + strconv.Itoa(limit) + "page=" + strconv.Itoa(int(total)/limit)
	}
	coursesResponse.Courses = courses
	coursesResponse.Links = append(coursesResponse.Links, prePage)
	coursesResponse.Links = append(coursesResponse.Links, nextPage)
	coursesResponse.CourseLength = int32(total)
	return
}

func (uc *coursesUseCase) GetCourse(fakeId string) (*Course.GetCourseResponse, error) {
	id, err := uc.h.Decode(fakeId)
	if err != nil {
		return nil, err
	}
	course, err := uc.repo.GetCourse(id)
	if err != nil {
		return nil, err
	}
	//if !course.IsPublish {
	//	return nil, common.NewCustomError(err, http.StatusNotFound, "This course has not published yet!")
	//}
	course.FakeId = uc.h.Encode(course.Id)
	//if !course.IsPaid {
	//	course.Price.Value = "free"
	//}
	var sections []*Course.Section
	for _, i := range course.Sections {
		var lectures []*Course.Lecture
		for _, j := range i.Lectures {
			var resource Course.Resource
			if j.IsFree {
				resource = Course.Resource{
					Url:      j.Resource.Url,
					Duration: j.Resource.Duration,
				}
			}
			j.FakeId = uc.h.Encode(j.Id)
			lectures = append(lectures, &Course.Lecture{
				Id:       j.FakeId,
				Title:    j.Title,
				Content:  j.Content,
				Status:   j.Status,
				IsFree:   j.IsFree,
				Resource: &resource,
			})
		}
		i.FakeId = uc.h.Encode(i.Id)
		sections = append(sections, &Course.Section{
			Id:               i.FakeId,
			Title:            i.Title,
			NumberOfLectures: int32(i.NumberOfLectures),
			Lectures:         lectures,
		})
	}
	var img Course.Image
	if course.Thumbnail != nil {
		img.Url = course.Thumbnail.Url
		img.Width = course.Thumbnail.Width
		img.Height = course.Thumbnail.Height
	}
	res := &Course.Course{
		Id:          course.FakeId,
		Title:       course.Title,
		Description: course.Description,
		Level:       course.Level,
		Language:    course.Language,
		Price: &Course.Price{
			Id:       strconv.Itoa(course.Price.Id),
			Value:    course.Price.Value,
			Currency: course.Price.Currency,
		},
		AvgRating: strconv.Itoa(int(course.Rating)),
		Thumbnail: &img,
		Instructor: &User.Instructor{
			Id: uc.h.Encode(course.InstructorID),
		},
		SubCategory: uc.h.Encode(course.SubcategoryId),
		IsPublish:   course.IsPublish,
		Sections:    sections,
		NumReviews:  strconv.Itoa(course.NumReviews),
	}
	fmt.Println(res)
	return &Course.GetCourseResponse{
		Course: res,
	}, nil
}

func (uc *coursesUseCase) GetAllCategories() (*Course.GetAllCategoriesResponse, error) {
	cate, err := uc.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	var res Course.GetAllCategoriesResponse

	for _, category := range cate {
		var subcategories []*Course.SubCategory
		for _, subcategory := range category.SubCategories {
			subcategories = append(subcategories, &Course.SubCategory{
				Id:         uc.h.Encode(subcategory.Id),
				Name:       subcategory.Name,
				CategoryId: uc.h.Encode(subcategory.CategoryId),
			})
		}
		res.Categories = append(res.Categories, &Course.Category{
			Id:            uc.h.Encode(category.Id),
			Name:          category.Name,
			Subcategories: subcategories,
		})
	}
	return &res, nil
}
