package pwd

import (
	"fmt"
	"testing"
)

// 对加密的函数进行测试
func TestHashPwd(t *testing.T) {
	fmt.Printf(HashPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$AxrxsqkYMiw5PF/.f6NBhOZrs9k01XFSe4EvUFmHqEBuAqL9OP7EG", "1234"))
}
