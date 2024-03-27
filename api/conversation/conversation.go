package conversation

import (
	"contact/api"
	proto "contact/api/qvbilam/contact/v1"
	"contact/global"
	"contact/resource"
	"contact/validate"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Get(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	cs, err := global.ContactConversationServerClient.Get(context.Background(), &proto.GetConversationRequest{
		UserId: userID,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	res := resource.ConversationsResource{}
	api.SuccessNotMessage(ctx, res.Resource(cs))
}

func Create(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.ConversationCreate{}
	if err := ctx.Bind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	_, err := global.ContactConversationServerClient.Create(context.Background(), &proto.UpdateConversationRequest{
		UserId:     userID,
		ObjectType: request.Type,
		ObjectId:   request.ObjectId,
	})
	if err != nil {
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

	_, err := global.ContactConversationServerClient.Delete(context.Background(), &proto.UpdateConversationRequest{
		Id:     int64(id),
		UserId: userID,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}
