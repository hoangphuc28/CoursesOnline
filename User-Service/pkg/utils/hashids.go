package utils

import "github.com/speps/go-hashids/v2"

type Hasher struct {
	HashID *hashids.HashID
}

func NewHasher(salt string, minLength int) *Hasher {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	hs, _ := hashids.NewWithData(hd)
	return &Hasher{hs}
}
func (h *Hasher) Encode(id int) (res string) {
	res, _ = h.HashID.Encode([]int{id})
	return
}
func (h *Hasher) Decode(str string) (res int, nil error) {
	s, err := h.HashID.DecodeWithError(str)
	if err != nil {
		return 0, err
	}
	res = s[0]
	return
}
