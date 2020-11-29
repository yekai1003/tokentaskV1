package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"tokentaskV1/bcos"

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
	err = bcos.Bcos_Register(uu.UserName, uu.PassWord)
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
	ok, err := bcos.Bcos_Login(uu.UserName, uu.PassWord)
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
	err = bcos.Bcos_Mint(mi.UserName, mi.Value)
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
	var task bcos.TaskInfoData
	err := c.Bind(&task)
	if err != nil {
		fmt.Println("Failed to Bind", err)
		resp.Code = TASK_PARAMERR
		return
	}
	fmt.Printf("Issue[%+v]\n", task)

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
	fmt.Println("begin call Eth_Issue")

	//操作数据库（区块链智能合约）
	err = bcos.Bcos_Issue(username.(string), passwd.(string), task.Desc, task.Bonus)
	if err != nil {
		fmt.Println("Failed to Eth_Issue", err)
		resp.Code = TASK_ETHERR
		return
	}
}

//任务修改接口
func Modify(c *gin.Context) {
	//组织响应消息
	resp := RespData{
		Code: TASK_OK,
	}
	defer ResponseData(c, &resp)
	//解析数据
	var task bcos.TaskInfoData
	err := c.Bind(&task)
	if err != nil {
		fmt.Println("Failed to Bind", err)
		resp.Code = TASK_PARAMERR
		return
	}
	fmt.Println("Modify:", task)

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
	err = bcos.Bcos_Update(username.(string), passwd.(string), task.Comment, int64(taskid), task.Status)
	if err != nil {
		fmt.Println("Failed to Eth_Update", err)
		resp.Code = TASK_ETHERR
		return
	}
}

//任务列表查询
//GET http://10.211.55.3:9090/tasklist?page=1
func TaskList(c *gin.Context) {
	//组织响应消息
	resp := RespData{
		Code: TASK_OK,
	}
	defer ResponseData(c, &resp)
	//解析数据
	page := c.Query("page")
	ipage, _ := strconv.Atoi(page)

	//查询当前所有任务
	tasks, err := bcos.QueryTask()
	if err != nil {
		resp.Code = TASK_ETHERR
		return
	}

	begin := (ipage - 1) * 10
	end := ipage * 10
	if end > len(tasks) {
		end = len(tasks)
	}

	ts := struct {
		Total int         `json:"total"`
		Data  interface{} `json:"data"`
	}{
		Total: len(tasks),
		Data:  tasks[begin:end],
	}
	resp.Data = ts
}
