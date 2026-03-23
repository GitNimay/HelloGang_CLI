package animation

import (
	"time"
)

// Player controls animation frame cycling
type Player struct {
	frames     []string
	currentIdx int
	interval   time.Duration
	ticker     *time.Ticker
}

// NewPlayer creates a new animation player
func NewPlayer(frames []string, interval time.Duration) *Player {
	return &Player{
		frames:     frames,
		currentIdx: 0,
		interval:   interval,
	}
}

// Next advances to the next frame and returns it
func (p *Player) Next() string {
	if len(p.frames) == 0 {
		return ""
	}
	frame := p.frames[p.currentIdx]
	p.currentIdx = (p.currentIdx + 1) % len(p.frames)
	return frame
}

// Current returns the current frame without advancing
func (p *Player) Current() string {
	if len(p.frames) == 0 {
		return ""
	}
	return p.frames[p.currentIdx]
}

// Reset resets the animation to the first frame
func (p *Player) Reset() {
	p.currentIdx = 0
}

// FrameCount returns the number of frames
func (p *Player) FrameCount() int {
	return len(p.frames)
}

// SetInterval changes the animation interval
func (p *Player) SetInterval(d time.Duration) {
	p.interval = d
}

// Interval returns the current interval
func (p *Player) Interval() time.Duration {
	return p.interval
}