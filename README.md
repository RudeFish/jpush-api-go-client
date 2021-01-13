jpush-api-go-client
===================

概述
----------------------------------- 
   这是JPush REST API 的 go 版本封装开发包,仅支持最新的REST API v3功能。
   REST API 文档：http://docs.jpush.cn/display/dev/Push-API-v3


使用  
----------------------------------- 
   go get github.com/ylywyn/jpush-api-go-client


推送流程  
----------------------------------- 
### 1.构建要推送的平台： jpushclient.Platform
	//Platform
	var pf jpushclient.Platform
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)
	pf.Add(jpushclient.WINPHONE)
	//pf.All()

### 2.构建接收听众： jpushclient.Audience
	//Audience
	var ad jpushclient.Audience
	s := []string{"t1", "t2", "t3"}
	ad.SetTag(s)
	id := []string{"1", "2", "3"}
	ad.SetID(id)
	//ad.All()

### 3.构建通知 jpushclient.Notice，或者消息： jpushclient.Message

	//Notice
	var notice jpushclient.Notice
	notice.SetAlert("alert_test")
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: "AndroidNotice"})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: "IOSNotice"})
	notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice"})
	  
	//jpushclient.Message
	var msg jpushclient.Message
	msg.Title = "Hello"
	msg.Content = "你是ylywn"

### 4.构建jpushclient.PayLoad
    payload := jpushclient.NewPushPayLoad()
    payload.SetPlatform(&pf)
    payload.SetAudience(&ad)
    payload.SetMessage(&msg)
    payload.SetNotice(&notice)


​      
### 5.构建PushClient，发出推送
	c := jpushclient.NewPushClient(secret, appKey)
	r, err := c.Send(bytes)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	} else {
		fmt.Printf("ok:%s", r)
	}

  

### 6.完整demo
    package main
    
    import (
    	"fmt"
    	"github.com/ylywyn/jpush-api-go-client"
    	jpushclient "github.com/RudeFish/jpush-api-go-client"
    	"github.com/astaxie/beego/logs"
    	"github.com/bitly/go-simplejson"
    )
    
    const (
    	appKey = "you jpush appkey"
    	secret = "you jpush secret"
    )
    
    func main() {
    	// 推送用户
    	registrationIds := []string{"..."}
    	// 推送内容
    	content := "Content text"
    	title := "Title text"
    
    	var notice jpushclient.Notice // 构建通知
    	var pf jpushclient.Platform   // 平台 ad jpushclient.Audience
    	var ad jpushclient.Audience   // 接收听众
    	notice.SetAlert("alert_test")
    
    	// 加载附加字段
    	extras := make(map[string]interface{}, 0)
    	//if msg.ExtrasJson != nil {
    	//	for _, v := range msg.ExtrasJson {
    	//		extras[v["key"]] = v["val"]
    	//	}
    	//}
    	iosMsg := make(map[string]interface{}, 0)
    	// 填充信息 - 安卓
    	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: content, Title: title, Extras: extras})
    	pf.Add(jpushclient.ANDROID)
    	// IOS
    	iosMsg["title"] = title
    	iosMsg["body"] = content
    	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: iosMsg, Badge: "+1", Extras: extras})
    	pf.Add(jpushclient.IOS)
    	
    	// 附加内容
    	var option jpushclient.Option
    	// True 表示推送生产环境，False 表示要推送开发环境； 如果不指定则为推送生产环境； IOS用
    	option.ApnsProduction = false
    	// =================================================设置厂商通道
    	thirdPartyChannel := make(map[string]interface{}, 0)
    	// ===============小米
    	xiaomi := make(map[string]interface{}, 0)
    	xiaomi["distribution"] = "secondary_push"
    	thirdPartyChannel["xiaomi"] = xiaomi
    	// ===============华为
    	huawei := make(map[string]interface{}, 0)
    	huawei["distribution"] = "secondary_push"
    	thirdPartyChannel["huawei"] = huawei
    	//// ===============魅族
    	//meizu := make(map[string]interface{}, 0)
    	//meizu["distribution"] = "secondary_push"
    	//thirdPartyChannel["meizu"] = meizu
    	// ===============oppo
    	oppo := make(map[string]interface{}, 0)
    	oppo["distribution"] = "secondary_push"
    	thirdPartyChannel["oppo"] = oppo
    	// ===============vivo
    	vivo := make(map[string]interface{}, 0)
    	vivo["distribution"] = "secondary_push"
    	thirdPartyChannel["vivo"] = vivo
    
    	option.SetThirdPartyChannel(thirdPartyChannel)
    	// ==================================================4.构建jpushclient.PayLoad
    	payload := jpushclient.NewPushPayLoad()
    	payload.SetOptions(&option)
    
    	payload.SetPlatform(&pf)   // 平台
    	payload.SetNotice(&notice) // 推送信息
    
    	// 设置听众
    	ad.SetID(registrationIds)
    	payload.SetAudience(&ad)
    
    	bytes, _ := payload.ToBytes()
    	logs.Info("极光消息发送参数:", string(bytes))
    	jPushclient := jpushclient.NewPushClient(secret, appKey)
    	res, err := jPushclient.Send(bytes)
    	if err == nil {
    		respMap := jsonToMap(res)
    		if respMap["error"] != nil {
    			fmt.Printf("极光Push消息，通过RegistrationId发送消息返回错误信息 err:%v\n", respMap)
    			return
    		} else if respMap["msg_id"] != nil {
    			fmt.Printf("发送成功,msg_id为%v\n", respMap["msg_id"].(string))
    		}
    	} else {
    		fmt.Printf("极光Push消息，通过RegistrationId发送消息返回状态错误 err:%v\n", err.Error())
    		return
    	}
    }
    
    func jsonToMap(value string) map[string]interface{} {
    	byteValue := []byte(value)
    	js, _ := simplejson.NewJson(byteValue)
    	nodes, _ := js.Map()
    	return nodes
    }

