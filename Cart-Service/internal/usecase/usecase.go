package usecase

import (
	"errors"
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Cart"
	"github.com/hoangphuc28/CoursesOnline/Cart-Service/config"
	"github.com/hoangphuc28/CoursesOnline/Cart-Service/internal/model"
	"github.com/hoangphuc28/CoursesOnline/Cart-Service/pkg/common"
	"github.com/hoangphuc28/CoursesOnline/Cart-Service/pkg/utils"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type CartRepo interface {
	FindDataWithCondition(condition map[string]any) (*model.Cart, error)
	AddItemToCart(item *model.CartCourse) error
	FindCartWithUser(userId int) (int, error)
	RemoveItemToCart(item *model.CartCourse) error
	ResetCart(cartId int) error
	NewCart(userId int) (*model.Cart, error)
}
type cartUsecase struct {
	config *config.Config
	repo   CartRepo
	h      *utils.Hasher
}

func NewCartUseCase(cf *config.Config, repo CartRepo, h *utils.Hasher) *cartUsecase {
	return &cartUsecase{cf, repo, h}
}
func (uc cartUsecase) GetCart(fakeId string) (*Cart.GetCartResponse, error) {
	id, err := uc.h.Decode(fakeId)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	cart, err := uc.repo.FindDataWithCondition(map[string]any{"user_id": id})
	if err != nil {
		return nil, err
	}
	totalCourse := 0

	if len(cart.Courses) == 0 {
		return &Cart.GetCartResponse{
			Cart: &Cart.Cart{
				Id:      uc.h.Encode(cart.Id),
				Courses: []*Cart.Course{},
			},
			TotalCourse: strconv.Itoa(totalCourse),
			Error:       nil,
		}, nil
	}
	res := Cart.GetCartResponse{
		Cart: &Cart.Cart{
			Id: uc.h.Encode(cart.Id),
		},
	}
	var totalPrice float64
	for _, item := range cart.Courses {
		v, _ := strconv.ParseFloat(item.Price.Value, 64)
		totalPrice += v
		totalCourse++
		var img Cart.Image
		if item.CourseThumbnail != nil {
			img = Cart.Image{
				Url:    item.CourseThumbnail.Url,
				Width:  item.CourseThumbnail.Width,
				Height: item.CourseThumbnail.Height,
			}
		}

		res.Cart.Courses = append(res.Cart.Courses, &Cart.Course{
			Id:           uc.h.Encode(item.Id),
			Title:        item.Title,
			Description:  item.CourseDescription,
			Level:        item.CourseLevel,
			Price:        item.Price.Value,
			Discount:     item.CourseDiscount,
			Currency:     item.Price.Currency,
			Duration:     item.CourseDuration,
			Status:       item.CourseStatus,
			Rating:       item.CourseRating,
			Thumbnail:    &img,
			InstructorId: uc.h.Encode(item.InstructorID),
		})
	}
	res.Cart.Currency = cart.Courses[0].CourseCurrency
	res.TotalCourse = strconv.Itoa(totalCourse)
	res.Cart.TotalPrice = strconv.FormatFloat(math.Round(totalPrice*100)/100, 'f', -1, 64)
	return &res, nil
}

func (uc cartUsecase) AddToCart(userId string, courseId string) error {
	userIdDecoded, err := uc.h.Decode(userId)
	cart, err := uc.repo.FindDataWithCondition(map[string]any{"user_id": userIdDecoded})
	if cart == nil {
		cart, err = uc.repo.NewCart(userIdDecoded)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return common.ErrInternal(err)
	}
	newCourseId, err := uc.h.Decode(courseId)
	if err != nil {
		return common.ErrInternal(err)
	}
	item := &model.CartCourse{
		CartId:   cart.Id,
		CourseId: newCourseId,
	}
	if err = uc.repo.AddItemToCart(item); err != nil {
		if err == gorm.ErrDuplicatedKey {
			return common.NewCustomError(errors.New("course has already been"), "Course has already been!")
		}
	}
	return nil
}

func (uc cartUsecase) RemoveItem(cartId string, courseId string) error {
	newCartId, err := uc.h.Decode(cartId)

	if err != nil {
		return common.ErrInternal(err)
	}
	newCourseId, err := uc.h.Decode(courseId)
	if err != nil {
		return common.ErrInternal(err)
	}
	item := &model.CartCourse{
		CartId:   newCartId,
		CourseId: newCourseId,
	}
	if err = uc.repo.RemoveItemToCart(item); err != nil {
		if err == gorm.ErrDuplicatedKey {
			return common.NewCustomError(errors.New("course has not been in cart"), "Course has already been added!")
		}
	}
	return nil
}

func (uc cartUsecase) ResetCart(cartId string) error {
	decodedCartId, err := uc.h.Decode(cartId)
	if err != nil {
		return err
	}
	if err = uc.repo.ResetCart(decodedCartId); err != nil {
		return err
	}
	return nil
}
func (uc cartUsecase) NewCart(userId string) error {
	decodedUserId, err := uc.h.Decode(userId)
	if err != nil {
		return err
	}
	if _, err = uc.repo.NewCart(decodedUserId); err != nil {
		return err
	}
	return nil
}
