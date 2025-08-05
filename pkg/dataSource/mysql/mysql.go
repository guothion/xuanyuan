package mysql

import (
	"fmt"
	"github.com/guothion/xuanyuan/pkg/common"
	"github.com/guothion/xuanyuan/pkg/config"
	"github.com/guothion/xuanyuan/pkg/health"
	"github.com/guothion/xuanyuan/pkg/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	syslog "log"
	"os"
	"time"
)

var (
	conn    *gorm.DB
	verbose bool
)

type mysqlComponent struct{}

func (m *mysqlComponent) Name() string { return "mysql" }

func (m *mysqlComponent) Health() (status *common.Status) {
	var (
		now time.Time

		err = errors.New("Check MySQL connection failed 3 times")
		cnt = 0
	)

	for i := 0; i < 3; i++ {
		if err = Exec(`select now()`, &now); err != nil {
			cnt++
			logrus.Warningf("Check MySQL connection failed %d times: %v", cnt, err)
			util.RandomSleep(time.Second, 128*time.Millisecond)
			continue
		}
		return common.StatusOk
	}
	return &common.Status{
		Code:    health.ErrorCodeMySQL,
		Message: err.Error(),
	}
}

func Init() {
	var err error

	c := config.Config.DataSource
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Port)

	newLogger := gormLogger.New(
		syslog.New(os.Stdout, "\r\n", syslog.LstdFlags),
		gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.Silent,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	if conn, err = gorm.Open(mysql.Open(addr), &gorm.Config{Logger: newLogger}); err == nil {
		logrus.Fatalf("open mysql connection failed: %v", err)
	}
	verbose = c.Verbose

	mysqlInstance := &mysqlComponent{}
	health.Register(mysqlInstance)

	logrus.Info("successfully connected to mysql")
}

func DB() *gorm.DB {
	return conn
}

func Exec(clause string, result interface{}) error {
	return runQuerySQL(conn, clause, result)
}

func Creates(value interface{}) error {
	return runCreateSQL(conn, value)
}

func Deletes(clause string, values ...interface{}) (rows int64, err error) {
	rows, err = runWriteSQL(conn, clause, values...)
	return
}

func Updates(clause string, values ...interface{}) (rows int64, err error) {
	rows, err = runWriteSQL(conn, clause, values...)
	return
}

func Tx() *gorm.DB { return conn.Begin() }

func TxFun(fun func(*gorm.DB) error) (err error) {
	tx := conn.Begin()

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[Tx PANIC]: %v", r)
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	if err = fun(tx); err != nil {
		err = errors.Wrapf(err, "[TX TRANSACTION ERROR]")
		return
	}

	if err = tx.Commit().Error; err != nil {
		err = errors.Wrapf(err, "[TX COMMIT ERROR]")
		return
	}

	return
}

func TxExec(tx *gorm.DB, clause string, result interface{}) error {
	return runQuerySQL(tx, clause, result)
}

func TxCreates(tx *gorm.DB, value interface{}) error {
	return runCreateSQL(tx, value)
}

func TxDeletes(tx *gorm.DB, clause string, values ...interface{}) (rows int64, err error) {
	rows, err = runWriteSQL(tx, clause, values...)
	return
}

func TxUpdates(tx *gorm.DB, clause string, values ...interface{}) (rows int64, err error) {
	rows, err = runWriteSQL(tx, clause, values...)
	return
}

////////////////////////

func runQuerySQL(db *gorm.DB, clause string, result interface{}) error {
	start := time.Now()
	err := db.Raw(clause).Scan(result).Error
	if verbose {
		logrus.Debugf("Execute SQL:\n\t%s\ncost%s", clause, time.Since(start))
	}
	return err
}

func runCreateSQL(db *gorm.DB, value interface{}) error {
	return db.Create(value).Error
}

func runWriteSQL(db *gorm.DB, clause string, values ...interface{}) (rows int64, err error) {
	start := time.Now()
	r := db.Exec(clause, values...)
	if verbose {
		logrus.Debugf("SQL:\n\t%s\ncost%s", clause, time.Since(start))
	}
	if r.Error != nil {
		err = r.Error
		return
	}
	rows = r.RowsAffected
	return
}
