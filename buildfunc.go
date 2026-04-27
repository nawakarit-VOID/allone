// Copyright (c) 2026 Nawakarit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package main

import (
	"os"
	"os/exec"
	"text/template"

	"fyne.io/fyne/v2/widget"
)

func showMsg(msg string) {
	// ใช้ dialog ถ้าต้องการ popup
	// dialog.ShowInformation("แจ้งเตือน", msg, w)
	// แต่ตัวอย่างนี้ขอใช้ print
	println(msg)
}

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
			output.SetText("🚀 opened terminal: " + c[0])
			return
		}
	}

	output.SetText("❌ no terminal found")
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
			output.SetText("🚀 opened terminal: " + c[0])
			return
		}
	}

	output.SetText("❌ no terminal found")
}

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
			output.SetText("🚀 opened terminal: " + c[0])
			return
		}
	}

	output.SetText("❌ no terminal found")
}
