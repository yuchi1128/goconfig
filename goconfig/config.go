package goconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func Load(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("設定ファイル '%s' を開けませんでした: %w", filePath, err)
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case ".json":
		// JSONデコーダーを作成し、デコードする
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(v); err != nil {
			return fmt.Errorf("JSONファイル '%s' のデコードに失敗しました: %w", filePath, err)
		}
	case ".yaml", ".yml":
		// YAMLデコーダーを作成し、デコードする
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(v); err != nil {
			return fmt.Errorf("YAMLファイル '%s' のデコードに失敗しました: %w", filePath, err)
		}
	default:
		return fmt.Errorf("サポートされていないファイル形式です: '%s' ( .json, .yaml, .yml のみサポート)", ext)
	}

	// 成功した場合は nil を返す
	return nil
}