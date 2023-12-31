package main

import (
	"fmt"
	"kalami/libs"
	"os"
	"path"
	"strconv"
	"strings"

	"golang.org/x/sys/windows/registry"
)

var Provinces string = `
北京(BeiJing)
上海(ShangHai)
天津(TianJin)
重庆(ChongQing)
香港(XiangGang)
澳门(Aomen)
安徽(AnHui)
福建(FuJian)
广东(GuangDong)
广西(GuangXi)
贵州(GuiZhou)
甘肃(GanSu)
海南(HaiNan)
河北(HeBei)
河南(HeNan)
黑龙江(HeiLongJiang)
湖北(HuBei)
湖南(HuNan)
吉林(JiLin)
江苏(JiangSu)
江西(JiangXi)
辽宁(LiaoNing)
内蒙古(NeiMengGu)
宁夏(NingXia)
青海(QingHai)
陕西(ShanXi)
山西(ShanXi)
山东(ShanDong)
四川(SiChuan)
台湾(TaiWan)
西藏(XiZang)
新疆(XinJiang)
云南(YunNan)
浙江(ZheJiang)
`

var Cities string = `
香港(XiangGang)
澳门(AoMen)
安庆(AnQing)，亳州(BoZhou)，蚌埠(BangBu)，池州(ChiZhou)，巢湖(ChaoHu)，滁州(ChuZhou)，阜阳(FuYang)，合肥(HeFei)，淮南(HuaiNan)，淮北(HuaiBei)，黄山(HuangShan)，六安(LiuAn)，马鞍山(MaAnShan)，宿州(SuZhou)，铜陵(TongLing)，芜湖(WuHu)，宣城(XuanCheng)
福州(FuZhou)，龙岩(LongYan)，南平(NanPing)，宁德(NingDe)，莆田(PuTian)，泉州(QuanZhou)，三明(SanXing)，厦门(XiaMen)，漳州(ZhangZhou)
潮州(ChaoZhou)，东莞(DongGuan)，佛山(FoShan)，广州(GuangZhou)，河源(HeYuan)，惠州(HuiZhou)，江门(JiangMen)，揭阳(JieYang)，梅州(MeiZhou)，茂名(MaoMing)，清远(QingYuan)，深圳(ShenZhen)，汕头(ShanTou)，韶关(ShaoGuan)，汕尾(ShanWei)，云浮(YunFu)，阳江(YangJiang)，珠海(ZhuHai)，中山(ZhongShan)，肇庆(ZhaoQing)，湛江(ZhanXiang)
北海(BeiHai)，百色(BaiSe)，崇左(ChongZuo)，防城港(FangChengGang)，桂林(GuiLi)，贵港(GuiGang)，贺州(HeZhou)，河池(HeChi)，柳州(LiuZhou)，来宾(LaiBin)，南宁(NanNing)，钦州(QinZhou)，梧州(WuZhou)，玉林(YuLin)
安顺(AnShun)，贵阳(GuiYang)，六盘水(LiuPanShui)，遵义(ZunYi)
白银(BaiYin)，金昌(JinChang)，嘉峪关(JiaYuGuan)，酒泉(JiuQuan)，兰州(LanZhou)，庆阳(QingYang)，天水(TianShui)，武威(WuWei)，张掖(ZhangYe)
海口(HaiKou)，三亚(SanYa)
保定(BaoDing)，承德(ChengDe)，沧州(CangZhou)，邯郸(HanDan)，衡水(HengShui)，廊坊(LangFang)，秦皇岛(QinHuangDao)，石家庄(ShiJiaZhuang)，唐山(TangShan)，邢台(XingTai)，张家口(ZhangJiaKou)
安阳(AnYang)，焦作(JiaoZuo)，开封(KaiFeng)，洛阳(LuoYang)，鹤壁(HeBi)，漯河(LuoHe)，南阳(NanYang)，平顶山(PingDingShan)，濮阳(PuYang)，三门峡(SanMenXia)，商丘(SHangQiu)，新乡(XinXiang)，许昌(XuChang)，信阳(XinYang)，郑州(ZhengZhou)，周口(ZhouKou)，驻马店(ZhuMaDian)
大庆(DaQing)，哈尔滨(HaErBin)，鹤岗(HeGang)，黑河(HeiHe)，鸡西(JiXi)，佳木斯(JiaMuSi)，牡丹江(MuDanJiang)，齐齐哈尔(QiQiHaEr)，七台河(QiTaiHe)，绥化(SuiHua)，双鸭山(ShuangYaShan)，伊春(YiChun)
鄂州(E'Zhou)，黄石(HuangShi)，黄冈(HuangGang)，荆州(JingZhou)，荆门(JingMen)，十堰(ShiYan)，武汉(WuHan)，襄阳(XiangYang)，孝感(XiaoGan)，咸宁(XianNing)，宜昌(YiChang)
长沙(ChangSha)，常德(ChangDe)，郴州(ChenZhou)，衡阳(HengYang)，怀化(HuaiHua)，娄底(LouDi)，邵阳(ShaoYang)，湘潭(xiangTan)，岳阳(YueYang)，益阳(YiYang)，永州(YongZhou)，株洲(ZhuZhou)，张家界(ZhangJiaJie)
白山(BaiShan)，白城(BaiCheng)，长春(ChangChun)，吉林(JiLin)，辽源(LiaoYuan)，四平(SiPing)，松原(SongYuan)，通化(TongHua)
常州(ChangZhou)，淮安(HuaiAn)，连云港(LianYunGang)，南京(NanJing)，南通(NanTong)，苏州(SuZhou)，宿迁(SuQian)，泰州(TaiZhou)，无锡(WuXi)，徐州(XuZhou)，盐城(YanCheng)，扬州(YangZhou)，镇江(ZhenJiang)
抚州(FuZhou)，赣州(GanZhou)，景德镇(JingDeZhen)，九江(JiuJiang)，吉安(JiAn)，南昌(NanChang)，萍乡(PingXiang)，上饶(ShangRao)，新余(XinYu)，鹰潭(YingTan)，宜春(YiChun)
鞍山(AnShan)，本溪(BenXi)，朝阳(ChaoYang)，大连(DaLian)，丹东(DanDong)，抚顺(FuShun)，阜新(FuXin)，葫芦岛(HuLuDao)，锦州(JinZhou)，辽阳(LiaoYang)，盘锦(PanJin)，沈阳(ShenYang)，铁岭(TieLing)，营口(YingKou)
包头(BaoTou)，赤峰(ChiFeng)，鄂尔多斯(E'ErDuoSi)，呼和浩特(HuHeHaoTe)，通辽(TongLiao)，乌海(WuHai)
固原(GuYuan)，吴忠(WuZhong)，银川(YingChuan)
西宁(XiNing)
安康(AnKang)，宝鸡(BaoJi)，汉中(HanZhong)，商洛(ShangLuo)，铜川(TongChuan)，渭南(WeiNan)，西安(Xi'An)，延安(YaNan)，咸阳(XianYang)，榆林(YuLin)
长治(ChangZhi)，大同(DaTong)，晋城(JinCheng)，临汾(LinFen)，朔州(ShuoZhou)，太原(TaiYuan)，忻州(XinZhou)，阳泉(YangQuan)，运城(YunCheng)
滨州(BinZhou)，东营(DongYing)，德州(DeZhou)，菏泽(HeZe)，济南(JiNan)，济宁(JiNing)，莱芜(LaiWu)，临沂(LinYi)，聊城(LiaoCheng)，青岛(QingDao)，日照(RiZhao)，泰安(TaiAn)，潍坊(WeiFang)，威海(WeiHai)，烟台(YanTai)，淄博(ZiBo)，枣庄(ZaoZhuang)
巴中(BaZhong)，成都(ChengDu)，德阳(DeYang)，达州(DaZhou)，广元(GuangYuan)，广安(GuangAn)，泸州(LuZhou)，乐山(LeShan)，绵阳(MianYang)，眉山(MeiShan)，内江(NeiJiang)，南充(NanChong)，攀枝花(PanZhiHua)，遂宁(SuiNing)，雅安(YaAn)，宜宾(YiBin)，自贡(ZiGong)，资阳(Ziyang)
高雄(GaoXiong)，基隆(JiLong)，嘉义(JiaYi)，台北(TaiBei)，台中(TaiZhong)，新竹(XinZhu)
拉萨(LaSa)
克拉玛依(KeLaMaYi)，乌鲁木齐(WuLuMuQi)
保山(BaoShan)，昆明(KunMing)，玉溪(YuXi)，昭通(ZhaoTong)
杭州(HangZhou)，湖州(HuZhou)，嘉兴(JiaXing)，金华(JinHua)，丽水(LiShui)，宁波(NingBo)，衢州(QuZhou)，绍兴(ShaoXing)，台州(TaiZhou)，温州(WenZhou)，舟山(ZhouShan)
`

func GetWeChatPath() string {
	keys := libs.GetUserPath(``)

	installPath := ""
	fileSavePath := ""
	for _, key := range keys {
		if strings.Contains(key, "WeChat") {
			regKey, err := registry.OpenKey(registry.USERS, key, registry.READ)
			if err != nil {
				// fmt.Printf("get lang key failed: %s\n", err)
				continue
			}
			defer regKey.Close()

			_installPath, _, err := regKey.GetStringValue("InstallPath")
			if err == nil {
				installPath = _installPath
			}

			_fileSavePath, _, err := regKey.GetStringValue("FileSavePath")
			if err == nil {
				fileSavePath = _fileSavePath
			}
		}
	}

	if fileSavePath == "" && installPath != "" {
		fileSavePath = fmt.Sprintf("%s\\locales\\WeChat Files\\", installPath)
	}

	if !libs.PathExists(fileSavePath) {
		fileSavePath = libs.GetHomeDir() + "\\Documents\\WeChat Files\\"
	}

	if !libs.PathExists(fileSavePath) {
		return ""
	}

	return fileSavePath
}

// 当前电脑登录过的账号
func GetWeChatUsers() []string {
	userIds := []string{}

	fileSavePath := GetWeChatPath()
	files, err := os.ReadDir(fileSavePath)
	if err != nil {
		return userIds
	}

	for _, f := range files {
		fileName := f.Name()
		if strings.Contains(fileName, "wxid") {
			userIds = append(userIds, fileName)
		}
	}

	return userIds
}

func GetUserInfo(wechatPath, userPath string) {
	accountFile := path.Join(wechatPath, userPath, "config\\AccInfo.dat")
	accountFile = strings.ReplaceAll(accountFile, "/", "\\")
	accountFile = strings.ReplaceAll(accountFile, "\\\\", "\\")

	dt, err := os.ReadFile(accountFile)
	if err != nil {
		return
	}

	str := libs.Byte2String(dt, libs.UTF8)
	str = str[strings.Index(str, "wxid"):]

	info := ""
	for _, num := range str {
		char := fmt.Sprintf("%c", num)

		// fmt.Println(num, char)

		if num == 26 {
			info += "`"
		} else {
			info += string(char)
		}
	}

	var wxid string
	var account string
	var logo string
	var name string

	var province string
	var city string

	dts := strings.Split(info, "`")

	// fmt.Println()
	// for _, misc := range dts {
	// 	fmt.Println(misc[0], misc[1], misc[2], misc[3], misc[4], misc[5], "--", misc)
	// }
	// fmt.Println()

	for _, misc := range dts {
		if strings.Contains(misc, "wxid") {
			wxid = misc
			// fmt.Printf("wxid: %s\n", wxid)
		}
		if strings.Contains(strings.ToLower(Provinces), strings.ToLower(misc)) {
			province = misc
			// fmt.Printf("province: %s\n", province)
		}
		if strings.Contains(strings.ToLower(Cities), strings.ToLower(misc)) {
			city = misc
			// fmt.Printf("city: %s\n", city)
		}
		if len(misc) > 6 && len(misc) < 20 {
			firstNum, err := strconv.Atoi(string(misc[0]))
			if err == nil && firstNum >= 0 && firstNum <= 9 {
				account = misc
				// fmt.Printf("account: %s\n", account)
			}
		}
		if strings.Index(misc, "qlogo") > 0 && logo == "" {
			logo = "http" + strings.Split(misc, "http")[1]
			// fmt.Printf("logo: %s\n", logo)
		}
		if misc[1] == byte(8) && misc[3] == byte(18) && misc[0] == 4+misc[4] && misc[4] != 32 {
			name = misc[5:]
		}
	}

	fmt.Printf("wxid: %s, account: %s, name %s, province: %s, city: %s, logo: %s\n", wxid, account, name, province, city, logo)
}

func main() {
	wechatPath := GetWeChatPath()
	fmt.Println(wechatPath)

	userIds := GetWeChatUsers()
	fmt.Println(userIds)

	for _, userId := range userIds {
		GetUserInfo(wechatPath, userId)
	}
}
