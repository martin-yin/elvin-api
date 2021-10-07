package utils

import (
	"dancin-api/model"
	"github.com/gin-gonic/gin"
)

type RouterFunc func(context *gin.Context)
type ServiceFunc func(obj string, commonFiles *model.CommonFiles)

type Handles struct {
	RouterHandlers  map[string]RouterFunc
	ServiceHandlers map[string]ServiceFunc
}

// 创建一个新的空map
func NewHandles() *Handles {
	return &Handles{
		RouterHandlers:  make(map[string]RouterFunc),
		ServiceHandlers: make(map[string]ServiceFunc),
	}
}

func (h *Handles) RoutersHandlerRegister(routerFunc map[string]RouterFunc) {
	for index, item := range routerFunc {
		if h.RouterHandlers == nil {
			h.RouterHandlers = make(map[string]RouterFunc)
		}
		if _, exist := h.RouterHandlers[index]; exist {
			return
		}
		h.RouterHandlers[index] = item
	}
}

func (h *Handles) ServicesHandlerRegister(serviceFunc map[string]ServiceFunc) {
	for index, item := range serviceFunc {
		if h.ServiceHandlers == nil {
			h.ServiceHandlers = make(map[string]ServiceFunc)
		}
		if _, exist := h.ServiceHandlers[index]; exist {
			return
		}
		h.ServiceHandlers[index] = item
	}
}
