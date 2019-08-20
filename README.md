# Packer for split by ranges

### Dep
```bash
dep ensure -add github.com/gopereza/packer
```

### Use for batch operations

```golang
	assert.Equal(t, []Range{{0, 100}}, Pack(100, 200))
	assert.Equal(t, []Range{{0, 100}}, Pack(100, 500))

	assert.Equal(t, []Range{{0, 50}, {50, 100}}, Pack(100, 50))

	assert.Equal(t, []Range{{0, 30}, {30, 60}, {60, 90}, {90, 100}}, Pack(100, 30))
```

### Example
```golang
package main

import "github.com/gopereza/packer"

func main() {
	const (
		size = 1000000
		pack = 200
	)

	ids := make([]int, size)

	packs := packer.Pack(size, pack)

	for _, pack := range packs {
		update(ids[pack.From:pack.To])
	}
}

func update(ids []int) {
	// NOP
}
```
