package form

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// NoneFocusableButton delegate focus to another element when button was clicked
type NoneFocusableButton struct {
	*tview.Button

	focusable tview.Primitive
	click     func()
}

// NewNoneFocusableButton returns a NoneFocusableButton without a border.
func NewNoneFocusableButton(l string) *NoneFocusableButton {
	b := &NoneFocusableButton{
		Button: tview.NewButton(l),
	}

	b.SetDisabledStyle(tcell.StyleDefault.Background(tview.Styles.ContrastBackgroundColor).Foreground(tview.Styles.ContrastSecondaryTextColor))
	b.SetStyle(tcell.StyleDefault.Background(tview.Styles.ContrastBackgroundColor).Foreground(tview.Styles.PrimaryTextColor))
	return b
}

// Draw draws this primitive onto the screen.
func (b *NoneFocusableButton) SetFocusable(f tview.Primitive) *NoneFocusableButton {
	b.focusable = f
	return b
}

// Draw draws this primitive onto the screen.
func (b *NoneFocusableButton) SetClick(c func()) *NoneFocusableButton {
	b.click = c
	return b
}

// InputHandler returns nil. NoneFocusableButton has no default input handling.
func (b *NoneFocusableButton) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return nil
}

// Focus is called when this primitive receives focus.
func (b *NoneFocusableButton) Focus(delegate func(p tview.Primitive)) {
	delegate(b.focusable)
}

// HasFocus returns whether or not this primitive has focus.
func (b *NoneFocusableButton) HasFocus() bool {
	return false
}

// Blur is called when this primitive loses focus.
func (b *NoneFocusableButton) Blur() {

}

// MouseHandler returns nil. NoneFocusableButton has no default mouse handling.
func (b *NoneFocusableButton) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	return func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
		if action == tview.MouseLeftClick && b.InRect(event.Position()) {
			if b.click != nil {
				b.click()
			}
			consumed = true
		}
		return
	}
}

// PasteHandler returns nil. NoneFocusableButton has no default paste handling.
func (b *NoneFocusableButton) PasteHandler() func(pastedText string, setFocus func(p tview.Primitive)) {
	return nil
}
