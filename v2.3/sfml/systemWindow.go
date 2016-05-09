package sfml

// Used for types Window and RenderWindow
type SystemWindow interface {
	Close()
	IsOpen() bool
	GetSettings() *ContextSettings
	PollEvent() *Event
	WaitEvent() *Event
	GetPosition() Vector2i
	SetPosition(position Vector2i)
	GetSize() Vector2u
	SetSize(size Vector2u)
	SetTitle(title string)
	SetIcon(width, height uint, pixels []uint8)
	SetVisible(visible bool)
	SetMouseCursorVisible(visible bool)
	SetVerticalSyncEnabled(enabled bool)
	SetKeyRepeatEnabled(enabled bool)
	SetActive(active bool) bool
	RequestFocus()
	HasFocus() bool
	Display()
	SetFramerateLimit(limit uint)
	SetJoystickThreshold(treshold float32)
	GetSystemHandle() WindowHandle
}
