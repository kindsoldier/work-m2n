/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package main

import (
    "errors"
    "encoding/json"
    "bytes"
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
    "time"

    "github.com/gin-gonic/gin"

    "pmapp/pmmaster"
    "pmapp/pmcommon"
    "pmapp/pmconfig"
    "pmapp/pmlog"
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

type UUID = string

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

const (
    jrpcPath string = "/jrpc"

    rpcListBoardDescs       string  = "listBoardDescs"
    rpcGetDevicesInSquare   string  = "getDevicesInSquare"
    rpcGetBoardDesc         string  = "getBoardDesc"
    rpcSetBoardAttribute    string  = "setBoardAttribute"
)

func (this *Service) ServeWeb() error {
    var err error
    gin.DisableConsoleColor()
    gin.SetMode(gin.ReleaseMode)
    gin.DefaultWriter = os.Stdout

    router := gin.New()
    router.POST(jrpcPath, this.rpcFuncResolver)
    router.NoRoute(this.SendNoRoute)

    err = router.Run()    
    return err
}

func (this *Service) rpcFuncResolver(ctx *gin.Context) {
    var err error
    var request BaseRequest

    var requestBytes []byte
    if ctx.Request.Body != nil {
        requestBytes, _ = ioutil.ReadAll(ctx.Request.Body)
        ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBytes))
    }

    err = ctx.BindJSON(&request)
    if err != nil {
        responseBytes, _ := this.rpcMakeError(request.Id, errorParseError, err)
        ctx.Data(http.StatusOK, mimeApplicationJson, responseBytes)
        return
    }

    start := time.Now()
    switch request.Method {
        case rpcListBoardDescs:
            responseBytes, _ := this.rpcListBoardDescs(requestBytes)
            ctx.Data(http.StatusOK, mimeApplicationJson, responseBytes)
        case rpcGetBoardDesc:
            responseBytes, _ := this.rpcGetBoardDesc(requestBytes)
            ctx.Data(http.StatusOK, mimeApplicationJson, responseBytes)
        case rpcSetBoardAttribute:
            responseBytes, _ := this.rpcSetBoardAttribute(requestBytes)
            ctx.Data(http.StatusOK, mimeApplicationJson, responseBytes)
        case rpcGetDevicesInSquare:
            responseBytes, _ := this.rpcGetDevicesInSquare(requestBytes)
            ctx.Data(http.StatusOK, mimeApplicationJson, responseBytes)
        default:
            response, _ := this.rpcMakeError(request.Id, errorMethodNotFound, err)
            ctx.Data(http.StatusOK, mimeApplicationJson, response)
    }
    elapsed := time.Since(start)
    pmlog.LogDebug(request.Method, elapsed) 
    return
}

func (this *Service) rpcListBoardDescs(requestBytes []byte) ([]byte, error) {
    var err error
    var request BaseRequest
    err = json.Unmarshal(requestBytes, &request)
    if err != nil {
        return this.rpcMakeError(request.Id, errorParseError, err)
    }
    boards := this.Master.GetBoardDescs()
    return this.rpcMakeResult(request.Id, boards)
}


type GetBoardDescRequest struct {
    BaseRequest
    Params struct {
        BoardId         UUID       `json:"boardId"`
    } `json:"params"`
}
func (this *Service) rpcGetBoardDesc(requestBytes []byte) ([]byte, error) {
    var err error
    var request GetBoardDescRequest
    err = json.Unmarshal(requestBytes, &request)
    if err != nil {
        return this.rpcMakeError(request.Id, errorParseError, err)
    }
    params := request.Params
    board, err := this.Master.GetBoardDesc(params.BoardId)
    if err != nil {
        return this.rpcMakeError(request.Id, errorInternalError, err)
    }
    return this.rpcMakeResult(request.Id, board)
}

type SetBoardArrributeRequest struct {
    BaseRequest
    Params struct {
        BoardId         pmcommon.UUID       `json:"boardId"`
        AttributeId     pmcommon.UUID       `json:"attributeId"`
        Value           pmconfig.DValue     `json:"value"`
    } `json:"params"`
}
func (this *Service) rpcSetBoardAttribute(requestBytes []byte) ([]byte, error) {
    var err error
    var request SetBoardArrributeRequest
    err = json.Unmarshal(requestBytes, &request)
    if err != nil {
        return this.rpcMakeError(request.Id, errorParseError, err)
    }
    params := request.Params
    err = this.Master.SetBoardAttribute(params.BoardId, params.AttributeId, params.Value)
    if err != nil {
        return this.rpcMakeError(request.Id, errorInternalError, err)
    }
    return this.rpcMakeOk(request.Id)
}

type GetDevicesInSquareRequest struct {
    BaseRequest
    Params struct {
        LatiMin     float64         `json:"latiMin"`
        LatiMax     float64         `json:"latiMax"`
        LongiMin    float64         `json:"longiMin"`
        LongiMax    float64         `json:"longiMax"`
    } `json:"params"`
}
func (this *Service) rpcGetDevicesInSquare(requestBytes []byte) ([]byte, error) {
    var err error
    var request GetDevicesInSquareRequest
    err = json.Unmarshal(requestBytes, &request)
    if err != nil {
        return this.rpcMakeError(request.Id, errorParseError, err)
    }
    params := request.Params
    boards := this.Master.GetDevicesInSquare(params.LatiMin, params.LatiMax, params.LongiMin, params.LongiMax)
    if err != nil {
        return this.rpcMakeError(request.Id, errorInternalError, err)
    }
    return this.rpcMakeResult(request.Id, boards)
}

func (this *Service) rpcMakeError(reqId string, errorCode int, funcErr error) ([]byte, error) {
    var err error
    if funcErr == nil {
        funcErr = errors.New("undefined")
    }
    responseError := ResponseError{
        Code:       errorCode,
        Message:    fmt.Sprintf("%s", funcErr),
    }
    response := Response{
        JSONRPC:    jsonrpcId,
        Id:         reqId,
        Error:      responseError,
    }
    resBytes, err := json.MarshalIndent(response, "", "    ")
    return resBytes, err
}

func (this *Service) rpcMakeResult(reqId string, result interface{}) ([]byte, error) {
    var err error
    response := Response{
        JSONRPC:    jsonrpcId,
        Id:         reqId,
        Result:     result,
    }
    resBytes, err := json.MarshalIndent(response, "", "    ")
    return resBytes, err
}

func (this *Service) rpcMakeOk(reqId string) ([]byte, error) {
    var err error
    response := Response{
        JSONRPC:    jsonrpcId,
        Id:         reqId,
    }
    resBytes, err := json.MarshalIndent(response, "", "    ")
    return resBytes, err
}

func (this *Service) SendNoRoute(ctx *gin.Context) {
    response, _ := this.rpcMakeError(emptyRequestId, errorMethodNotFound, errors.New("route not found"))
    ctx.Data(http.StatusOK, mimeApplicationJson, response)

}

func main() {
    app := NewService()
    app.LoadBoards() 
    app.ServeWeb()
}
//EOF
