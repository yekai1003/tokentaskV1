package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"tokentaskV1/blocks"

	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"
)

type User struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type MintInfo struct {
	UserName string `json:"username"`
	Value    int64  `json:"value"`
}

type TaskInfo struct {
	Task_ID string `json:"task_id"`
	Issuer  string `json:"issuer"`
	Worker  string `json:"task_user"`
	Bonus   int64  `json:"bonus"`
	Desc    string `json:"desc"`
	Comment string `json:"comment"`
	Status  uint8  `json:"status"`
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

	//session处理
	store := ginsession.FromContext(c)
	store.Set("username", uu.UserName)
	store.Set("passwd", uu.PassWord)
	err = store.Save()
	if err != nil {
		c.AbortWithError(500, err)
		resp.Code = TASK_UNKNOWNERR
		return
	}

}

func Mint(c *gin.Context) {
	//组织响应消息
	resp := RespData{
		Code: TASK_OK,
	}
	defer ResponseData(c, &resp)
	//解析数据
	var mi MintInfo
	err := c.Bind(&mi)
	if err != nil {
		fmt.Println("Failed to Bind", err)
		resp.Code = TASK_PARAMERR
		return
	}
	fmt.Println(mi)
	//操作数据库（区块链智能合约）
	err = blocks.Eth_Mint(mi.UserName, mi.Value)
	if err != nil {
		fmt.Println("Failed to Eth_Mint", err)
		resp.Code = TASK_ETHERR
		return
	}
}

func Issue(c *gin.Context) {
	//组织响应消息
	resp := RespData{
		Code: TASK_OK,
	}
	defer ResponseData(c, &resp)
	//解析数据
	var task TaskInfo
	err := c.Bind(&task)
	if err != nil {
		fmt.Println("Failed to Bind", err)
		resp.Code = TASK_PARAMERR
		return
	}
	fmt.Println(task)

	//从session中获取对应信息
	store := ginsession.FromContext(c)
	username, ok := store.Get("username")
	if !ok {
		c.AbortWithStatus(404)
		return
	}
	passwd, ok := store.Get("passwd")
	if !ok {
		c.AbortWithStatus(404)
		return
	}

	//操作数据库（区块链智能合约）
	err = blocks.Eth_Issue(username.(string), passwd.(string), task.Desc, task.Bonus)
	if err != nil {
		fmt.Println("Failed to Eth_Issue", err)
		resp.Code = TASK_ETHERR
		return
	}
}

//任务修改接口
func Update(c *gin.Context) {
	//组织响应消息
	resp := RespData{
		Code: TASK_OK,
	}
	defer ResponseData(c, &resp)
	//解析数据
	var task TaskInfo
	err := c.Bind(&task)
	if err != nil {
		fmt.Println("Failed to Bind", err)
		resp.Code = TASK_PARAMERR
		return
	}
	fmt.Println(task)

	//从session中获取对应信息
	store := ginsession.FromContext(c)
	username, ok := store.Get("username")
	if !ok {
		c.AbortWithStatus(404)
		return
	}
	passwd, ok := store.Get("passwd")
	if !ok {
		c.AbortWithStatus(404)
		return
	}

	//操作数据库（区块链智能合约）
	//Eth_Update(who, pass, comment string, taskID int64, status uint8)
	taskid, _ := strconv.Atoi(task.Task_ID)
	err = blocks.Eth_Update(username.(string), passwd.(string), task.Comment, int64(taskid), task.Status)
	if err != nil {
		fmt.Println("Failed to Eth_Issue", err)
		resp.Code = TASK_ETHERR
		return
	}
}
