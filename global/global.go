package global

import (
	proto "contact/api/qvbilam/contact/v1"
	messageProto "contact/api/qvbilam/message/v1"
	userProto "contact/api/qvbilam/user/v1"
	"contact/config"
	ut "github.com/go-playground/universal-translator"
)

var (
	Trans                     ut.Translator // 表单验证
	ServerConfig              *config.ServerConfig
	ContactFriendServerClient proto.FriendClient
	ContactGroupServerClient  proto.GroupClient
	UserServerClient          userProto.UserClient
	MessageServerClient       messageProto.MessageClient
)
