package main

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/util/gconv"
)

type Task struct {
	Name string
	Cron struct {
		Pattern string
		Command string
		Enable  bool
	}
}

func init() {
	g.Log().SetFlags(glog.F_TIME_STD)
}

type UserInfo struct {
	Name   string `bson:"name"`
	Age    uint16 `bson:"age"`
	Weight uint32 `bson:"weight"`
}

type DataTicker struct {
	Pair     string  `bson:"pair"`
	Open     float64 `bson:"open"`
	Close    float64 `bson:"close"`
	High     float64 `bson:"high"`
	Low      float64 `bson:"low"`
	Vol      float64 `bson:"vol"`
	Amount   float64 `bson:"amount"`
	Count    int64   `bson:"count"`
	Increase float64 `bson:"increase"`
	Date     uint64  `bson:"date"`
}

func main() {
	runTasks()
}

func runTasks() {
	tasks, err := loadTasks()
	if err != nil {
		g.Log().Fatal(gctx.New(), err)
	}
	for _, t := range tasks {
		if !t.Cron.Enable {
			continue
		}
		task := t
		_, err = gcron.AddSingleton(gctx.New(), task.Cron.Pattern, func(ctx context.Context) {
			result, err := gproc.ShellExec(task.Cron.Command)
			g.Log().Cat(task.Name).Info(context.Background(), result, err)
		}, task.Name)
		if err != nil {
			g.Log().Cat(task.Name).Error(gctx.New(), err)
		}
	}
	gcron.Start()
	select {}
}

func loadTasks() (tasks []Task, err error) {
	adapter := g.Cfg().GetAdapter().(*gcfg.AdapterFile)
	adapter.SetFileName("cron.yaml")
	data, err := g.Cfg().Get(gctx.New(), "tasks")
	if err != nil {
		return
	}
	err = gconv.Scan(data, &tasks)
	if err != nil {
		return
	}
	return
}
