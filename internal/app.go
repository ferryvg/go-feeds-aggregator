package internal

type App interface {
	Boot() error
	Stop() error
}
