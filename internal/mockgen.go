package internal

//go:generate mockgen -destination=./mock/presentation/mock/flusher_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/flusher/presentation Flusher
//go:generate mockgen -destination=./mock/presentation/mock/repo_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/repo/presentation Repo

//go:generate mockgen -destination=./mock/slide/mock/flusher_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/flusher/slide Flusher
//go:generate mockgen -destination=./mock/slide/mock/repo_mock.go -package=mock github.com/ozoncp/ocp-presentation-api/internal/repo/slide Repo
