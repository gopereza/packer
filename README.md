# Packer for split by ranges

### use for batch operations

```golang
	assert.Equal(t, []Range{{0, 100}}, Pack(100, 200))
	assert.Equal(t, []Range{{0, 100}}, Pack(100, 500))

	assert.Equal(t, []Range{{0, 50}, {50, 100}}, Pack(100, 50))

	assert.Equal(t, []Range{{0, 30}, {30, 60}, {60, 90}, {90, 100}}, Pack(100, 30))
```