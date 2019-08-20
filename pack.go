package packer

func Pack(sourceSize, batch int) []Range {
	if sourceSize > batch {
		packs := sourceSize / batch

		length := packs
		end := sourceSize%batch > 0
		if end {
			length += 1
		}

		result := make([]Range, 0, length)

		from := 0
		to := batch
		for i := 0; i < packs; i++ {
			result = append(result, Range{from, to})

			from, to = to, to+batch
		}

		if end {
			result = append(result, Range{from, sourceSize})
		}

		return result
	}

	return []Range{{0, sourceSize}}
}
