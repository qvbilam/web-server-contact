package friend

import (
	"contact/api"
	proto "contact/api/qvbilam/contact/v1"
	"contact/global"
	"contact/validate"
	"context"
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

	users, err := global.ContactFriendServerClient.GetApply(context.Background(), &proto.UpdateFriendApplyRequest{UserId: userID})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}
	api.SuccessNotMessage(ctx, users)
}

func Friends(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	users, err := global.ContactFriendServerClient.Get(context.Background(), &proto.SearchFriendRequest{UserId: userID})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotMessage(ctx, users)
}

func Update(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	request := validate.FriendUpdate{}
	if err := ctx.Bind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	if _, err := global.ContactFriendServerClient.Update(context.Background(), &proto.UpdateFriendRequest{
		Id:     int64(id),
		UserId: userID,
		Remark: request.Remark,
	}); err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func Delete(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	if _, err := global.ContactFriendServerClient.Delete(context.Background(), &proto.UpdateFriendRequest{
		Id:     int64(id),
		UserId: userID,
	}); err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}
