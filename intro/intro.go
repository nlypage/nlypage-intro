package intro

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"time"
)

const delay = time.Millisecond * 700

func Start() {
	area, _ := pterm.DefaultArea.WithCenter().Start()
	plum := pterm.NewRGB(221, 160, 221)
	logo, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("N", pterm.NewStyle(pterm.FgLightMagenta)),
		putils.LettersFromStringWithRGB("LYPAGE", plum)).Srender()
	github := plum.Sprint("github.com/")
	logoUnderText := pterm.Sprintf("Starting %s project", pterm.LightMagenta("Awesome"))

	text := pterm.Sprintf("%s\n%s\n%s", github, logo, logoUnderText)

	withoutAnyPages := plum.Sprint("Without any pages")
	area.Update(pterm.DefaultBox.
		WithTopPadding(1).
		WithLeftPadding(5).
		WithRightPadding(5).
		WithBottomPadding(1).
		WithHorizontalString("═").
		WithVerticalString("║").
		WithBottomLeftCornerString("╗").
		WithBottomRightCornerString("╔").
		WithTopLeftCornerString("╝").
		WithTopRightCornerString("╚").
		WithBoxStyle(pterm.NewStyle(pterm.FgLightMagenta)).
		Sprint(text))
	time.Sleep(delay)
	area.Update(pterm.DefaultBox.
		WithTopPadding(1).
		WithLeftPadding(5).
		WithRightPadding(5).
		WithBottomPadding(1).
		WithHorizontalString("═").
		WithVerticalString("║").
		WithBottomLeftCornerString("╗").
		WithBottomRightCornerString("╔").
		WithTopLeftCornerString("╝").
		WithTopRightCornerString("╚").
		WithBoxStyle(pterm.NewStyle(pterm.FgLightMagenta)).
		WithTitle(withoutAnyPages).
		WithTitleTopCenter().
		Sprint(text))
	time.Sleep(delay)
	area.Update(pterm.DefaultBox.
		WithTopPadding(1).
		WithLeftPadding(5).
		WithRightPadding(5).
		WithBottomPadding(1).
		WithHorizontalString("═").
		WithVerticalString("║").
		WithBottomLeftCornerString("╗").
		WithBottomRightCornerString("╔").
		WithTopLeftCornerString("╝").
		WithTopRightCornerString("╚").
		WithBoxStyle(pterm.NewStyle(pterm.FgLightMagenta)).
		WithTitle(withoutAnyPages).
		WithTitleTopRight().
		Sprint(text))
	time.Sleep(delay)
	area.Update(pterm.DefaultBox.
		WithTopPadding(1).
		WithLeftPadding(5).
		WithRightPadding(5).
		WithBottomPadding(1).
		WithHorizontalString("═").
		WithVerticalString("║").
		WithBottomLeftCornerString("╗").
		WithBottomRightCornerString("╔").
		WithTopLeftCornerString("╝").
		WithTopRightCornerString("╚").
		WithBoxStyle(pterm.NewStyle(pterm.FgLightMagenta)).
		WithTitle(withoutAnyPages).
		WithTitleBottomRight().
		Sprint(text))
	time.Sleep(delay)
	area.Update(pterm.DefaultBox.
		WithTopPadding(1).
		WithLeftPadding(5).
		WithRightPadding(5).
		WithBottomPadding(1).
		WithHorizontalString("═").
		WithVerticalString("║").
		WithBottomLeftCornerString("╗").
		WithBottomRightCornerString("╔").
		WithTopLeftCornerString("╝").
		WithTopRightCornerString("╚").
		WithBoxStyle(pterm.NewStyle(pterm.FgLightMagenta)).
		WithTitle(withoutAnyPages).
		WithTitleBottomCenter().
		Sprint(text))
	time.Sleep(delay)
	area.Update(pterm.DefaultBox.
		WithTopPadding(1).
		WithLeftPadding(5).
		WithRightPadding(5).
		WithBottomPadding(1).
		WithHorizontalString("═").
		WithVerticalString("║").
		WithBottomLeftCornerString("╗").
		WithBottomRightCornerString("╔").
		WithTopLeftCornerString("╝").
		WithTopRightCornerString("╚").
		WithBoxStyle(pterm.NewStyle(pterm.FgLightMagenta)).
		WithTitle(withoutAnyPages).
		WithTitleBottomLeft().
		Sprint(text))
	time.Sleep(delay)
	area.Update(pterm.DefaultBox.
		WithTopPadding(1).
		WithLeftPadding(5).
		WithRightPadding(5).
		WithBottomPadding(1).
		WithHorizontalString("═").
		WithVerticalString("║").
		WithBottomLeftCornerString("╗").
		WithBottomRightCornerString("╔").
		WithTopLeftCornerString("╝").
		WithTopRightCornerString("╚").
		WithBoxStyle(pterm.NewStyle(pterm.FgLightMagenta)).
		WithTitle(withoutAnyPages).
		WithTitleTopLeft().
		Sprint(text))
	time.Sleep(delay)
	area.Update(pterm.DefaultBox.
		WithTopPadding(1).
		WithLeftPadding(5).
		WithRightPadding(5).
		WithBottomPadding(1).
		WithHorizontalString("═").
		WithVerticalString("║").
		WithBottomLeftCornerString("╗").
		WithBottomRightCornerString("╔").
		WithTopLeftCornerString("╝").
		WithTopRightCornerString("╚").
		WithBoxStyle(pterm.NewStyle(pterm.FgLightMagenta)).
		WithTitle(withoutAnyPages).
		WithTitleTopCenter().
		Sprint(text))
	time.Sleep(delay)
	area.Clear()
}
