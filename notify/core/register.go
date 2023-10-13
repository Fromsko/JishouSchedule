package core

import (
	"notify/utils"
)

type Server struct {
	TemplateID string
	RegisterServerList []RegisterServer
}

type RegisterServer struct {
	NickName string // 订阅者 别名
	UserID   string // 订阅者 ID
}

func (receiver RegisterServer) Notify(schedule Service) {
	schedule.Template.ToUser = receiver.UserID
	// 推送
	schedule.SendMsg(schedule.TempInfo, func(resp string) {
		utils.Log.Info("❤️ 成功推送给 => " + receiver.NickName)
	})
}

// NewRegister 注册一个服务
func NewRegister(s *Server) (Schedule *ScheduleService, Serve *Service) {
	Schedule = &ScheduleService{}

	for _, user := range s.RegisterServerList {
		Schedule.Subscribe(user)
	}

	Serve = InitServe(s.TemplateID)

	return Schedule, Serve
}
