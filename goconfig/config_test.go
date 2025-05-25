package goconfig

import (
	"os"
	"testing"
)

type TestConfig struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"database"`
	ApiKey    string `json:"api_key"`
	DebugMode bool   `json:"debug_mode"`
}

// テスト用のJSONファイルを作成する関数
func createTestConfigFile(t *testing.T, content string) string {
	t.Helper()
	file, err := os.CreateTemp("", "test_config_*.json")
	if err != nil {
		t.Fatalf("一時ファイルの作成に失敗しました: %v", err)
	}
	filePath := file.Name()
	_, err = file.WriteString(content)
	if err != nil {
		file.Close()
		os.Remove(filePath)
		t.Fatalf("一時ファイルへの書き込みに失敗しました: %v", err)
	}
	file.Close()
	return filePath
}

func TestLoad_Success(t *testing.T) {
	jsonContent := `{
		"database": {
			"host": "localhost",
			"port": 5432,
			"user": "admin",
			"password": "secretpassword"
		},
		"api_key": "YOUR_API_KEY_HERE",
		"debug_mode": true
	}`
	filePath := createTestConfigFile(t, jsonContent)
	defer os.Remove(filePath)

	var cfg TestConfig
	err := Load(filePath, &cfg)

	if err != nil {
		t.Fatalf("Load() で予期せぬエラーが発生しました: %v", err)
	}

	if cfg.Database.Host != "localhost" {
		t.Errorf("Database.Host の値が不正です: got %v want %v", cfg.Database.Host, "localhost")
	}
	if cfg.Database.Port != 5432 {
		t.Errorf("Database.Port の値が不正です: got %v want %v", cfg.Database.Port, 5432)
	}
	if cfg.ApiKey != "YOUR_API_KEY_HERE" {
		t.Errorf("ApiKey の値が不正です: got %v want %v", cfg.ApiKey, "YOUR_API_KEY_HERE")
	}
	if !cfg.DebugMode {
		t.Errorf("DebugMode の値が不正です: got %v want %v", cfg.DebugMode, true)
	}
}

func TestLoad_FileNotFound(t *testing.T) {
	var cfg TestConfig
	err := Load("non_existent_file.json", &cfg)

	if err == nil {
		t.Fatalf("存在しないファイルを指定した場合にエラーが返されませんでした")
	}
}

func TestLoad_InvalidJson(t *testing.T) {
	jsonContent := `{ "database": { "host": "localhost", }`
	filePath := createTestConfigFile(t, jsonContent)
	defer os.Remove(filePath)

	var cfg TestConfig
	err := Load(filePath, &cfg)

	if err == nil {
		t.Fatalf("不正なJSONを読み込んだ場合にエラーが返されませんでした")
	}
}