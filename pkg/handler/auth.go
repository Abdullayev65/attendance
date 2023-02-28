package hendler

import (
	"github.com/Abdullayev65/attendance/pkg/ioPut"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var userAdd ioPut.UserAdd
	c.Bind(&userAdd)
	if ok, fieldName := userAdd.Valid(); !ok {
		c.String(400, "invalid user field ["+fieldName+"]")
		return
	}
	user := userAdd.MapToUser()
	err := h.Service.UserAdd(user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &user)
}
func (h *Handler) SignIn(c *gin.Context) {
	var sign ioPut.Sign
	c.Bind(&sign)
	user, err := h.Service.UserByUsername(sign.Username)
	if err != nil || user.Password != sign.Password {
		c.String(http.StatusBadRequest, "username or password wrong")
		return
	}
	token, err := h.TokenJWT.Generate(user.ID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
