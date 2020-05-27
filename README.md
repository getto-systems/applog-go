# applog

golang: simple logger

status: production ready

```golang
import (
	"fmt"
	"log"
	"os"

	"github.com/getto-systems/applog-go"
)

output := log.New(os.Stdout, "", 0)
entry := func(level string, message string) string {
	return fmt.Sprintf("%s: %s", level, message)
}

logger := NewDebugLogger(output, entry)

logger.Debug("log message")
// output: DEBUG: log message
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
logger := NewDebugLogger(output, entry)
```

- output has `Println(v ...interface{})` method : such as `*log.Logger`
- entry is formatting function : `func(level string, message string) string`


### leveled Logger

- DebugLogger : logging AUDIT, ERROR, WARN, INFO, DEBUG
- InfoLogger : logging AUDIT, ERROR, WARN, INFO
- WarnLogger : logging AUDIT, ERROR, WARN
- ErrorLogger : logging AUDIT, ERROR


## License

[MIT](LICENSE) license.

Copyright &copy; shun-fix9
