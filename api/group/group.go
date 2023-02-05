package group

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

func Create(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.GroupCreate{}
	if err := ctx.Bind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	_, err := global.ContactGroupServerClient.Create(context.Background(), &proto.UpdateGroupRequest{
		UserId:    userID,
		Name:      request.Name,
		Avatar:    request.Avatar,
		Cover:     request.Cover,
		Introduce: request.Introduce,
	})

	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func Search(ctx *gin.Context) {
	// todo
	api.SuccessNotContent(ctx)
}

func Mine(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	groups, err := global.ContactGroupServerClient.Mine(context.Background(), &proto.SearchGroupRequest{
		UserId: userID,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	res := resource.GroupsResource{}

	api.SuccessNotMessage(ctx, res.Resource(groups))
}

func Members(ctx *gin.Context) {
	//uID, _ := ctx.Get("userId")
	//userID := uID.(int64)

	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	members, err := global.ContactGroupServerClient.Members(context.Background(), &proto.SearchGroupMemberRequest{
		GroupId: int64(id),
	})

	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}
	res := resource.GroupMembersResource{}
	api.SuccessNotMessage(ctx, res.Resource(members))
}

func Join(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	_, err := global.ContactGroupServerClient.Join(context.Background(), &proto.UpdateGroupMemberRequest{
		GroupId:    int64(id),
		UserId:     userID,
		OperatorId: userID,
	})

	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func Quit(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	_, err := global.ContactGroupServerClient.Quit(context.Background(), &proto.UpdateGroupMemberRequest{
		GroupId:    int64(id),
		UserId:     userID,
		OperatorId: userID,
	})

	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}
