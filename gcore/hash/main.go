package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

type Hasher struct {
	FilePath string
}

func NewHasher(args []string) *Hasher {
	filePath := ""
	fmt.Println(args)
	if len(args) > 1 {
		filePath = args[1]
		fmt.Println("filePath:", filePath)
	}
	return &Hasher{
		FilePath: filePath,
	}
}

func (h *Hasher) CalculateHash() string {
	ok, err := h.ifFile()
	if err != nil {
		return "未知错误，我的天呐！"
	}
	if ok {
		result, err := h.calculateFileMD5()
		if err != nil {
			return "计算错误"
		}
		return result
	} else {
		return "不支持非文件类型"
	}
}

func (h *Hasher) ifFile() (bool, error) {
	fileInfo, err := os.Stat(h.FilePath)
	if err != nil {
		return false, err
	}
	return !fileInfo.IsDir(), nil
}

func (h *Hasher) calculateFileMD5() (string, error) {
	file, err := os.Open(h.FilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := md5.New()
	buffer := make([]byte, 2*1024*1024) // 2MB
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			_, _ = hash.Write(buffer[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
	}
	md5sum := hash.Sum(nil)
	return hex.EncodeToString(md5sum), nil
}
