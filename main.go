package RouterForGuild

import (
	"RouterForGuild/global"
	"RouterForGuild/global/terminal"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

// init 初始化程序
func init() {
	global.InitLog()

	log.Info("项目地址：https://github.com/Kyomotoi/RouterForGuild")
	log.Info("当前版本：" + global.Version())
	log.Info("交流群：567297659")

	log.Info("正在导入设置...")
	log.Info("即将使用 config.yml 内的配置启动 Router")
	conf, err := global.ConfigDealer()
	if err != nil {
		if os.IsNotExist(err) {
			log.Warning("检测为初次启动，已于同目录下生成 config.yml，请配置并重新启动！")
			genErr := global.ConfigGener()
			if genErr != nil {
				log.Error("无法创建文件：config.yml，请确认是否给足系统权限")
			}
		} else {
			log.Warning("处理 config.yml 失败")
		}

		log.Warning("将于5秒后退出")
		time.Sleep(time.Second * 5)
		os.Exit(1)
	}

	if conf.Debug {
		log.Debug("DEBUG 已启用")
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if terminal.RunningByDoubleClick() {
		log.Warning("不建议直接双击运行本程序，这将导致一些非可预料后果，请通过*控制台*启动本程序")
		log.Warning("将等待10秒后启动")
		time.Sleep(time.Second * 10)
	}

	log.Info("アトリは、高性能ですから！")
}

func main() {
	// to do...
}
