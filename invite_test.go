package invite

import (
	"testing"
)

func TestGenerator_Encode(t *testing.T) {
	g := NewGenerator[uint32](CHARSET, 6)
	t.Log("最大支持ID:", g.MaxSupportID())

	test := func(id uint32) bool {
		code, e := g.Encode(id)
		if e != nil {
			t.Error(id, e)
			return false
		}
		t.Logf("ID:%d code:%s", id, code)
		nid := g.Decode(code)
		//t.Logf("解析邀请码结果：code:%s id:%d 是否相等:%t", code, nid, id == nid)
		//t.Log("=========================")
		if nid != id {
			t.Error(id, nid)
			return false
		}
		return true
	}

	var _min, _max uint32 = 0, 20
	for id := _min; id <= _max; id++ {
		if !test(id) {
			return
		}
	}
}
