// +build windows,386,cgo

package cqp

// #include "cq.h"
import "C"

// 这个要在init函数内被设置
var AppID string

//export _appinfo
func _appinfo() *C.char { return C.CString("9," + AppID) }

//export _on_enable
func _on_enable() int32 {
	if Enable == nil {
		return 0
	}
	return Enable()
}

//export _on_disable
func _on_disable() int32 {
	if Disable == nil {
		return 0
	}
	return Disable()
}

//export _on_private_msg
func _on_private_msg(subType, msgID int32, fromQQ int64, msg *C.char, font int32) int32 {
	if PrivateMsg == nil {
		return 0
	}
	return PrivateMsg(subType, msgID, fromQQ, C.GoString(msg), font)
}

// Enable 在插件启动时被调用
// Disable 在插件禁用时被调用
var Enable, Disable func() int32

// PrivateMsg 在收到私聊消息时被调用
// - subType	子类型，11/来自好友 1/来自在线状态 2/来自群 3/来自讨论组
// - msgId		消息ID
// - fromQQ		来源QQ
// - msg		消息内容
// - font		字体
// 返回非零值,消息将被拦截,最高优先不可拦截
var PrivateMsg func(subType, msgID int32, fromQQ int64, msg string, font int32) int32
