// Copyright (c) 2026 Nawakarit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"text/template"

	"fyne.io/fyne/v2/widget"
)

func openBuildTerminal(projectPath, script string, output *widget.Entry) {
	command := "cd -- " + strconv.Quote(projectPath) + " && chmod +x " + strconv.Quote(script) + " && ./" + strconv.Quote(script) + "; exec bash"
	commands := [][]string{}
	if terminal := os.Getenv("TERMINAL"); terminal != "" {
		commands = append(commands, []string{terminal, "-e", "bash", "-lc", command})
	}
	commands = append(commands,
		[]string{"gnome-terminal", "--", "bash", "-lc", command},
		[]string{"x-terminal-emulator", "-e", "bash", "-lc", command},
		[]string{"konsole", "-e", "bash", "-lc", command},
		[]string{"xfce4-terminal", "-e", "bash", "-lc", command},
		[]string{"mate-terminal", "--", "bash", "-lc", command},
		[]string{"alacritty", "-e", "bash", "-lc", command},
		[]string{"kitty", "bash", "-lc", command},
		[]string{"foot", "bash", "-c", command},
		[]string{"wezterm", "start", "--", "bash", "-lc", command},
		[]string{"tilix", "-e", "bash", "-lc", command},
		[]string{"st", "-e", "bash", "-lc", command},
	)

	for _, candidate := range commands {
		if _, err := exec.LookPath(candidate[0]); err != nil {
			continue
		}
		if err := exec.Command(candidate[0], candidate[1:]...).Start(); err == nil {
			output.SetText("✅️ opened terminal: " + candidate[0])
			return
		}
	}
	output.SetText("🔴️ no terminal found; install a terminal or set TERMINAL")
}

// ============================================================================
// ฟังชั้น build Icons
// ============================================================================
func runScriptbuildIcons(projectPath string, output *widget.Entry) {
	openBuildTerminal(projectPath, "buildicons.sh", output)
}

// ============================================================================
// Flatpak
// ============================================================================
// ============================================================================
// ฟังชั้น gen + run template
// ============================================================================
func generateFile(tmplPath, outputPath string, data AppConfig) error {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	//projectPath
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}

// ============================================================================
// ฟังชั้น build เป็นไฟล์ flatpak
// ============================================================================
func runScriptbuildflatpak(projectPath string, output *widget.Entry) {
	openBuildTerminal(projectPath, "buildflatpak.sh", output)
}

// ============================================================================
// ฟังชั้น build เป็น install flatpak
// ============================================================================
func runScripinstallflatpak(projectPath string, output *widget.Entry) {
	openBuildTerminal(projectPath, "buildinstall.sh", output)
}

// ============================================================================
// Appimagetool
// ============================================================================
// ============================================================================
// .image
// ============================================================================
func copyAppImageTool(projectPath string) error {
	src := "./appimagetool/appimagetool-x86_64.AppImage"
	dst := filepath.Join(projectPath, "appimagetool-x86_64.AppImage")

	// ถ้ามีอยู่แล้ว → ไม่ต้อง copy
	if _, err := os.Stat(dst); err == nil {
		return nil
	}

	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, data, 0755)
}

// ============================================================================
// build image
// ============================================================================
func packimage(projectPath string, output *widget.Entry) {
	openBuildTerminal(projectPath, "buildimage.sh", output)
}

// test
func showMsg(msg string) {
	// ใช้ dialog ถ้าต้องการ popup
	// dialog.ShowInformation("แจ้งเตือน", msg, w)
	// แต่ตัวอย่างนี้ขอใช้ print
	println(msg)
}

// ============================================================================
// EXE
// ============================================================================
// ============================================================================
// ฟังชั้น build Scriptbuild EXE
// ============================================================================
func buildexe(projectPath string, output *widget.Entry) {
	openBuildTerminal(projectPath, "buildexe.sh", output)
}

// ============================================================================
// EXE
// ============================================================================
// ============================================================================
// ฟังชั้น build Scriptbuild EXE
// ============================================================================
func clearFile(projectPath string, output *widget.Entry) {
	openBuildTerminal(projectPath, "clear.sh", output)
}
