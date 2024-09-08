package main

import "context"

type UsePointsAsDiscount struct {
	UserID int
	Points int
}

type UsePointsAsDiscountHandler struct {
	userRepository userRepository
}

type userRepository interface {
	UsePointsForDiscount(ctx context.Context, userID int, point int) error
}

func NewUsePointsAsDiscountHandler(
	userRepository userRepository,
) UsePointsAsDiscountHandler {
	return UsePointsAsDiscountHandler{
		userRepository: userRepository,
	}
}

func (h UsePointsAsDiscountHandler) Handle(ctx context.Context, cmd UsePointsAsDiscount) error {
	return h.userRepository.UsePointsForDiscount(ctx, cmd.UserID, cmd.Points)
}
