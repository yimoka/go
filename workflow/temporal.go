package workflow

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/yimoka/go/config"
	"go.temporal.io/sdk/client"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
)

type TemporalLogger struct {
	logger *log.Helper
}

func (l *TemporalLogger) Debug(msg string, keyvals ...interface{}) {
	l.logger.Debug(msg, keyvals)
}

func (l *TemporalLogger) Info(msg string, keyvals ...interface{}) {
	l.logger.Info(msg, keyvals)
}

func (l *TemporalLogger) Warn(msg string, keyvals ...interface{}) {
	l.logger.Warn(msg, keyvals)
}

func (l *TemporalLogger) Error(msg string, keyvals ...interface{}) {
	l.logger.Error(msg, keyvals)
}

func GetTemporalLogger(helper *log.Helper) tlog.Logger {
	return &TemporalLogger{logger: helper}
}

func GetTemporalClientOptions(config *config.Config) client.Options {
	f := config.Workflow
	opt := client.Options{}
	if f == nil {
		return opt
	}
	if f.HostPort != "" {
		opt.HostPort = f.HostPort
	}
	if f.Namespace != "" {
		opt.Namespace = f.Namespace
	}
	if f.Identity != "" {
		opt.Identity = f.Identity
	}
	if f.DisableErrorCodeMetricTags {
		opt.DisableErrorCodeMetricTags = f.DisableErrorCodeMetricTags
	}
	return opt
}

func GetTemporalWorkerOptions(workerName string, config *config.Config) worker.Options {
	opt := worker.Options{}
	f := config.Workflow
	if f == nil {
		return opt
	}
	w, ok := f.Workers[workerName]
	if !ok || w == nil {
		return opt
	}
	if w.MaxConcurrentActivityExecutionSize != 0 {
		opt.MaxConcurrentActivityExecutionSize = int(w.MaxConcurrentActivityExecutionSize)
	}
	if w.WorkerActivitiesPerSecond != 0 {
		opt.WorkerActivitiesPerSecond = float64(w.WorkerActivitiesPerSecond)
	}
	if w.MaxConcurrentLocalActivityExecutionSize != 0 {
		opt.MaxConcurrentLocalActivityExecutionSize = int(w.MaxConcurrentLocalActivityExecutionSize)
	}
	if w.WorkerLocalActivitiesPerSecond != 0 {
		opt.WorkerLocalActivitiesPerSecond = float64(w.WorkerLocalActivitiesPerSecond)
	}
	if w.TaskQueueActivitiesPerSecond != 0 {
		opt.TaskQueueActivitiesPerSecond = float64(w.TaskQueueActivitiesPerSecond)
	}
	if w.MaxConcurrentActivityTaskPollers != 0 {
		opt.MaxConcurrentActivityTaskPollers = int(w.MaxConcurrentActivityTaskPollers)
	}
	if w.MaxConcurrentWorkflowTaskExecutionSize != 0 {
		opt.MaxConcurrentWorkflowTaskExecutionSize = int(w.MaxConcurrentWorkflowTaskExecutionSize)
	}
	if w.MaxConcurrentWorkflowTaskPollers != 0 {
		opt.MaxConcurrentWorkflowTaskPollers = int(w.MaxConcurrentWorkflowTaskPollers)
	}
	if w.EnableLoggingInReplay {
		opt.EnableLoggingInReplay = w.EnableLoggingInReplay
	}
	if w.StickyScheduleToStartTimeout != 0 {
		opt.StickyScheduleToStartTimeout = time.Microsecond * time.Duration(w.StickyScheduleToStartTimeout)
	}
	if w.WorkerStopTimeout != 0 {
		opt.WorkerStopTimeout = time.Microsecond * time.Duration(w.WorkerStopTimeout)
	}
	if w.EnableSessionWorker {
		opt.EnableSessionWorker = w.EnableSessionWorker
	}
	if w.MaxConcurrentSessionExecutionSize != 0 {
		opt.MaxConcurrentSessionExecutionSize = int(w.MaxConcurrentSessionExecutionSize)
	}
	if w.DisableWorkflowWorker {
		opt.DisableWorkflowWorker = w.DisableWorkflowWorker
	}
	if w.LocalActivityWorkerOnly {
		opt.LocalActivityWorkerOnly = w.LocalActivityWorkerOnly
	}
	if w.Identity != "" {
		opt.Identity = w.Identity
	}
	if w.DeadlockDetectionTimeout != 0 {
		opt.DeadlockDetectionTimeout = time.Microsecond * time.Duration(w.DeadlockDetectionTimeout)
	}
	if w.MaxHeartbeatThrottleInterval != 0 {
		opt.MaxHeartbeatThrottleInterval = time.Microsecond * time.Duration(w.MaxHeartbeatThrottleInterval)
	}
	if w.DefaultHeartbeatThrottleInterval != 0 {
		opt.DefaultHeartbeatThrottleInterval = time.Microsecond * time.Duration(w.DefaultHeartbeatThrottleInterval)
	}
	if w.DisableEagerActivities {
		opt.DisableEagerActivities = w.DisableEagerActivities
	}
	if w.MaxConcurrentEagerActivityExecutionSize != 0 {
		opt.MaxConcurrentEagerActivityExecutionSize = int(w.MaxConcurrentEagerActivityExecutionSize)
	}
	if w.DisableRegistrationAliasing {
		opt.DisableRegistrationAliasing = w.DisableRegistrationAliasing
	}
	if w.BuildID != "" {
		opt.BuildID = w.BuildID
	}
	if w.UseBuildIDForVersioning {
		opt.UseBuildIDForVersioning = w.UseBuildIDForVersioning
	}
	return opt
}
