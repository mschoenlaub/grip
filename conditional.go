package grip

import (
	"os"

	"github.com/coreos/go-systemd/journal"
)

// Conditional logging methods, which take two arguments, a boolean,
// and a message argument. Messages can be strings, Objects that
// implement the MessageComposter interface or errors. If the
// threshold level is met, and the message to log is not an empty
// string, then it logs the resolved message.

func (self *Journaler) conditionalSend(priority journal.Priority, conditional bool, message interface{}) {
	if !conditional || priority > self.thresholdLevel {
		return
	}

	self.genericSend(priority, message)
	return

}

func (self *Journaler) conditionalSendPanic(priority journal.Priority, conditional bool, message interface{}) {
	if !conditional || priority > self.thresholdLevel {
		return
	}

	msg := self.genericSend(priority, message)
	panic(msg)
}

func (self *Journaler) conditionalSendFatal(priority journal.Priority, conditional bool, message interface{}) {
	if !conditional || priority > self.thresholdLevel {
		return
	}

	self.genericSend(priority, message)
	os.Exit(1)
}

func (self *Journaler) DefaultWhen(conditional bool, message interface{}) {
	self.conditionalSend(self.defaultLevel, conditional, message)
}
func DefaultWhen(conditional bool, message interface{}) {
	std.DefaultWhen(conditional, message)
}
func (self *Journaler) DefaultWhenln(conditional bool, msg ...interface{}) {
	self.DefaultWhen(conditional, NewDefaultMessage(msg))
}
func DefaultWhenln(conditional bool, msg ...interface{}) {
	std.DefaultWhenln(conditional, msg...)
}
func (self *Journaler) DefaultWhenf(conditional bool, msg string, args ...interface{}) {
	self.DefaultWhen(conditional, NewFormatedMessage(msg, args))
}
func DefaultWhenf(conditional bool, msg string, args ...interface{}) {
	std.DefaultWhenf(conditional, msg, args...)
}

func (self *Journaler) EmergencyWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriEmerg, conditional, message)
}
func EmergencyWhen(conditional bool, message interface{}) {
	std.EmergencyWhen(conditional, message)
}
func (self *Journaler) EmergencyWhenln(conditional bool, msg ...interface{}) {
	self.EmergencyWhen(conditional, NewDefaultMessage(msg))
}
func EmergencyWhenln(conditional bool, msg ...interface{}) {
	std.EmergencyWhenln(conditional, msg...)
}
func (self *Journaler) EmergencyWhenf(conditional bool, msg string, args ...interface{}) {
	self.EmergencyWhen(conditional, NewFormatedMessage(msg, args))
}
func EmergencyWhenf(conditional bool, msg string, args ...interface{}) {
	std.EmergencyWhenf(conditional, msg, args...)
}
func (self *Journaler) EmergencyPanicWhen(conditional bool, msg interface{}) {
	self.conditionalSendPanic(journal.PriEmerg, conditional, msg)
}
func (self *Journaler) EmergencyPanicWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendPanic(journal.PriEmerg, conditional, NewDefaultMessage(msg))
}
func (self *Journaler) EmergencyPanicWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendPanic(journal.PriEmerg, conditional, NewFormatedMessage(msg, args))
}
func (self *Journaler) EmergencyFatalWhen(conditional bool, msg interface{}) {
	self.conditionalSendFatal(journal.PriEmerg, conditional, msg)
}
func (self *Journaler) EmergencyFatalWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendFatal(journal.PriEmerg, conditional, NewDefaultMessage(msg))
}
func (self *Journaler) EmergencyFatalWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendFatal(journal.PriEmerg, conditional, NewFormatedMessage(msg, args))
}
func EmergencyPanicWhen(conditional bool, msg interface{}) {
	std.conditionalSendPanic(journal.PriEmerg, conditional, msg)
}
func EmergencyPanicWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendPanic(journal.PriEmerg, conditional, NewDefaultMessage(msg))
}
func EmergencyPanicWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendPanic(journal.PriEmerg, conditional, NewFormatedMessage(msg, args))
}
func EmergencyFatalWhen(conditional bool, msg interface{}) {
	std.conditionalSendFatal(journal.PriEmerg, conditional, msg)
}
func EmergencyFatalWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendFatal(journal.PriEmerg, conditional, NewDefaultMessage(msg))
}
func EmergencyFatalWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendFatal(journal.PriEmerg, conditional, NewFormatedMessage(msg, args))
}

func (self *Journaler) AlertWhen(condit4ional bool, message interface{}) {
	self.conditionalSend(journal.PriAlert, conditional, message)
}
func AlertWhen(conditional bool, message interface{}) {
	std.AlertWhen(conditional, message)
}
func (self *Journaler) AlertWhenln(conditional bool, msg ...interface{}) {
	self.AlertWhen(conditional, NewDefaultMessage(msg))
}
func AlertWhenln(conditional bool, msg ...interface{}) {
	std.AlertWhenln(conditional, msg...)
}
func (self *Journaler) AlertWhenf(conditional bool, msg string, args ...interface{}) {
	self.AlertWhen(conditional, NewFormatedMessage(msg, args))
}
func AlertWhenf(conditional bool, msg string, args ...interface{}) {
	std.AlertWhenf(conditional, msg, args...)
}
func (self *Journaler) AlertPanicWhen(conditional bool, msg interface{}) {
	self.conditionalSendPanic(journal.PriAlert, conditional, msg)
}
func (self *Journaler) AlertPanicWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendPanic(journal.PriAlert, conditional, NewDefaultMessage(msg))
}
func (self *Journaler) AlertPanicWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendPanic(journal.PriAlert, conditional, NewFormatedMessage(msg, args))
}
func (self *Journaler) AlertFatalWhen(conditional bool, msg interface{}) {
	self.conditionalSendFatal(journal.PriAlert, conditional, msg)
}
func (self *Journaler) AlertFatalWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendFatal(journal.PriAlert, conditional, NewDefaultMessage(msg))
}
func (self *Journaler) AlertFatalWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendFatal(journal.PriAlert, conditional, NewFormatedMessage(msg, args))
}
func AlertPanicWhen(conditional bool, msg interface{}) {
	std.conditionalSendPanic(journal.PriAlert, conditional, msg)
}
func AlertPanicWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendPanic(journal.PriAlert, conditional, NewDefaultMessage(msg))
}
func AlertPanicWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendPanic(journal.PriAlert, conditional, NewFormatedMessage(msg, args))
}
func AlertFatalWhen(conditional bool, msg interface{}) {
	std.conditionalSendFatal(journal.PriAlert, conditional, msg)
}
func AlertFatalWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendFatal(journal.PriAlert, conditional, NewDefaultMessage(msg))
}
func AlertFatalWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendFatal(journal.PriAlert, conditional, NewFormatedMessage(msg, args))
}

func (self *Journaler) CriticalWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriCrit, conditional, message)
}
func CriticalWhen(conditional bool, message interface{}) {
	std.CriticalWhen(conditional, message)
}
func (self *Journaler) CriticalWhenln(conditional bool, msg ...interface{}) {
	self.CriticalWhen(conditional, NewDefaultMessage(msg))
}
func CriticalWhenln(conditional bool, msg ...interface{}) {
	std.CriticalWhenln(conditional, msg...)
}
func (self *Journaler) CriticalWhenf(conditional bool, msg string, args ...interface{}) {
	self.CriticalWhen(conditional, NewFormatedMessage(msg, args))
}
func CriticalWhenf(conditional bool, msg string, args ...interface{}) {
	std.CriticalWhenf(conditional, msg, args...)
}
func (self *Journaler) CriticalPanicWhen(conditional bool, msg interface{}) {
	self.conditionalSendPanic(journal.PriCrit, conditional, msg)
}
func (self *Journaler) CriticalPanicWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendPanic(journal.PriCrit, conditional, NewDefaultMessage(msg))
}
func (self *Journaler) CriticalPanicWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendPanic(journal.PriCrit, conditional, NewFormatedMessage(msg, args))
}
func (self *Journaler) CriticalFatalWhen(conditional bool, msg interface{}) {
	self.conditionalSendFatal(journal.PriCrit, conditional, msg)
}
func (self *Journaler) CriticalFatalWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendFatal(journal.PriCrit, conditional, NewDefaultMessage(msg))
}
func (self *Journaler) CriticalFatalWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendFatal(journal.PriCrit, conditional, NewFormatedMessage(msg, args))
}
func CriticalPanicWhen(conditional bool, msg interface{}) {
	std.conditionalSendPanic(journal.PriCrit, conditional, msg)
}
func CriticalPanicWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendPanic(journal.PriCrit, conditional, NewDefaultMessage(msg))
}
func CriticalPanicWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendPanic(journal.PriCrit, conditional, NewFormatedMessage(msg, args))
}
func CriticalFatalWhen(conditional bool, msg interface{}) {
	std.conditionalSendFatal(journal.PriCrit, conditional, msg)
}
func CriticalFatalWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendFatal(journal.PriCrit, conditional, NewDefaultMessage(msg))
}
func CriticalFatalWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendFatal(journal.PriCrit, conditional, NewFormatedMessage(msg, args))
}

func (self *Journaler) ErrorWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriErr, conditional, message)
}
func ErrorWhen(conditional bool, message interface{}) {
	std.ErrorWhen(conditional, message)
}
func (self *Journaler) ErrorWhenln(conditional bool, msg ...interface{}) {
	self.ErrorWhen(conditional, NewDefaultMessage(msg))
}
func ErrorWhenln(conditional bool, msg ...interface{}) {
	std.ErrorWhenln(conditional, msg...)
}
func (self *Journaler) ErrorWhenf(conditional bool, msg string, args ...interface{}) {
	self.ErrorWhen(conditional, NewFormatedMessage(msg, args))
}
func ErrorWhenf(conditional bool, msg string, args ...interface{}) {
	std.ErrorWhenf(conditional, msg, args...)
}
func (self *Journaler) ErrorPanicWhen(conditional bool, msg interface{}) {
	self.conditionalSendPanic(journal.PriErr, conditional, msg)
}
func (self *Journaler) ErrorPanicWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendPanic(journal.PriErr, conditional, NewDefaultMessage(msg))
}
func (self *Journaler) ErrorPanicWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendPanic(journal.PriErr, conditional, NewFormatedMessage(msg, args))
}
func (self *Journaler) ErrorFatalWhen(conditional bool, msg interface{}) {
	self.conditionalSendFatal(journal.PriErr, conditional, msg)
}
func (self *Journaler) ErrorFatalWhenln(conditional bool, msg ...interface{}) {
	self.conditionalSendFatal(journal.PriErr, conditional, NewDefaultMessage(msg))
}
func (self *Journaler) ErrorFatalWhenf(conditional bool, msg string, args ...interface{}) {
	self.conditionalSendFatal(journal.PriErr, conditional, NewFormatedMessage(msg, args))
}
func ErrorPanicWhen(conditional bool, msg interface{}) {
	std.conditionalSendPanic(journal.PriErr, conditional, msg)
}
func ErrorPanicWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendPanic(journal.PriErr, conditional, NewDefaultMessage(msg))
}
func ErrorPanicWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendPanic(journal.PriErr, conditional, NewFormatedMessage(msg, args))
}
func ErrorFatalWhen(conditional bool, msg interface{}) {
	std.conditionalSendFatal(journal.PriErr, conditional, msg)
}
func ErrorFatalWhenln(conditional bool, msg ...interface{}) {
	std.conditionalSendFatal(journal.PriErr, conditional, NewDefaultMessage(msg))
}
func ErrorFatalWhenf(conditional bool, msg string, args ...interface{}) {
	std.conditionalSendFatal(journal.PriErr, conditional, NewFormatedMessage(msg, args))
}

func (self *Journaler) WarningWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriWarning, conditional, message)
}
func WarningWhen(conditional bool, message interface{}) {
	std.WarningWhen(conditional, message)
}
func (self *Journaler) WarningWhenln(conditional bool, msg ...interface{}) {
	self.WarningWhen(conditional, NewDefaultMessage(msg))
}
func WarningWhenln(conditional bool, msg ...interface{}) {
	std.WarningWhenln(conditional, msg...)
}
func (self *Journaler) WarningWhenf(conditional bool, msg string, args ...interface{}) {
	self.WarningWhen(conditional, NewFormatedMessage(msg, args))
}
func WarningWhenf(conditional bool, msg string, args ...interface{}) {
	std.WarningWhenf(conditional, msg, args...)
}

func (self *Journaler) NoticeWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriNotice, conditional, message)
}
func NoticeWhen(conditional bool, message interface{}) {
	std.NoticeWhen(conditional, message)
}
func (self *Journaler) NoticeWhenln(conditional bool, msg ...interface{}) {
	self.NoticeWhen(conditional, NewDefaultMessage(msg))
}
func NoticeWhenln(conditional bool, msg ...interface{}) {
	std.NoticeWhenln(conditional, msg...)
}
func (self *Journaler) NoticeWhenf(conditional bool, msg string, args ...interface{}) {
	self.NoticeWhen(conditional, NewFormatedMessage(msg, args))
}
func NoticeWhenf(conditional bool, msg string, args ...interface{}) {
	std.NoticeWhenf(conditional, msg, args...)
}

func (self *Journaler) InfoWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriInfo, conditional, message)
}
func InfoWhen(conditional bool, message interface{}) {
	std.InfoWhen(conditional, message)
}
func (self *Journaler) InfoWhenln(conditional bool, msg ...interface{}) {
	self.InfoWhen(conditional, NewDefaultMessage(msg))
}
func InfoWhenln(conditional bool, msg ...interface{}) {
	std.InfoWhenln(conditional, msg...)
}
func (self *Journaler) InfoWhenf(conditional bool, msg string, args ...interface{}) {
	self.InfoWhen(conditional, NewFormatedMessage(msg, args))
}
func InfoWhenf(conditional bool, msg string, args ...interface{}) {
	std.InfoWhenf(conditional, msg, args...)
}

func (self *Journaler) DebugWhen(conditional bool, message interface{}) {
	self.conditionalSend(journal.PriDebug, conditional, message)
}
func DebugWhen(conditional bool, message interface{}) {
	std.DebugWhen(conditional, message)
}
func (self *Journaler) DebugWhenln(conditional bool, msg ...interface{}) {
	self.DebugWhen(conditional, NewDefaultMessage(msg))
}
func DebugWhenln(conditional bool, msg ...interface{}) {
	std.DebugWhenln(conditional, msg...)
}
func (self *Journaler) DebugWhenf(conditional bool, msg string, args ...interface{}) {
	self.DebugWhen(conditional, NewFormatedMessage(msg, args))
}
func DebugWhenf(conditional bool, msg string, args ...interface{}) {
	std.DebugWhenf(conditional, msg, args...)
}