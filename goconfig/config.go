package goconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("設定ファイル '%s' を開けませんでした: %w", filePath, err)
	}
	defer file.Close()

	// JSONデコーダーを作成し、ファイルの内容をデコードする
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(v); err != nil {
		return fmt.Errorf("設定ファイル '%s' のデコードに失敗しました: %w", filePath, err)
	}

	return nil
}