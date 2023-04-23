package api

import (
	"DiscordRolesBot/internal/async_work"
	"DiscordRolesBot/pkg/base_resp"
	error_code "DiscordRolesBot/pkg/err_code"
	"fmt"
	"github.com/gin-gonic/gin"
)

func BindDiscord(c *gin.Context) {
	walletAddress := c.Param("walletAddress")
	bsp := error_code.NewBaseResp()
	if len(walletAddress) == 0 {
		bsp.SetMsg(error_code.ParamsError, "walletAddress is nil")
		base_resp.JsonResponse(c, bsp, "")
		return
	}

	if err := async_work.AddRoles(walletAddress); err != nil {
		fmt.Println(err.Error())
	}

	base_resp.JsonResponse(c, bsp, "")
}
