# component world

This module is a component package, 
to use it in your project:

1. you should add `liuggio/hello:v1` in your require section of your `component.json`
2. run `component install`
3. use in your code like a library:

```
import "liuggio/hello"

hi, err := hello()

fmt.Print(hi, err)
```

3. Run your script.

