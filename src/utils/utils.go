package utils

import (
	"os"
	"path/filepath"
	strings "strings"
	"text/template"

	"github.com/Dencyuman/logvista-server/config"
)

// 文字列中のダミー改行文字\\nを改行文字に置き換える
func ReplaceDummyLf(str string) string {
	return strings.ReplaceAll(str, "\\n", "\n")
}

// string配列から空文字を除外する
func FilterEmptyStrings(strings []string) []string {
	var result []string
	for _, s := range strings {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}

// 指定されたディレクトリ内の最初のJSファイルを探す
func FindFirstJSFile(staticDirPath string) (string, error) {
	var firstJSFile string
	found := false

	err := filepath.Walk(staticDirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != staticDirPath {
			// staticディレクトリ以外のディレクトリはスキップ
			return filepath.SkipDir
		}
		if !found && !info.IsDir() && strings.HasSuffix(info.Name(), ".js") {
			firstJSFile = path
			found = true
			return filepath.SkipDir // 走査を停止
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	return firstJSFile, nil
}

// テンプレートjs内の{{.VITE_API_URL}}に.env内のVITE_API_URLを埋め込んだjsファイルを生成する
func GenerateJSFileFromTemplate(tmplPath string, outputPath string) error {
	input, err := os.ReadFile(tmplPath)
	if err != nil {
		return err
	}
	tmpl, err := template.New("index").Parse(string(input))
	if err != nil {
		return err
	}
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	err = tmpl.Execute(outputFile, struct{ VITE_API_URL string }{config.AppConfig.ViteApiUrl})
	if err != nil {
		return err
	}
	return nil
}
