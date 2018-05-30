package tailf

import (
	"fmt"
	"github.com/hpcloud/tail"
	G "gopractice/day11/logagent/g"
	"github.com/astaxie/beego/logs"
	"time"
)


type TextMsg struct {
	Msg string
	Topic string
}

type TailObj struct {
	tail *tail.Tail
	conf G.CollectConf
}

type TailObjManager struct {
	tailSlice []*TailObj
	MsgChan chan *TextMsg //消息管道
}

var (
	tailObjMgr *TailObjManager
)



func InitTail(conf []G.CollectConf, chanSize int) (err error) {
	if len(conf) == 0 {
		err = fmt.Errorf("invalid config for log collect, conf: %v", conf)
		return
	}

	tailObjMgr = &TailObjManager{
		MsgChan: make(chan *TextMsg, chanSize),
	}

	for _,v := range conf {
		obj := &TailObj{
			conf: v,
		}

		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen: true,
			Follow: true,
			Location: &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll: true,
		})

		if errTail != nil {
			err = errTail
			return
		}

		obj.tail = tails

		tailObjMgr.tailSlice = append(tailObjMgr.tailSlice, obj)

		go readFromTail(obj)

	}
	return
}


func readFromTail(tailobj *TailObj) {
	for true {
		line, ok := <- tailobj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen, filename: %s\n", tailobj.tail.Filename)
			time.Sleep(100*time.Millisecond)
			continue
		}
		textMsg := &TextMsg{
			Msg: line.Text,
			Topic: tailobj.conf.Topic,
		}

		tailObjMgr.MsgChan <- textMsg
	}
}

func GetOneLine() (msg *TextMsg) {
	msg = <- tailObjMgr.MsgChan
	return
}