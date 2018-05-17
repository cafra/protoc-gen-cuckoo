package usi

type Msg struct {
	Type   int64
	Length int64
	Value  []byte
}
type USinterface interface {
	Start() error           // 必须初始化server才可以调用
	Stop() error            // stop 以后执行清理系统资源并exit(1)
	GetMasterAddr() string  // server未start并且master未选择返回错误
	SendMsg(msg *Msg) error // rugosa未初始化send返回错误
	GetMsgCh() chan *Msg    // 一旦创建server便可以获取channel接口
}
