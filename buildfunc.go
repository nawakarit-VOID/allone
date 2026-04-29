// Copyright (c) 2026 Nawakarit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"fyne.io/fyne/v2/widget"
)

// ============================================================================
// ฟังชั้น build Icons
// ============================================================================
func runScriptbuildIcons(projectPath string, output *widget.Entry) {

	commands := [][]string{ //ใช้ imagemagick
		{"gnome-terminal", "--", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildicons.sh && ./buildicons.sh; exec bash"},
		{"x-terminal-emulator", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildicons.sh && ./buildicons.sh; exec bash"},
		{"konsole", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildicons.sh && ./buildicons.sh; exec bash"},
		{"xfce4-terminal", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildicons.sh && ./buildicons.sh; exec bash"},
	}

	for _, c := range commands {
		cmd := exec.Command(c[0], c[1:]...)
		err := cmd.Start()
		if err == nil {
			output.SetText("✅️ opened terminal: " + c[0])
			return
		}
	}

	output.SetText("🔴️ no terminal found")
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

	commands := [][]string{
		{"gnome-terminal", "--", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildflatpak.sh && ./buildflatpak.sh; exec bash"},
		{"x-terminal-emulator", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildflatpak.sh && ./buildflatpak.sh; exec bash"},
		{"konsole", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildflatpak.sh && ./buildflatpak.sh; exec bash"},
		{"xfce4-terminal", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildflatpak.sh && ./buildflatpak.sh; exec bash"},
	}

	for _, c := range commands {
		cmd := exec.Command(c[0], c[1:]...)
		err := cmd.Start()
		if err == nil {
			output.SetText("✅️ opened terminal: " + c[0])
			return
		}
	}

	output.SetText("🔴️ no terminal found")
}

// ============================================================================
// ฟังชั้น build เป็น install flatpak
// ============================================================================
func runScripinstallflatpak(projectPath string, output *widget.Entry) {

	commands := [][]string{
		{"gnome-terminal", "--", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildinstall.sh && ./buildinstall.sh; exec bash"},
		{"x-terminal-emulator", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildinstall.sh && ./buildinstall.sh; exec bash"},
		{"konsole", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildinstall.sh && ./buildinstall.sh; exec bash"},
		{"xfce4-terminal", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildinstall.sh && ./buildinstall.sh; exec bash"},
	}

	for _, c := range commands {
		cmd := exec.Command(c[0], c[1:]...)
		err := cmd.Start()
		if err == nil {
			output.SetText("✅️ opened terminal: " + c[0])
			return
		}
	}

	output.SetText("🔴️ no terminal found")
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

	commands := [][]string{

		{"gnome-terminal", "--", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildimage.sh && ./buildimage.sh; exec bash"},
		{"x-terminal-emulator", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildimage.sh && ./buildimage.sh; exec bash"},
		{"konsole", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildimage.sh && ./buildimage.sh; exec bash"},
		{"xfce4-terminal", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildimage.sh && ./buildimage.sh; exec bash"},
	}

	for _, c := range commands {
		cmd := exec.Command(c[0], c[1:]...)
		err := cmd.Start()
		if err == nil {
			output.SetText("✅️ opened terminal: " + c[0])
			return
		}
	}

	output.SetText("🔴️ no terminal found")
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

	commands := [][]string{
		{"gnome-terminal", "--", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildexe.sh && ./buildexe.sh; exec bash"},
		{"x-terminal-emulator", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildexe.sh && ./buildexe.sh; exec bash"},
		{"konsole", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildexe.sh && ./buildexe.sh; exec bash"},
		{"xfce4-terminal", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildexe.sh && ./buildexe.sh; exec bash"},
	}

	for _, c := range commands {
		cmd := exec.Command(c[0], c[1:]...)
		err := cmd.Start()
		if err == nil {
			output.SetText("✅️ opened terminal: " + c[0])
			return
		}
	}

	output.SetText("🔴️ no terminal found")
}

// ============================================================================
// EXE
// ============================================================================
// ============================================================================
// ฟังชั้น build Scriptbuild EXE
// ============================================================================
func clearFile(projectPath string, output *widget.Entry) {

	commands := [][]string{
		{"gnome-terminal", "--", "bash", "-c", "cd '" + projectPath + "' && chmod +x clear.sh && ./clear.sh; exec bash"},
		{"x-terminal-emulator", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x clear.sh && ./clear.sh; exec bash"},
		{"konsole", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x clear.sh && ./clear.sh; exec bash"},
		{"xfce4-terminal", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x clear.sh && ./clear.sh; exec bash"},
	}

	for _, c := range commands {
		cmd := exec.Command(c[0], c[1:]...)
		err := cmd.Start()
		if err == nil {
			output.SetText("✅️ opened terminal: " + c[0])
			return
		}
	}

	output.SetText(" 🔴️ no terminal found")
}
