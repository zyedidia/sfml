package sfml

//#include <SFML/Window.h>
import "C"

const (
	KeyUnknown   KeyCode = C.sfKeyUnknown
	KeyA         KeyCode = C.sfKeyA
	KeyB         KeyCode = C.sfKeyB
	KeyC         KeyCode = C.sfKeyC
	KeyD         KeyCode = C.sfKeyD
	KeyE         KeyCode = C.sfKeyE
	KeyF         KeyCode = C.sfKeyF
	KeyG         KeyCode = C.sfKeyG
	KeyH         KeyCode = C.sfKeyH
	KeyI         KeyCode = C.sfKeyI
	KeyJ         KeyCode = C.sfKeyJ
	KeyK         KeyCode = C.sfKeyK
	KeyL         KeyCode = C.sfKeyL
	KeyM         KeyCode = C.sfKeyM
	KeyN         KeyCode = C.sfKeyN
	KeyO         KeyCode = C.sfKeyO
	KeyP         KeyCode = C.sfKeyP
	KeyQ         KeyCode = C.sfKeyQ
	KeyR         KeyCode = C.sfKeyR
	KeyS         KeyCode = C.sfKeyS
	KeyT         KeyCode = C.sfKeyT
	KeyU         KeyCode = C.sfKeyU
	KeyV         KeyCode = C.sfKeyV
	KeyW         KeyCode = C.sfKeyW
	KeyX         KeyCode = C.sfKeyX
	KeyY         KeyCode = C.sfKeyY
	KeyZ         KeyCode = C.sfKeyZ
	KeyNum0      KeyCode = C.sfKeyNum0
	KeyNum1      KeyCode = C.sfKeyNum1
	KeyNum2      KeyCode = C.sfKeyNum2
	KeyNum3      KeyCode = C.sfKeyNum3
	KeyNum4      KeyCode = C.sfKeyNum4
	KeyNum5      KeyCode = C.sfKeyNum5
	KeyNum6      KeyCode = C.sfKeyNum6
	KeyNum7      KeyCode = C.sfKeyNum7
	KeyNum8      KeyCode = C.sfKeyNum8
	KeyNum9      KeyCode = C.sfKeyNum9
	KeyEscape    KeyCode = C.sfKeyEscape
	KeyLControl  KeyCode = C.sfKeyLControl
	KeyLShift    KeyCode = C.sfKeyLShift
	KeyLAlt      KeyCode = C.sfKeyLAlt
	KeyLSystem   KeyCode = C.sfKeyLSystem
	KeyRControl  KeyCode = C.sfKeyRControl
	KeyRShift    KeyCode = C.sfKeyRShift
	KeyRAlt      KeyCode = C.sfKeyRAlt
	KeyRSystem   KeyCode = C.sfKeyRSystem
	KeyMenu      KeyCode = C.sfKeyMenu
	KeyLBracket  KeyCode = C.sfKeyLBracket
	KeyRBracket  KeyCode = C.sfKeyRBracket
	KeySemiColon KeyCode = C.sfKeySemiColon
	KeyComma     KeyCode = C.sfKeyComma
	KeyPeriod    KeyCode = C.sfKeyPeriod
	KeyQuote     KeyCode = C.sfKeyQuote
	KeySlash     KeyCode = C.sfKeySlash
	KeyBackSlash KeyCode = C.sfKeyBackSlash
	KeyTilde     KeyCode = C.sfKeyTilde
	KeyEqual     KeyCode = C.sfKeyEqual
	KeyDash      KeyCode = C.sfKeyDash
	KeySpace     KeyCode = C.sfKeySpace
	KeyReturn    KeyCode = C.sfKeyReturn
	KeyBack      KeyCode = C.sfKeyBack
	KeyTab       KeyCode = C.sfKeyTab
	KeyPageUp    KeyCode = C.sfKeyPageUp
	KeyPageDown  KeyCode = C.sfKeyPageDown
	KeyEnd       KeyCode = C.sfKeyEnd
	KeyHome      KeyCode = C.sfKeyHome
	KeyInsert    KeyCode = C.sfKeyInsert
	KeyDelete    KeyCode = C.sfKeyDelete
	KeyAdd       KeyCode = C.sfKeyAdd
	KeySubtract  KeyCode = C.sfKeySubtract
	KeyMultiply  KeyCode = C.sfKeyMultiply
	KeyDivide    KeyCode = C.sfKeyDivide
	KeyLeft      KeyCode = C.sfKeyLeft
	KeyRight     KeyCode = C.sfKeyRight
	KeyUp        KeyCode = C.sfKeyUp
	KeyDown      KeyCode = C.sfKeyDown
	KeyNumpad0   KeyCode = C.sfKeyNumpad0
	KeyNumpad1   KeyCode = C.sfKeyNumpad1
	KeyNumpad2   KeyCode = C.sfKeyNumpad2
	KeyNumpad3   KeyCode = C.sfKeyNumpad3
	KeyNumpad4   KeyCode = C.sfKeyNumpad4
	KeyNumpad5   KeyCode = C.sfKeyNumpad5
	KeyNumpad6   KeyCode = C.sfKeyNumpad6
	KeyNumpad7   KeyCode = C.sfKeyNumpad7
	KeyNumpad8   KeyCode = C.sfKeyNumpad8
	KeyNumpad9   KeyCode = C.sfKeyNumpad9
	KeyF1        KeyCode = C.sfKeyF1
	KeyF2        KeyCode = C.sfKeyF2
	KeyF3        KeyCode = C.sfKeyF3
	KeyF4        KeyCode = C.sfKeyF4
	KeyF5        KeyCode = C.sfKeyF5
	KeyF6        KeyCode = C.sfKeyF6
	KeyF7        KeyCode = C.sfKeyF7
	KeyF8        KeyCode = C.sfKeyF8
	KeyF9        KeyCode = C.sfKeyF9
	KeyF10       KeyCode = C.sfKeyF10
	KeyF11       KeyCode = C.sfKeyF11
	KeyF12       KeyCode = C.sfKeyF12
	KeyF13       KeyCode = C.sfKeyF13
	KeyF14       KeyCode = C.sfKeyF14
	KeyF15       KeyCode = C.sfKeyF15
	KeyPause     KeyCode = C.sfKeyPause
	KeyCount     KeyCode = C.sfKeyCount
)

type KeyCode int

func KeyboardIsKeyPressed(key KeyCode) bool {
	return goBool(C.sfKeyboard_isKeyPressed(C.sfKeyCode(key)))
}
