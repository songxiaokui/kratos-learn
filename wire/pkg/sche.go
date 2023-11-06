package pkg

import (
	"fmt"
	"github.com/google/wire"
)

// 调度器

type Schedule interface {
	Start() // 启动调度任务,定时回调任务
}

type LinuxJob struct {
}

func (l *LinuxJob) Start() {
	fmt.Println("linux job running")
}

type LocalJob struct {
}

func (l *LocalJob) Start() {
	fmt.Println("local job running")
}

func NewSchedule(c Config) Schedule {
	switch c.Sche.Type {
	case "linux":
		return &LinuxJob{}
	case "local":
		return &LocalJob{}
	}
	return Schedule(nil)
}

var ProviderScheduleSet = wire.NewSet(NewSchedule)
