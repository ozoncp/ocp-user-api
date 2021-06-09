package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-user-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-user-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/alarm_mock.go -package=mocks github.com/ozoncp/ocp-user-api/internal/alarm Alarm
