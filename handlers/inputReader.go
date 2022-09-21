package handlers

import (
	"fmt"

	"github.com/dilungasr/reuse/models"
	"github.com/dilungasr/reuse/utils"
)

func ReadFromInput() {
	// extensions
	fmt.Print("Extensions (use spaces to separate): ")
	extsInputs := utils.SpacedStringToSlice(utils.ScanString())

	// ignore
	fmt.Print("Files/folders to ignore (use spaces to separate): ")
	ignore := utils.SpacedStringToSlice(utils.ScanString())

	// rep
	fmt.Println("Define your replacements (in old: new pair per line)")
	fmt.Println("Enter q to finish:")
	changesInputs := utils.ScanStringPerLine()
	//del
	fmt.Print("Files/folders to delete (use spaces to separate): ")
	del := utils.SpacedStringToSlice(utils.ScanString())

	// run
	fmt.Println("Commands to run (sequentially - line by line)")
	fmt.Println("Enter q to finish:")
	run := utils.ScanStringPerLine()

	models.ProjectChanges.AddExts(extsInputs)
	models.ProjectChanges.AddIgnore(ignore)
	models.ProjectChanges.AddRepsFromInput(changesInputs)
	models.ProjectChanges.AddRun(run)
	models.ProjectChanges.AddDel(del)
}
