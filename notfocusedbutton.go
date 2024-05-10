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
	label tview.TextView
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
		label:         *tview.NewTextView().SetLabelWidth(3).SetTextAlign(tview.AlignCenter).SetLabel(l),
		style:         tcell.StyleDefault.Background(tview.Styles.ContrastBackgroundColor).Foreground(tview.Styles.PrimaryTextColor),
		disabledStyle: tcell.StyleDefault.Background(tview.Styles.ContrastBackgroundColor).Foreground(tview.Styles.ContrastSecondaryTextColor),
	}

	b.SetDisabled(false)

	return b
}

// SetStyle sets the style of the button used when it is not focused.
func (b *NoneFocusableButton) SetStyle(style tcell.Style) *NoneFocusableButton {
	b.style = style
	return b
}

// SetDisabledStyle sets the style of the button used when it is disabled.
func (b *NoneFocusableButton) SetDisabledStyle(style tcell.Style) *NoneFocusableButton {
	b.disabledStyle = style
	return b
}

// SetDisabled sets whether or not the button is disabled. Disabled buttons
// cannot be activated.
//
// If the button is part of a form, you should set focus to the form itself
// after calling this function to set focus to the next non-disabled form item.
func (b *NoneFocusableButton) SetDisabled(disabled bool) *NoneFocusableButton {
	st := b.style
	if disabled {
		st = b.disabledStyle
	}

	_, bg, _ := st.Decompose()

	b.label.SetBackgroundColor(bg)
	b.label.SetTextStyle(st)
	b.label.SetDisabled(disabled)

	b.disabled = disabled
	return b
}

// IsDisabled returns whether or not the button is disabled.
func (b *NoneFocusableButton) IsDisabled() bool {
	return b.disabled
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

// Draw draws this primitive onto the screen.
func (b *NoneFocusableButton) Draw(screen tcell.Screen) {
	b.label.Draw(screen)
}

// GetRect returns the current position of the rectangle, x, y, width, and
// height.
func (b *NoneFocusableButton) GetRect() (int, int, int, int) {
	return b.label.GetRect()
}

// SetRect sets a new position of the primitive. Note that this has no effect
// if this primitive is part of a layout (e.g. Flex, Grid) or if it was added
// like this:
//
//	application.SetRoot(p, true)
func (b *NoneFocusableButton) SetRect(x, y, width, height int) {
	b.label.SetRect(x, y, width, height)
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
		if action == tview.MouseLeftClick && b.label.InRect(event.Position()) {
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
