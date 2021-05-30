package config


import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

var (
	ErrCallbackNotImp    = errors.New("callback function not implement")
	ErrGroupNotExist     = errors.New("rainbow group not exist")
	ErrRainbowEnvNotExit = errors.New("rainbow env not exist")
)

// Rainbow
type Rainbow struct {
	connectionString string
	appID            string
	groupName        string
	//configAPI        *confapi.ConfAPI
	//configOptions    []types.AssignGetOption
	viper            *viper.Viper
}

// NewRainbow 创建七彩石配置实例
func NewRainbow(connectionString string, appID string, groupName string, userID string, userKey string) (*Rainbow, error) {
	fmt.Printf("[Rainbow New] ConnectionString: %s, appId: %s, groupName: %s\n", connectionString, appID, groupName)
	if connectionString == "" || appID == "" || groupName == "" {
		return nil, ErrRainbowEnvNotExit
	}
	rainbow := &Rainbow{
		connectionString: connectionString,
		appID:            appID,
		groupName:        groupName,
	}

	//rainbow.configOptions = getOpts
	return rainbow, nil
}




