package internal

//go:generate mockgen -destination=./mock/presentation/flusher_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/flusher/presentation Flusher
//go:generate mockgen -destination=./mock/presentation/repo_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/repo/presentation Repo
//go:generate mockgen -destination=./mock/presentation/saver_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/saver/presentation Saver
//go:generate mockgen -destination=./mock/presentation/alarm_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/util Alarm

//go:generate mockgen -destination=./mock/slide/flusher_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/flusher/slide Flusher
//go:generate mockgen -destination=./mock/slide/repo_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/repo/slide Repo
//go:generate mockgen -destination=./mock/slide/alarm_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/util Alarm
