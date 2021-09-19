package logic

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogic_RegisterNewUser(t *testing.T) {
	usersRegistered, usersOnline := newFakeMetrics()
	l := New(usersRegistered, usersOnline)

	assert.Equal(t, float64(0), testutil.ToFloat64(usersRegistered))
	assert.Equal(t, float64(0), testutil.ToFloat64(usersOnline))

	l.RegisterNewUser()

	assert.Equal(t, float64(1), testutil.ToFloat64(usersRegistered))
	assert.Equal(t, float64(0), testutil.ToFloat64(usersOnline))
}

func TestLogic_LoginUser(t *testing.T) {
	usersRegistered, usersOnline := newFakeMetrics()
	l := New(usersRegistered, usersOnline)

	assert.Equal(t, float64(0), testutil.ToFloat64(usersRegistered))
	assert.Equal(t, float64(0), testutil.ToFloat64(usersOnline))

	l.LoginUser()

	assert.Equal(t, float64(0), testutil.ToFloat64(usersRegistered))
	assert.Equal(t, float64(1), testutil.ToFloat64(usersOnline))

	l.LoginUser()

	assert.Equal(t, float64(0), testutil.ToFloat64(usersRegistered))
	assert.Equal(t, float64(2), testutil.ToFloat64(usersOnline))
}

func TestLogic_LogoutUser(t *testing.T) {
	usersRegistered, usersOnline := newFakeMetrics()
	l := New(usersRegistered, usersOnline)

	assert.Equal(t, float64(0), testutil.ToFloat64(usersRegistered))
	assert.Equal(t, float64(0), testutil.ToFloat64(usersOnline))

	l.LogoutUser()

	assert.Equal(t, float64(0), testutil.ToFloat64(usersRegistered))
	assert.Equal(t, float64(-1), testutil.ToFloat64(usersOnline))

	l.LogoutUser()

	assert.Equal(t, float64(0), testutil.ToFloat64(usersRegistered))
	assert.Equal(t, float64(-2), testutil.ToFloat64(usersOnline))
}

func newFakeMetrics() (prometheus.Counter, prometheus.Gauge) {
	usersRegistered := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "users_registered",
		})

	usersOnline := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "users_online",
		})

	return usersRegistered, usersOnline
}
