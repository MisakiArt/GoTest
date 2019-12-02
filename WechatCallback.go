package main
import (
	"encoding/xml"
	"fmt"
)

//var dataStr string
//var mid string
//var data string
//var openid string
//var appId string
//var callTime string
//var keyword string
//var eventName string
//var eventKey string
//var msgType string


type weChatData struct {
	FromUserName string `xml:"FromUserName"`
	ToUserName string `xml:"ToUserName"`
	CreateTime string `xml:"CreateTime"`
	Content string `xml:"Content"`
	Event string `xml:"Event"`
	EventKey string `xml:"EventKey"`
	MsgType string `xml:"MsgType"`
	MenuId string `xml:"MenuId"`
}

func getData()  {
	str :=`<xml><ToUserName><![CDATA[gh_588b9e8cd3cb]]></ToUserName>
	<FromUserName><![CDATA[{$openid}]]></FromUserName>
	<CreateTime>1542956683</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[subscribe]]></Event>
	<EventKey><![CDATA[]]></EventKey>
	</xml>`
	dealCallback(33,str)

}

func  dealCallback(mid int,data string)  {
	initData(mid,data)


}

func initData(mid int, data string){
	if data != "" {
		//dataStr =data
		xmlByte := []byte(data)
		var weChatD = new(weChatData)
		_ = xml.Unmarshal(xmlByte,weChatD)
		fmt.Println(weChatD.FromUserName)
		fmt.Println(weChatD.ToUserName)
	}


}

func deal()  {


}

