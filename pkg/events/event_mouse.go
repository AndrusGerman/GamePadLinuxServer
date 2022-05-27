package events

import (
	"game_pad_linux_server/pkg/utils"
	"math"
	"time"

	"github.com/bendahl/uinput"
)

// maximum mouse speed
var pxMovementMax float64 = 12

func (ctx *Events) UsedMouse(mouse uinput.Mouse) {
	switch ctx.Mode {
	case 3:
		pxEvent = nil
	case 1:
		var x = pxMovementMax * float64(ctx.ValueX)
		var y = pxMovementMax * float64(ctx.ValueY)

		// Is Negative
		var xNegative = x < 0
		var yNegative = y < 0

		var xPercent = (x / pxMovementMax)
		var yPercent = (y / pxMovementMax)

		// set min speed percent 30%
		// the minimum speed cannot be less than 30 percent of the initial value
		if utils.GetPositive(xPercent) < 0.3 {
			xPercent = utils.ReturnValueInSRC(0.3, xPercent)
		}

		if utils.GetPositive(yPercent) < 0.3 {
			yPercent = utils.ReturnValueInSRC(0.3, yPercent)
		}

		// Percent speed
		x = (x * xPercent)
		y = (y * yPercent)

		// fixed Negative
		if x > 0 && xNegative {
			x = x * -1
		}
		if y > 0 && yNegative {
			y = y * -1
		}

		mouse.Move(
			int32(math.Round(x)),
			int32(math.Round(y)),
		)
	}
}

// Need refactor
var pxEvent *Events
var move = true
var universalMouse uinput.Mouse

// reset move mouse
func clearMove() {
	go func() {
		time.Sleep(6 * time.Millisecond)
		move = true
	}()
}

func init() {
	go func() {
		for {
			if universalMouse != nil {
				if pxEvent != nil {
					pxEvent.UsedMouse(universalMouse)
				}
			}
			time.Sleep(16 * time.Millisecond)
		}
	}()
}

func (ctx *Events) ManagerMouse(mouse uinput.Mouse) {
	switch ctx.Mode {
	case 3:
		pxEvent = nil
	case 1:
		if move {
			move = false
			clearMove()
		}
		universalMouse = mouse
		pxEvent = ctx
	}
}
