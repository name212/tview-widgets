package form

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// NoneFocusableButton implements the Primitive interface with an empty background and optional
// elements such as a border and a title. NoneFocusableButton itself does not hold any content
// but serves as the superclass of all other primitives. Subclasses add their
// own content, typically (but not necessarily) keeping their content within the
// NoneFocusableButton's rectangle.
//
// NoneFocusableButton provides a number of utility functions available to all primitives.
//
// See https://github.com/rivo/tview/wiki/NoneFocusableButton for an example.
type NoneFocusableButton struct {
	*tview.TextView
	// The button's style (when deactivated).
	style tcell.Style

	// The button's style (when disabled).
	disabledStyle tcell.Style

	disabled bool

	focusable tview.Primitive
	click     func()
}

// NewNoneFocusableButton returns a NoneFocusableButton without a border.
func NewNoneFocusableButton(l string) *NoneFocusableButton {
	b := &NoneFocusableButton{
		TextView:      tview.NewTextView().SetTextAlign(tview.AlignCenter).SetLabel(l),
		style:         tcell.StyleDefault.Background(tview.Styles.ContrastBackgroundColor).Foreground(tview.Styles.PrimaryTextColor),
		disabledStyle: tcell.StyleDefault.Background(tview.Styles.ContrastBackgroundColor).Foreground(tview.Styles.ContrastSecondaryTextColor),
	}

	_, bg, _ := b.style.Decompose()

	b.SetTextStyle(b.style)
	b.SetBackgroundColor(bg)

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
