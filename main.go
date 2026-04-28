// Copyright (c) 2026 Nawakarit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package main

import (
	"embed"
	_ "embed"
	"os"
	"path/filepath"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type AppConfig struct {
	Name        string
	AppID       string
	Command     string
	Categories  string
	Summary     string
	Description string
	License     string
	Developer   string
	Date        string
	TimeEntry   string
	Version     string
	DesUpdate1  string
	DesUpdate2  string
	DesUpdate3  string
	Owner       string
	NameRepo    string
	NamePix1    string
	NamePix2    string
	NamePix3    string
	NamePix4    string
	NamePix5    string
}

// โหลด icon
func loadIcon(size int) fyne.Resource {
	var file string
	switch {
	case size >= 512:
		file = "icons/icon-512.png" ///ที่อยู่
	case size >= 256:
		file = "icons/icon-256.png"
	case size >= 128:
		file = "icons/icon-128.png"
	default:
		file = "icons/icon-64.png"
	}
	data, _ := iconFS.ReadFile(file)
	return fyne.NewStaticResource(file, data)
}

//go:embed icons/*
var iconFS embed.FS

//go:embed assets/font/Itim-Regular.ttf
var fontItim []byte
var myFont = fyne.NewStaticResource("Itim-Regular.ttf", fontItim)

// ============================================================================
// main
// ============================================================================
func main() {

	a := app.NewWithID("com.nawakarit.allone")
	icons := loadIcon(64) //เอา data มาใช้
	a.SetIcon(icons)
	w := a.NewWindow("allone")
	w.SetIcon(icons)

	a.Settings().SetTheme(&MyTheme{})

	// inputs
	name := widget.NewEntry()
	name.SetPlaceHolder("*App Name")

	appID := widget.NewEntry()
	appID.SetPlaceHolder("*com.example.app")

	command := widget.NewEntry()
	command.SetPlaceHolder("*binary name")

	categories := widget.NewEntry()
	categories.SetPlaceHolder("*Utility;")

	catmenu := widget.NewCheckGroup(
		[]string{
			"Utility",
			"Development",
			"Game",
			"Graphics",
			"Network",
			"Office",
			"Audio",
			"Video",
			"System"},
		func(selected []string) {
			if len(selected) == 0 {
				categories.SetText("ยังไม่ได้เลือก")
				return
			}
			categories.SetText(strings.Join(selected, ";") + ";")
		},
	)

	summary := widget.NewEntry()
	summary.SetPlaceHolder("*Short summary - คุณบัติของแอพ")

	description := widget.NewMultiLineEntry()

	description.SetPlaceHolder("*Description - รายละเอียดของแอพ")
	description.SetMinRowsVisible(6)

	developer := widget.NewEntry()
	developer.SetPlaceHolder("by Your name")

	date := widget.NewEntry()
	date.SetPlaceHolder("YYYY-MM-DD")

	timeEntry := widget.NewEntry()
	timeEntry.SetPlaceHolder("HH:MM")

	version := widget.NewEntry()
	version.SetPlaceHolder("*V เช่น 1.0.0")

	desUpdate1 := widget.NewEntry()
	desUpdate1.SetPlaceHolder("*อัพเดท 1")
	desUpdate2 := widget.NewEntry()
	desUpdate2.SetPlaceHolder("*อัพเดท 2")
	desUpdate3 := widget.NewEntry()
	desUpdate3.SetPlaceHolder("*อัพเดท 3")

	owner := widget.NewEntry()
	owner.SetText("nawakarit-VOID")
	owner.SetPlaceHolder("*ชื่อเจ้าของ Github [Owner]")

	nameRepo := widget.NewEntry()
	nameRepo.SetPlaceHolder("*Repository")

	namePix1 := widget.NewEntry()
	namePix1.SetPlaceHolder("*1.เฉพาะชื่อรูป (png, วางข้าง main)")

	namePix2 := widget.NewEntry()
	namePix2.SetPlaceHolder("*2.เฉพาะชื่อรูป (png, วางข้าง main)")

	namePix3 := widget.NewEntry()
	namePix3.SetPlaceHolder("*3.เฉพาะชื่อรูป (png, วางข้าง main)")

	namePix4 := widget.NewEntry()
	namePix4.SetPlaceHolder("*4.เฉพาะชื่อรูป (png, วางข้าง main)")

	namePix5 := widget.NewEntry()
	namePix5.SetPlaceHolder("*5.เฉพาะชื่อรูป (png, วางข้าง main)")

	// log box
	logBox := widget.NewMultiLineEntry()
	logBox.SetPlaceHolder("Logs will appear here...")
	logBox.Wrapping = fyne.TextWrapWord

	// ============================================================================
	// เลือกแฟ้มเป้าหมาย
	// ============================================================================
	projectPath := ""
	labelSelectProject := widget.NewLabel("🔴️ เลือกโฟลเดอร์")
	logSelectProject := widget.NewEntry()

	selectBtn := widget.NewButton("Select Project", func() {
		g := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri == nil {
				return
			}
			projectPath = uri.Path()

			labelSelectProject.SetText("✅️ เลือกโฟลเดอร์")
			logBox.SetText(projectPath)
			logSelectProject.SetText(projectPath)

		}, w)

		g.Resize(fyne.NewSize(800, 600))
		g.Show()
	})
	// ============================================================================
	// Generate scrip Icons
	// ============================================================================
	genscripiconsBtn := widget.NewButton("scrip Icons", func() {
		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}
		cfg := AppConfig{}
		generateFile("templates/tmp_icons/buildicons.tmpl",
			filepath.Join(projectPath, "buildicons.sh"), cfg) //เอา scrip build ออกมาไว้นอกแฟ้ม flatpak
		logBox.SetText("✅️ Generated File - - buildicons - -")
	})
	// ============================================================================
	// Build Icons **ใช้ imagemagick
	// ============================================================================
	buildIconsBtn := widget.NewButton("Build Icons", func() {

		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}

		//  run script
		go runScriptbuildIcons(projectPath, logBox)

		logBox.SetText("✅️ Build started in terminal...")
	})
	// ============================================================================
	// test ด่วน
	// ============================================================================
	exBtn := widget.NewButton("Ex.", func() {
		name.SetText("Music_Player")
		appID.SetText("com.xxx.Music_Player")
		command.SetText("Music_Player")
		categories.SetText("Utility;Audio;")
		summary.SetText("Test Music_Player")
		description.SetText("test and Music_Player")
		developer.SetText("nawakarit")
		version.SetText("999.999.999")
		desUpdate1.SetText("ad go func")
		desUpdate2.SetText("ad SCR")
		desUpdate3.SetText("ad icons")
		owner.SetText("nawakarit-VOID")
		nameRepo.SetText("Music_Player")
		namePix1.SetText("test_2026-04-06_21-06-09")
		namePix2.SetText("test_2026-04-06_21-07-08")
		namePix3.SetText("test_2026-04-06_21-07-18")
		namePix4.SetText("test_2026-04-06_21-07-08")
		namePix5.SetText("test_2026-04-06_21-06-09")

		logBox.SetText("✅️ Example now")
	})
	// ============================================================================
	// Reset EX
	// ============================================================================
	resetExBtn := widget.NewButton("ResetEx.", func() {
		name.SetText("")
		appID.SetText("com.nawakarit.")
		command.SetText("")
		categories.SetText("")
		summary.SetText("")
		description.SetText("")
		developer.SetText("nawakarit")
		version.SetText("")
		desUpdate1.SetText("")
		desUpdate2.SetText("")
		desUpdate3.SetText("")
		owner.SetText("nawakarit-VOID")
		nameRepo.SetText("")
		namePix1.SetText("")
		namePix2.SetText("")
		namePix3.SetText("")
		namePix4.SetText("")
		namePix5.SetText("")

		logBox.SetText("✅️ Reset example")
	})
	// ============================================================================
	// AppimageTool
	// ============================================================================
	// ============================================================================
	// coppy image master to project
	// ============================================================================
	coppyimagebtn := widget.NewButton("Coppy image", func() {
		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}
		go copyAppImageTool(projectPath)
	})
	// ============================================================================
	// Generate scrip Appimage Btn
	// ============================================================================
	labelScripAppimage := widget.NewLabel("🔴️")
	scripimageBtn := widget.NewButton("Scrip Appimage", func() {
		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}
		cfg := AppConfig{
			Name:       name.Text,
			Command:    command.Text,
			Categories: categories.Text,
		}
		generateFile("templates/tmp_image/buildimage.tmpl",
			filepath.Join(projectPath, "buildimage.sh"), cfg) //เอา scrip build appimage ออกมาไว้นอกแฟ้ม flatpak

		logBox.SetText("✅️ Generated Scrip AppimageTool")
		labelScripAppimage.SetText("✅️")
	})
	// ============================================================================
	// build Appimage
	// ============================================================================
	buildimageBtn := widget.NewButton("Run Build", func() {

		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")

			return
		}

		//  copy appimagetool (ถ้ามีฟังก์ชันนี้)
		err := copyAppImageTool(projectPath)
		if err != nil {
			logBox.SetText("🔴️ คัดลอกล้มเหลว : " + err.Error())
			return
		}

		//  run script
		go runbuildimage(projectPath, logBox)

		logBox.SetText("✅️ Build started in terminal...")
	})
	// ============================================================================
	// Flatpak
	// ============================================================================
	// ============================================================================
	// Generate scrip flatpak Btn
	// ============================================================================
	genscripflatpakBtn := widget.NewButton("Generate scrip Folder", func() {

		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}

		cfg := AppConfig{
			Name:        name.Text,
			AppID:       appID.Text,
			Command:     command.Text,
			Categories:  categories.Text,
			Summary:     summary.Text,
			Description: description.Text,
			License:     "GPL-3.0-or-later",
			Developer:   developer.Text,
			Date:        date.Text,
			TimeEntry:   timeEntry.Text,
			Version:     version.Text,
			DesUpdate1:  desUpdate1.Text,
			DesUpdate2:  desUpdate2.Text,
			DesUpdate3:  desUpdate3.Text,
			Owner:       owner.Text,
			NameRepo:    nameRepo.Text,
			NamePix1:    namePix1.Text,
			NamePix2:    namePix2.Text,
			NamePix3:    namePix3.Text,
			NamePix4:    namePix4.Text,
			NamePix5:    namePix5.Text,
		}

		flatpakPath := projectPath + "/" + "flatpak"
		os.MkdirAll(flatpakPath, 0755)

		generateFile("templates/tmp_flatpak/desktop.tmpl",
			filepath.Join(flatpakPath, cfg.AppID+".desktop"), cfg)

		generateFile("templates/tmp_flatpak/manifest.tmpl",
			filepath.Join(flatpakPath, cfg.AppID+".json"), cfg)

		generateFile("templates/tmp_flatpak/metainfo.tmpl",
			filepath.Join(flatpakPath, cfg.AppID+".metainfo.xml"), cfg)

		generateFile("templates/tmp_flatpak/buildflatpak.tmpl",
			filepath.Join(projectPath, "buildflatpak.sh"), cfg) //เอา scrip build ออกมาไว้นอกแฟ้ม flatpak

		generateFile("templates/tmp_flatpak/buildinstall.tmpl",
			filepath.Join(projectPath, "buildinstall.sh"), cfg)

		logBox.SetText("✅️ Generated File Flatpak - - and - - ✅️ File Scrip Build Flatpak\n")
	})

	// ============================================================================
	// ปุ่ม Build flatpak
	// ============================================================================
	buildflatpakBtn := widget.NewButton("7 - Run Build Flatpak", func() {

		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}

		//  run script
		go runScriptbuildflatpak(projectPath, logBox)

		logBox.SetText("✅️ Build started in terminal...")
	})

	// ============================================================================
	// ปุ่ม Install
	// ============================================================================
	installBtn := widget.NewButton("8 - Install Flatpak", func() {

		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}
		go runScripinstallflatpak(projectPath, logBox)

		logBox.SetText("✅️ Install started in terminal...")
	})

	// ============================================================================
	// ปุ่มเพิ่มวัน เวลา
	// ============================================================================
	nowBtn := widget.NewButton("เวลาปัจจุบัน", func() {
		now := time.Now()

		date.SetText(now.Format("2006-01-02"))
		timeEntry.SetText(now.Format("15:04"))
	})

	// ============================================================================
	// จัดหน้า..มัน
	// ============================================================================
	// สร้างพื้นที่แสดงเนื้อหาหลัก (ด้านขวา)
	contentArea := container.NewStack()

	// ฟังก์ชันเปลี่ยนหน้า
	setContent := func(content fyne.CanvasObject) {
		contentArea.Objects = []fyne.CanvasObject{content}
		contentArea.Refresh()
	}

	// หน้าแรก (ต้อนรับ)
	welcome := widget.NewLabel("")
	setContent(welcome)

	// ปุ่ม .image
	btnimage := widget.NewButton("Image", func() {
		/*	buttons := container.NewGridWithColumns(2,
			widget.NewButton("ตัวเลือก Alpha", func() { showMsg("เลือก Alpha") }),
			widget.NewButton("ตัวเลือก Beta", func() { showMsg("เลือก Beta") }),
			widget.NewButton("ตัวเลือก Gamma", func() { showMsg("เลือก Gamma") }),
			widget.NewButton("ตัวเลือก Delta", func() { showMsg("เลือก Delta") }),
		) */
		setContent(container.NewVBox(
			widget.NewLabel("AppimageTool"),
			name, command,
			categories,
			catmenu,
			//buttons,
			coppyimagebtn, scripimageBtn, buildimageBtn,
		))

	})

	// ปุ่ม btnflatpak
	btnflatpak := widget.NewButton("Flatpak", func() {

		flatpak := container.NewScroll(

			container.NewVBox(

				container.NewGridWithColumns(2, name, appID),
				container.NewGridWithColumns(2, command, developer),
				categories,
				catmenu,
				container.NewGridWithColumns(2, version),
				container.NewGridWithColumns(3, date, timeEntry, nowBtn),
				summary, description,

				container.NewGridWithColumns(3, desUpdate1, desUpdate2, desUpdate3),
				container.NewGridWithColumns(2, owner, nameRepo),
				namePix1,
				namePix2,
				namePix3,
				namePix4,
				namePix5,
				genscripflatpakBtn,

				container.NewCenter(widget.NewLabel("6 - ตรวจเช็คไฟล์ XML ก่อน")),

				buildflatpakBtn, installBtn,
			))
		setContent(container.NewBorder(
			widget.NewLabel("Flatpak"),
			nil,
			nil,
			nil,
			flatpak,
		))
	})

	// เมนูด้านซ้าย
	leftMenu := container.NewBorder(

		container.NewVBox(

			widget.NewLabelWithStyle("เมนูหลัก", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			//widget.NewSeparator(),
			selectBtn,
			container.NewHScroll(logSelectProject),
			container.NewGridWithColumns(2, genscripiconsBtn, buildIconsBtn),
			container.NewGridWithColumns(2, exBtn, resetExBtn),
			//btnA,
			//btnB,
			//btnC,
			//btnD,
			//		labelSelectProject.SetText("✅️"+projectPath),

			btnimage,
			btnflatpak,
			widget.NewLabel("logBox"),
			//widget.NewSeparator(),
			logBox,
			widget.NewLabel("List"),
			labelSelectProject,
			labelScripAppimage,
		),
		container.NewGridWrap(fyne.NewSize(200, 40), widget.NewButton("🔴️ ออก", func() { a.Quit() })),
		nil,
		nil,
		nil,
	)
	// จัด layout แบบ Border (ซ้าย : ขวา)
	// ไม่มีเส้นแบ่ง ไม่มีพื้นหลังแยก
	mainContainer := container.NewBorder(
		nil,
		nil,
		leftMenu,
		nil,
		contentArea)

	w.SetContent(mainContainer)
	w.Resize(fyne.NewSize(850, 850))
	w.ShowAndRun()
}

/*// --- สร้างเมนูด้านซ้าย ---
// ปุ่ม A (แสดงรายละเอียด)
btnA := widget.NewButton("📁 รายละเอียด A", func() {
	detail := widget.NewCard("รายละเอียด A",
		"ข้อมูลเพิ่มเติม",
		widget.NewLabel("นี่คือรายละเอียดของเมนู A\nสามารถเพิ่มข้อความหรือ input ได้"))
	setContent(detail)
})

// ปุ่ม B (มีปุ่มย่อยในเนื้อหา)
btnB := widget.NewButton("⚙️ ตั้งค่า B", func() {
	subBtn1 := widget.NewButton("ตัวเลือกที่ 1", func() {
		widget.NewLabel("เลือก 1")
	})
	subBtn2 := widget.NewButton("ตัวเลือกที่ 2", func() {
		widget.NewLabel("เลือก 2")
	})
	subForm := container.NewVBox(
		widget.NewLabel("เลือกการทำงานเพิ่มเติม:"),
		subBtn1,
		subBtn2,
		widget.NewSeparator(),
		widget.NewLabel("หรือกรอกข้อมูล:"),
		widget.NewEntry(),
	)
	setContent(subForm)
})

// ปุ่ม C (แสดงฟอร์ม)
btnC := widget.NewButton("📝 ฟอร์ม C", func() {
	form := widget.NewForm(
		widget.NewFormItem("ชื่อ", widget.NewEntry()),
		widget.NewFormItem("อีเมล", widget.NewEntry()),
	)
	form.SubmitText = "บันทึก"
	form.OnSubmit = func() {
		setContent(widget.NewLabel("บันทึกสำเร็จ!"))
	}
	setContent(form)
})

// ปุ่ม D (แสดงปุ่มย่อยหลายปุ่ม)
btnD := widget.NewButton("🔘 เมนู D (ปุ่มย่อย)", func() {
	buttons := container.NewGridWithColumns(2,
		widget.NewButton("ตัวเลือก Alpha", func() { showMsg("เลือก Alpha") }),
		widget.NewButton("ตัวเลือก Beta", func() { showMsg("เลือก Beta") }),
		widget.NewButton("ตัวเลือก Gamma", func() { showMsg("เลือก Gamma") }),
		widget.NewButton("ตัวเลือก Delta", func() { showMsg("เลือก Delta") }),
	)
	setContent(container.NewVBox(
		widget.NewLabel("เมนูย่อยของ D:"),
		buttons,
	))
})
*/
