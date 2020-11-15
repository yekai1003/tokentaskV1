package routes

import (
	"fmt"
	"net/http"

	"tokentaskV1/blocks"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type RespData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	TASK_OK         = "0"
	TASK_PARAMERR   = "1"
	TASK_LOGINERR   = "2"
	TASK_DBERR      = "3"
	TASK_ETHERR     = "4"
	TASK_UNKNOWNERR = "5"
)

var task_text = map[string]string{
	TASK_OK:         "成功",
	TASK_PARAMERR:   "参数错误",
	TASK_LOGINERR:   "登录失败",
	TASK_DBERR:      "数据库错误",
	TASK_ETHERR:     "以太坊操作失败",
	TASK_UNKNOWNERR: "未知错误",
}

func taskMsgByCode(code string) string {
	str, ok := task_text[code]
	if !ok {
		return task_text[TASK_UNKNOWNERR]
	}
	return str
}

func ResponseData(c *gin.Context, d *RespData) {
	d.Msg = taskMsgByCode(d.Code)
	c.JSON(http.StatusOK, d)
}

func Register(c *gin.Context) {
	//组织响应消息
	resp := RespData{
		Code: TASK_OK,
	}
	defer ResponseData(c, &resp)
	//解析数据
	var uu User
	err := c.Bind(&uu)
	if err != nil {
		fmt.Println("Failed to Bind", err)
		resp.Code = TASK_PARAMERR
		return
	}
	fmt.Println(uu)
	//操作数据库（区块链智能合约）
	err = blocks.Eth_Register(uu.UserName, uu.PassWord)
	if err != nil {
		fmt.Println("Failed to Eth_Register", err)
		resp.Code = TASK_ETHERR
		return
	}
}

func Login(c *gin.Context) {
	//组织响应消息
	resp := RespData{
		Code: TASK_OK,
	}
	defer ResponseData(c, &resp)
	//解析数据
	var uu User
	err := c.Bind(&uu)
	if err != nil {
		fmt.Println("Failed to Bind", err)
		resp.Code = TASK_PARAMERR
		return
	}
	fmt.Println(uu)
	//操作数据库（区块链智能合约）
	// resp.Code = TASK_LOGINERR
	// if uu.UserName == "yekai" && uu.PassWord == "123" {
	// 	resp.Code = TASK_OK
	// }
	ok, err := blocks.Eth_Login(uu.UserName, uu.PassWord)
	if err != nil {
		fmt.Println("Failed to Eth_Login", err)
		resp.Code = TASK_ETHERR
		return
	}
	if !ok {
		fmt.Println("Failed to Login, password or user err")
		resp.Code = TASK_LOGINERR
		return
	}
}
