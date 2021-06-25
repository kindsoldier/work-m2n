/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package main

import (
    "errors"
    "fmt"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"

    "pmapp/pmmaster"
    "pmapp/pmcommon"
    "pmapp/pmconfig"
)


const (
    mimeApplicationJson string  = "application/json"
    jsonrpcId           string  = "2.0"
    funcAddName         string  = "add"
    emptyRequestId      string  = ""
    
    errorDefaultError   int     = 1
    errorParseError     int     = -32700 	// Invalid JSON was received by the server.
    errorInvalidRequest int     = -32600 	// The JSON sent is not a valid Request object.
    errorMethodNotFound int     = -32601 	// The method does not exist / is not available.
    errorInvalidParams  int     = -32602 	// Invalid method parameter(s).
    errorInternalError  int     = -32603 	// Internal JSON-RPC error.
)

type BaseRequest struct {
    JSONRPC     string          `json:"jsonrpc"`
    Id          string          `json:"id"`
    Method      string          `json:"method"`
}

type  ResponseError struct  {
    Code    int                 `json:"code,omitempty"`
    Message string              `json:"message,omitempty"`
} 

type Response struct {
    JSONRPC     string          `json:"jsonrpc"`
    Id          string          `json:"id,omitempty"`
    Error       interface{}     `json:"error,omitempty"`
    Result      interface{}     `json:"result,omitempty"`
}

type Service struct {
    Master *pmmaster.BMaster
}

func NewService() *Service {
    var app Service
    app.Master = pmmaster.NewBMaster()
    return &app
}

func (this *Service) LoadBoards() error {
    var err error
    err = this.Master.LoadBoards()
    if err != nil {
        return err
    }
    return err
}

func (this *Service) ServeWeb() error {
    var err error
    gin.DisableConsoleColor()
    //gin.SetMode(gin.ReleaseMode)
    gin.DefaultWriter = os.Stdout

    router := gin.New()

    router.GET("/board/objects/list", this.ListBoardObjects)
    router.POST("/board/attribute/set", this.SetBoardAttribute)
    router.POST("/board/object/get", this.GetBoardObject)

    router.NoRoute(this.SendNoRoute)

    err = router.Run()    
    return err
}

func (this *Service) ListBoardObjects(ctx *gin.Context) {
    boards := this.Master.GetBoardObjects()
    
    SendResult(ctx, emptyRequestId, boards)
}

type SetArrributeRequest struct {
    BaseRequest
    Params struct {
        BoardId         pmcommon.UUID       `json:"boardId"`
        AttributeId     pmcommon.UUID       `json:"attributeId"`
        Value           pmconfig.DValue     `json:"value"`
    } `json:"params"`
}

func (this *Service) SetBoardAttribute(ctx *gin.Context) {
    var err error
    var request SetArrributeRequest
    ctx.BindJSON(&request)
    params := request.Params
    err = this.Master.SetBoardAttribute(params.BoardId, params.AttributeId, params.Value) 
    if err != nil {
        SendError(ctx, emptyRequestId, errorDefaultError, err)
        return
    }
    SendOk(ctx)
}

type GetBoardRequest struct {
    BaseRequest
    Params struct {
        BoardId         UUID       `json:"boardId"`
    } `json:"params"`
}

func (this *Service) GetBoardObject(ctx *gin.Context) {
    var err error
    var request GetBoardRequest
    ctx.BindJSON(&request)

    params := request.Params
    board, err := this.Master.GetBoardObject(params.BoardId) 
    if err != nil {
        SendError(ctx, emptyRequestId, errorDefaultError, err)
        return
    }
    SendResult(ctx, emptyRequestId, board)
}


func (this *Service) SendNoRoute(ctx *gin.Context) {
    SendError(ctx, emptyRequestId, errorMethodNotFound, errors.New("route not found"))
}

func SendError(ctx *gin.Context, requestId string, errorCode int, err error) {
    if err == nil {
        err = errors.New("undefined")
    }
    responseError := ResponseError{
        Code:       errorCode,
        Message:    fmt.Sprintf("%s", err),
    }
    response := Response{
        JSONRPC:    jsonrpcId,
        Id:         requestId,
        Error:      responseError,
    }
    ctx.IndentedJSON(http.StatusOK, response)
}

func SendResult(ctx *gin.Context, requestId string, result interface{}) {
    response := Response{
        JSONRPC:    jsonrpcId,
        Id:         requestId,
        Result: result,
    }
    ctx.IndentedJSON(http.StatusOK, &response)
}

func SendOk(ctx *gin.Context) {
    response := Response{
        JSONRPC:    jsonrpcId,
    }
    ctx.JSON(http.StatusOK, response)
}

func main() {
    app := NewService()
    app.LoadBoards() 
    app.ServeWeb()
}
//EOF
