# zap-examples

uber-go/zap logger examples

## override default config

Example of overriding default config, if you are fine with default, just use `logger, _ := zap.NewProduction()` to
instantiate logger.

```go
func NewZapConfig(level zapcore.Level) zap.Config {

    config := zap.NewProductionConfig()
    config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
    config.EncoderConfig.TimeKey = "timestamp"
    config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    config.Level.SetLevel(level)
    return config
}
```

## examples

### Simple log

```go
logger, _ := NewZapConfig(zapcore.InfoLevel).Build()
defer logger.Sync()

logger.Info("hello")
who := "world"
logger.Warn(fmt.Sprintf("hello %s", who))
```

Output
```json
{"level":"INFO","timestamp":"2021-08-05T11:41:06.069+0100","caller":"zap-examples/main.go:14","msg":"hello"}
{"level":"WARN","timestamp":"2021-08-05T11:41:06.070+0100","caller":"zap-examples/main.go:16","msg":"hello world"}
```

### Log with fields/context

```go
logger, _ := NewZapConfig(zapcore.InfoLevel).Build()
defer logger.Sync()

url := "test.com"
logger.Info("failed to fetch URL",
    zap.String("url", url),
    zap.Int("attempt", 3),
    zap.Duration("backoff", time.Second),
)
```

Output
```json
{"level":"INFO","timestamp":"2021-08-05T11:45:50.029+0100","caller":"zap-examples/main.go:19","msg":"failed to fetch URL","url":"test.com","attempt":3,"backoff":1}
```

### Log with global/parent fields/context

```go
logger, _ := NewZapConfig(zapcore.InfoLevel).Build()
defer logger.Sync()

// create child logger with structured context
logger = logger.With(zap.String("app", "zap-examples"))
logger.Info("failed to fetch URL",
    zap.Int("attempt", 3),
)
```

Output
```json
{"level":"INFO","timestamp":"2021-08-05T11:49:34.845+0100","caller":"zap-examples/main.go:17","msg":"failed to fetch URL","app":"zap-examples","attempt":3}
```

### Sugared logger

Sugared logger can be simpler to use, because it includes `pritf` style API.

```go
logger, _ := NewZapConfig(zapcore.InfoLevel).Build()
defer logger.Sync()

sugar := logger.Sugar()
// no need to call zap.Int(...) function
sugar.Infow("failed to fetch URL",
    "attempt", 3,
)

url := "test"
// no need to call fmt.Sprintf
sugar.Infof("Failed to fetch URL: %s", url)
```

Output
```json
{"level":"INFO","timestamp":"2021-08-05T11:53:44.552+0100","caller":"zap-examples/main.go:22","msg":"failed to fetch URL","attempt":3}
{"level":"INFO","timestamp":"2021-08-05T11:53:44.552+0100","caller":"zap-examples/main.go:27","msg":"Failed to fetch URL: test"}
```

### Global logger

Global logger is NOT recommended to use. Logger should be either passed around as an argument, or set a variable.
Default global logger is set to `no-op` (it never writes any logs). If you still want/need to use global logger,
you need to replace global one:

```go
logger, _ := NewZapConfig(zapcore.InfoLevel).Build()
defer logger.Sync()

// without this line, global logger would still be set to no-op logger
zap.ReplaceGlobals(logger)
// L() return logger and S() returns sugared logger
zap.L().Info("hello")
```

Output
```json
{"level":"INFO","timestamp":"2021-08-05T11:58:51.942+0100","caller":"zap-examples/main.go:14","msg":"hello"}
```
