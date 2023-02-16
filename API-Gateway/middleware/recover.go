package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangphuc28/CoursesOnline/API-Gateway/pkg/common"
)

func (m *MiddleareManager) Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//if appErr, ok := err.(*pb.ErrorResponse); ok {
				//	fmt.Println("check")
				//	ctx.AbortWithStatusJSON(int(appErr.Code), appErr)
				//	panic(err)
				//}
				if appErr, ok := err.(*common.AppError); ok {
					ctx.AbortWithStatusJSON(int(appErr.StatusCode), appErr)
					panic(err)
				}

				//ctx.AbortWithStatusJSON(500, &pb.ErrorResponse{
				//	Code:    500,
				//	Message: "Internal Server error",
				//})
			}
		}()
		ctx.Next()
	}
}
