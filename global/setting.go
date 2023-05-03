package global

import (
	"blogService/pkg/logger"
	"blogService/pkg/setting"
)

var (
	// 配置
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	// 日志
	Logger *logger.Logger
)
