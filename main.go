package main

import (
	"log"

	asr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr/v20190614"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func main() {

	// 实例化一个认证对象，入参需要传入腾讯云账户 secretId、secretKey
	credential := common.NewCredential("#secretId", "#secretKey")

	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 5
	cpf.Debug = true

	client, _ := asr.NewClient(credential, "ap-shanghai", cpf)

	request := asr.NewCreateAsyncRecognitionTaskRequest()

	params := "{\"EngineModelType\":\"16k_zh\",\"ChannelNum\":1,\"ResTextFormat\":0,\"SourceType\":0,\"Url\":\"https://asr-audio-1300466766.cos.ap-nanjing.myqcloud.com/test16k.wav\"}"

	err := request.FromJsonString(params)

	if err != nil {
		log.Println(err)
	}

	response, err := client.CreateAsyncRecognitionTask(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		log.Printf("An API error has returned: %s", err)
		return
	}

	// 非 SDK 异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		log.Println(err)
	}
	// 打印返回的 json 字符串
	log.Printf("%s", response.ToJsonString())
}
