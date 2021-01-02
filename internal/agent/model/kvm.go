package model

import "encoding/xml"

type KvmDomain struct {
	XMLName       xml.Name `xml:"domain"`
	Text          string   `xml:",chardata"`
	Type          string   `xml:"type,attr"`
	Name          string   `xml:"name"`
	Uuid          string   `xml:"uuid"`
	Memory        string   `xml:"memory"`
	CurrentMemory string   `xml:"currentMemory"`
	Vcpu          string   `xml:"vcpu"`
	KvmOs         struct {
		Text    string `xml:",chardata"`
		KvmType struct {
			Text    string `xml:",chardata"`
			Arch    string `xml:"arch,attr"`
			Machine string `xml:"machine,attr"`
		} `xml:"type"`
		KvmBoot []struct {
			Text string `xml:",chardata"`
			Dev  string `xml:"dev,attr"`
		} `xml:"boot"`
		KvmBootmenu struct {
			Text   string `xml:",chardata"`
			Enable string `xml:"enable,attr"`
		} `xml:"bootmenu"`
	} `xml:"os"`
	KvmFeatures struct {
		Text string `xml:",chardata"`
		Acpi string `xml:"acpi"`
		Apic string `xml:"apic"`
		Pae  string `xml:"pae"`
	} `xml:"features"`
	KvmClock struct {
		Text   string `xml:",chardata"`
		Offset string `xml:"offset,attr"`
	} `xml:"clock"`
	OnPoweroff string `xml:"on_poweroff"`
	OnReboot   string `xml:"on_reboot"`
	OnCrash    string `xml:"on_crash"`
	KvmDevices struct {
		Text     string `xml:",chardata"`
		Emulator string `xml:"emulator"`
		KvmDisk  []struct {
			Text      string `xml:",chardata"`
			Type      string `xml:"type,attr"`
			Device    string `xml:"device,attr"`
			KvmDriver struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
				Type string `xml:"type,attr"`
			} `xml:"driver"`
			KvmSource struct {
				Text string `xml:",chardata"`
				File string `xml:"file,attr"`
			} `xml:"source"`
			KvmTarget struct {
				Text string `xml:",chardata"`
				Dev  string `xml:"dev,attr"`
				Bus  string `xml:"bus,attr"`
			} `xml:"target"`
			KvmAlias struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"alias"`
			Readonly string `xml:"readonly"`
		} `xml:"disk"`
		KvmController []struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Index      string `xml:"index,attr"`
			KvmAddress struct {
				Text     string `xml:",chardata"`
				Type     string `xml:"type,attr"`
				Domain   string `xml:"domain,attr"`
				Bus      string `xml:"bus,attr"`
				Slot     string `xml:"slot,attr"`
				Function string `xml:"function,attr"`
			} `xml:"address"`
		} `xml:"controller"`
		KvmInterface struct {
			Text   string `xml:",chardata"`
			Type   string `xml:"type,attr"`
			KvmMac struct {
				Text    string `xml:",chardata"`
				Address string `xml:"address,attr"`
			} `xml:"mac"`
			KvmSource struct {
				Text   string `xml:",chardata"`
				Bridge string `xml:"bridge,attr"`
			} `xml:"source"`
			KvmTarget struct {
				Text string `xml:",chardata"`
				Dev  string `xml:"dev,attr"`
			} `xml:"target"`
			KvmAlias struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"alias"`
			KvmAddress struct {
				Text     string `xml:",chardata"`
				Type     string `xml:"type,attr"`
				Domain   string `xml:"domain,attr"`
				Bus      string `xml:"bus,attr"`
				Slot     string `xml:"slot,attr"`
				Function string `xml:"function,attr"`
			} `xml:"address"`
		} `xml:"interface"`
		KvmSerial struct {
			Text      string `xml:",chardata"`
			Type      string `xml:"type,attr"`
			KvmTarget struct {
				Text string `xml:",chardata"`
				Port string `xml:"port,attr"`
			} `xml:"target"`
		} `xml:"serial"`
		KvmConsole struct {
			Text      string `xml:",chardata"`
			Type      string `xml:"type,attr"`
			KvmTarget struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Port string `xml:"port,attr"`
			} `xml:"target"`
		} `xml:"console"`
		KvmInput []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Bus  string `xml:"bus,attr"`
		} `xml:"input"`
		KvmGraphics struct {
			Text       string `xml:",chardata"`
			Type       string `xml:"type,attr"`
			Port       string `xml:"port,attr"`
			Autoport   string `xml:"autoport,attr"`
			AttrListen string `xml:"listen,attr"`
			Keymap     string `xml:"keymap,attr"`
			KvmListen  struct {
				Text    string `xml:",chardata"`
				Type    string `xml:"type,attr"`
				Address string `xml:"address,attr"`
			} `xml:"listen"`
		} `xml:"graphics"`
		KvmSound struct {
			Text       string `xml:",chardata"`
			Model      string `xml:"model,attr"`
			KvmAddress struct {
				Text     string `xml:",chardata"`
				Type     string `xml:"type,attr"`
				Domain   string `xml:"domain,attr"`
				Bus      string `xml:"bus,attr"`
				Slot     string `xml:"slot,attr"`
				Function string `xml:"function,attr"`
			} `xml:"address"`
		} `xml:"sound"`
		KvmVideo struct {
			Text     string `xml:",chardata"`
			KvmModel struct {
				Text  string `xml:",chardata"`
				Type  string `xml:"type,attr"`
				Vram  string `xml:"vram,attr"`
				Heads string `xml:"heads,attr"`
			} `xml:"model"`
			KvmAddress struct {
				Text     string `xml:",chardata"`
				Type     string `xml:"type,attr"`
				Domain   string `xml:"domain,attr"`
				Bus      string `xml:"bus,attr"`
				Slot     string `xml:"slot,attr"`
				Function string `xml:"function,attr"`
			} `xml:"address"`
		} `xml:"video"`
		KvmMemballoon struct {
			Text       string `xml:",chardata"`
			Model      string `xml:"model,attr"`
			KvmAddress struct {
				Text     string `xml:",chardata"`
				Type     string `xml:"type,attr"`
				Domain   string `xml:"domain,attr"`
				Bus      string `xml:"bus,attr"`
				Slot     string `xml:"slot,attr"`
				Function string `xml:"function,attr"`
			} `xml:"address"`
		} `xml:"memballoon"`
	} `xml:"devices"`
}
