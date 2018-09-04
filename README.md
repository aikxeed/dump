# dump
Dump Go variables in a pretty format

## Usage Example
```golang
package main

import "github.com/aikxeed/dump"

type example struct {
	num int
	str string
}

func main() {
	num := 42
	str := "test"
	arr := [3]int{1, 2, 3}
	sli := []string{"go", "is", "cool"}
	m := map[string]int{"hello": 42, "world": 0}
	s := example{num: 1, str: "abc"}

	dump.Dump(num, str, arr, sli, m, s)
}
```

Output

```json
[
  {
    "type": "int",
    "value": "42"
  },
  {
    "type": "string",
    "value": "test"
  },
  {
    "type": "[3]int",
    "value": "[1 2 3]"
  },
  {
    "type": "[]string",
    "value": "[go is cool]"
  },
  {
    "type": "map[string]int",
    "value": "map[hello:42 world:0]"
  },
  {
    "type": "main.example",
    "value": "{num:1 str:abc}"
  }
]
```

## License
MIT