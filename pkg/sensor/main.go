package sensor

import (
	"github.com/capsule8/capsule8/pkg/config"
	"github.com/capsule8/capsule8/pkg/services"
	"github.com/golang/glog"
)

// Main is the main entrypoint for the sensor
func Main() {
	manager := services.NewServiceManager()
	if len(config.Global.ProfilingAddr) > 0 {
		service := services.NewProfilingService(
			config.Global.ProfilingAddr)
		manager.RegisterService(service)
	}
	if len(config.Sensor.ServerAddr) > 0 {
		service, err := NewTelemetryService(config.Sensor.ServerAddr)
		if err != nil {
			glog.Fatalf("Could not create telemetry service: %s", err)
		}
		manager.RegisterService(service)
	}

	manager.Run()
}
