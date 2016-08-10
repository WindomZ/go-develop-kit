package mutex

type MutexDebug struct {
	MutexSafe
	debug
}

func NewMutexDebug() *MutexDebug {
	return new(MutexDebug)
}

func (m *MutexDebug) Lock() {
	m.MutexSafe.Lock()
	m.debug.set()
}

func (m *MutexDebug) Unlock() {
	m.debug.reset()
	m.MutexSafe.Unlock()
}

func (m *MutexDebug) LockedDuration() int64 {
	if m.IsLocked() {
		return m.debug.Duration()
	}
	return 0
}

func (m *MutexDebug) Log() string {
	return m.debugString()
}
