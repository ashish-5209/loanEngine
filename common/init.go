/*
 * clubs all the helpers that will be used in entire project.
 */
package common

import "loanEngine/config"

func init() {
	Logger = NewSugaredLogger(config.Environment, "/var/log/loanEngine/error.log")
	CriticalLogger = NewCriticalLogger(config.Environment, "/var/log/loanEngine/critical.log") // for storing critical logs with stack trace
}
