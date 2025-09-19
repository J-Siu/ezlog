/*
Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

/*
ezlog - A simple log mapping module

0 Emerg
1 Alert
2 Crit
3 Err
4 Warning
5 Notice
6 Info
7 Debug
8 Trace
9 Disable
*/
package ezlog

type Level int8

const (
	EmergLevel Level = iota
	AlertLevel
	CritLevel
	ErrLevel
	WarningLevel
	NoticeLevel
	InfoLevel
	DebugLevel
	TraceLevel
	Disabled
)

var (
	logLevel = Level(Disabled)
	Logger   = new(Log)
)

type LogFunc func(msg *string)

type Log struct {
	emerg   LogFunc
	alert   LogFunc
	crit    LogFunc
	err     LogFunc
	warning LogFunc
	notice  LogFunc
	info    LogFunc
	debug   LogFunc
	trace   LogFunc
}

func (log *Log) SetEmerg(f LogFunc) {
	log.emerg = f
}

func (log *Log) SetAlert(f LogFunc) {
	log.alert = f
}

func (log *Log) SetCrit(f LogFunc) {
	log.crit = f
}

func (log *Log) SetErr(f LogFunc) {
	log.err = f
}

func (log *Log) SetWarning(f LogFunc) {
	log.warning = f
}

func (log *Log) SetNotice(f LogFunc) {
	log.notice = f
}

func (log *Log) SetInfo(f LogFunc) {
	log.info = f
}

func (log *Log) SetTrace(f LogFunc) {
	log.trace = f
}

func (log *Log) SetDebug(f LogFunc) {
	log.debug = f
}

func Emerg(msg string) {
	if logLevel >= EmergLevel && Logger.emerg != nil {
		Logger.emerg(&msg)
	}
}
func Alert(msg string) {
	if logLevel >= AlertLevel && Logger.alert != nil {
		Logger.alert(&msg)
	}
}
func Crit(msg string) {
	if logLevel >= CritLevel && Logger.crit != nil {
		Logger.crit(&msg)
	}
}
func Err(msg string) {
	if logLevel >= ErrLevel && Logger.err != nil {
		Logger.err(&msg)
	}
}
func Warning(msg string) {
	if logLevel >= WarningLevel && Logger.warning != nil {
		Logger.warning(&msg)
	}
}
func Notice(msg string) {
	if logLevel >= NoticeLevel && Logger.notice != nil {
		Logger.notice(&msg)
	}
}
func Info(msg string) {
	if logLevel >= InfoLevel && Logger.info != nil {
		Logger.info(&msg)
	}
}
func Debug(msg string) {
	if logLevel >= DebugLevel && Logger.debug != nil {
		Logger.debug(&msg)
	}
}
func Trace(msg string) {
	if logLevel >= TraceLevel && Logger.trace != nil {
		Logger.trace(&msg)
	}
}

func LogLevel() Level {
	return logLevel
}

func SetLogLevel(level Level) {
	logLevel = level
}

func EmergP(msg *string) {
	if logLevel >= EmergLevel && Logger.emerg != nil {
		Logger.emerg(msg)
	}
}
func AlertP(msg *string) {
	if logLevel >= AlertLevel && Logger.alert != nil {
		Logger.alert(msg)
	}
}
func CritP(msg *string) {
	if logLevel >= CritLevel && Logger.crit != nil {
		Logger.crit(msg)
	}
}
func ErrP(msg *string) {
	if logLevel >= ErrLevel && Logger.err != nil {
		Logger.err(msg)
	}
}
func WarningP(msg *string) {
	if logLevel >= WarningLevel && Logger.warning != nil {
		Logger.warning(msg)
	}
}
func NoticeP(msg *string) {
	if logLevel >= NoticeLevel && Logger.notice != nil {
		Logger.notice(msg)
	}
}
func InfoP(msg *string) {
	if logLevel >= InfoLevel && Logger.info != nil {
		Logger.info(msg)
	}
}
func DebugP(msg *string) {
	if logLevel >= DebugLevel && Logger.debug != nil {
		Logger.debug(msg)
	}
}
func TraceP(msg *string) {
	if logLevel >= TraceLevel && Logger.trace != nil {
		Logger.trace(msg)
	}
}
