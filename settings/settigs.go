package settings

// 使用viper 管理文件
import (
	"fmt"
	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config") // 指定配置文件名称(不需要带后缀)
	viper.SetConfigType("yaml")   // 指定配置文件类型
	viper.AddConfigPath("./")     // 指定配置文件路径
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper.ReadInConfig() failed", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了....")
	})
	return
}
