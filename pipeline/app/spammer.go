package main

func RunPipeline(cmds ...cmd) {
	// Implement me!
}

// in - string
// out - User
func SelectUsers(in, out chan any) {
	// Implement me!
}

// in - User
// out - MsgID
func SelectMessages(in, out chan any) {
	// Implement me!
}

// in - MsgID
// out - MsgData
func CheckSpam(in, out chan any) {
	// Implement me!
}

// in - MsgData
// out - string
func CombineResults(in, out chan any) {
	// Implement me!
}
