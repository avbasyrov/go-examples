package logic

import "github.com/prometheus/client_golang/prometheus"

type Logic struct {
	newUsersCounter  prometheus.Counter
	usersOnlineGauge prometheus.Gauge
}

func New(newUsersCounter prometheus.Counter, usersOnlineGauge prometheus.Gauge) *Logic {
	return &Logic{
		newUsersCounter:  newUsersCounter,
		usersOnlineGauge: usersOnlineGauge,
	}
}

func (l *Logic) RegisterNewUser() {
	l.newUsersCounter.Inc()
}

func (l *Logic) LoginUser() {
	l.usersOnlineGauge.Inc()
}

func (l *Logic) LogoutUser() {
	l.usersOnlineGauge.Dec()
}
