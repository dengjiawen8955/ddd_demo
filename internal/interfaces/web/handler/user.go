package handler

import (
	"ddd_demo/internal/application"
	"ddd_demo/internal/domain"
	"ddd_demo/internal/interfaces/web/middleware"
	"ddd_demo/internal/interfaces/web/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserApp application.UserAppInterface
}

func NewUserHandler(userApp application.UserAppInterface) *UserHandler {
	return &UserHandler{
		UserApp: userApp,
	}
}

func (u *UserHandler) Login(c *gin.Context) {
	var err error
	req := &domain.C2S_Login{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 转化为领域对象 + 参数验证
	loginParams, err := req.ToDomain()

	// 调用应用层
	user, err := u.UserApp.Login(loginParams)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Ok(c, user)
}

// UserInfo 获取用户信息
func (u *UserHandler) UserInfo(c *gin.Context) {
	userIDStr := c.GetString(middleware.UserIDKey)

	userID, err := domain.NewUserID(userIDStr)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	userInfo, err := u.UserApp.Get(userID)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 返回用户信息
	response.Ok(c, userInfo)
}

// Register 注册
func (u *UserHandler) Register(c *gin.Context) {
	var err error
	req := &domain.C2S_Register{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 转化为领域对象 + 参数验证
	registerParams, err := req.ToDomain()

	// 调用应用层
	user, err := u.UserApp.Register(registerParams)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Ok(c, user)
}
