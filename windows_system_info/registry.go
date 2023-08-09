package main

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

// see: https://www.cnblogs.com/woshimrf/p/language-code-lcid.html
var LangCode map[string]string = map[string]string{
	"0401": "Arabic - Saudi Arabia",
	"0402": "Bulgarian",
	"0403": "Catalan",
	"0404": "Chinese - Taiwan",
	"0405": "Czech",
	"0406": "Danish",
	"0407": "German - Germany",
	"0408": "Greek",
	"0409": "English - United States",
	"0410": "Italian - Italy",
	"0411": "Japanese",
	"0412": "Korean",
	"0413": "Dutch - Netherlands",
	"0414": "Norwegian - Bokml",
	"0415": "Polish",
	"0416": "Portuguese - Brazil",
	"0417": "Raeto-Romance",
	"0418": "Romanian - Romania",
	"0419": "Russian",
	"0420": "Urdu",
	"0421": "Indonesian",
	"0422": "Ukrainian",
	"0423": "Belarusian",
	"0424": "Slovenian",
	"0425": "Estonian",
	"0426": "Latvian",
	"0427": "Lithuanian",
	"0428": "Tajik",
	"0429": "Farsi - Persian",
	"0430": "Sesotho (Sutu)",
	"0431": "Tsonga",
	"0432": "Setsuana",
	"0433": "Venda",
	"0434": "Xhosa",
	"0435": "Zulu",
	"0436": "Afrikaans",
	"0437": "Georgian",
	"0438": "Faroese",
	"0439": "Hindi",
	"0440": "Kyrgyz - Cyrillic",
	"0441": "Swahili",
	"0442": "Turkmen",
	"0443": "Uzbek - Latin",
	"0444": "Tatar",
	"0445": "Bengali - India",
	"0446": "Punjabi",
	"0447": "Gujarati",
	"0448": "Oriya",
	"0449": "Tamil",
	"0450": "Mongolian",
	"0451": "Tibetan",
	"0452": "Welsh",
	"0453": "Khmer",
	"0454": "Lao",
	"0455": "Burmese",
	"0456": "Galician",
	"0457": "Konkani",
	"0458": "Manipuri",
	"0459": "Sindhi",
	"0460": "Kashmiri",
	"0461": "Nepali",
	"0462": "Frisian - Netherlands",
	"0464": "Filipino",
	"0466": "Edo",
	"0470": "Igbo - Nigeria",
	"0474": "Guarani - Paraguay",
	"0476": "Latin",
	"0477": "Somali",
	"0481": "Maori",
	"0801": "Arabic - Iraq",
	"0804": "Chinese - China",
	"0807": "German - Switzerland",
	"0809": "English - Great Britain",
	"0810": "Italian - Switzerland",
	"0813": "Dutch - Belgium",
	"0814": "Norwegian - Nynorsk",
	"0816": "Portuguese - Portugal",
	"0818": "Romanian - Moldova",
	"0819": "Russian - Moldova",
	"0843": "Uzbek - Cyrillic",
	"0845": "Bengali - Bangladesh",
	"0850": "Mongolian",
	"1001": "Arabic - Libya",
	"1004": "Chinese - Singapore",
	"1007": "German - Luxembourg",
	"1009": "English - Canada",
	"1115": "Sinhala",
	"1401": "Arabic - Algeria",
	"1404": "Chinese - Macau SAR",
	"1407": "German - Liechtenstein",
	"1409": "English - New Zealand",
	"1801": "Arabic - Morocco",
	"1809": "English - Ireland",
	"2001": "Arabic - Oman",
	"2009": "English - Jamaica",
	"2401": "Arabic - Yemen",
	"2409": "English - Caribbean",
	"2801": "Arabic - Syria",
	"2809": "English - Belize",
	"3001": "Arabic - Lebanon",
	"3009": "English - Zimbabwe",
	"3401": "Arabic - Kuwait",
	"3409": "English - Phillippines",
	"3801": "Arabic - United Arab Emirates",
	"4001": "Arabic - Qatar",
	"4009": "English - India",
}

// 系统默认语言
func GetSystemLanguage() string {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Nls\Language`, registry.READ)
	if err != nil {
		// fmt.Printf("get lang key failed: %s\n", err)
		return ""
	}
	defer key.Close()

	v, _, err := key.GetStringValue("Default")
	if err != nil {
		// fmt.Printf("get lang value failed: %s\n", err)
		return ""
	}
	// fmt.Printf("lang %s (type %d)\n", v, vt)

	if lang, ok := LangCode[v]; ok {
		return lang
	} else {
		return ""
	}
}

// 获取主板信息
func GetBiosStat() string {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\DESCRIPTION\System\BIOS`, registry.READ)
	if err != nil {
		fmt.Printf("open key failed: %s\n", err)
		return ""
	}
	defer key.Close()

	info, err := key.Stat()
	if err != nil {
		fmt.Printf("get key stat failed: %s\n", err)
		return ""
	}
	// fmt.Printf("key info %v\n", info)

	valueName, err := key.ReadValueNames(int(info.ValueCount))
	if err != nil {
		fmt.Printf("get value name failed: %s\n", err)
		return ""
	}
	for _, name := range valueName {
		v, _, err := key.GetStringValue(name)
		if err == nil {
			fmt.Println(name, v)
		}
	}

	return ""
}
