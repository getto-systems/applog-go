# applog

golang: simple logger

status: production ready : version 2

```golang
package applog

import (
	"log"
	"os"
)

output := log.New(os.Stdout, "", 0)

debugLogger := NewDebugLogger(output)

debugLogger.Always("always: log message")
debugLogger.Info("info: log message")
debugLogger.Debug("debug: log message")
// output: always: log message
// info: log message
// debug: log message


infoLogger := NewInfoLogger(output)

infoLogger.Always("always: log message")
infoLogger.Info("info: log message")
infoLogger.Debug("debug: log message")
// output: always: log message
// info: log message


quietLogger := NewQuietLogger(output)

quietLogger.Always("always: log message")
quietLogger.Info("info: log message")
quietLogger.Debug("debug: log message")
// output: always: log message
```


###### Table of Contents

- [Requirements](#Requirements)
- [Usage](#Usage)
- [License](#License)

## Requirements

- golang: 1.14


## Usage

### initialize Logger

```golang
logger := NewDebugLogger(output)
logger := NewInfoLogger(output)
logger := NewQuietLogger(output)
```

- output has `Println(v ...interface{})` method : such as `*log.Logger`


### leveled Logger

- DebugLogger : logging Always, Info, Debug message
- InfoLogger : logging Always, Info message
- QuietLogger : logging Always message


## License

[MIT](LICENSE) license.

Copyright &copy; shun-fix9
