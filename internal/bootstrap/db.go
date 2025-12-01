package bootstrap

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/guothion/xuanyuan/internal/global"
	"github.com/guothion/xuanyuan/internal/model"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() *gorm.DB {
	switch global.App.Config.DataSource.Driver {
	case "mysql":
		return initMySqlGorm()
	default:
		return initMySqlGorm()
	}
}

// Session 创建一个带有trace_id的数据库会话
func Session(ctx context.Context) *gorm.DB {
	db := global.App.DB

	// 从上下文中获取trace_id
	traceID, ok := ctx.Value("trace_id").(string)
	if ok && traceID != "" {
		// 将trace_id添加到数据库会话中
		return db.WithContext(ctx).Session(&gorm.Session{
			Logger: db.Logger.LogMode(logger.Info),
		}).Set("trace_id", traceID)
	}

	return db.WithContext(ctx)
}

func initMySqlGorm() *gorm.DB {
	dbConfig := global.App.Config.DataSource

	if dbConfig == nil {
		logrus.Errorf("database config is nil")
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Schema, dbConfig.Charset)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
		//Logger:                                   getGormLogger(), // 使用自定义 logger
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		DryRun: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",                                // table name prefix, table for `User` would be `t_users`
			SingularTable: false,                             // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,                             // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	}); err != nil {
		logrus.Errorf("mysql connect failed, err:", err)
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		// 连接池基本配置
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)

		// Session相关配置
		if dbConfig.ConnMaxLifetime > 0 {
			sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime) * time.Second)
		}
		if dbConfig.ConnMaxIdleTime > 0 {
			sqlDB.SetConnMaxIdleTime(time.Duration(dbConfig.ConnMaxIdleTime) * time.Second)
		}

		// 初始化表
		initMySqlTables(db)
		sessionDB := db.Session(&gorm.Session{
			// ✅ 开启预编译语句缓存（提升高频查询性能）
			PrepareStmt: dbConfig.PrepareStmt,
			// ✅ 自动刷新缓存（配合 PrepareStmt）
			// (GORM v2 默认开启)
			// ❌ 不要在这里设置 Context！因为不能全局共享一个 ctx
			Logger: getGormLogger(),
		})
		return sessionDB
	}
}

//func initPostgresGorm() *gorm.DB {}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel
	switch global.App.Config.DataSource.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}
	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                            // 慢 SQL 阈值
		LogLevel:                  logMode,                                           // 日志级别
		IgnoreRecordNotFoundError: false,                                             // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !global.App.Config.DataSource.EnableFileLogWriter, // 禁用彩色打印
	})
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.App.Config.DataSource.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.DataSource.LogFileName,
			MaxSize:    global.App.Config.Log.Size,
			MaxBackups: global.App.Config.Log.Backups,
			MaxAge:     global.App.Config.Log.Age,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

func initMySqlTables(db *gorm.DB) {
	// 我们在这里初始化表格
	err := db.Set("gorm:table _options", "ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(
		&model.User{},
		&model.Step{},
		&model.Process{},
		&model.ProcessLogs{},
		&model.StallApplication{},
		&model.Stalls{},
		&model.StallImages{},
		&model.StallCategories{},
		&model.Favorites{},
		&model.Comments{},
		&model.Notifications{},
		&model.AdminLogs{},
	)

	if err != nil {
		logrus.Errorln("migrate table failed", err)
		os.Exit(0)
	}
	// 下面是我们执行的sql语句，比如我们可以设置一些默认值等
	//db.Exec()
}
