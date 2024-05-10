package form

import (
	"image"

	"github.com/gdamore/tcell/v2"
	. "github.com/rivo/tview"
)

// FormScrollable allows you to combine multiple one-line form elements into a vertical
// or horizontal layout. FormScrollable elements include types such as InputField or
// Checkbox. These elements can be optionally followed by one or more buttons
// for which you can define form-wide actions (e.g. Save, Clear, Cancel).
//
// See https://github.com/rivo/tview/wiki/FormScrollable for an example.
type FormScrollable struct {
	*Box

	// The items of the form (one row per item).
	items []FormItem

	// The buttons of the form.
	buttons []*Button

	// If set to true, instead of position items and buttons from top to bottom,
	// they are positioned from left to right.
	horizontal bool

	// The alignment of the buttons.
	buttonsAlign int

	// The number of empty cells between items.
	itemPadding int

	// The index of the item or button which has focus. (Items are counted first,
	// buttons are counted last.) This is only used when the form itself receives
	// focus so that the last element that had focus keeps it.
	focusedElement int

	// The label color.
	labelColor tcell.Color

	// The background color of the input area.
	fieldBackgroundColor tcell.Color

	// The text color of the input area.
	fieldTextColor tcell.Color

	// The style of the buttons when they are not focused.
	buttonStyle tcell.Style

	// The style of the buttons when they are focused.
	buttonActivatedStyle tcell.Style

	// The style of the buttons when they are disabled.
	buttonDisabledStyle tcell.Style

	// The last (valid) key that wsa sent to a "finished" handler or -1 if no
	// such key is known yet.
	lastFinishedKey tcell.Key

	// An optional function which is called when the user hits Escape.
	cancel func()

	// Scroll buttons
	upScrollButton   *NoneFocusableButton
	downScrollButton *NoneFocusableButton
}

// NewFormScrollable returns a new form.
func NewFormScrollable() *FormScrollable {
	box := NewBox().SetBorderPadding(1, 1, 1, 1)

	f := &FormScrollable{
		Box:                  box,
		itemPadding:          1,
		labelColor:           Styles.SecondaryTextColor,
		fieldBackgroundColor: Styles.ContrastBackgroundColor,
		fieldTextColor:       Styles.PrimaryTextColor,
		buttonStyle:          tcell.StyleDefault.Background(Styles.ContrastBackgroundColor).Foreground(Styles.PrimaryTextColor),
		buttonActivatedStyle: tcell.StyleDefault.Background(Styles.PrimaryTextColor).Foreground(Styles.ContrastBackgroundColor),
		buttonDisabledStyle:  tcell.StyleDefault.Background(Styles.ContrastBackgroundColor).Foreground(Styles.ContrastSecondaryTextColor),
		lastFinishedKey:      tcell.KeyTab, // To skip over inactive elements at the beginning of the form.

		downScrollButton: NewNoneFocusableButton("\u2193"),
		upScrollButton:   NewNoneFocusableButton("\u2191"),
	}

	onNext := func() {
		all := len(f.items) + len(f.buttons)
		var nn func(int)
		nn = func(next int) {
			if next > 0 {
				f.upScrollButton.SetDisabled(false)
			}

			if next >= all-1 {
				f.downScrollButton.SetDisabled(true)
			}

			if next >= all {
				return
			}

			if next < len(f.items) {
				if _, ok := f.GetFormItem(next).(*TextView); ok {
					nn(next + 1)
					return
				}
			}

			f.SetFocus(next)
		}

		nn(f.focusedElement + 1)

	}

	f.downScrollButton.SetFocusable(f).SetClick(onNext).SetDisabled(false)

	onBack := func() {
		var bb func(int)
		bb = func(prev int) {
			if prev == 0 {
				f.upScrollButton.SetDisabled(true)
			}

			if prev < 0 {
				return
			}

			f.downScrollButton.SetDisabled(false)

			if prev < len(f.items) {
				if _, ok := f.GetFormItem(prev).(*TextView); ok {
					bb(prev - 1)
					return
				}
			}

			f.SetFocus(prev)
		}

		bb(f.focusedElement - 1)

	}

	f.upScrollButton.SetFocusable(f).SetClick(onBack).SetDisabled(true)

	return f
}

// SetItemPadding sets the number of empty rows between form items for vertical
// layouts and the number of empty cells between form items for horizontal
// layouts.
func (f *FormScrollable) SetItemPadding(padding int) *FormScrollable {
	f.itemPadding = padding
	return f
}

// SetHorizontal sets the direction the form elements are laid out. If set to
// true, instead of positioning them from top to bottom (the default), they are
// positioned from left to right, moving into the next row if there is not
// enough space.
func (f *FormScrollable) SetHorizontal(horizontal bool) *FormScrollable {
	f.horizontal = horizontal
	return f
}

// SetLabelColor sets the color of the labels.
func (f *FormScrollable) SetLabelColor(color tcell.Color) *FormScrollable {
	f.labelColor = color
	return f
}

// SetFieldBackgroundColor sets the background color of the input areas.
func (f *FormScrollable) SetFieldBackgroundColor(color tcell.Color) *FormScrollable {
	f.fieldBackgroundColor = color
	return f
}

// SetFieldTextColor sets the text color of the input areas.
func (f *FormScrollable) SetFieldTextColor(color tcell.Color) *FormScrollable {
	f.fieldTextColor = color
	return f
}

// SetButtonsAlign sets how the buttons align horizontally, one of AlignLeft
// (the default), AlignCenter, and AlignRight. This is only
func (f *FormScrollable) SetButtonsAlign(align int) *FormScrollable {
	f.buttonsAlign = align
	return f
}

// SetButtonBackgroundColor sets the background color of the buttons. This is
// also the text color of the buttons when they are focused.
func (f *FormScrollable) SetButtonBackgroundColor(color tcell.Color) *FormScrollable {
	f.buttonStyle = f.buttonStyle.Background(color)
	f.buttonActivatedStyle = f.buttonActivatedStyle.Foreground(color)
	return f
}

// SetButtonTextColor sets the color of the button texts. This is also the
// background of the buttons when they are focused.
func (f *FormScrollable) SetButtonTextColor(color tcell.Color) *FormScrollable {
	f.buttonStyle = f.buttonStyle.Foreground(color)
	f.buttonActivatedStyle = f.buttonActivatedStyle.Background(color)
	return f
}

// SetButtonStyle sets the style of the buttons when they are not focused.
func (f *FormScrollable) SetButtonStyle(style tcell.Style) *FormScrollable {
	f.buttonStyle = style
	return f
}

// SetButtonActivatedStyle sets the style of the buttons when they are focused.
func (f *FormScrollable) SetButtonActivatedStyle(style tcell.Style) *FormScrollable {
	f.buttonActivatedStyle = style
	return f
}

// SetButtonDisabledStyle sets the style of the buttons when they are disabled.
func (f *FormScrollable) SetButtonDisabledStyle(style tcell.Style) *FormScrollable {
	f.buttonDisabledStyle = style
	return f
}

// SetFocus shifts the focus to the form element with the given index, counting
// non-button items first and buttons last. Note that this index is only used
// when the form itself receives focus.
func (f *FormScrollable) SetFocus(index int) *FormScrollable {
	var current, future int
	for itemIndex, item := range f.items {
		if itemIndex == index {
			future = itemIndex
		}
		if item.HasFocus() {
			current = itemIndex
		}
	}
	for buttonIndex, button := range f.buttons {
		if buttonIndex+len(f.items) == index {
			future = buttonIndex + len(f.items)
		}
		if button.HasFocus() {
			current = buttonIndex + len(f.items)
		}
	}
	var focus func(p Primitive)
	focus = func(p Primitive) {
		p.Focus(focus)
	}
	if current != future {
		if current >= 0 && current < len(f.items) {
			f.items[current].Blur()
		} else if current >= len(f.items) && current < len(f.items)+len(f.buttons) {
			f.buttons[current-len(f.items)].Blur()
		}
		if future >= 0 && future < len(f.items) {
			focus(f.items[future])
		} else if future >= len(f.items) && future < len(f.items)+len(f.buttons) {
			focus(f.buttons[future-len(f.items)])
		}
	}
	f.focusedElement = future
	return f
}

// AddTextArea adds a text area to the form. It has a label, an optional initial
// text, a size (width and height) referring to the actual input area (a
// fieldWidth of 0 extends it as far right as possible, a fieldHeight of 0 will
// cause it to be [DefaultFormFieldHeight]), and a maximum number of bytes of
// text allowed (0 means no limit).
//
// The optional callback function is invoked when the content of the text area
// has changed. Note that especially for larger texts, this is an expensive
// operation due to technical constraints of the [TextArea] primitive (every key
// stroke leads to a new reallocation of the entire text).
func (f *FormScrollable) AddTextArea(label, text string, fieldWidth, fieldHeight, maxLength int, changed func(text string)) *FormScrollable {
	if fieldHeight == 0 {
		fieldHeight = DefaultFormFieldHeight
	}
	textArea := NewTextArea().
		SetLabel(label).
		SetSize(fieldHeight, fieldWidth).
		SetMaxLength(maxLength)
	if text != "" {
		textArea.SetText(text, true)
	}
	if changed != nil {
		textArea.SetChangedFunc(func() {
			changed(textArea.GetText())
		})
	}
	f.items = append(f.items, textArea)
	return f
}

// AddTextView adds a text view to the form. It has a label and text, a size
// (width and height) referring to the actual text element (a fieldWidth of 0
// extends it as far right as possible, a fieldHeight of 0 will cause it to be
// [DefaultFormFieldHeight]), a flag to turn on/off dynamic colors, and a flag
// to turn on/off scrolling. If scrolling is turned off, the text view will not
// receive focus.
func (f *FormScrollable) AddTextView(label, text string, fieldWidth, fieldHeight int, dynamicColors, scrollable bool) *FormScrollable {
	if fieldHeight == 0 {
		fieldHeight = DefaultFormFieldHeight
	}
	textArea := NewTextView().
		SetLabel(label).
		SetSize(fieldHeight, fieldWidth).
		SetDynamicColors(dynamicColors).
		SetScrollable(scrollable).
		SetText(text)
	f.items = append(f.items, textArea)
	return f
}

// AddInputField adds an input field to the form. It has a label, an optional
// initial value, a field width (a value of 0 extends it as far as possible),
// an optional accept function to validate the item's value (set to nil to
// accept any text), and an (optional) callback function which is invoked when
// the input field's text has changed.
func (f *FormScrollable) AddInputField(label, value string, fieldWidth int, accept func(textToCheck string, lastChar rune) bool, changed func(text string)) *FormScrollable {
	f.items = append(f.items, NewInputField().
		SetLabel(label).
		SetText(value).
		SetFieldWidth(fieldWidth).
		SetAcceptanceFunc(accept).
		SetChangedFunc(changed))
	return f
}

// AddPasswordField adds a password field to the form. This is similar to an
// input field except that the user's input not shown. Instead, a "mask"
// character is displayed. The password field has a label, an optional initial
// value, a field width (a value of 0 extends it as far as possible), and an
// (optional) callback function which is invoked when the input field's text has
// changed.
func (f *FormScrollable) AddPasswordField(label, value string, fieldWidth int, mask rune, changed func(text string)) *FormScrollable {
	if mask == 0 {
		mask = '*'
	}
	f.items = append(f.items, NewInputField().
		SetLabel(label).
		SetText(value).
		SetFieldWidth(fieldWidth).
		SetMaskCharacter(mask).
		SetChangedFunc(changed))
	return f
}

// AddDropDown adds a drop-down element to the form. It has a label, options,
// and an (optional) callback function which is invoked when an option was
// selected. The initial option may be a negative value to indicate that no
// option is currently selected.
func (f *FormScrollable) AddDropDown(label string, options []string, initialOption int, selected func(option string, optionIndex int)) *FormScrollable {
	f.items = append(f.items, NewDropDown().
		SetLabel(label).
		SetOptions(options, selected).
		SetCurrentOption(initialOption))
	return f
}

// AddCheckbox adds a checkbox to the form. It has a label, an initial state,
// and an (optional) callback function which is invoked when the state of the
// checkbox was changed by the user.
func (f *FormScrollable) AddCheckbox(label string, checked bool, changed func(checked bool)) *FormScrollable {
	f.items = append(f.items, NewCheckbox().
		SetLabel(label).
		SetChecked(checked).
		SetChangedFunc(changed))
	return f
}

// AddImage adds an image to the form. It has a label and the image will fit in
// the specified width and height (its aspect ratio is preserved). See
// [Image.SetColors] for a description of the "colors" parameter. Images are not
// interactive and are skipped over in a form. The "width" value may be 0
// (adjust dynamically) but "height" should generally be a positive value.
func (f *FormScrollable) AddImage(label string, image image.Image, width, height, colors int) *FormScrollable {
	f.items = append(f.items, NewImage().
		SetLabel(label).
		SetImage(image).
		SetSize(height, width).
		SetAlign(AlignTop, AlignLeft).
		SetColors(colors))
	return f
}

// AddButton adds a new button to the form. The "selected" function is called
// when the user selects this button. It may be nil.
func (f *FormScrollable) AddButton(label string, selected func()) *FormScrollable {
	f.buttons = append(f.buttons, NewButton(label).SetSelectedFunc(selected))
	return f
}

// GetButton returns the button at the specified 0-based index. Note that
// buttons have been specially prepared for this form and modifying some of
// their attributes may have unintended side effects.
func (f *FormScrollable) GetButton(index int) *Button {
	return f.buttons[index]
}

// RemoveButton removes the button at the specified position, starting with 0
// for the button that was added first.
func (f *FormScrollable) RemoveButton(index int) *FormScrollable {
	f.buttons = append(f.buttons[:index], f.buttons[index+1:]...)
	return f
}

// GetButtonCount returns the number of buttons in this form.
func (f *FormScrollable) GetButtonCount() int {
	return len(f.buttons)
}

// GetButtonIndex returns the index of the button with the given label, starting
// with 0 for the button that was added first. If no such label was found, -1
// is returned.
func (f *FormScrollable) GetButtonIndex(label string) int {
	for index, button := range f.buttons {
		if button.GetLabel() == label {
			return index
		}
	}
	return -1
}

// Clear removes all input elements from the form, including the buttons if
// specified.
func (f *FormScrollable) Clear(includeButtons bool) *FormScrollable {
	f.items = nil
	if includeButtons {
		f.ClearButtons()
	}
	f.focusedElement = 0
	return f
}

// ClearButtons removes all buttons from the form.
func (f *FormScrollable) ClearButtons() *FormScrollable {
	f.buttons = nil
	return f
}

// AddFormItem adds a new item to the form. This can be used to add your own
// objects to the form. Note, however, that the Form class will override some
// of its attributes to make it work in the form context. Specifically, these
// are:
//
//   - The label width
//   - The label color
//   - The background color
//   - The field text color
//   - The field background color
func (f *FormScrollable) AddFormItem(item FormItem) *FormScrollable {
	f.items = append(f.items, item)
	return f
}

// GetFormItemCount returns the number of items in the form (not including the
// buttons).
func (f *FormScrollable) GetFormItemCount() int {
	return len(f.items)
}

// GetFormItem returns the form item at the given position, starting with index
// 0. Elements are referenced in the order they were added. Buttons are not
// included.
func (f *FormScrollable) GetFormItem(index int) FormItem {
	return f.items[index]
}

// RemoveFormItem removes the form element at the given position, starting with
// index 0. Elements are referenced in the order they were added. Buttons are
// not included.
func (f *FormScrollable) RemoveFormItem(index int) *FormScrollable {
	f.items = append(f.items[:index], f.items[index+1:]...)
	return f
}

// GetFormItemByLabel returns the first form element with the given label. If
// no such element is found, nil is returned. Buttons are not searched and will
// therefore not be returned.
func (f *FormScrollable) GetFormItemByLabel(label string) FormItem {
	for _, item := range f.items {
		if item.GetLabel() == label {
			return item
		}
	}
	return nil
}

// GetFormItemIndex returns the index of the first form element with the given
// label. If no such element is found, -1 is returned. Buttons are not searched
// and will therefore not be returned.
func (f *FormScrollable) GetFormItemIndex(label string) int {
	for index, item := range f.items {
		if item.GetLabel() == label {
			return index
		}
	}
	return -1
}

// GetFocusedItemIndex returns the indices of the form element or button which
// currently has focus. If they don't, -1 is returned respectively.
func (f *FormScrollable) GetFocusedItemIndex() (formItem, button int) {
	index := f.focusIndex()
	if index < 0 {
		return -1, -1
	}
	if index < len(f.items) {
		return index, -1
	}
	return -1, index - len(f.items)
}

// SetCancelFunc sets a handler which is called when the user hits the Escape
// key.
func (f *FormScrollable) SetCancelFunc(callback func()) *FormScrollable {
	f.cancel = callback
	return f
}

// Draw draws this primitive onto the screen.
func (f *FormScrollable) Draw(screen tcell.Screen) {
	f.Box.DrawForSubclass(screen, f)

	// Determine the actual item that has focus.
	if index := f.focusIndex(); index >= 0 {
		f.focusedElement = index
	}

	// Determine the dimensions.
	x, y, width, height := f.GetInnerRect()
	topLimit := y
	bottomLimit := y + height
	rightLimit := x + width
	startX := x

	// Find the longest label.
	var maxLabelWidth int
	for _, item := range f.items {
		labelWidth := TaggedStringWidth(item.GetLabel())
		if labelWidth > maxLabelWidth {
			maxLabelWidth = labelWidth
		}
	}
	maxLabelWidth++ // Add one space.

	// Calculate positions of form items.
	type position struct{ x, y, width, height int }
	positions := make([]position, len(f.items)+len(f.buttons))
	var (
		focusedPosition position
		lineHeight      = 1
	)
	for index, item := range f.items {
		// Calculate the space needed.
		labelWidth := TaggedStringWidth(item.GetLabel())
		var itemWidth int
		if f.horizontal {
			fieldWidth := item.GetFieldWidth()
			if fieldWidth <= 0 {
				fieldWidth = DefaultFormFieldWidth
			}
			labelWidth++
			itemWidth = labelWidth + fieldWidth
		} else {
			// We want all fields to align vertically.
			labelWidth = maxLabelWidth
			itemWidth = width
		}
		itemHeight := item.GetFieldHeight()
		if itemHeight <= 0 {
			itemHeight = DefaultFormFieldHeight
		}

		// Advance to next line if there is no space.
		if f.horizontal && x+labelWidth+1 >= rightLimit {
			x = startX
			y += lineHeight + 1
			lineHeight = itemHeight
		}

		// Update line height.
		if itemHeight > lineHeight {
			lineHeight = itemHeight
		}

		// Adjust the item's attributes.
		if x+itemWidth >= rightLimit {
			itemWidth = rightLimit - x
		}
		item.SetFormAttributes(
			labelWidth,
			f.labelColor,
			f.GetBackgroundColor(),
			f.fieldTextColor,
			f.fieldBackgroundColor,
		)

		// Save position.
		positions[index].x = x
		positions[index].y = y
		positions[index].width = itemWidth
		positions[index].height = itemHeight
		if item.HasFocus() {
			focusedPosition = positions[index]
		}

		// Advance to next item.
		if f.horizontal {
			x += itemWidth + f.itemPadding
		} else {
			y += itemHeight + f.itemPadding
		}
	}

	// How wide are the buttons?
	buttonWidths := make([]int, len(f.buttons))
	buttonsWidth := 0
	for index, button := range f.buttons {
		w := TaggedStringWidth(button.GetLabel()) + 4
		buttonWidths[index] = w
		buttonsWidth += w + 1
	}
	buttonsWidth--

	// Where do we place them?
	if !f.horizontal && x+buttonsWidth < rightLimit {
		if f.buttonsAlign == AlignRight {
			x = rightLimit - buttonsWidth
		} else if f.buttonsAlign == AlignCenter {
			x = (x + rightLimit - buttonsWidth) / 2
		}

		// In vertical layouts, buttons always appear after an empty line.
		if f.itemPadding == 0 {
			y++
		}
	}

	// Calculate positions of buttons.
	for index, button := range f.buttons {
		space := rightLimit - x
		buttonWidth := buttonWidths[index]
		if f.horizontal {
			if space < buttonWidth-4 {
				x = startX
				y += lineHeight + 1
				space = width
				lineHeight = 1
			}
		} else {
			if space < 1 {
				break // No space for this button anymore.
			}
		}
		if buttonWidth > space {
			buttonWidth = space
		}
		button.SetStyle(f.buttonStyle).
			SetActivatedStyle(f.buttonActivatedStyle).
			SetDisabledStyle(f.buttonDisabledStyle)

		buttonIndex := index + len(f.items)
		positions[buttonIndex].x = x
		positions[buttonIndex].y = y
		positions[buttonIndex].width = buttonWidth
		positions[buttonIndex].height = 1

		if button.HasFocus() {
			focusedPosition = positions[buttonIndex]
		}

		x += buttonWidth + 1
	}

	// Determine vertical offset based on the position of the focused item.
	var offset int
	if focusedPosition.y+focusedPosition.height > bottomLimit {
		offset = focusedPosition.y + focusedPosition.height - bottomLimit
		if focusedPosition.y-offset < topLimit {
			offset = focusedPosition.y - topLimit
		}
	}

	// Draw items.
	for index, item := range f.items {
		// Set position.
		y := positions[index].y - offset
		height := positions[index].height
		item.SetRect(positions[index].x, y, positions[index].width, height)

		// Is this item visible?
		if y+height <= topLimit || y >= bottomLimit {
			continue
		}

		// Draw items with focus last (in case of overlaps).
		if item.HasFocus() {
			defer item.Draw(screen)
		} else {
			item.Draw(screen)
		}
	}

	// Draw buttons.
	for index, button := range f.buttons {
		// Set position.
		buttonIndex := index + len(f.items)
		y := positions[buttonIndex].y - offset
		height := positions[buttonIndex].height
		button.SetRect(positions[buttonIndex].x, y, positions[buttonIndex].width, height)

		// Is this button visible?
		if y+height <= topLimit || y >= bottomLimit {
			continue
		}

		// Draw button.
		button.Draw(screen)
	}

	const scrollBtnWidth = 1
	const scrollBtnHeight = 1

	_, _, ww, hh := f.GetRect()

	f.upScrollButton.SetRect(ww-scrollBtnWidth, 0, scrollBtnWidth, scrollBtnHeight)
	f.upScrollButton.Draw(screen)

	f.downScrollButton.SetRect(ww-scrollBtnWidth, hh-1, scrollBtnWidth, scrollBtnHeight)
	f.downScrollButton.Draw(screen)
}

// Focus is called by the application when the primitive receives focus.
func (f *FormScrollable) Focus(delegate func(p Primitive)) {
	// Hand on the focus to one of our child elements.
	if f.focusedElement < 0 || f.focusedElement >= len(f.items)+len(f.buttons) {
		f.focusedElement = 0
	}
	var handler func(key tcell.Key)
	handler = func(key tcell.Key) {
		if key >= 0 {
			f.lastFinishedKey = key
		}
		switch key {
		case tcell.KeyTab, tcell.KeyEnter:
			f.focusedElement++
			f.Focus(delegate)
		case tcell.KeyBacktab:
			f.focusedElement--
			if f.focusedElement < 0 {
				f.focusedElement = len(f.items) + len(f.buttons) - 1
			}
			f.Focus(delegate)
		case tcell.KeyEscape:
			if f.cancel != nil {
				f.cancel()
			} else {
				f.focusedElement = 0
				f.Focus(delegate)
			}
		default:
			if key < 0 && f.lastFinishedKey >= 0 {
				// Repeat the last action.
				handler(f.lastFinishedKey)
			}
		}
	}

	// Track whether a form item has focus.
	var itemFocused bool

	// Set the handler and focus for all items and buttons.
	for index, button := range f.buttons {
		button.SetExitFunc(handler)
		if f.focusedElement == index+len(f.items) {
			if button.IsDisabled() {
				f.focusedElement++
				if f.focusedElement >= len(f.items)+len(f.buttons) {
					f.focusedElement = 0
				}
				continue
			}

			itemFocused = true
			func(b *Button) { // Wrapping might not be necessary anymore in future Go versions.
				defer delegate(b)
			}(button)
		}
	}
	for index, item := range f.items {
		item.SetFinishedFunc(handler)
		if f.focusedElement == index {
			itemFocused = true
			func(i FormItem) { // Wrapping might not be necessary anymore in future Go versions.
				defer delegate(i)
			}(item)
		}
	}

	// If no item was focused, focus the form itself.
	if !itemFocused {
		f.Box.Focus(delegate)
	}
}

// HasFocus returns whether or not this primitive has focus.
func (f *FormScrollable) HasFocus() bool {
	if f.focusIndex() >= 0 {
		return true
	}
	return f.Box.HasFocus()
}

// focusIndex returns the index of the currently focused item, counting form
// items first, then buttons. A negative value indicates that no containeed item
// has focus.
func (f *FormScrollable) focusIndex() int {
	for index, item := range f.items {
		if item.HasFocus() {
			return index
		}
	}
	for index, button := range f.buttons {
		if button.HasFocus() {
			return len(f.items) + index
		}
	}
	return -1
}

// MouseHandler returns the mouse handler for this primitive.
func (f *FormScrollable) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
	return f.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
		// At the end, update f.focusedElement and prepare current item/button.
		defer func() {
			if consumed {
				index := f.focusIndex()
				if index >= 0 {
					f.focusedElement = index
				}
			}
		}()

		// Determine items to pass mouse events to.
		for _, item := range f.items {
			// Exclude TextView items from mouse-down events as they are
			// read-only items and thus should not be focused.
			if _, ok := item.(*TextView); ok && action == MouseLeftDown {
				continue
			}

			consumed, capture = item.MouseHandler()(action, event, setFocus)
			if consumed {
				return
			}
		}
		for _, button := range f.buttons {
			consumed, capture = button.MouseHandler()(action, event, setFocus)
			if consumed {
				return
			}
		}

		consumed, capture = f.upScrollButton.MouseHandler()(action, event, setFocus)
		if consumed {
			return
		}

		consumed, capture = f.downScrollButton.MouseHandler()(action, event, setFocus)
		if consumed {
			return
		}

		// A mouse down anywhere else will return the focus to the last selected
		// element.
		if action == MouseLeftDown && f.InRect(event.Position()) {
			f.Focus(setFocus)
			consumed = true
		}

		return
	})
}

// InputHandler returns the handler for this primitive.
func (f *FormScrollable) InputHandler() func(event *tcell.EventKey, setFocus func(p Primitive)) {
	return f.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Primitive)) {
		for _, item := range f.items {
			if item != nil && item.HasFocus() {
				if handler := item.InputHandler(); handler != nil {
					handler(event, setFocus)
					return
				}
			}
		}

		for _, button := range f.buttons {
			if button.HasFocus() {
				if handler := button.InputHandler(); handler != nil {
					handler(event, setFocus)
					return
				}
			}
		}
	})
}

// PasteHandler returns the handler for this primitive.
func (f *FormScrollable) PasteHandler() func(pastedText string, setFocus func(p Primitive)) {
	return f.WrapPasteHandler(func(pastedText string, setFocus func(p Primitive)) {
		for _, item := range f.items {
			if item != nil && item.HasFocus() {
				if handler := item.PasteHandler(); handler != nil {
					handler(pastedText, setFocus)
					return
				}
			}
		}

		for _, button := range f.buttons {
			if button.HasFocus() {
				if handler := button.PasteHandler(); handler != nil {
					handler(pastedText, setFocus)
					return
				}
			}
		}
	})
}
