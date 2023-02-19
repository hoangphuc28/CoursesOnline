package usecase

import "github.com/Zhoangp/User-Service/pb"

func (uc *userUseCase) GetProfile(userId string) (*pb.GetProfileResponse, error) {
	userIdDecoded, err := uc.h.Decode(userId)
	if err != nil {
		return nil, err
	}
	res, err := uc.userRepo.FindDataWithCondition(map[string]any{"id": userIdDecoded})
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileResponse{
		Information: &pb.User{
			Id:          res.FakeId,
			FirstName:   res.FirstName,
			LastName:    res.LastName,
			Email:       res.Email,
			PhoneNumber: res.Phone,
			Address:     res.Address,
			Avatar: &pb.Picture{
				Url: res.Avatar.Url,
			},
		},
	}, nil
}
