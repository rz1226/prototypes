package tickwarn

import (
	"fmt"
	"time"
)

//时不时的刺激一下，如果超过某个时间段没有收到刺激，则会报警
/**
warn := NewTickWarn(300，100 ) //五分钟没有tick，则报警，报警频率最高100秒一次

*/

type TickWarn struct {
	MinTickDuration int64 //事件低于这个频率就报警  单位秒
	MinWarnDuration int64 //报警周期不高于这里 单位秒

	LastTimeTick int64 //最后一次事件时间  时间戳
	LastTimeWarn int64 //最后一次报警时间 时间戳

	WarnFunc func() //报警函数，报警的时候会触发它
}

func NewTickWarn(minTickDuration, minWarnDuration int64, f func()) *TickWarn {
	t := &TickWarn{}
	t.MinTickDuration = minTickDuration
	t.MinWarnDuration = minWarnDuration
	//初始值防止一启动就误报
	t.LastTimeWarn = time.Now().Unix()

	t.WarnFunc = f
	t.run()
	return t
}

func (t *TickWarn) Tick() {
	fmt.Println("tick", time.Now().Format("2006-01-02 15:04:05"))
	t.LastTimeTick = time.Now().Unix()

}

func (t *TickWarn) warn() bool {
	now := time.Now().Unix()
	if now-t.LastTimeWarn < t.MinWarnDuration {
		return false
	}
	if now-t.LastTimeTick < t.MinTickDuration {
		return false
	}

	t.WarnFunc()
	t.LastTimeWarn = now
	return true

}

func (t *TickWarn) run() {
	//先运行一下事件防止误报
	t.Tick()
	f := func() {
		for {
			time.Sleep(time.Millisecond * 500)
			t.warn()
		}
	}
	go f()
}
