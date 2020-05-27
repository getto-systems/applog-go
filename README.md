# applog

golang: simple logger

status: production ready


###### Table of Contents

- [Requirements](#Requirements)
- [Usage](#Usage)
- [License](#License)

## Requirements

- golang: 1.14


## Usage

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

logger := applog.NewDebugLogger(output, entry)

logger.Debug("message") // => DEBUG: message
```


## License

[MIT](LICENSE) license.

Copyright &copy; shun-fix9
