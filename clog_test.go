package clog

import "testing"

func TestClog(t *testing.T) {
    logger := NewConsoleLogger("ConsoleLogger: ")
    logger.Printf("This is normal.")
    logger.Success("This is a success!")
    logger.Error("[ERROR] you should check it out: %d", -1)
    logger.Printf("This should be normal, again.")
}
