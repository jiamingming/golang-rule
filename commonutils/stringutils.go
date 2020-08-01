package commonutils

import "bytes"

func ContactString(item ...string) string {

	var item_buffer bytes.Buffer
	for _, v := range item {
		item_buffer.WriteString(v)
	}
	return item_buffer.String()
}
