// Copyright (c) 2026 Nawakarit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package main

import (
	"embed"
	_ "embed"
	"os"
	"path/filepath"
	"sort"
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

	//exe
	CompanyName string
	Fileversion string
	Years       string
	Licenseexe  string
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
	/*
		catmenu := container.NewVScroll(widget.NewCheckGroup(
			[]string{
				"AudioVideo",
				"Audio",
				"Video",
				"Player",
				"Recorder",
				"Mixer",
				"Development",
				"IDE",
				"Debugger",
				"VersionControl",
				"Graphics",
				"2DGraphics",
				"RasterGraphics",
				"VectorGraphics",
				"Photography",
				"Network",
				"WebBrowser",
				"Email",
				"IRCClient",
				"FileTransfer",
				"Utility",
				"Office",
				"Game",
				"System"},

			func(selected []string) {
				if len(selected) == 0 {
					categories.SetText("ยังไม่ได้เลือก")
					return
				}
				categories.SetText(strings.Join(selected, ";") + ";")
			},
		),
		)
	*/
	// เก็บ state
	selected := make(map[string]bool)

	// parent map
	parentMap := map[string]string{
		"Audio":    "AudioVideo",
		"Video":    "AudioVideo",
		"Player":   "AudioVideo",
		"Recorder": "AudioVideo",
		"Mixer":    "AudioVideo",

		"IDE":            "Development",
		"Debugger":       "Development",
		"VersionControl": "Development",

		"2DGraphics":     "Graphics",
		"RasterGraphics": "Graphics",
		"VectorGraphics": "Graphics",
		"Photography":    "Graphics",

		"WebBrowser":   "Network",
		"Email":        "Network",
		"IRCClient":    "Network",
		"FileTransfer": "Network",
	}

	// อัปเดต output
	updateCategories := func() {
		set := make(map[string]bool)

		for k, v := range selected {
			if v {
				set[k] = true
				if p, ok := parentMap[k]; ok {
					set[p] = true
				}
			}
		}

		if len(set) == 0 {
			categories.SetText("")
			return
		}

		var result []string
		for k := range set {
			result = append(result, k)
		}

		sort.Strings(result)
		categories.SetText(strings.Join(result, ";") + ";")
	}

	// สร้าง checkbox
	makeCheck := func(name string) *widget.Check {
		return widget.NewCheck(name, func(checked bool) {
			selected[name] = checked
			updateCategories()
		})
	}

	summary := widget.NewEntry()
	summary.SetPlaceHolder("*Short summary - คุณบัติของแอพ")

	description := widget.NewMultiLineEntry()

	description.SetPlaceHolder("*Description - รายละเอียดของแอพ")
	//description.SetMinRowsVisible(6)

	developer := widget.NewEntry()
	developer.SetPlaceHolder("by Your name")

	date := widget.NewEntry()
	date.SetPlaceHolder("YYYY-MM-DD")

	timeEntry := widget.NewEntry()
	timeEntry.SetPlaceHolder("HH:MM")

	version := widget.NewEntry()
	version.SetPlaceHolder("*Version เช่น 1.0.0 (Dot)")

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

	//exe
	companyName := widget.NewEntry()
	companyName.SetText("Nawakarit")
	companyName.SetPlaceHolder("*ชื่อบริษัท")

	fileversion := widget.NewEntry()
	fileversion.SetText("1,1,1,1")
	fileversion.SetPlaceHolder("*version (exe) เช่น 1,1,1,1 (Comma)")

	years := widget.NewEntry()
	years.SetPlaceHolder("*20XX")
	years1 := container.NewGridWrap(fyne.NewSize(100, 35), years)

	month := widget.NewEntry()
	month.SetPlaceHolder("*01-12")
	month1 := container.NewGridWrap(fyne.NewSize(80, 35), month)

	days := widget.NewEntry()
	days.SetPlaceHolder("*01-31")
	days1 := container.NewGridWrap(fyne.NewSize(80, 35), days)

	licenseexe := widget.NewEntry()
	licenseexe.SetText("GNU General Public License v3.0")
	licenseexe.SetPlaceHolder("*ใส่ประเภท license *ถ้าต้องการ")

	// log box
	logBox := widget.NewMultiLineEntry()
	logBox.SetPlaceHolder("Logs will appear here...")
	logBox.Wrapping = fyne.TextWrapWord

	// ============================================================================
	// เลือกแฟ้มเป้าหมาย
	// ============================================================================
	projectPath := ""
	labelSelectProject := widget.NewLabel("🔴️ เลือกโฟลเดอร์")
	//labelSelectProject.SetText("✅️ เลือกโฟลเดอร์")
	logSelectProject := widget.NewEntry()

	selectBtn := widget.NewButton("Select Project", func() {
		g := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri == nil {
				return
			}
			projectPath = uri.Path()

			go func() {
				cfg := AppConfig{}
				generateFile("templates/clear/clear.tmpl",
					filepath.Join(projectPath, "clear.sh"), cfg) //

				fyne.Do(func() { // 🚪 เข้าประตู UI
					labelSelectProject.SetText("✅️ เลือกโฟลเดอร์")
					logBox.SetText(projectPath)
					logSelectProject.SetText(projectPath)

				})
			}()
		}, w)

		g.Resize(fyne.NewSize(800, 600))
		g.Show()
	})
	// ============================================================================
	// Generate scrip Icons
	// ============================================================================
	labelScripIcons := widget.NewLabel("🔴️ Scrip Icons")

	genscripiconsBtn := widget.NewButton("Scrip Icons", func() {
		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}
		cfg := AppConfig{}
		generateFile("templates/tmp_icons/buildicons.tmpl",
			filepath.Join(projectPath, "buildicons.sh"), cfg) //เอา scrip build ออกมาไว้นอกแฟ้ม flatpak
		logBox.SetText("✅️ Generated File - - buildicons - -")
		labelScripIcons.SetText("✅️ Scrip Icons")
	})
	// ============================================================================
	// Build Icons **ใช้ imagemagick
	// ============================================================================
	labelBuildIcons := widget.NewLabel("🔴️ Icons")

	buildIconsBtn := widget.NewButton("Build Icons", func() {

		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}

		//  run script
		go runScriptbuildIcons(projectPath, logBox)

		logBox.SetText("✅️ Build started in terminal...")
		labelBuildIcons.SetText("✅️ Icons")
	})
	// ============================================================================
	// test ด่วน
	// ============================================================================
	exBtn := widget.NewButton("Ex.", func() {
		//	go func() {
		//	fyne.DoAndWait(func() {

		name.SetText("Music_Player")
		appID.SetText("com.xxx.Music_Player")
		command.SetText("Music_Player")
		categories.SetText("AudioVideo;Audio;")
		summary.SetText("Test Music_Player")
		description.SetText("test and Music_Player")
		developer.SetText("nawakarit")
		version.SetText("1.1.1")
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
		//exe
		companyName.SetText("Nawakarit")
		licenseexe.SetText("GNU General Public License v3.0")
		fileversion.SetText("1,1,1,1")
		days.SetText("1")
		month.SetText("11")
		years.SetText("2026")

		logBox.SetText("✅️ Example now")

	})
	//	}()
	//	})
	// ============================================================================
	// Reset
	// ============================================================================
	resetBtn := widget.NewButton("Reset", func() {
		//	go func() {
		//		fyne.DoAndWait(func() {

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
		//exe
		companyName.SetText("")
		licenseexe.SetText("")
		fileversion.SetText("")
		days.SetText("")
		month.SetText("")
		years.SetText("")

		logBox.SetText("✅️ Reset")
	})
	//	}()
	//	})

	// ============================================================================
	// AppimageTool
	// ============================================================================
	// ============================================================================
	// coppy image master to project
	// ============================================================================
	labelCoppyimage := widget.NewLabel("🔴️ Coppy Appimage")

	coppyimagebtn := widget.NewButton("Coppy Appimage", func() {
		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}
		go func() { //จะไม่ใช้ go func ก็ได้ //ทำเบื่องหลังเร็วขึ้น*มั้ง
			copyAppImageTool(projectPath)
			fyne.Do(func() { //กันฟังชั้น go ทำ ui พัง
				logBox.SetText("✅️ Coppy Appimage")
				labelCoppyimage.SetText("✅️ Coppy Appimage")
			})
		}()
	})
	// ============================================================================
	// Generate scrip Appimage Btn
	// ============================================================================
	labelScripAppimage := widget.NewLabel("🔴️ Scrip Appimage")

	scripimageBtn := widget.NewButton("Scrip Appimage", func() {
		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}
		go func() {
			cfg := AppConfig{
				Name:       name.Text,
				Command:    command.Text,
				Categories: categories.Text,
			}
			generateFile("templates/tmp_image/buildimage.tmpl",
				filepath.Join(projectPath, "buildimage.sh"), cfg) //เอา scrip build appimage ออกมาไว้นอกแฟ้ม flatpak
			fyne.Do(func() {
				logBox.SetText("✅️ Scrip Appimage")
				labelScripAppimage.SetText("✅️ Scrip Appimage")
			})
		}()
	})
	// ============================================================================
	// pack Appimage
	// ============================================================================
	labelpackimage := widget.NewLabel("🔴️ Pack Image")

	packimageBtn := widget.NewButton("Pack Image", func() {

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

		go func() {
			packimage(projectPath, logBox)
			fyne.Do(func() { //
				logBox.SetText("✅️ Pack Image started in terminal...")
				labelpackimage.SetText("✅️ Pack Image")
			})
		}()
	})
	// ============================================================================
	// Flatpak
	// ============================================================================
	// ============================================================================
	// Generate scrip flatpak Btn
	// ============================================================================
	labelGeneratescripflatpak := widget.NewLabel("🔴️ Scrip flatpak")

	genscripflatpakBtn := widget.NewButton("Generate scrip flatpak", func() {

		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}
		go func() {
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
			fyne.Do(func() { //
				logBox.SetText("✅️ Scrip flatpak")
				labelGeneratescripflatpak.SetText("✅️ Scrip flatpak")
			})
		}()
	})

	// ============================================================================
	// ปุ่ม Build flatpak
	// ============================================================================
	labelPackFlatpak := widget.NewLabel("🔴️ Pack Flatpak")

	buildflatpakBtn := widget.NewButton("Pack Flatpak", func() {

		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}

		go func() {
			runScriptbuildflatpak(projectPath, logBox)

			fyne.Do(func() { // เข้าประตู UI
				logBox.SetText("✅️ Pack Flatpak started in terminal...")
				labelPackFlatpak.SetText("✅️ Pack Flatpak")
			})
		}()
	})

	// ============================================================================
	// ปุ่ม Install
	// ============================================================================
	labelInstallFlatpak := widget.NewLabel("🔴️ ติดตั้ง Flatpak")

	installBtn := widget.NewButton("Install Flatpak", func() {

		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}
		go func() {
			runScripinstallflatpak(projectPath, logBox)
			fyne.Do(func() { //
				logBox.SetText("✅️ ติดตั้ง Flatpak started in terminal...")
				labelInstallFlatpak.SetText("✅️ ติดตั้ง Flatpak")
			})
		}()
	})

	// ============================================================================
	// ปุ่มเพิ่มวัน เวลา
	// ============================================================================
	labelTime := widget.NewLabel("🔴️ เวลาปัจจุบัน")
	dates := widget.NewLabel("")

	nowBtn := widget.NewButton("เวลาปัจจุบัน", func() {
		now := time.Now()

		date.SetText(now.Format("2006-01-02"))
		dates.SetText(now.Format("เวลา 15:04\nวันที่ 02/01/2006"))

		timeEntry.SetText(now.Format("15:04"))
		labelTime.SetText("✅️ เวลาปัจจุบัน")

		years.SetText(now.Format("2006")) //ปี
		month.SetText(now.Format("01"))   //เดือน
		days.SetText(now.Format("02"))    //วัน
	})

	// ============================================================================
	// Generate scrip EXE Btn
	// ============================================================================
	labelscripEXE := widget.NewLabel("🔴️ scrip EXE")

	genscripexeBtn := widget.NewButton("Generate scrip EXE", func() {

		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}
		go func() {
			cfg := AppConfig{
				Name:    name.Text,
				AppID:   appID.Text,
				Version: version.Text,
				//exe
				CompanyName: companyName.Text,
				Fileversion: fileversion.Text,
				Years:       years.Text,
				Licenseexe:  licenseexe.Text,
			}
			generateFile("templates/tmp_exe/app.rc.tmpl",
				filepath.Join(projectPath, "app.rc"), cfg) //เอา scrip build ออกมาไว้นอกแฟ้ม

			generateFile("templates/tmp_exe/buildexe.tmpl",
				filepath.Join(projectPath, "buildexe.sh"), cfg)

			generateFile("templates/tmp_exe/FyneApp.toml.tmpl",
				filepath.Join(projectPath, "FyneApp.toml"), cfg)

			fyne.Do(func() { //
				logBox.SetText("✅ Generated scrip exe")
				labelscripEXE.SetText("✅️ scrip EXE")
			})
		}()
	})
	// ============================================================================
	// ปุ่ม Build EXE
	// ============================================================================
	labelBuildEXE := widget.NewLabel("🔴️ Build EXE")

	buildexe := widget.NewButton("Build EXE", func() {

		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}
		go func() {
			buildexe(projectPath, logBox)
			fyne.Do(func() { //
				logBox.SetText("✅ Build started in terminal...")
				labelBuildEXE.SetText("✅️ Build EXE")
			})
		}()
	})
	// ============================================================================
	// ลบไฟล์ build ทั้งหมด
	// ============================================================================
	labelClear := widget.NewLabel("🔴️ Clear")

	clearBtn := widget.NewButton("Clear", func() {
		if projectPath == "" {
			logBox.SetText("🔴️ โปรดเลือกโฟลเดอร์โปรเจค")
			return
		}
		go func() {
			clearFile(projectPath, logBox)
			fyne.Do(func() { //
				logBox.SetText("✅ Clear started in terminal...")
				labelClear.SetText("✅️ Clear")
			})
		}()
	})
	// ============================================================================
	// จัดหน้า..มัน *****************************************************************************************
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

	// ============================================================================
	// Check categories
	// ============================================================================
	catmenu := container.NewVScroll(container.NewVBox(

		widget.NewLabel("🎵 Audio / Video"),
		//makeCheck("AudioVideo"),
		makeCheck("Audio"),
		makeCheck("Video"),
		makeCheck("Player"),
		makeCheck("Recorder"),
		makeCheck("Mixer"),

		widget.NewSeparator(),

		widget.NewLabel("💻 Development"),
		//makeCheck("Development"),
		makeCheck("IDE"),
		makeCheck("Debugger"),
		makeCheck("VersionControl"),

		widget.NewSeparator(),

		widget.NewLabel("🎨 Graphics"),
		//makeCheck("Graphics"),
		makeCheck("2DGraphics"),
		makeCheck("RasterGraphics"),
		makeCheck("VectorGraphics"),
		makeCheck("Photography"),

		widget.NewSeparator(),

		widget.NewLabel("🌐 Network"),
		//makeCheck("Network"),
		makeCheck("WebBrowser"),
		makeCheck("Email"),
		makeCheck("IRCClient"),
		makeCheck("FileTransfer"),

		widget.NewSeparator(),

		widget.NewLabel("📦 Other"),
		makeCheck("Utility"),
		makeCheck("Office"),
		makeCheck("Game"),
		makeCheck("System"),
	))

	// ============================================================================
	// Appimage
	// ============================================================================
	// ปุ่ม .image
	btnimage := widget.NewButton("Image", func() {
		image := container.NewScroll(
			container.NewVBox(
				name, command,
				categories,
				container.NewGridWithColumns(2,
					container.NewVScroll(catmenu),
					container.NewVBox(coppyimagebtn, scripimageBtn, packimageBtn)),
			))
		setContent(container.NewBorder(
			widget.NewLabel("AppimageTool"),
			nil,
			nil,
			nil,
			image,
		))
	})
	// ============================================================================
	// Flatpak
	// ============================================================================
	// ปุ่ม btnflatpak
	btnflatpak := widget.NewButton("Flatpak", func() {
		flatpak := container.NewScroll(
			container.NewVBox(
				container.NewGridWithColumns(2,
					container.NewVBox(
						name,
						command,
						nameRepo,
						appID,
						owner,
						developer,
						version,
						summary,
						description,
						desUpdate1,
						desUpdate2,
						desUpdate3,
					),
					container.NewBorder(categories, nil, nil, nil, catmenu),
				),
				container.NewGridWithColumns(3, date, timeEntry, nowBtn),
				namePix1,
				container.NewGridWithColumns(2, namePix2, genscripflatpakBtn),
				container.NewGridWithColumns(2, namePix3, container.NewCenter(widget.NewLabel("ตรวจเช็คไฟล์ XML ก่อน"))),
				container.NewGridWithColumns(2, namePix4, buildflatpakBtn),
				container.NewGridWithColumns(2, namePix5, installBtn),
			))
		setContent(container.NewBorder(
			widget.NewLabel("Flatpak"),
			nil,
			nil,
			nil,
			flatpak,
		))
	})

	// ============================================================================
	// EXE
	// ============================================================================
	// ปุ่ม EXE
	btnEXE := widget.NewButton("EXE", func() {
		EXE := container.NewScroll(
			container.NewVBox(
				name,
				appID,
				companyName,
				licenseexe,
				version,
				fileversion,
				container.NewCenter(container.NewHBox(widget.NewLabel("วันที่ "), days1, widget.NewLabel("เดือน "), month1, widget.NewLabel("ปี "), years1)),
				nowBtn,
				genscripexeBtn,
				buildexe,
			))
		setContent(container.NewBorder(
			widget.NewLabel("EXE"),
			nil,
			nil,
			nil,
			EXE,
		))
	})

	// ============================================================================
	// เมนูหลัก
	// ============================================================================
	// เมนูด้านซ้าย
	// ซ้ายย่อย //label*
	labellist := container.NewVBox(labelSelectProject, logSelectProject)

	labelicons := container.NewHBox(labelScripIcons, labelBuildIcons)

	labelappimage := container.NewVBox(
		container.NewHBox(labelCoppyimage, labelScripAppimage),
		container.NewHBox(labelpackimage))

	labelflatpak := container.NewVBox(
		container.NewHBox(labelGeneratescripflatpak, labelPackFlatpak),
		labelInstallFlatpak)

	labeldatetimenow := container.NewVBox(
		container.NewHBox(labelTime, dates))

	labelEXE := container.NewVBox(
		container.NewHBox(labelscripEXE, labelBuildEXE),
	)

	// ซ้ายหลัก
	leftMenu := container.NewBorder(
		container.NewVBox(
			widget.NewLabelWithStyle("เมนูหลัก", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			selectBtn,
			container.NewGridWithColumns(2, genscripiconsBtn, buildIconsBtn),
			container.NewGridWithColumns(2, exBtn, resetBtn),
			container.NewGridWithColumns(2, btnimage, btnflatpak),
			btnEXE,
			logBox,
		),
		//container.NewGridWrap(fyne.NewSize(250, 40), widget.NewButton("🔴️ ออก", func() { a.Quit() })),
		container.NewGridWithColumns(2,
			clearBtn,
			widget.NewButton("🔴️ ออก", func() { a.Quit() })),

		nil,
		nil,
		container.NewVScroll(
			container.NewVBox(
				widget.NewCard("List", "", labellist),
				widget.NewCard("Icons", "", labelicons),
				widget.NewCard("AppimageTool", "", labelappimage),
				widget.NewCard("Flatpak", "", labelflatpak),
				widget.NewCard("Time", "", labeldatetimenow),
				widget.NewCard("EXE", "", labelEXE),
				widget.NewCard("Clear", "", labelClear),
			)),
	)

	// จัด layout แบบ Border (ซ้าย : ขวา) จัดหน้ามัน
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
