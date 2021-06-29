package internal

//go:generate mockgen -destination=./ocp-presentation-api/mock/alarm_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/common/alarm Alarm
//go:generate mockgen -destination=./ocp-slide-api/mock/alarm_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/common/alarm Alarm
//go:generate mockgen -destination=./ocp-presentation-api/mock/clock_wrapper_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/common/clockwrapper ClockWrapper
//go:generate mockgen -destination=./ocp-slide-api/mock/clock_wrapper_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/common/clockwrapper ClockWrapper

//go:generate mockgen -destination=./ocp-presentation-api/mock/flusher_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/flusher Flusher
//go:generate mockgen -destination=./ocp-presentation-api/mock/repo_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/repo Repo
//go:generate mockgen -destination=./ocp-presentation-api/mock/saver_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/saver Saver

//go:generate mockgen -destination=./ocp-slide-api/mock/flusher_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/flusher Flusher
//go:generate mockgen -destination=./ocp-slide-api/mock/repo_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/repo Repo
