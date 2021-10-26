//go:build !windows
// +build !windows

/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package main

import (
	"os"

	process_worker "github.com/eolinker/eosc/process-worker"

	"github.com/eolinker/eosc"

	admin_open_api "github.com/eolinker/eosc/modules/admin-open-api"
	"github.com/eolinker/eosc/process-master/admin"

	"github.com/eolinker/eosc/eoscli"
	"github.com/eolinker/eosc/log"
	"github.com/eolinker/eosc/process"
)

func init() {
	registerInnerExtenders()
	admin.Register("/api/", admin_open_api.CreateHandler())
	process.Register(eosc.ProcessWorker, process_worker.Process)
	process.Register(eosc.ProcessMaster, ProcessMaster)
	process.Register(eosc.ProcessHelper, ProcessHelper)
}

func main() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		log.Error("main recover error: ", err)
	//	}
	//	log.Close()
	//}()
	if process.Run() {
		//log.Close()
		return
	}
	//if env.IsDebug() {
	//	if process.RunDebug(eosc.ProcessMaster) {
	//		log.Info("debug done")
	//	} else {
	//		log.Warn("debug not exist")
	//	}
	//	//log.Close()
	//	return
	//}
	app := eoscli.NewApp()
	app.Default()
	err := app.Run(os.Args)
	if err != nil {
		log.Error(err)
	}
	log.Close()
}
