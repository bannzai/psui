package ui

// USER               PID  %CPU %MEM      VSZ    RSS   TT  STAT STARTED      TIME COMMAND
// bannzai     37493   0.0  0.0  4280196   1588   ??  Ss   11:41PM   0:00.00 /Users/hiroseyuudai/.vim/plugged/LanguageClient-neovim/bin/languageclient
type process struct {
	user    string // USER
	pid     string // PID
	cpu     string // %CPU
	mem     string // %MEM
	vsz     string // VSZ
	rss     string // RSS
	tt      string // TT
	stat    string // STAT
	started string // STARTED
	time    string // TIME
	command string // COMMAND
}

type processTable struct {
	*tView.Table
}

func NewProcessTable() processTable {
	table := processTable{}
}
