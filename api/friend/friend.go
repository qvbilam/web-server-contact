package friend

import (
	"contact/api"
	proto "contact/api/qvbilam/contact/v1"
	"contact/global"
	"contact/validate"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Apply(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.FriendApply{}
	if err := ctx.Bind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	fmt.Println(global.ContactFriendServerClient)

	_, err := global.ContactFriendServerClient.Apply(context.Background(), &proto.UpdateFriendApplyRequest{
		UserId:      request.UserID,
		ApplyUserId: userID,
		Content:     request.Content,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func Agree(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	paramApplyId := ctx.Param("id")
	applyId, _ := strconv.Atoi(paramApplyId)

	_, err := global.ContactFriendServerClient.ApplyAgree(context.Background(), &proto.UpdateFriendApplyRequest{
		Id:     int64(applyId),
		UserId: userID,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func Reject(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	paramApplyId := ctx.Param("id")
	applyId, _ := strconv.Atoi(paramApplyId)

	_, err := global.ContactFriendServerClient.ApplyReject(context.Background(), &proto.UpdateFriendApplyRequest{
		Id:     int64(applyId),
		UserId: userID,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func News(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	fmt.Printf("userId: %d\n", userID)

	users, err := global.ContactFriendServerClient.GetApply(context.Background(), &proto.UpdateFriendApplyRequest{UserId: userID})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}
	api.SuccessNotMessage(ctx, users)
}

func Friends(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}
