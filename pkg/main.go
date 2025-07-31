package main

import (
	"github.com/guothion/xuanyuan/pkg/api"
	"github.com/guothion/xuanyuan/pkg/common"
	"github.com/guothion/xuanyuan/pkg/config"
	"github.com/guothion/xuanyuan/pkg/dataSource"
)

func main() {
	config.Init()
	common.Init()
	dataSource.Init()
	api.Server()
}
