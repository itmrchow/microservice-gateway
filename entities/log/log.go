package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var (
	logger        *zerolog.Logger
	loggerMutex   sync.RWMutex
	LOG_LEVEL_STR = ""
	LOG_OUTPUT    = ""
	LOG_FILE      = ""
	LOG_DIR       = ""
	SERVER_NAME   = ""

	logsMap sync.Map
)

func InitLog() {
	// init env
	LOG_LEVEL_STR = viper.GetString("log_level")
	LOG_OUTPUT = viper.GetString("log_output")
	LOG_FILE = viper.GetString("log_file")
	LOG_DIR = viper.GetString("log_dir")
	SERVER_NAME = viper.GetString("server_name")

	// set global log level
	logLevel, err := zerolog.ParseLevel(LOG_LEVEL_STR)
	if err != nil {
		log.Fatal().Err(err).Msgf("log level init error: %s", LOG_LEVEL_STR)
	}
	zerolog.SetGlobalLevel(logLevel)

	// set global log time
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// set default logger
	setLogger()

	// check log rotation
	if LOG_OUTPUT == "file" {
		go checkLogRotation()
	}

	// 等待2秒讓log初始化完成
	time.Sleep(2 * time.Second)

	log.Info().Msgf("log init success")
}

func Logger() *zerolog.Logger {
	loggerMutex.RLock()
	defer loggerMutex.RUnlock()
	return logger
}

func Info() *zerolog.Event {
	return Logger().Info()
}

func Warn() *zerolog.Event {
	return Logger().Warn()
}

func Err(err error) *zerolog.Event {
	return Logger().Err(err)
}

func Fatal() *zerolog.Event {
	return Logger().Fatal()
}

func setLogger() {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	lg := log.Logger.Output(getLogWriter())
	logger = &lg
}

func getLogWriter() io.Writer {

	if LOG_OUTPUT == "stdout" {
		return zerolog.ConsoleWriter{Out: os.Stdout}
	} else if LOG_OUTPUT == "file" {
		return fileWriter(SERVER_NAME, LOG_DIR, LOG_FILE)
	}

	return nil
}

type LogT struct {
	lastTime string
	file     *os.File
}

// fileWriter: 寫入log到指定檔案
func fileWriter(serverName, logDir, logFile string) io.Writer {
	timeStr := time.Now().Format("20060102")

	// cache
	logT, ok := logsMap.Load(serverName)
	if ok && logT.(*LogT) != nil && timeStr == logT.(*LogT).lastTime {
		return logT.(*LogT).file
	}
	if ok && logT.(*LogT) != nil {
		go func() {
			time.Sleep(5 * time.Second) // 給予 5 秒緩衝時間
			logT.(*LogT).file.Close()
		}()
	}

	// 確保目錄存在
	filepath := fmt.Sprintf("%s/%s.%s.%s", logDir, serverName, timeStr, logFile)
	if err := os.MkdirAll(filepath[:strings.LastIndex(filepath, "/")], 0775); err != nil {
		log.Fatal().Err(err).Msgf("create log directory error: %s", filepath)
	}
	ff, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0775)
	if err != nil {
		log.Fatal().Err(err).Msgf("new log file error: %s", filepath)
	}

	logNode := &LogT{lastTime: timeStr, file: ff}
	logsMap.Store(serverName, logNode)

	return ff
}

// checkLogRotation: 檢查是否換日 init logger
func checkLogRotation() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {

		timeStr := time.Now().Format("20060102")

		// cache
		logT, ok := logsMap.Load(SERVER_NAME)
		if ok && logT.(*LogT) != nil && timeStr != logT.(*LogT).lastTime {
			setLogger()
		}
	}
}
