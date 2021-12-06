package sensitive

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var sendatapath string
var sentivedir = "sensitive"
var sensitiveTxt = "sensitive.txt"

func SernsitiveInit(dataPath string) {



	personSensitiveWords()
	go initSystemSensitiveWords()
}


func mkdriDht(dataPath string) {
	sendatapath = dataPath
	senTextPath := filepath.Join(dataPath, sentivedir)
	_, err := os.Stat(senTextPath)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		err := os.Mkdir(senTextPath, os.ModePerm)
		if err != nil {
			logrus.Error("error", err)
		}
	}
}


func SensitiveSave(sensitiveWords string) error {
	return writeSensitiveToFile(sensitiveWords)
}


func QuerySensitive() (string, error) {
	sentiByte, err := readSensitiveWords()
	if err != nil {
		logrus.Error("error", err)
	}
	if sentiByte != nil {
		return string(sentiByte), err
	} else {
		return "", nil
	}

}


func writeSensitiveToFile(sensitiveWords string) error {
	btTextPath := filepath.Join(sendatapath, sentivedir, sensitiveTxt)
	f, err := os.OpenFile(btTextPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		if !strings.Contains(err.Error(), "The system cannot find the file specified") && !strings.Contains(err.Error(), "no such file or directory") {
			logrus.Error("file not exist", err)
			return err
		} else {
			f, err = os.Create(btTextPath)
			if err != nil {
				logrus.Error("error", err)
				return err
			}
		}
	}
	defer f.Close()
	n, err := f.Write([]byte(sensitiveWords))
	if err == nil && n < len([]byte(sensitiveWords)) {
		err = io.ErrShortWrite
	}
	if err != nil {
		logrus.Error("error", err)
		return err
	}
	return nil
}


func readSensitiveWords() ([]byte, error) {







	return []byte(personDefaultWord), nil





}

const systemword = "考前答案 万科 家宝 辛灏年 陈胜 紧掏 紧淘 锦淘 锦掏 紧套 藏独 学生静坐 人体炸弹 温家保 炸药 代考 温家堡 造反 共产党 温总 恽小华 黄疽 胡进套 温家饱 黄JU HUANG菊 HUANGJU huang菊 黄ju huangju 绝食 静坐 声援 请愿 八九六 八九 观音法门 升达 郭玉闪 成杰 余辉 车殿光 秦高潮 王克勤 张振刚 董昕 王学永 李宇静 褚玉光 刘志华 宗顺留 庄公惠 朱振中 朱兆良 朱增泉 朱永新 朱相远 朱文泉 朱启 朱佩玲 朱培康 朱铭 朱丽兰 周子玉 周正庆 周玉清 周永康 周宜兴 周小川 周铁农 周绍熹 周坤仁 周伯华 仲兆隆 钟起煌 征鹏 赵展岳 赵勇 赵燕 赵喜明 赵龙 赵乐际 赵金铎 赵地 章祥荪 张左己 张中伟 张志坚 张芝庭 张云川 张毓茂 张佑才 张永珍 张学忠 张学东 张绪武 张新时 张肖 张吾乐 张文岳 张维庆 张廷翰 张涛 张思卿 张圣坤 张榕明 张庆黎 张洽 张平 张美兰 张梅颖 张龙俊 张立昌 张克辉 张俊九 张继禹 张怀西 张宏伟 张国祥 张工 张耕 张高丽 张帆 张发强 张定发 张德邻 张德江 张大宁 张春贤 张承芬 张宝文 张宝顺 张宝明 张柏林 翟泰丰 扎汗·俄马尔 曾荫权 曾宪梓 曾培炎 曾华 袁行霈 袁驷 袁汉民 喻林祥 俞正声 俞正 俞泽猷 俞云波 余国春 于珍 于幼军 于均波 尤仁 一诚 叶小文 叶少兰 叶如棠 叶青 叶连松 叶朗 叶大年 姚志彬 姚湘成 姚守拙 杨振杰 杨永良 杨兴富 杨孙西 杨岐 杨俊文 杨景宇 杨晶 杨国庆 杨国梁 杨德清 杨春兴 杨崇汇 杨长槐 杨邦杰 杨柏龄 阳安江 阎洪臣 严义埙 许仲林 许智宏 许志琴 许永跃 许克敏 许嘉璐 许柏年 徐自强 徐志纯 徐至展 徐展堂 徐永清 徐荣凯 徐麟祥 徐匡迪 徐鸿道 徐光春 徐冠华 徐更生 徐才厚 邢世忠 邢军 信春鹰 谢佑卿 谢生林 谢丽娟 谢克昌 萧灼基 向巴平措 夏赞忠 夏培度 夏家骏 习近平 西纳 武连元 伍增荣 伍淑清 伍绍祖 吴祖强 吴正德 吴贻弓 吴新涛 吴蔚然 吴双战 吴润忠 吴明熹 吴敬琏 吴基传 吴国祯 吴光正 吴光宇 吴冠中 吴德馨 吴爱英 乌云其木格 乌日图 魏复盛 卫留成 韦家能 王佐书 王祖训 王忠禹 王兆国 王占 王云龙 王云坤 王永炎 王英凡 王以铭 王耀华 王学萍 王旭东 王先琼 王维城 王涛 王太岚 王宋大 王生铁 王少阶 王忍之 王全书 王钦敏 王岐山 王宁生 王明明 王珉 王梦奎 王蒙 王梅祥 王茂润 王茂林 王良漙 王立平 王力平 王乐泉 王巨禄 王金山 王建民 王怀远 王鸿举 王恒丰 王鹤龄 王国发 王广宪 王光谦 王刚 王东明 王东江 王东 王大中 王选 汪啸风 汪恕诚 汪纪戎 汪光焘 万学远 万学文 万钢 万鄂湘 瓦哈甫·苏来曼 图道多吉 童傅 田玉科 田期玉 田岚 田成平 陶伯钧 汤洪高 索丽生 孙优贤 孙英 孙晓群 孙文盛 孙淑义 孙金龙 孙家正 孙淦 隋明太 苏荣 苏纪兰 宋秀岩 宋瑞祥 宋平顺 宋金升 宋宝瑞 舒圣佑 舒惠国 石宗源 石秀诗 石万鹏 石四箴 石广生 盛华仁 圣辉 沈辛荪 沈春耀 邵奇惠 邵华泽 邵鸿 任玉岭 任文燕 任启兴 任茂东 任克礼 曲钦岳 秦玉琴 乔晓阳 乔清晨 钱运录 钱景仁 启功 齐续春 彭钊 彭小枫 庞丽娟 潘霞 潘贵玉 潘蓓蕾 帕巴拉·格列朗杰 欧阳明高 钮茂生 倪岳峰 倪润峰 倪国熙 南振中 墨文川 莫时仁 闵智亭 闵乃本 孟建柱 毛增华 毛如柏 买买提明·阿不都热依木 马忠臣 马志伟 马永伟 马万祺 马庆生 马启智 马凯 罗清泉 罗豪才 栾恩杰 吕祖善 路甬祥 路明 陆锡蕾 陆浩 陆兵 卢展工 卢瑞华 卢荣景 卢强 卢光琇 卢登华 卢邦正 柳斌 刘仲藜 刘忠德 刘志忠 刘志军 刘政奎 刘镇武 刘振伟 刘泽民 刘云山 刘元仁 刘永好 刘迎龙 刘应明 刘亦铭 刘延东 刘廷焕 刘书田 刘胜玉 刘绍先 刘璞 刘明祖 刘民复 刘积斌 刘珩 刘鹤章 刘汉铨 刘光复 刘冬冬 刘大响 刘炳森 刘柏年 刘淇 令狐安 林兆枢 林文漪 林强 列确 廖锡龙 廖晖 梁振英 梁荣欣 梁绮萍 梁金泉 梁国扬 梁光烈 厉有为 厉以宁 厉无畏 李重庵 李至伦 李肇星 李兆焯 李泽钜 李源潮 李元正 李勇武 李雅芳 李学举 李新良 李铁映 李树文 李世济 李盛霖 李慎明 李乾元 李奇生 李其炎 李明豫 李敏宽 李良辉 李连宁 李金明 李金华 李建国 李继耐 李慧珍 李贵鲜 李赣骝 李德洙 李从军 李慈君 李春亭 李承淑 李成玉 李昌鉴 李宝祥 李蒙 黎乐民 冷溶 雷鸣球 雷蕾 孔小均 克尤木·巴吾东 靖志远 靳尚谊 金异 金日光 金人庆 金鲁贤 金烈 金开诚 金基鹏 金炳华 蒋正华 蒋以任 蒋树声 姜颖 姜笑琴 姜恩柱 贾志杰 贾军 季允石 吉佩定 霍英东 霍达 黄智权 黄小晶 黄孟复 黄丽满 黄康生 黄璜 黄华华 黄光汉 黄关从 黄格胜 胡彦林 胡贤生 胡康生 胡光宝 胡富国 胡德平 胡彪 侯义斌 洪绂曾 贺一诚 贺旻 贺铿 何柱国 何晔晖 何添发 何鲁丽 何厚铧 何鸿燊 何椿霖 郝建秀 韩忠朝 韩正 韩寓群 韩喜凯 韩生贵 韩汝琦 韩启德 韩大建 郭廷标 郭树言 郭金龙 郭凤莲 郭东坡 郭伯雄 郭炳湘 桂世镛 顾秀莲 谷建芬 苟建丽 龚学平 龚世萍 龚谷成 高占祥 高强 高洪 高国才 甘子钊 甘宇平 傅志寰 傅铁山 傅杰 傅家祥 傅惠民 符廷贵 奉恒高 冯培恩 冯健亲 方祖岐 方兆祥 方兆本 方新 范徐丽泰 范长龙 范宝俊 多吉才让 杜宜瑾 杜铁环 杜青林 窦瑞华 董建华 丁石孙 丁人林 丁光训 邓伟志 邓朴方 邓成城 邓昌友 德哇仓 刀述仁 刀美兰 戴证良 戴相龙 丛斌 储波 程誌青 程贻举 程津培 程安东 陈宗兴 陈政立 陈章良 陈永棋 陈益群 陈宜瑜 陈耀邦 陈勋儒 陈心昭 陈士能 陈绍基 陈清泰 陈清华 陈难先 陈明德 陈凌孚 陈良宇 陈奎元 陈抗甫 陈俊亮 陈军 陈建生 陈建国 陈佳贵 陈佳洱 陈辉光 陈虹 陈昊苏 陈广元 陈广文 陈高华 陈德铭 陈德敏 陈昌智 陈炳德 陈邦柱 曹圣洁 曹其真 曹刚川 曹伯纯 薄熙来 包叙定 白立忱 白克明 白恩培 巴金 安启元 艾丕善 出售答案 邵长良 费良勇 刘西峰 李新德 浦志强 郭飞熊 郭飞雄 杨茂东 李大同 袁伟时 施亮 安魂网 杜湘成 杨文武 许万平 夏逸陶 杨斌 唐淳风 卢雪松 胡绩伟 朱成虎 退党 游行 真善忍 江棋生 王通智 孙连桂 伊济源 赵铁斌 吴镇南 王维民 邓可人 宗凤鸣 红魂网站 法会 茳澤民 回良玉 王传平 徐金芝 昝爱宗 刘晓竹 夏春荣 杜导斌 张先玲 丁家班 何家栋 焦国标 集体上访 孟令伟 蔣彥永 同胞书 蒋彦永 吴仪 民主墙 张博涵 六四 共军 共匪 共党 河殇 贝领 暴政 高自联 反共 多维 民主潮 民运 疆独 达赖 澤民 朱镕基 李鹏 吴官正 罗干 李长春 黄菊 曾庆红 贾庆林 吴邦国 胡锦涛 苏晓康 李旺阳 李禄 严家其 温元凯 万润南 江公子 江锦恒 洪吟 张京生 方励之 刘晓波 张祖桦 侯杰 吕加平 任畹町 张晓平 杨子立 王建军 周国强 陈子明 方觉 马强 何德普 王美茹 华惠棋 刘凤钢 刘军宁 陈小雅 谢小庆 章虹 鲍彤 鲍筒 丁子霖 肖爱玲 傅怡彬 魏京生 王丹 何清涟 华岳 李志绥 柴玲 暴乱 陶驷驹 江责民 法輪 转法轮 法轮 李大师 李宏治 李宏志 明Hui 慧网 明慧 府谷 骚乱 拉萨事件 310事件 马明哲 关六如 周鸿陵 李伙田 李汉田 李秋田 钟志根 李连玉 邳州 段桂青 段桂清 死亡笔记 梁定邦 清明桥 令计划 覃志刚 艾仰华 陈晓铭 周天勇 动乱 胡耀邦 大法 台独 文化大革命 文革 学运 学联 学潮 天葬 一中一台 镇压 两个中国 法轮功 法轮大法好 谢燕益 告同胞书 新唐人 关贵敏 白雪 JIANZEMIN LIPENG FALUN 江主席 LIHONGZH jiangzemin 小B 他妈的 国民党 新唐人电视台 PowertotheFalunGong BrothersinArms theUndergroundResistance FalunGong UndergroundResistanceArms BrothersTIBET tibet 预审查 请看广传海外的报导 鞍钢 鞍山钢铁 重庆特钢厂 重特钢 封杀 主持 解说 版主 魔兽 南水北调 三峡 LIHONGZHI 一党专治 一党专制 人事 多党 民主 暴动 天安门 季先 小油条 一党制 延昌雅士 卫寺狂 遗珠 南露 老夫永乐 塞北云中鹤1 XJG351 假货 出售 贩假 卖假 嘴脸 骗子 小丑 那日贵 蚁力神 助考 枪手 台海 台湾 立委 选举 谢燕益 告同胞书 张铁钧 戚晓春 姜敏 陈瑞武 苏显达 张佳慧 大小百合 江泽民 李洪志 法抡 李瑞环 赵化勇 王八 周恩来 邓小平 李宏治 毛泽东 江青 林彪 变态 法伦 丁关根 毛主席 西藏 打倒 拉登 告全体网民书 孙大午 朱胜文 民进党 罗明 张长明 胡恩 平反 六·四 生者与死者 天安门母亲 曹思源 茅于轼 娄义 先审 新闻管制 舆论钳制 2.23会议 蒋彦勇 社会公器 冯骥 赵安 线人 分家在十月 波波娃 诺夫 上街 集会 牛羊瘟疫 老刀 中宣部 讨伐 正邪大决战 救度世人 赵紫阳 总书记 大参考 大纪元郑重声明 001工程 共和国之辉 聚集 抗议 堵截 堵路 天安门广场 爱国者同盟 反日 反曰 抵制 日货 曰货游行 抗日 抗曰 日本 曰本 保钓 钓鱼岛 使馆 领事馆 示威 五四 五月四日 5.4 天鹅绒行动 颜色革命 天鹅绒 张戎 《毛：不为人知的故事》 东突 圣战 太石村 太*石*村 泰石村 鱼窝头镇 罢官筹委会 羊皮狼 拉法叶舰 拉案 罢课 许万平 人事调整 王江波 春运 神职 龙滩水电站 郑茂清 邓天生 黎国如 田青 程宝山 程宝山中将 程中将 CHENG中将 肖副团长 肖白 肖大校 二炮 石占明 羊倌 国旗 沙桐 朱环 林程 女主持 黄健翔 刘建宏 段喧 张斌 冬日那 韩乔生 佐藤琢磨 婴儿汤 婴儿煲汤 四川朱昱 占地 拆迁 小胡 李咏 孙悦 伍思凯 举报 黄色网站 法院 好处费 政府 郑培民 陈水扁 阿扁 赵建铭 吕秀莲 苏贞昌 罢扁 巴拉圭 下课 倭国 主持人 下岗 李文娟 彝 穆斯林 回族 蛮 伊斯兰 王红旗 藏独 成吉思汗 蒙古 转世 外逃 占领 嘎玛巴 藏传佛教 陈光诚 周金伙 傅先财 施丹 道歉 黄贱翔 护短 侮辱 霸道 斑竹 主义 贱翔 周稚舜 中央电视台 汉奸 温家宝 秦裕 祝均一 高智晟 程翔 赵岩 六合彩 垃圾 央视 经纬 十七大 人事变动 削权 筹备 高知 团系 接班 胡温 整肃官场 省市换班 第五代 冒起 人事小组 宋平 七种人 政治局会议 筹备领导组 李晓峰 蔡欢 法轮功 自焚 雅昌 百度古董 百家 宋什么 上访 领导 柴璐 李红 预审 大纪元 年轻 刘少奇 朱德 彭德怀 刘伯承 陈毅 贺龙 聂荣臻 徐向前 罗荣桓 叶剑英 李大钊 陈独秀 孙中山 孙文 孙逸仙 陈云 尉健行 李岚清 唐家璇 华建敏 陈至立 贺国强 李登辉 连战 宋楚瑜 郁慕明 蒋介石 蒋中正 蒋经国 马英九 布什 布莱尔 小泉纯一郎 萨马兰奇 安南 阿拉法特 普京 默克尔 克林顿 里根 尼克松 林肯 杜鲁门 赫鲁晓夫 列宁 斯大林 马克思 恩格斯 金正日 金日成 萨达姆 胡志明 西哈努克 希拉克 撒切尔 阿罗约 曼德拉 卡斯特罗 富兰克林 华盛顿 艾森豪威尔 拿破仑 亚历山大 路易 拉姆斯菲尔德 劳拉 鲍威尔 布朗 奥巴马 梅德韦杰夫 潘基文 本拉登 奥马尔 达赖喇嘛 张春桥 姚文元 王洪文 东条英机 希特勒 墨索里尼 冈村秀树 冈村宁次 高丽朴 沃尔开西 赖昌星 马加爵 班禅 额尔德尼 山本五十六 热比娅 鲁迅 莎士比亚 爱因斯坦 牛顿 哥白尼 哥伦布 达芬奇 米开朗基罗 毕加索 梵高 海明威 斯蒂芬霍金 六四运动 美国之音 密宗 民国 摩门教 纳粹 南华早报 南蛮 明慧网 起义 亲民党 瘸腿帮 人民报 法轮大法 台独分子 台联 台湾民国 台湾独立 太子党 天安门事件 屠杀 小泉 新党 新疆独立 新疆国 西藏国 西藏独立 一党专政 一贯道 圆满 政变 政治 政治反对派 政治犯 中共 中国民主党 中国之春 苏家屯 基地组织 塔利班 东亚病夫 支那 高治联 藏青会 藏妇会 台湾岛国 台湾分裂 西藏分裂 新疆分裂 核工业基地 核武器 铀 原子弹 氢弹 导弹 核潜艇 小参考 国内动态详情 道教 佛教 佛祖 释迦牟尼 如来 阿弥陀佛 观世音 普贤 文殊 地藏 回教 穆罕默德 升天 圣母 耶和华 耶稣 真主安拉 白莲教 天主教 基督教 东正教 真理教 走向圆满 黄大仙 风水 跳大神 神汉 神婆 大卫教 阎王 黑白无常 牛头马面 谋杀 杀人 吸毒 贩毒 赌博 拐卖 走私 卖淫 监狱 抢劫 先奸后杀 下注 押大 押小 抽头 坐庄 赌马 赌球 筹码 老虎机 轮盘赌 安非他命 大麻 可卡因 海洛因 冰毒 摇头丸 杜冷丁 鸦片 罂粟 迷幻药 白粉 嗑药 陈天桥 16大 十六大 64事件 六四学潮 64运动 89事件 中国共产党 共产 共产主义 中共中央政治局 届中央政治局委 中央政治局 反党 反人民 反动 反动派 反革命 打倒中国 打倒中国共产党 打倒中华人民共和国 党中央 国家主席 党主席 台湾帝国 台湾公投 台湾共产党 台湾共和国 台湾国 台湾总统 保卫台湾 倒台 两国论 全国人大 陈总统 阿扁万岁 东突暴动 东突独立 a-lun-gong FLG 发伦 藏独会 藏独分子 藏独分裂 藏独势力 真世界 膏药旗 根正苗红 黑五类 红卫兵 皇军 共荣圈 国共合作 国军 国民党万岁 国民政府 四人帮 靖国神社 没有人权 批林批孔 政府无能 政治避难 政治打击 政治封锁 政治压迫 四清 抢夺圣火 社会主义 中国人民解放军 中国人民志愿军 中华人民共和国 中南海 国务院 革命 公安局 政治局 暴行 朝鲜族 大东亚 大东亚共荣圈 大跃进 大字报 抵制日货 独裁 法西斯 浮夸风 腐败 复辟 密传 和平演变 一边一国 戒急用忍 军国主义 垮台 民國 中华民国 中華民國 青天白日 青天白日旗 政协 人民大会堂 总理 委员 常委 西藏喇嘛 资本主义 指导员 武力统一 台湾党 走资帮 天皇 天神 上山下乡 言论自由 新义安 毛澤東 鄧小平 江泽明 江澤民 江core 蔡启芳 蔡和森 曹庆泽 阿沛?阿旺晋美 蔡英文 陈博志 蔡庆林 陈伯达 陈建铭 薄一波 陈慕华 陈定南 陈锡联 陈菊 陈唐山 成克杰 陈丕显 陈永贵 姜春云 陈希同 邓发 董必武 邓颖超 谷牧 邓力群 杜正胜 董文华 韩杼滨 洪兴 顾顺章 郝伯村 韩光 何勇 韩天石 胡启立 黄永生 胡乔木 黄克诚 瞿秋白 华国锋 纪登奎 康生 凯丰 李登柱 黄仲生 李德生 李雪峰 李俊毅 李立三 李维汉 傅作义 李先念 傅全有 廖承志 李作鹏 林佳龙 林益世 卢福坦 林信义 林祖涵 刘华清 连惠心 林伯渠 刘丽英 刘文雄 刘澜涛 陆定一 马国瑞 倪志福 彭冲 彭佩云 彭真 祁培文 钱其琛 乔冠华 乔石 宋庆龄 宋任穷 苏兆征 苏振华 王从吾 王汉斌 王鹤寿 王稼祥 王克 王金平 王震 韦国清 郑宝清 秦基伟 朱立伦 邹家华 李克强 张廷发 张万年 张闻天 张震 章孝勇 章孝严 赵洪祝 曾志郎 杨白冰 张博雅 邱会作 邱太三 杨得志 杨尚昆 杨勇 姚依林 叶菊兰 叶群 任弼时 任建新 翁金珠 习仲勋 卓伯源 谭绍文 汪东兴 田弘茂 田纪云 许世友 谭震林 耿飚 谭政 吴德 许财利 吴桂贤 伍世文 向忠发 颜庆章 于永波 余秋里 谢长廷 谢非 谢深山 袁纯清 唐骏 李红志 川岛芳子 塞福昂 艾则孜 扎卡维 乌兰夫 三个代表 一党 专政 洪志 红志 洪智 红智 法论 法沦 发轮 发论 发沦 轮功 轮公 轮攻 沦功 沦公 沦攻 论攻 论功 论公 伦攻 伦功 伦公 朱容基 警察 民警 公安 大盖帽 武警 黑社会 交警 消防队 刑警 公款 首长 书记 城管 李远哲 司法警官 高干 人大 文字狱 宋祖英 骗局 猫肉 吸储 张五常 张丕林 空难 专制 三個代表 一黨 多黨 專政 大紀元 紅志 紅智 法論 法淪 法倫 發輪 發論 發淪 發倫 輪功 輪公 輪攻 淪功 淪公 淪攻 論攻 論功 論公 倫攻 倫功 倫公 民運 台獨 李鵬 天安門 朱鎔基 李長春 李瑞環 胡錦濤 臺灣獨立 藏獨 西藏獨立 疆獨 新疆獨立 大蓋帽 黑社會 消防隊 夜總會 媽個 首長 書記 腐敗 暴動 暴亂 李遠哲 高幹 李嵐清 黃麗滿 於幼軍 文字獄 騙局 貓肉 吸儲 張五常 張丕林 空難 溫家寶 吳邦國 曾慶紅 黃菊 羅幹 賈慶林 專制 八老 巴赫 白立朴 白梦 白皮书 鲍戈 北大三角地论坛 北韩 北京当局 北京之春 北美自由论坛 博讯 蔡崇国 曹长青 常劲 陈炳基 陈蒙 陈破空 陈小同 陈宣良 陈一谘 程凯 程铁军 程真 迟浩田 持不同政见 赤匪 赤化 春夏自由论坛 大纪元新闻网 大纪园 大家论坛 大史 大史记 大史纪 大中国论坛 大中华论坛 大众真人真事 弹劾 登辉 邓笑贫 迪里夏提 地下教会 地下刊物 第四代 电视流氓 丁元 东北独立 东方时空 东南西北论谈 东社 东土耳其斯坦 东西南北论坛 独夫 独立台湾会 杜智富 屙民 俄国 发愣 发正念 反封锁技术 反腐败论坛 反攻 反人类 反社会 方舟子 飞扬论坛 斐得勒 分家在 分裂 粉饰太平 风雨神州 风雨神州论坛 封从德 冯东海 冯素英 佛展千手法 付申奇 傅申奇 高官 高文谦 高薪养廉 高瞻 戈扬 鸽派 歌功颂德 个人崇拜 工自联 功法 共狗 关卓中 贯通两极法 广闻 郭罗基 郭平 郭岩华 国家安全 国家机密 国贼 韩东方 韩联潮 红灯区 红色恐怖 宏法 洪传 洪哲胜 胡紧掏 胡锦滔 胡锦淘 胡景涛 胡平 胡总书记 护法 华通时事论坛 华夏文摘 华语世界论坛 华岳时事论坛 黄慈萍 黄祸 黄菊 黄翔 回民暴动 悔过书 鸡毛信文汇 姬胜德 积克馆 基督 贾廷安 贾育台 建国党 江八点 江流氓 江罗 江绵恒 江戏子 江则民 江泽慧 江贼 江贼民 江折民 江猪 江猪媳 将则民 僵贼 僵贼民 讲法 酱猪媳 揭批书 金尧如 锦涛 禁看 经文 开放杂志 看中国 邝锦文 劳动教养所 劳改 劳教 老江 老毛 黎安友 李红痔 李洪宽 李兰菊 李老师 李录 李少民 李淑娴 李文斌 李小朋 李小鹏 李月月鸟 李总理 李总统 连胜德 联总 廉政大论坛 炼功 梁擎墩 两岸关系 两岸三地论坛 两会 两会报道 两会新闻 林保华 林长盛 林樵清 林慎立 凌锋 刘宾深 刘宾雁 刘刚 刘国凯 刘俊国 刘凯中 刘千石 刘青 刘山青 刘士贤 刘文胜 刘永川 流亡 龙虎豹 陆委会 吕京花 抡功 轮大 罗礼诗 马大维 马良骏 马三家 马时敏 卖国 毛厕洞 毛贼东 美国参考 蒙独 蒙古独立 绵恒 民联 民意 民意论坛 民阵 民猪 民族矛盾 莫伟强 木犀地 木子论坛 南大自由论坛 闹事 倪育贤 你说我说论坛 潘国平 泡沫经济 迫害 祁建 齐墨 钱达 钱国梁 抢粮记 亲美 钦本立 秦晋 轻舟快讯 庆红 全国两会 热站政论网 人民内情真相 人民真实 人民之声论坛 人权 瑞士金融大学 善恶有报 上海帮 上海孤儿院 邵家健 神通加持法 沈彤 盛雪 石戈 事实独立 双十节 水扁 税力 司马晋 司马璐 司徒华 斯诺 四川独立 宋书元 苏绍智 台盟 台湾狗 台湾建国运动组织 台湾青年独立联盟 台湾政论区 台湾自由联盟 汤光中 唐柏桥 唐捷 滕文生 天怒 童屹 统独 统独论坛 统战 万晓东 汪岷 王宝森 王炳章 王策 王超华 王辅臣 王涵万 王沪宁 王军涛 王力雄 王瑞林 王润生 王若望 王希哲 王秀丽 王冶坪 网特 魏新生 吴百益 吴方城 吴弘达 吴宏达 吴仁华 吴学灿 吴学璨 吾尔开希 五不 伍凡 项怀诚 项小吉 肖强 邪恶 谢选骏 谢中之 新闻封锁 新语丝 信用危机 邢铮 熊炎 熊焱 修炼 徐邦秦 徐水良 许家屯 薛伟 学习班 学自联 雪山狮子 严家祺 阎明复 央视内部晚会 杨怀安 杨建利 杨巍 杨月清 杨周 姚月谦 义解 异见人士 异议人士 易丹轩 易志熹 尹庆民 由喜贵 于大海 于浩成 余英时 舆论 舆论反制 宇明网 远志明 岳武 在十月 则民 择民 泽民 贼民 张伯笠 张钢 张宏堡 张健 张林 张伟国 张昭富 张志清 赵海青 赵南 赵品潞 赵晓微 哲民 真相 真象 争鸣论坛 正见网 郑义 正义党论坛 89风波 89風波 9ping 9评 9坪 9評 fa轮 fa輪 FL功 freetibet GONG党 GONG黨 HEIDEROOSJESAND LOTENUNITEDTEBET hyperballad-tibet jiuping UnitedTibet 大际元 大際元 邓矮子 鄧矮子 狄玉明 法0功 法lun 法X功 共独 共獨 共贼 共賊 激流china 激流cn 金菩提 金菩提上师 金菩提上師 金麒麟 九ping 九评 九凭 九評 龙应台 龍應臺 泣血的春天 全球華人聯歡晚會 仁布其 神韵艺术团 神韻藝術團 生者和死者 生者與死者 四神記 透视中国 透視中國 退黨 亡党 亡黨 伪火 偽火 无产阶级专政 无界 吾尔凯希 希望之声 希望之聲 邪党 邪黨 邪教 新搪人 一九八九年春夏之交的政治风波 張宏堡 中功 周刊纪事 周刊紀事 周晓川 自由亚洲 自由亞洲 自由之门 自由之門 激流中国 激流中國 东楼 東樓 郭文贵"

const personDefaultWord = "1|-|淫 碼 乳 陰 姦 裸 码 阴 肛 妓 茎 尻 鸡 嫖 屎 逼 插 屄 操 屌 屙 舔 肏 婚外情 车震 骚穴 骚妈 骚母 骚妹 骚姐 骚妇 骚女 巨骚 骚妻 风骚 骚洞 骚浪 骚图 发骚 够骚 骚B 骚动 骚货 骚劲 骚叫 骚媚 骚情 骚热 骚声 骚水 骚味 骚痒 骚幽 骚状 骚逼 傻逼 阴唇 肉穴 肉洞 肉棍 毛片 交而不泄 先肾后心 9浅1深 BJ Blow Job CAR SEX G点 高潮 KJ Y染色体 師母 房中 慕男症 上下其手 after-play 包二奶 性爱 阿姨 爱抚 爱侣 爱女 爱慰 爱液 爱欲 yanjiaoshequ 安抚 按揉 按柔 按压 暗红 昂奋 傲然挺立 扒开 巴氏腺 拔出 把玩 白带恶臭 白嫩 白浊 摆布 摆动 摆弄 摆脱 扳开 伴侣 半遮半露 膀胱 棒棒 蚌唇 胞漏疮 包覆 包皮 剥开 保持 保精 饱胀 暴露 暴涨 暴胀 爆射 贝肉 闭经 壁肉 臂部 变粗 拨开阴毛 拨弄 勃发 勃起 性交 不举 不泄 不育 擦拭 采精 采阴补阳 侧臀 叉开 叉我 缠抱 缠绵 颤动 颤抖 长驱直入 长兄 潮红 潮湿 撑爆 撑破 撑涨 撑胀感 持续 耻部 耻毛 耻丘 炽热 充血 冲刺 抽擦 抽出 抽搐 抽打 抽捣 抽动 抽离 抽了 抽弄 抽送 抽送着 抽缩 初血 初夜 出水 出血 矗立 触动 触摸 触碰 处男 处女 喘叫 床事 床戏 吹弹欲破 吹萧 垂软 春洞 春宫 春情 春心 春药 唇瓣 唇缝 唇间 唇片 唇肉 唇舌 纯熟 蠢蠢欲动 戳穿 戳入 戳穴 雌二醇 雌激素 慈母 刺激 次数 丛毛 粗暴 粗长 粗粗 粗大 粗黑 粗红 粗鲁 粗硬 粗涨 粗壮 促使 窜动 摧残 催情 搓蹭 搓捏 搓弄 搓揉 搓柔 搓玩 搓著 搓着 打飞机 打炮 打泡 打手枪 大波 大抽 大奶 大娘 大肉 大泄 大穴 带状沟 蛋蛋 蛋子 荡妇 荡叫 荡声 档部 倒骑 登床 低潮期 低嚎 滴出 底裤 颠鸾倒凤 叼住 调逗 调经 调情 调戏 爹爹 叠股 顶紧 顶进 顶弄 顶破 顶送 顶我 顶住 动情区 动欲区 动作 洞洞 洞开 洞口 洞穴 洞眼 抖颤 逗弄 独生女 肚脐 短粗 多毛 多情 多汁 多姿 多睾 恶露 儿媳妇 耳垂 耳磨鬓擦 发颤 发春 发抖 发浪 发麻 发情 发热 发泄 发痒 发涨 翻动 翻搅 翻弄 芳香 房事 房室伤 房中七经 房中术 房中之术 放荡 放在 飞溅 飞燕 肥大 肥美 肥奶 肥翘 肥润 肥臀 肥穴 焚身 粉白 粉臂 粉额 粉汗微出 粉颊 粉嫩 粉腮 粉舌 粉头 粉腿 粉臀 粉腰 丰肥 丰隆 丰硕 丰臀 丰腴 风流 夫妻 伏在 服囊肉膜 抚爱 抚摸 抚模 抚摩 抚捏 抚弄 抚揉 抚玩 抚慰 抚著 抚着 俯弄 干过炮 缸交 高胀 根毛 供精 公公 共浴 沟缝 狗交 狗爬 箍住 鼓胀 骨感 骨盆 骨盆腔 股沟 刮宫 刮官 乖巧 乖肉 冠状沟 光洁无毛 光溜溜 光脱脱 广东疮 龟腾 龟头 鬼交 跪骑于 跪臀位 跪姿 滚动 滚热 滚烫 滚圆 含春 含弄 含入 含吮 含咬 含住 含着 好棒 好爽 好性 耗精伤气 呵痒 和谐 合拢 黑洞 黑黑的阴毛 黑毛 黑色的阴毛 狠干 横冲直撞 红唇 红颊 红润 喉交 侯龙涛 后入 后庭 后戏 呼呼 壶腹部 花瓣 花苞 花唇 花蕾 花蜜 花蕊 花芯 花心 花穴 滑出 滑到 滑动 滑抚 滑进 滑溜 滑美 滑嫩 滑入 滑润 滑湿 滑爽 滑顺 滑下 滑向 滑粘 欢爱 欢吟 欢愉 欢悦 环境 缓进速出 缓慢 唤起 秽疮 秽物 会阴 魂飞魄散 浑圆 混圆 活塞 火辣 火热 火柱 获得 击打 肌肉 饥渴 激情 激射 激素 急抽 急喘 挤捏 挤压 技巧 悸动 夹紧 夹著 夹住 夹着 假装 尖挺 奸弄 奸虐 奸辱 奸尸 奸我 将嘴套至 交缠 交合 交欢 交颈 交配 交融 交媾 交媾素 骄躯 骄穴 娇喘 娇哼 娇呼 娇叫 娇媚 娇嫩 娇娘 娇躯 娇容 娇软 娇弱 娇声 娇态 娇啼 娇小 娇笑 娇艳 娇吟 搅弄 脚交 较好 叫床 叫声 接触 接吻 洁阴法 解脲脲原体 筋肉 津液 紧小 紧咬 紧窄 精巢 精虫 精阜 精关失固 精浆 精满自溢 精门 精门开 精母 精囊 精水 精索 精脱 精细胞 精液 精元 精原 精原核 精种 精子 经期 经血 经验 静香 境界 痉挛 痉脔 久战 九浅一深 菊孔 咀唇 举高双腿 巨棒 巨根 巨棍 巨炮 巨枪 巨物 剧烈 撅起 撅着 绝经 俊逸 开苞 揩擦 亢奋 亢进 渴望 啃咬 控制射精 抠摸 抠弄 抠挖 口爆 口唇 口含 口交 口中 扣弄 狂抽 狂捣 狂干 狂热 狂吻 狂泄 狂舐 窥探 困境 括约肌 括约肌间沟 来潮 来经 来搔抚 浪喘 浪妇 浪哼 浪货 浪叫 浪劲 浪媚 浪女 浪态 浪穴 浪样 浪吟 浪语 蕾苞 类菌质体 冷阴症 连炮几炮 恋母 恋童 两腿之间 撩拨 撩动 撩开 撩乱 撩弄 撩起 林醉 淋巴管 淋巴结 淋病 淋菌 淋漓 淋球菌 淋证 凌乱 凌辱 灵肉 令女人春心荡漾 令女性春心荡漾 流出 流到 流溢 龙根 隆起 搂抱 露出 露阴 窥阴 陆玄霜 卵巢 卵子 乱抽 乱蹬 乱顶 乱伦 乱摸 乱揉 乱淌 轮暴 轮奸 妈咪 麻酥 麻酥酥 麻痒 马眼 满胀 满足 曼妙 毛囊 毛茸茸 美体 美腿 美臀 美香 美穴 妹夫 妹妹 妹子 媚唇 媚功 媚力 媚娘 媚肉 媚术 媚态 媚笑 媚艳 媚液 闷哼 猛颤 猛冲 猛抽 猛喘 猛刺 猛干 猛搅 猛男 猛挺 猛撞 梦交 梦失精 梦泄精 梦遗 迷情 秘贝 秘部 秘处 秘唇 秘洞 秘缝 秘穴 泌出 蜜唇 蜜洞 蜜壶 蜜肉 蜜桃 蜜穴 蜜液 蜜意 蜜汁 密处 密洞 密合 密窥 密穴 密汁 绵软 妙目 敏感带 名器 命根子 摸抠 摸摸 摸捏 摸弄 摸揉 摸他 摸玩 摸我 摸向 摸着 模式 磨擦 磨搽 磨搓 磨弄 磨穴 摩擦 那话 那话儿 奶房 奶尖 奶水 奶头 奶子 奶頭 男方膝立位 男根 男跪女後 男欢女爱 男精女血 男女 男上式 内睾提肌 嫩白 嫩红 嫩脸 嫩嫩 嫩肉 嫩舌 嫩爽 嫩腿 嫩臀 嫩娃 嫩穴 腻滑 逆行射精 娘亲 尿道口 浓精 浓密的阴毛 浓热 浓浊 弄弄 弄破 弄湿 弄穴女畅男欢 女方跪臀位 女器 女前男后 女人的BB 女上式 女上位 女臀 女卧男立式 女下位 女性 女阴 女优 性冷淡 虐待 排出 排过精 排精 排卵 排卵期 排卵日 排入 排射 排泄 抛浪 泡彦 泡浴 培养 配偶 喷出 喷发 喷射 喷泄 喷涌 屁股 屁门 屁穴 屁眼 频度 频繁 频率 平滑 破处 破瓜 破坏 破身 破贞 迫进 葡萄胎 漆黑的阴毛 奇痒 骑乘位 起性 器具 气氛 气淋 乾妈 乾姊 前戏 前戏 潜欲 浅出浅入 腔内 腔肉 强暴 强奸 鞘膜腔 翘起 翘臀 俏丽 俏脸 俏眼 切除子宫 亲热 情色 情穴 情欲 求欢 取悦 去能因子 去吮 泉涌 全根 缺乏 燃烧 让女人春心荡漾 让女性春心荡漾 热滚 热烘烘 热乎乎 热浆 热辣辣 肉~棒 肉瓣 肉棒 肉蚌 肉贝 肉壁 肉臂 肉搏 肉搏战 肉帛 肉肠 肉虫 肉唇 肉袋 肉弹 肉道 肉豆 肉缝 肉感 肉根 肉沟 肉冠 肉核 肉乎乎 肉壶 肉紧 肉具 肉孔 肉粒 肉门 肉膜 肉腔 肉丘 肉球 肉圈 肉色 肉身 肉体 肉团 肉臀 肉香 肉芽 肉芽肿 肉牙儿 肉眼 肉欲 肉柱 肉嘟嘟 肉襞 蠕动 如醉如痴 入浴 软掉 软了 软毛 软绵 软绵绵 软肉 润滑 润湿 塞进 塞入 搔弄 搔痒 色欲 杀精 伤精 上床 上翘 上上下下 上位 上下蠕动 上压下顶 少妇 射出 射到 射精 射了 射向 射液 伸入 身体 身无寸缕 身子 深喉 深入浅出 深吻 婶婶 生精 生母 生殖 受精 受孕 瘦小 兽奸 兽交 双唇 双峰 双管齐下 双胯 双奶 双腿 双臀 双子宫 爽滑 爽劲 爽快 爽死 爽透 水多 水淋淋 吮了 吮奶 吮吻 吮咬 吮著 吮着 吮舐 硕大 硕壮 撕开 撕裂 撕破 私部 私处 死精 死精症 死精子症 松弛 松软 耸动 耸起 酥到 酥淋 酥麻 酥熔 酥软 酥爽 酥酥 酥胸 酥痒 素女经 素燕 酸软 酸痒 锁精术 胎盘 瘫软 探入 探索 堂哥 堂妹 堂嫂 烫热 套动 套紧 套弄 套上 提枪 提睾筋膜 体味 体位 体香 体液 替我 剃掉 甜蜜 挑拨 挑逗 挑弄 铁硬 停经 挺立的性器 挺起 同性恋 童男 捅 捅进 捅了 痛快 偷汉 偷欢 偷窥 推油 腿儿 腿缝 腿根 腿间 臀部 臀瓣 臀部 臀洞 臀缝 脱光 脱裤 外流 外阴 玩六九 玩摸 玩弄 完事 晚期流产 万艾可 温存 吻遍 吻摸 吻向 稳定 我射了 握缩感 无精子症 无毛 无排卵性月经 无月经 无睾症 五淋 五征五欲 舞奴 戏弄 细菌性阴道病 峡部 下贱 下流 下身 下体 下阴 先射 香唇 香滑 香肌 香肩 香津 香嫩 香软 香臀 香涎 销魂 小便 小唇 小弟第 小弟弟 小洞 小缝 小核 小咀 小脸 小屁眼 小巧玲珑 小舌 小腿 小雄 小穴 小腰 小姨 小嘴 挟紧双腿 泄出 泄洪 泄精 泄了 泄射 泄身 泄欲 泻出 泻了 行房 性器 性欲 性奋 性感 性黄金 性虐 性侵犯 性情 性生活 性事 性妄想症 性戏 性腺 性信号 性行为 性招式 胸罩 虚脱 蓄精 穴壁 穴唇儿 穴道 穴洞 穴缝 穴门 穴肉 穴水 穴心 穴裡 穴穴 穴眼 穴痒 血精 血性精液 血睾 血睾屏障 艳妇 阳根 阳精 阳具 阳水 阳萎 阳物 养精 咬扯 野合 冶荡 腋尾 夜御数女 液体 一泻千里 衣原体 遗精 溢精 殷红 阴□ 阴壁 阴埠 阴道 阴蒂 阴洞 阴缝 阴阜 阴沟 阴垢 阴核 阴户 阴精 阴径 阴亏血燥证 阴毛 阴囊 阴肉 阴虱病 阴水 阴庭 阴穴 阴液 隐睾症 樱唇 应召 营造 迎合 盈满 硬邦邦 硬梆梆 硬绑绑 硬立 硬热 硬挺 硬物 硬硬 硬涨 右臀 诱惑 又稠又粘 又粗又短 又肥又厚 又美又嫩 又细又嫩 又肿又大 幼嫩 幼稚型子宫 玉腿 玉棒 玉背 玉穴 玉液 玉指 玉柱 欲火 欲望 早泄 造爱 增粗 张开双腿 张开小嘴 胀得难受 胀硬 阵阵快感 整根 肢体 壮神鞭 自慰 自渎 纵欲 撸着 攥住 啜著 噘起 噙住 後洞 後进 後穴 陽具 掰开 胴体 裆部 舐吻 舐著 舐着 酡红 蹂躏 稣胸 av貼圖 japansweet adult 16dy-图库 獸交 爱女人 拔出来 吃精 调教 黄色电影 迷奸 色猫 幼圖 中文搜性网 鷄巴 學生妹 999日本妹 幹炮 摸阴蒂 金鳞岂是池中物 掰穴 白虎 蕩妹 粉穴 干穴 口活 狼友 风艳阁 兽欲 菊花蕾 大力抽送 anal 肉蒲团 极品波霸 销魂洞 成人网站 一夜欢 偷窥图片 性奴 incest 奸幼 中年美妇 喷精 脱内裤 菊花洞 抠穴 露B 母子奸情 色界 爱图公园 色色五月天 鹿城娱乐 幼香阁 隐窝窝 美幼 97sese 日本AV女优 美女走光 33bbb走光 成人论坛 就去诱惑 BlowJobs 丽春苑 窝窝客 银民吧 亚洲色 碧香亭 爱色cc 宾馆女郎 好色cc 成人图片 大胆少女 要色色 洗肠射尿 騷浪 射爽 蘚鮑 無毛穴 舌头穴 食精 肉便器 母奸 発射 幹砲 性佣 爽穴 嫩鲍 金毛穴 体奸 爆草 a4u 酥穴 厕所盗摄 打洞 盗撮 薄码 少修正 成人片 穴爽 99bb g点 tw18 asiasex teen sexy 欢欢娱乐时空 裤袜 买春 妹妹阴毛 免费成人网站 免费偷窥网 免费A片 摩洛客 色区 色书库 射颜 性虎 性饥渴 幼交 嫩女 噴精 群交亂舞 色狐狸网址 色狼论坛 色狼小说 湿穴 好嫩 亂倫 hardcore amateur 色诱 秘裂 采花堂 亚洲性虐 熟女 裹本 欢乐性今宵 性愛 学生妹 炮友之家 花花公子 hotsex porn 小姐打飞机 按摩棒 震动棒 奸情 被干 阴缔 人妻 内射 欲仙欲浪 吞精 成人午夜场 买春堂 性之站 成人社区 群交 夫妻俱乐部 诱色uu 就去色色 熟妇 mm美图 走光偷拍 77bbb 虎骑 咪咪图片 成人导航 深爱色色 厕所偷拍 成人A片 夫妻多p 我就色 释欲 你色吗 裙内偷拍 色97爱 丝诱 色色婷婷 美体艳姿 熟母 肉丝裤袜 打野炮 就去日 爱幼阁 夫妻3p 日本熟母 大胆出位 快乐AV 强制浣肠 凸肉优 撅起大白腚 双龙入洞 美女吞精 开包 寂寞自摸 警奴 极品奶妹 惹火身材 暴力虐待 扉之阴 丁字裤 暴奸 无修正 国产AV 东热空姐 粉B 我就去色 亮穴 幼男 肉箫 白浆 性战 朝天穴 公媳乱 慰春情 嫩奶 a4y 品穴 一本道 羞耻母 艳照 三P 露毛 紧穴 露点 18禁 g片 荡女 露穴 迷药 无码 现代情色小说 性息 艳情小说 阴部特写 幼女 玉蒲团 玉女心经 援助交易 中国成人论坛 穴图 要射了 一夜性网 婬乱军团 H动漫 交换夫妻 丝袜 亚洲情色网 火辣图片 疯狂抽送 性虎色网 色迷城 petgirl 成人百强 品色堂 色窝窝 巨奶 群魔色舞 扒穴 六月联盟 55sss偷拍区 张筱雨 xiao77 极品黑丝 丝袜写真 天天情色 成人小说 成人文学 222se图片 偷拍 厕奴 人体摄影 东北xx网 玛雅网 成人bt 周六性吧 裙下风光 嘻游中国 御の二代目 足交 肮脏美学 亚洲有码 欲仙欲死 丝袜高跟 原味丝袜 针孔偷拍 拳交 迫奸 品香堂 北京xx网 虐奴 欧美无套 炮友 双飞 素人娘 前凸后翘 色色成人 胯下 呻吟 护士诱惑 群奸 极品白虎 曲线消魂 西洋美女 湿身诱惑 火辣写真 动漫色图 粉红穴 经典炮图 跨下 流蜜汁 放尿 入穴一游 丰唇艳姬 国产偷拍 雪腿玉胯 潮喷 颜骑 夫妻乱交 穴海 杂交 口射 多人轮 奶挺臀翘 浴尿 発妻 姫辱 嫩缝 暴干 母子交欢 足脚交 柔阴术 相奸 桃园蜜洞 二穴中出 奴畜抄 连続失禁 玩穴 [hz] (hz) [av] (av) [sm] (sm) sm 顶级 绝版 大片 限制 白痴 sb 赢盘 你妈 黄色网站 失身 AIDS aids Aids DICK dick Dick penis sex SM 爱滋 梅毒 吹喇叭 吹箫 机八 机巴 机掰 滥交 乱交 三陪 色情 伟哥 颜射 一夜情 换妻 Bitch cao FUCK Fuck kao NMD NND SHIT SUCK Suck tnnd K他命 笨蛋 草你 册那 侧那 测拿 蠢猪 废物 干她妈 干妳 干你 狗娘养的 贱货 贱人 烂人 老土 妈比 妈的 马的 妳老母的 妳娘的 破鞋 仆街 去她妈 去妳的 去妳妈 去你的 去死 去他妈 日你 赛她娘 赛妳娘 赛你娘 赛他娘 傻B 傻子 上妳 上你 神经病 王八蛋 我日 乡巴佬 骑你 湿了 骑他 骑她 欠骑 欠人骑 来爽我 干他 干她 干死 干爆 干机 机叭 臭机 烂鸟 览叫 摸咪咪 PENIS BITCH BLOWJOB KISSMYAS 干入 爽你 干干 干X 他干 干它 干牠 干您 干汝 干林 干尼 我咧干 干勒 干我 干到 干啦 干爽 欠干 狗干 我干 来干 轮干 轮流干 干一干 援交 奸暴 再奸 我奸 奸你 奸他 奸她 奸一奸 臭西 烂逼 大血比 戳你 逼你老母 挨球 草拟妈 卖逼 日死 奶娘 他娘 她娘 逼毛 叼你 渣波波 嫩b 真木加美 清纯 伦理 艳星 绝色 床上功夫 強暴 陈冠希 铃木麻奈美 星崎未来 东京热 菅野亚梨沙 吉岡美穗 红音 原千寻 沙雪 H动画 hgame H游戏 粉嫩小洞 放入春药 shit 牛B 傻b 老母 卖吗 同性恋 他妈的 做爱 白痴 缺德 赢盘 封杀 流产 傻比 包二奶 魔兽 婊子 少儿不宜 咪咪 酷刑 强迫 三級 三级 A片 A级 被虐 被迫 被逼 口技 緊縛 幼幼 女優 强歼 色友 蒲团 女女 喷尿 坐交 慰安妇 色狼 捅你 捅我 波霸 偷情 制服 亚热 走光 自摸 捆绑 潮吹 群射 臭作 薄格 調教 近親 連發無修正 尿尿 喷水 小泽玛莉亚 武腾兰 武藤兰 饭岛爱 小泽圆 長瀨愛 川島和津實 粉嫩的小沟 小澤園 飯島愛 星崎未來 及川奈央 松金洋子 波多野結衣 小泽玛利亚 苍井空 朝河蘭 夕樹舞子 松岛枫 大澤惠 金澤文子 三浦愛佳 伊東 武藤蘭 叶子楣 舒淇 翁虹 麻衣 櫻井 風花 星野桃 宝来 桜田 藤井彩 小森美王 平井 桃井望 榊彩弥 桜井 中条美華 大久保玲 松下 青木友梨 深田涼子 愛内萌 姫島瑠璃香 長瀬愛 中野千夏 春菜 望月 岡崎美女 宮下杏奈 加藤 日野美沙 北島優 夜勤病栋 避孕套 卖春 夜总会 坐台 婊子 情妇 喷潮 跳蛋 啪啪 asia_fox 禦姐 美鮑 金瓶梅 發情 约炮 精盆 少婦 3P 狂草 大胸 大逼 摸JB 迷干 情_色 乱X 人兽 姓欲 內射 東京熱 开房 女僕 恋足 通奸 肉偿 摸逼 抠逼"
