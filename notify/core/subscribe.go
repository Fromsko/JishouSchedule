package core

// Subscriber 订阅者接口
type Subscriber interface {
	Notify(schedule Service)
}

// ScheduleService 课表推送服务
type ScheduleService struct {
	subscribers []Subscriber
}

// Subscribe 添加订阅者
func (s *ScheduleService) Subscribe(subscriber ...Subscriber) {
	s.subscribers = append(s.subscribers, subscriber...)
}

// Unsubscribe 移除订阅者
func (s *ScheduleService) Unsubscribe(subscriber Subscriber) {
	for i, sub := range s.subscribers {
		if sub == subscriber {
			s.subscribers = append(s.subscribers[:i], s.subscribers[i+1:]...)
			break
		}
	}
}

// PushSchedule 推送课表信息给所有订阅者
func (s *ScheduleService) PushSchedule(schedule Service) {
	for _, sub := range s.subscribers {
		sub.Notify(schedule)
	}
}
