package common

import (
	"errors"
)

type Spider interface {
	SpiderUrl(url string) error
}

func NewSpider(from string) (Spider, error) {
	DEBUG("new spider xxx")
	switch from {
	case "booktxt":
		return new(BookTextSpider), nil
	default:
		return nil, errors.New("系统暂未处理该类型的配置文件")
	}
}