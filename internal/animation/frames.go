package animation

// Dancing figure animation frames
var DancingFrames = []string{
	`
    ╭───────────────────╮
    │   🎉 Welcome! 🎉   │
    ╰───────────────────╯
          \o/
           |
          / \
`,
	`
    ╭───────────────────╮
    │   🎉 Welcome! 🎉   │
    ╰───────────────────╯
           o/
          /|
          / \
`,
	`
    ╭───────────────────╮
    │   🎉 Welcome! 🎉   │
    ╰───────────────────╯
          \o
          /|
           \
`,
	`
    ╭───────────────────╮
    │   🎉 Welcome! 🎉   │
    ╰───────────────────╯
           o/
          /|
          / \
`,
}

// Waving hand animation frames
var WavingFrames = []string{
	`
    ╭───────────────────╮
    │   👋 Hello! 👋    │
    ╰───────────────────╯
          🙌
    `,
	`
    ╭───────────────────╮
    │   👋 Hello! 👋    │
    ╰───────────────────╯
          ✋
    `,
	`
    ╭───────────────────╮
    │   👋 Hello! 👋    │
    ╰───────────────────╯
          🖐️
    `,
}

// Simple robot animation frames
var RobotFrames = []string{
	`
    ╭───────────────────────╮
    │   🤖 HELLOGANG 🤖    │
    ╰───────────────────────╯
        ╔═══════════╗
        ║  ●     ●  ║
        ║     ▲     ║
        ║  ╰───────╯ ║
        ╚═══════════╝
           │  │  │
    `,
	`
    ╭───────────────────────╮
    │   🤖 HELLOGANG 🤖    │
    ╰───────────────────────╯
        ╔═══════════╗
        ║  ○     ○  ║
        ║     ◆     ║
        ║  ╰───────╯ ║
        ╚═══════════╝
           │  │  │
    `,
	`
    ╭───────────────────────╮
    │   🤖 HELLOGANG 🤖    │
    ╰───────────────────────╯
        ╔═══════════╗
        ║  ●     ●  ║
        ║     ▼     ║
        ║  ╰───────╯ ║
        ╚═══════════╝
           │  │  │
    `,
}

// Simple loading spinner characters
var SpinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// GetAnimationFrames returns frames for a named animation
func GetAnimationFrames(name string) []string {
	switch name {
	case "dancing":
		return DancingFrames
	case "waving":
		return WavingFrames
	case "robot":
		return RobotFrames
	default:
		return DancingFrames
	}
}

// GetSpinnerFrame returns a spinner frame at the given index
func GetSpinnerFrame(index int) string {
	return SpinnerFrames[index%len(SpinnerFrames)]
}