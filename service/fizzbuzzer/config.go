package fizzbuzzer

type Config struct {
	Backend string
	Routine *RoutineConfig
	Basic   *BasicConfig
	Split   *SplitConfig
}

type RoutineConfig struct {
	MaxConcurrentRoutines int // Field to set up define go routines limiter
}

type SplitConfig struct {
	MaxConcurrentRoutines int // Field to set up define go routines limiter
}

type BasicConfig struct{}
