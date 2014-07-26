package colorPrinter

import (
	"fmt"
	color "github.com/daviddengcn/go-colortext"
)

var (
	ErrorStyle = color.Red
	// InfoStyle    = color.Blue
	DebugStyle   = color.Blue
	WarningStyle = color.Yellow
	SuccessStyle = color.Green
)

func PrintfError(fomat string, params ...interface{}) {
	PrintfColor(ErrorStyle, fomat, params...)
}

// func PrintfInfo(fomat string, params ...interface{}) {
// 	PrintfColor(InfoStyle, fomat, params...)
// }

func PrintfDebug(fomat string, params ...interface{}) {
	PrintfColor(DebugStyle, fomat, params...)
}

func PrintfWarning(fomat string, params ...interface{}) {
	PrintfColor(WarningStyle, fomat, params...)
}

func PrintfSuccess(fomat string, params ...interface{}) {
	PrintfColor(SuccessStyle, fomat, params...)
}

func PrintfColor(style color.Color, fomat string, params ...interface{}) {
	color.ChangeColor(style, true, color.None, false)
	fmt.Printf(fomat, params...)
	color.ResetColor()
}
