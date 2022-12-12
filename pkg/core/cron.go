package core

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"sync"

	"es-cron/logger"
	"github.com/robfig/cron/v3"
)

var CronJob = make(map[string]*cron.Cron)

type cronHandler struct {
	Name      string
	CronCycle string
	Shell     string
	Cron      *cron.Cron
}

func NewCronHandler(name, cronCycle, shell string) *cronHandler {
	h := &cronHandler{
		Name:      name,
		CronCycle: cronCycle,
		Shell:     shell,
	}
	cr := cron.New()
	CronJob[name] = cr
	h.Cron = cr
	logger.LG.Info(fmt.Sprintf("注册shell命令 Name:%v 执行周期:%v Shell:%v", h.Name, h.CronCycle, h.Shell))
	return h
}

func (cr *cronHandler) Start() error {
	id, err := cr.Cron.AddFunc(cr.CronCycle, func() {
		if err := cr.StartCommand(); err != nil {
			logger.LG.ErrorWithErr("shell error", err)
		}
	})
	logger.LG.Info(fmt.Sprintf("任务注册成功 任务ID %d\n", id))
	cr.Cron.Start()
	return err
}

func (cr *cronHandler) StartCommand() error {
	logger.LG.Info(fmt.Sprintf("%s 执行 %s", cr.Name, cr.Shell))
	c := exec.Command("bash", cr.Shell) // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	if err = c.Start(); err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				fmt.Println(err)
				return
			}
			fmt.Println(readString)
		}
	}()
	wg.Wait()
	logger.LG.Info(fmt.Sprintf("%s 任务执行完成", cr.Name))
	return err
}
