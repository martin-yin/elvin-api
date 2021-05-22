package utils

import (
	"danci-api/model"
	"github.com/gin-gonic/gin"
)

type RouterFunc func(context *gin.Context)
type ServiceFunc func(obj string, publicFiles *model.PublicFiles)

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
func (h *Handles) RoutersHandlerRegister(key string, f RouterFunc) {
	if h.RouterHandlers == nil {
		h.RouterHandlers = make(map[string]RouterFunc)
	}
	if _, exist := h.RouterHandlers[key]; exist {
		return
	}
	h.RouterHandlers[key] = f
}
func (h *Handles) ServicesHandlerRegister(key string, f ServiceFunc) {
	if h.ServiceHandlers == nil {
		h.ServiceHandlers = make(map[string]ServiceFunc)
	}
	if _, exist := h.ServiceHandlers[key]; exist {
		return
	}
	h.ServiceHandlers[key] = f
}
