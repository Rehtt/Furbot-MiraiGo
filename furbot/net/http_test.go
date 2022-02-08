/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/2/8 16:39
 */

package net

import (
	"fmt"
	"testing"
)

func TestHttpGet(t *testing.T) {
	data, err := httpGet("https://rehtt.com", nil, nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(data))
}
