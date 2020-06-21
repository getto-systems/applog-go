# applog

golang: simple logger

status: production ready

```golang
package applog

import (
	"log"
	"os"
)

output := log.New(os.Stdout, "", 0)

logger := NewDebugLogger(output)

logger.Debug("debug: log message")
// output: debug: log message
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
