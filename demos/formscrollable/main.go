// Demo code for the Form primitive.
package main

import (
	form "github.com/name212/tview-widgets"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	form := form.NewFormScrollable().
		AddInputField("First name", "", 20, nil, nil).
		AddInputField("Last name", "", 20, nil, nil).
		AddTextArea("Address", "", 40, 0, 0, nil).
		AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
		AddTextView("Notes", "This is just a demo.\nYou can enter whatever you wish.", 40, 2, true, false).
		AddCheckbox("Age 18+", false, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}
