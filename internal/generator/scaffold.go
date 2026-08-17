package generator

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Rez209/golug/templates"
)

const (
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

type TemplateData struct {
	ServiceName string
	Port        string
	Module      string
}

func GenerateService(serviceName string, port string, lang string, module string) error {
	data := TemplateData{
		ServiceName: serviceName,
		Port:        port,
		Module:      module, 
	}

	if err := os.MkdirAll(serviceName, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %v", serviceName, err)
	}

	fmt.Printf(Yellow+"⏳ Generating templates for %s..."+Reset+"\n", strings.ToUpper(lang))

	err := fs.WalkDir(templates.FS, lang, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("templates for language '%s' not found: %v", lang, err)
		}
		if d.IsDir() {
			return nil
		}

		relPath := strings.TrimPrefix(path, lang+"/")
		targetPath := strings.TrimSuffix(relPath, ".tmpl")
		targetPath = strings.Replace(targetPath, "cmd/app", "cmd/"+serviceName, 1)

		finalFilePath := filepath.Join(serviceName, targetPath)
		finalDir := filepath.Dir(finalFilePath)

		if err := os.MkdirAll(finalDir, 0755); err != nil {
			return err
		}

		content, err := templates.FS.ReadFile(path)
		if err != nil {
			return err
		}

		tmpl, err := template.New(targetPath).Parse(string(content))
		if err != nil {
			return err
		}

		f, err := os.Create(finalFilePath)
		if err != nil {
			return err
		}
		defer func() { _ = f.Close() }()

		if err := tmpl.Execute(f, data); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	gitCmd := exec.Command("git", "init")
	gitCmd.Dir = serviceName
	if err := gitCmd.Run(); err != nil {
		fmt.Printf(Yellow+"⚠️ Warning: failed to initialize Git: %v"+Reset+"\n", err)
	} else {
		fmt.Println(Blue + "📦 Git repository initialized successfully!" + Reset)
	}

	fmt.Printf(Green+"✅ Successfully generated %s project for %s!"+Reset+"\n", strings.ToUpper(lang), serviceName)
	return nil
}
