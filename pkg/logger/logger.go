package logger

import "log"

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {

	log.Printf(format, args...)

}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {

	log.Printf(format, args...)

}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {

	log.Printf(format, args...)

}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {

	log.Printf(format, args...)

}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {

	log.Fatalf(format, args...)

}
