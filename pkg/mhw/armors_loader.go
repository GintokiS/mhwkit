package mhw

import (
	"errors"
	"fmt"
	"strings"
)

func (dm *dataManager) loadArmors() {
	// https://www.mhchinese.wiki/equipments
	const rawData = `
0 无 0 0 0

# 蛮颚
# 泥鱼

#----------------------------------------
@rare 5

@bonus 无
@resistences 0 -3 3 0 0
@set 飛雷龍α
1 飛雷龍頭盔α 				0 0 0 體術 +2
2 飛雷龍鎧甲α 				0 0 0 跳躍鐵人 +1 迴避距離UP +1
3 飛雷龍腕甲α 				0 0 0 迴避距離UP +2
4 飛雷龍腰甲α 				0 0 0 雷屬性攻擊強化 +2 雷耐性 +1
5 飛雷龍護腿α 				0 0 0 雷耐性 +1 體術 +1 
@set 飛雷龍β
1 飛雷龍頭盔β 				2 0 0 體術 +1
2 飛雷龍鎧甲β 				2 0 0 跳躍鐵人 +1
3 飛雷龍腕甲β 				2 0 0 迴避距離UP +1
4 飛雷龍腰甲β 				1 1 0 雷屬性攻擊強化 +1
5 飛雷龍護腿β 				1 1 0 雷耐性 +1
@bonus 无
@resistences -3 2 0 1 0
@set 浮空龍α
1 浮空龍頭盔α 				0 0 0 耐力急速回復 +2
2 浮空龍鎧甲α 				0 0 0 騎乘名人 +1 砲術 +1
3 浮空龍腕甲α 				0 0 0 砲術 +2
4 浮空龍腰甲α 				0 0 0 精靈加護 +2
5 浮空龍護腿α 				0 0 0 耐力急速回復 +1 風壓耐性 +1
@set 浮空龍β
1 浮空龍頭盔β 				2 0 0 耐力急速回復 +1
2 浮空龍鎧甲β 				1 0 0 騎乘名人 +1
3 浮空龍腕甲β 				1 0 0 砲術 +1
4 浮空龍腰甲β 				1 0 0 精靈加護 +1
5 浮空龍護腿β 				1 1 0 風壓耐性 +1
@bonus 无
@resistences 2 0 -2 0 -3
@set 雌火龍α
1 雌火龍頭盔α 				0 0 0 體力增強 +2 毒屬性強化 +1
2 雌火龍鎧甲α 				0 0 0 植生學 +2 體力回復量UP +1
3 雌火龍腕甲α 				0 0 0 毒屬性強化 +2
4 雌火龍腰甲α 				0 0 0 毒耐性 +2 體力增強 +1
5 雌火龍護腿α 				0 0 0 體力回復量UP +2 
@set 雌火龍β
1 雌火龍頭盔β 				1 0 0 體力增強 +1 毒屬性強化 +1
2 雌火龍鎧甲β 				1 0 0 植生學 +1 體力回復量UP +1
3 雌火龍腕甲β 				1 0 0 毒屬性強化 +1
4 雌火龍腰甲β 				1 0 0 毒耐性 +1 體力增強 +1
5 雌火龍護腿β 				1 0 0 體力回復量UP +1

#----------------------------------------
@rare 6
@bonus 櫻火龍之奧秘
@resistences 3 0 0 -3 -4
@set 火龍心α
1 火龍心頭盔α 				0 0 0 迴避距離UP +1 毒屬性強化 +1
2 火龍心鎧甲α 				0 0 0 整備 +2
3 火龍心腕甲α 				0 0 0 風壓耐性 +2
4 火龍心腰甲α 				0 0 0 匠 +1 毒耐性 +1
5 火龍心護腿α 				0 0 0 整備 +1 風壓耐性 +1
@set 火龍心β
1 火龍心頭盔β 				2 0 0 迴避距離UP +1
2 火龍心鎧甲β 				1 1 0 整備 +1
3 火龍心腕甲β 				2 0 0 風壓耐性 +1
4 火龍心腰甲β 				1 0 0 匠 +1
5 火龍心護腿β 				2 0 0 整備 +1
@bonus 无
@resistences 0 2 0 -1 -2
@set 骨鎚龍α
1 骨鎚龍頭盔α 				0 0 0 耐震 +2
2 骨鎚龍鎧甲α 				0 0 0 睡眠屬性強化 +1 裂傷耐性 +1
3 骨鎚龍腕甲α 				0 0 0 睡眠屬性強化 +2 防禦性能 +1
4 骨鎚龍腰甲α 				0 0 0 防禦性能 +2
5 骨鎚龍護腿α 				0 0 0 裂傷耐性 +2
@set 骨鎚龍β
1 骨鎚龍頭盔β 				2 0 0 耐震 +1
2 骨鎚龍鎧甲β 				1 0 0 睡眠屬性強化 +1
3 骨鎚龍腕甲β 				1 0 0 睡眠屬性強化 +1 防禦性能 +1
4 骨鎚龍腰甲β 				2 0 0 防禦性能 +1
5 骨鎚龍護腿β 				1 0 0 裂傷耐性 +1

#----------------------------------------
@rare 5
@bonus 无
@set 无
@resistences 3 0 0 0 0
1 突擊龍角α					0 0 0 減輕膽怯 +1 爆破異常狀態的耐性 +1
1 突擊龍角β					1 0 0 減輕膽怯 +1
@resistences 2 0 -3 0 2
2 酸翼龍斗篷α					0 0 0 熱傷害無效 +1 防御力DOWN耐性 +1
2 酸翼龍斗篷β					1 0 0 熱傷害無效 +1
@bonus 无
@resistences 2 0 -3 0 2
@set 岩賊龍α
1 岩賊龍頭盔α 				0 0 0 爆破異常狀態的耐性 +2
2 岩賊龍鎧甲α 				0 0 0 爆破屬性強化 +2
3 岩賊龍腕甲α 				0 0 0 炸彈客 +2
4 岩賊龍腰甲α 				0 0 0 砲術 +2 爆破異常狀態的耐性 +1
5 岩賊龍護腿α 				0 0 0 砲彈裝填數UP +1 爆破屬性強化 +1
@set 岩賊龍β
1 岩賊龍頭盔β 				1 0 0 爆破異常狀態的耐性 +1
2 岩賊龍鎧甲β 				1 0 0 爆破屬性強化 +1
3 岩賊龍腕甲β 				1 0 0 炸彈客 +1
4 岩賊龍腰甲β 				1 0 0 砲術 +2
5 岩賊龍護腿β 				1 0 0 砲彈裝填數UP +1

#----------------------------------------
@rare 6
@bonus 熔山龍的奧秘
@resistences 4 -3 -1 -2 -3
@set 鑄島熔岩α
1 鑄島熔岩頭盔α 				0 0 0 匠 +1 爆破屬性強化 +1
2 鑄島熔岩鎧甲α 				0 0 0 爆破屬性強化 +2 風壓耐性 +1
3 鑄島熔岩腕甲α 				0 0 0 屬性解放／裝填擴充 +1 炸彈客 +1
4 鑄島熔岩腰甲α 				0 0 0 不屈 +1 耐震 +1
5 鑄島熔岩護腿α 				0 0 0 炸彈客 +1 耳塞 +1
@set 鑄島熔岩β
1 鑄島熔岩頭盔β 				1 0 0 匠 +1
2 鑄島熔岩鎧甲β 				2 0 0 爆破屬性強化 +2
3 鑄島熔岩腕甲β 				1 0 0 屬性解放／裝填擴充 +1
4 鑄島熔岩腰甲β 				2 0 0 不屈 +1
5 鑄島熔岩護腿β 				2 0 0 炸彈客 +2
@bonus 无
@resistences 0 -1 -2 3 0
@set 礦石α
1 礦石頭盔α 					0 0 0 冰耐性 +2 耳塞 +1
2 礦石鎧甲α 					0 0 0 耳塞 +2 冰屬性攻擊強化 +1
3 礦石腕甲α 					0 0 0 防禦 +2 冰屬性攻擊強化 +1
4 礦石腰甲α 					0 0 0 砲彈裝填數UP +1 防禦 +1
5 礦石護腿α 					0 0 0 砲擊手 +1 冰耐性 +1
@set 礦石β
1 礦石頭盔β 					2 0 0 冰耐性 +2
2 礦石鎧甲β 					2 0 0 耳塞 +1 冰屬性攻擊強化 +1
3 礦石腕甲β 					1 0 0 防禦 +2
4 礦石腰甲β 					1 0 0 砲彈裝填數UP +1
5 礦石護腿β 					1 0 0 砲擊手 +1
@bonus 无
@resistences -2 0 3 -1 0
@set 鑄鐵α
1 鑄鐵頭盔α 					0 0 0 雷屬性攻擊強化 +2
2 鑄鐵鎧甲α 					0 0 0 雷耐性 +2
3 鑄鐵腕甲α 					0 0 0 體力增強 +2 雷屬性攻擊強化 +1
4 鑄鐵腰甲α 					0 0 0 風壓耐性 +2
5 鑄鐵護腿α 					0 0 0 防禦 +2 風壓耐性 +1
@set 鑄鐵β
1 鑄鐵頭盔β 					1 0 0 雷屬性攻擊強化 +1
2 鑄鐵鎧甲β 					1 0 0 雷耐性 +1
3 鑄鐵腕甲β 					1 0 0 體力增強 +2
4 鑄鐵腰甲β 					3 0 0 風壓耐性 +1
5 鑄鐵護腿β 					2 0 0 防禦 +2
@bonus 无
@resistences 1 0 2 0 2
@set 旅團α
1 旅團頭盔α 					0 0 0 回復速度 +2 吹笛名人 +1
2 旅團鎧甲α 					2 0 0 精靈加護 +2
3 旅團腕甲α 					0 0 0 砲擊手 +2 精靈加護 +1
4 旅團腰甲α 					0 0 0 地質學 +2 回復速度 +1
5 旅團護腿α 					1 0 0 研究者 +1 地質學 +1
@set 旅團β
1 旅團頭盔β 					1 0 0 回復速度 +2
2 旅團鎧甲β 					2 2 0 精靈加護 +1
3 旅團腕甲β 					1 0 0 砲擊手 +2
4 旅團腰甲β 					1 0 0 地質學 +2
5 旅團護腿β 					3 0 0 研究者 +1
@bonus 无
@set 无
@resistences 2 2 2 2 2
1 艾路頭套α 					0 0 0 指示隨從 +3
@resistences 2 2 2 2 2
1 知性眼鏡α					1 0 0 看破 +2
@bonus 无
@resistences 3 -3 1 -2 2
@set 熔岩龍α
1 熔岩龍頭盔α 				0 0 0 火屬性攻擊強化 +2 收刀術 +1
2 熔岩龍鎧甲α 				0 0 0 火屬性攻擊強化 +2 熱傷害無效 +1
3 熔岩龍腕甲α 				2 0 0 滑走強化 +1 廣域化 +1
4 熔岩龍腰甲α 				0 0 0 收刀術 +2 體力回復量UP +1
5 熔岩龍護腿α 				1 0 0 散彈・剛射強化 +1 火屬性攻擊強化 +1
@set 熔岩龍β
1 熔岩龍頭盔β 				1 1 0 火屬性攻擊強化 +1 收刀術 +1
2 熔岩龍鎧甲β 				1 1 0 火屬性攻擊強化 +2
3 熔岩龍腕甲β 				3 0 0 滑走強化 +1
4 熔岩龍腰甲β 				1 0 0 收刀術 +2
5 熔岩龍護腿β 				2 0 0 散彈・剛射強化 +1
@bonus 風漂龍的恩竉
@resistences -1 2 -3 3 0
@set 風漂龍α
1 風漂龍頭盔α 				0 0 0 精靈加護 +2 冰屬性攻擊強化 +1
2 風漂龍鎧甲α 				0 0 0 迴避性能 +2 精靈加護 +1
3 風漂龍腕甲α 				0 0 0 飛燕 +1 迴避性能 +1
4 風漂龍腰甲α 				0 0 0 冰屬性攻擊強化 +2 風壓耐性 +1
5 風漂龍護腿α 				0 0 0 冰耐性 +2 拔刀術【技】 +1
@set 風漂龍β
1 風漂龍頭盔β 				1 0 0 精靈加護 +2
2 風漂龍鎧甲β 				1 0 0 迴避性能 +2
3 風漂龍腕甲β 				3 0 0 飛燕 +1
4 風漂龍腰甲β 				2 0 0 冰屬性攻擊強化 +2
5 風漂龍護腿β 				3 0 0 冰耐性 +2
@bonus 慘爪龍的奧祕
@resistences 2 2 -2 -3 0
@set 慘爪龍α
1 慘爪龍頭盔α 				0 0 0 裂傷耐性 +2 看破 +1
2 慘爪龍鎧甲α 				0 0 0 砥石使用高速化 +2 裂傷耐性 +1
3 慘爪龍腕甲α 				0 0 0 體術 +1 看破 +1
4 慘爪龍腰甲α 				0 0 0 看破 +2 砥石使用高速化 +1
5 慘爪龍護腿α 				0 0 0 收刀術 +1 看破 +1
@set 慘爪龍β
1 慘爪龍頭盔β 				2 0 0 裂傷耐性 +2
2 慘爪龍鎧甲β 				1 0 0 砥石使用高速化 +2
3 慘爪龍腕甲β 				1 0 0 體術 +1
4 慘爪龍腰甲β 				2 0 0 看破 +2
5 慘爪龍護腿β 				1 0 0 收刀術 +1
@bonus 火龍奧祕
@resistences 3 1 -2 1 -3
@set 火龍α
1 火龍頭盔α 					0 0 0 攻擊 +2 火屬性攻擊強化 +1
2 火龍鎧甲α 					0 0 0 弱點特效 +2 火屬性攻擊強化 +1
3 火龍腕甲α 					0 0 0 火耐性 +2 攻擊 +1
4 火龍腰甲α 					0 0 0 火屬性攻擊強化 +2 集中 +1
5 火龍護腿α 					0 0 0 跳躍鐵人 +1 弱點特效 +1
@set 火龍β
1 火龍頭盔β 					1 0 0 攻擊 +2
2 火龍鎧甲β 					1 0 0 弱點特效 +2
3 火龍腕甲β 					1 0 0 火耐性 +2
4 火龍腰甲β 					2 0 0 火屬性攻擊強化 +2
5 火龍護腿β 					3 0 0 跳躍鐵人 +1

#----------------------------------------
@rare 7
@bonus 火龍奧祕
@resistences 3 2 3 -3 -4
@set 火龍魂α
1 火龍魂頭盔α 				0 0 0 超會心 +1 威嚇 +2
2 火龍魂鎧甲α 				1 0 0 屬性解放／裝填擴充 +1 威嚇 +1
3 火龍魂腕甲α 				1 0 0 集中 +1 風壓耐性 +1
4 火龍魂腰甲α 				0 0 0 風壓耐性 +2 貫通彈・龍之箭強化 +1
5 火龍魂護腿α 				0 0 0 集中 +1 攀岩者 +1 
@set 火龍魂β
1 火龍魂頭盔β 				1 1 0 超會心 +1
2 火龍魂鎧甲β 				2 0 0 屬性解放／裝填擴充 +1
3 火龍魂腕甲β 				2 0 0 集中 +1
4 火龍魂腰甲β 				3 0 0 風壓耐性 +2
5 火龍魂護腿β 				1 0 0 集中 +1

#----------------------------------------
@rare 6
@bonus 角龍之奧祕
@resistences 3 -2 0 -3 2
@set 角龍α
1 角龍頭盔α 					0 0 0 拔刀術【技】 +2 火場怪力 +1
2 角龍鎧甲α 					0 0 0 KO術 +2 火場怪力 +1
3 角龍腕甲α 					0 0 0 火場怪力 +2 耐震 +1
4 角龍腰甲α 					0 0 0 跑者 +2 拔刀術【技】 +1
5 角龍護腿α 					0 0 0 耐震 +2 KO術 +1
@set 角龍β
1 角龍頭盔β 					2 0 0 拔刀術【技】 +2
2 角龍鎧甲β 					2 0 0 KO術 +2
3 角龍腕甲β 					2 0 0 火場怪力 +2
4 角龍腰甲β 					2 0 0 跑者 +2
5 角龍護腿β 					1 1 0 耐震 +2

#----------------------------------------
#----------------------------------------
#----------------------------------------

#----------------------------------------
@rare 7
@bonus 角龍之奧祕
@resistences 2 -3 0 -4 2
@set 暴君角龍α
1 暴君角龍頭盔α				1 0 0 集中 +2
2 暴君角龍鎧甲α				1 0 0 怨恨 +2
3 暴君角龍腕甲α				0 0 0 集中 +1 跑者 +1
4 暴君角龍腰甲α				0 0 0 怨恨 +1 破壞王 +1
5 暴君角龍護腿α				1 0 0 通常彈・通常箭強化 +1 火場怪力 +1 
@set 暴君角龍β
1 暴君角龍頭盔β				3 0 0 集中 +1
2 暴君角龍鎧甲β				2 1 0 怨恨 +1
3 暴君角龍腕甲β				2 1 0 集中 +1
4 暴君角龍腰甲β				2 0 0 怨恨 +1
5 暴君角龍護腿β				2 0 0 通常彈・通常箭強化 +1 
@bonus 爆鎚龍的守護
@resistences 4 -3 1 -2 -2
@set 爆鎚龍α
1 爆鎚龍頭盔α 				3 0 0 雷耐性 +2
2 爆鎚龍鎧甲α 				1 0 0 破壞王 +2
3 爆鎚龍腕甲α 				1 0 0 防禦性能 +1 防禦 +1
4 爆鎚龍腰甲α 				1 1 0 破壞王 +1 雷耐性 +1
5 爆鎚龍護腿α 				1 0 0 防禦性能 +2
@set 爆鎚龍β
1 爆鎚龍頭盔β 				3 1 0 雷耐性 +1
2 爆鎚龍鎧甲β 				3 0 0 破壞王 +1
3 爆鎚龍腕甲β 				2 0 0 防禦性能 +1
4 爆鎚龍腰甲β 				2 1 0 破壞王 +1
5 爆鎚龍護腿β 				3 0 0 防禦性能 +1
@bonus 爆鱗龍的守護
@resistences 3 1 -4 -2 -2
@set 爆鱗龍α
1 爆鱗龍頭盔α					2 0 0 耳塞 +2 爆破屬性強化 +1
2 爆鱗龍鎧甲α					0 0 0 防禦性能 +2 爆破異常狀態的耐性 +2
3 爆鱗龍腕甲α					0 0 0 拔刀術【技】 +2 耳塞 +1
4 爆鱗龍腰甲α					1 0 0 耳塞 +2 跳躍鐵人 +1
5 爆鱗龍護腿α					1 0 0 炸彈客 +2 拔刀術【技】 +1
@set 爆鱗龍β
1 爆鱗龍頭盔β					3 1 0 耳塞 +2
2 爆鱗龍鎧甲β					1 1 0 防禦性能 +2
3 爆鱗龍腕甲β					3 0 0 拔刀術【技】 +2
4 爆鱗龍腰甲β					3 0 0 耳塞 +2
5 爆鱗龍護腿β					3 0 0 炸彈客 +2 
@bonus 无
@resistences 4 -2 0 0 0
@set 大馬士革α
1 大馬士革頭盔α				1 0 0 防禦 +2 砥石使用高速化 +1
2 大馬士革鎧甲α				1 0 0 集中 +2 防禦 +2
3 大馬士革腕甲α				0 0 0 匠 +1 砥石使用高速化 +1
4 大馬士革腰甲α				1 0 0 集中 +1 防御力DOWN耐性 +2
5 大馬士革護腿α				1 0 0 防禦 +2 防御力DOWN耐性 +1
@set 大馬士革β
1 大馬士革頭盔β 				3 0 0 防禦 +2
2 大馬士革鎧甲β 				1 1 1 集中 +2
3 大馬士革腕甲β 				1 0 0 匠 +1
4 大馬士革腰甲β 				1 1 1 集中 +1
5 大馬士革護腿β 				3 0 0 防禦 +2 
@bonus 无
@resistences -2 0 -1 -1 4
@set 杜賓α
1 杜賓頭盔α					0 0 0 屬性解放／裝填擴充 +1 龍耐性 +1
2 杜賓鎧甲α					0 0 0 攻擊 +2 最愛菇類 +1
3 杜賓腕甲α					0 0 0 耐力急速回復 +2 最愛菇類 +1
4 杜賓腰甲α					0 0 0 龍耐性 +2 攻擊 +1
5 杜賓護腿α					0 0 0 攻擊 +2 耐力急速回復 +1
@set 杜賓β
1 杜賓頭盔β					1 0 0 屬性解放／裝填擴充 +1
2 杜賓鎧甲β					3 0 0 攻擊 +2
3 杜賓腕甲β					1 1 0 耐力急速回復 +1 最愛菇類 +1
4 杜賓腰甲β					3 0 0 龍耐性 +2
5 杜賓護腿β					1 0 0 攻擊 +2
@bonus 无
@resistences 3 2 -2 3 -4
@set 死神α
1 死神首腦α					2 0 0 死裡逃生 +1 看破 +1
2 死神肌肉α					0 0 0 不屈 +1 屬性解放／裝填擴充 +1
3 死神雙手α					1 0 0 集中 +2 龍屬性攻擊強化 +1
4 死神臍帶α					0 0 0 龍屬性攻擊強化 +2 集中 +1
5 死神腳跟α					0 0 0 匠 +2 龍屬性攻擊強化 +2
@set 死神β
1 死神首腦β					3 0 0 死裡逃生 +1
2 死神肌肉β					3 0 0 不屈 +1
3 死神雙手β					1 1 0 集中 +2
4 死神臍帶β					1 1 0 龍屬性攻擊強化 +2
5 死神腳跟β					1 1 0 匠 +2
@bonus 无
@resistences -3 -1 -1 -1 2
@set 无
1 骷髏面罩α					2 0 0 匠 +1

#----------------------------------------
@rare 8
@bonus 滅盡龍之飢餓
@resistences 1 1 -3 1 -3
@set 戰紋α
1 戰紋頭盔α					1 0 0 精神抖擻 +2 攻擊 +1
2 戰紋鎧甲α					1 0 0 耐力急速回復 +2 挑戰者 +1
3 戰紋腕甲α					0 0 0 挑戰者 +2 攻擊 +1
4 戰紋腰甲α					0 0 0 攻擊 +2 耐力急速回復 +1
5 戰紋護腿α					1 0 0 精神抖擻 +1 挑戰者 +1
@set 戰紋β
1 戰紋頭盔β					1 1 0 精神抖擻 +2
2 戰紋鎧甲β					2 1 0 耐力急速回復 +2
3 戰紋腕甲β					1 0 0 挑戰者 +2
4 戰紋腰甲β					2 0 0 攻擊 +2
5 戰紋護腿β					2 0 0 精神抖擻 +1
@bonus 炎王龍之武技
@resistences 3 -3 1 -3 1
@set 帝王α
1 帝王皇冠α					0 0 0 力量解放 +2
2 帝王鎧甲α					0 0 0 特殊射擊強化 +2 力量解放 +1
3 帝王腕甲α					1 0 0 弱點特效 +2
4 帝王腰甲α					0 0 0 爆破屬性強化 +2 弱點特效 +1
5 帝王護腿α					0 0 0 力量解放 +2 爆破屬性強化 +1
@set 帝王β
1 帝王皇冠β					2 0 0 力量解放 +1
2 帝王鎧甲β					1 1 0 特殊射擊強化 +2
3 帝王腕甲β					3 0 0 弱點特效 +1
4 帝王腰甲β					2 0 0 爆破屬性強化 +2
5 帝王護腿β					1 0 0 力量解放 +2
@bonus 鋼龍的飛翔
@resistences 0 2 -3 4 -2
@set 鋼龍α
1 鋼龍強力α					0 0 0 冰屬性攻擊強化 +2 匠 +1
2 鋼龍恐懼α					0 0 0 匠 +2 集中 +1
3 鋼龍剛強α					0 0 0 迴避性能 +2 匠 +1
4 鋼龍安穩α					0 0 0 冰屬性攻擊強化 +2 迴避距離UP +1
5 鋼龍踏實α					0 0 0 迴避距離UP +2 匠 +1
@set 鋼龍β
1 鋼龍強力β					1 1 1 冰屬性攻擊強化 +2
2 鋼龍恐懼β					2 0 0 匠 +2
3 鋼龍剛強β					3 0 0 迴避性能 +2
4 鋼龍安穩β					1 1 0 冰屬性攻擊強化 +2
5 鋼龍踏實β					3 0 0 迴避距離UP +2
@bonus 屍套龍之命脈
@resistences -4 4 1 -1 -3
@set 烏爾德α
1 烏爾德頭盔α					0 0 0 無傷 +1 龍屬性攻擊強化 +2
2 烏爾德鎧甲α					1 0 0 回復速度 +2 龍屬性攻擊強化 +1
3 烏爾德腕甲α					1 0 0 無傷 +1 回復速度 +1
4 烏爾德腰甲α					1 0 0 無傷 +1 瘴氣耐性 +1
5 烏爾德護腿α					1 0 0 瘴氣耐性 +2 龍屬性攻擊強化 +1
@set 烏爾德β
1 烏爾德頭盔β					1 0 0 無傷 +1
2 烏爾德鎧甲β					1 1 0 回復速度 +2
3 烏爾德腕甲β					1 1 1 無傷 +1
4 烏爾德腰甲β					1 1 0 無傷 +1
5 烏爾德護腿β					1 1 0 瘴氣耐性 +2
@bonus 麒麟之恩寵
@resistences -3 -2 4 -2 2
@set 麒麟α
1 麒麟角α						0 0 0 跑者 +2 精靈加護 +1
2 麒麟服飾α					0 0 0 精靈加護 +2 雷屬性攻擊強化 +1
3 麒麟長腕甲α					0 0 0 雷屬性攻擊強化 +2 跑者 +1
4 麒麟腰環α					0 0 0 屬性異常狀態耐性 +1 屬性解放／裝填擴充 +1
5 麒麟護脛α					0 0 0 屬性解放／裝填擴充 +2 雷耐性 +2
@set 麒麟β
1 麒麟角β						1 0 0 跑者 +2
2 麒麟服飾β					1 0 0 精靈加護 +2
3 麒麟長腕甲β					1 1 0 雷屬性攻擊強化 +2
4 麒麟腰環β					1 1 1 屬性異常狀態耐性 +1
5 麒麟護脛β					1 1 0 屬性解放／裝填擴充 +2
@bonus 冥燈龍的神秘
@resistences -3 2 2 2 -4
@set 異種大型α
1 異種大型頭飾α				1 0 0 減輕膽怯 +1 特殊射擊強化 +1
2 異種大型皮α					1 0 0 強化持續 +2 屬性異常狀態耐性 +1
3 異種大型鋼爪α				1 0 0 減輕膽怯 +1 超會心 +1
4 異種大型脊椎α				2 0 0 屬性異常狀態耐性 +2 特殊射擊強化 +1
5 異種大型靴α					1 0 0 減輕膽怯 +1 強化持續 +1
@set 異種大型β
1 異種大型頭飾β				2 0 0 減輕膽怯 +1
2 異種大型皮β					1 1 0 強化持續 +2
3 異種大型鋼爪β				3 0 0 減輕膽怯 +1
4 異種大型脊椎β				2 1 0 屬性異常狀態耐性 +2
5 異種大型靴β					1 1 1 減輕膽怯 +1
@bonus 公會的指引
@resistences 0 0 0 0 0
@set 公會十字α
1 公會十字頭飾α				1 0 0 精靈加護 +2 雷耐性 +2
2 公會十字戰衣α				1 0 0 死裡逃生 +1 冰耐性 +2
3 公會十字腕甲α				1 0 0 整備 +1 水耐性 +2
4 公會十字腰甲α				1 0 0 不屈 +1 龍耐性 +2
5 公會十字靴α					1 0 0 飛燕 +1 火耐性 +2
@set 公會十字β
1 公會十字頭飾β				1 1 1 精靈加護 +2
2 公會十字戰衣β				3 0 0 死裡逃生 +1
3 公會十字腕甲β				3 0 0 整備 +1
4 公會十字腰甲β				1 1 1 不屈 +1
5 公會十字靴β					3 0 0 飛燕 +1
@bonus 无
@resistences 0 0 0 0 0
@set 无
1 龍王的獨眼α					3 0 0 弱點特效 +2
@bonus 調查團的指引
@resistences 2 2 2 2 2
@set 追蹤者α
1 追蹤者頭飾α					1 0 0 整備 +1 採集達人 +1
2 追蹤者服飾α					0 0 0 廣域化 +2 攀岩者 +1
3 追蹤者手套α					1 0 0 閃光強化 +1 剝取鐵人 +1
4 追蹤者皮帶α					1 0 0 最愛菇類 +1 搬運達人 +1
5 追蹤者長褲α					0 0 0 潛伏 +2 滑走強化 +1
@set 追蹤者β
1 追蹤者頭飾β					2 0 0 整備 +1
2 追蹤者服飾β					1 0 0 廣域化 +2
3 追蹤者手套β					2 0 0 閃光強化 +1
4 追蹤者皮帶β					2 0 0 最愛菇類 +1
5 追蹤者長褲β					2 0 0 潛伏 +2

#----------------------------------------
@rare 7
@bonus 无
@resistences 1 1 -3 2 -3
@set 殘虐α
1 殘虐頭盔α					1 0 0 破壞王 +1 匠 +1
2 殘虐鎧甲α					2 0 0 快吃 +1 匠 +1
3 殘虐腕甲α					0 0 0 力量解放 +2 破壞王 +1
4 殘虐腰甲α					1 0 0 力量解放 +1 匠 +1
5 殘虐護腿α					0 0 0 快吃 +2 破壞王 +1
@set 殘虐β
1 殘虐頭盔β					3 1 0 破壞王 +1
2 殘虐鎧甲β					3 0 0 快吃 +1
3 殘虐腕甲β					2 0 0 力量解放 +2
4 殘虐腰甲β					3 1 0 力量解放 +1
5 殘虐護腿β					1 1 0 快吃 +2

#----------------------------------------
@rare 8
@bonus 无
@resistences 4 -2 -3 4 -2
@set 絢輝龍鎧羅α
1 絢輝龍鎧羅頭盔α				1 0 0 屬性解放／裝填擴充 +1 挑戰者 +2
2 絢輝龍鎧羅鎧甲α				1 0 0 昏厥耐性 +2 超會心 +2
3 絢輝龍鎧羅腕甲α				1 0 0 屬性解放／裝填擴充 +1 強化持續 +2
4 絢輝龍鎧羅腰甲α				1 0 0 屬性解放／裝填擴充 +1 匠 +2
5 絢輝龍鎧羅護腿α				1 0 0 廣域化 +2 無傷 +2
@set 絢輝龍鎧羅β
1 絢輝龍鎧羅頭盔β				2 1 0 屬性解放／裝填擴充 +1 挑戰者 +1
2 絢輝龍鎧羅鎧甲β				3 1 0 昏厥耐性 +2 超會心 +1
3 絢輝龍鎧羅腕甲β				3 1 0 屬性解放／裝填擴充 +1 強化持續 +1
4 絢輝龍鎧羅腰甲β				3 1 0 屬性解放／裝填擴充 +1 匠 +1
5 絢輝龍鎧羅護腿β				2 1 0 廣域化 +2 無傷 +1
@bonus 无
@resistences 0 0 0 0 0
@set 无

#----------------------------------------
@rare 5
# 巴魯特α
# 潛水夫α
# 獨角仙后α
# 獨角仙后β

#----------------------------------------
@rare 8
@bonus 屍套龍之命脈
@resistences -4 4 1 -1 -3
@set 烏爾德λ
1 烏爾德頭盔λ					3 0 0 體力回復量UP +2 龍屬性攻擊強化 +1
2 烏爾德鎧甲λ					3 0 0 無傷 +1 體力回復量UP +1
3 烏爾德腕甲λ					3 2 0 龍屬性攻擊強化 +3
4 烏爾德腰甲λ					2 1 1 回復速度 +2 龍屬性攻擊強化 +1
5 烏爾德護腿λ					3 1 0 無傷 +1 回復速度 +1
@bonus 麒麟之恩寵
@resistences -3 -2 4 -2 2
@set 麒麟λ
1 麒麟角λ						2 0 0 屬性異常狀態耐性 +1 雷屬性攻擊強化 +3
2 麒麟服飾λ					1 1 0 屬性解放／裝填擴充 +2
3 麒麟長腕甲λ					2 2 0 屬性異常狀態耐性 +1 屬性解放／裝填擴充 +1
4 麒麟腰環λ					1 1 0 跑者 +2
5 麒麟護脛λ					3 0 0 屬性異常狀態耐性 +1 雷屬性攻擊強化 +2
@bonus 炎王龍之武技
@resistences 3 -3 1 -3 1
@set 帝王λ
1 帝王皇冠λ					2 2 0 看破 +2
2 帝王鎧甲λ					2 0 0 力量解放 +2 弱點特效 +1
3 帝王腕甲λ					3 0 0 看破 +3
4 帝王腰甲λ					1 1 0 力量解放 +3
5 帝王護腿λ					3 1 0 看破 +2
@bonus 无
@resistences 3 -3 3 -3 3
@set 但丁α
1 但丁假髮α					1 0 0 弱點特效 +1 看破 +1
2 但丁風衣α					3 0 0 迴避性能 +1 看破 +1
3 但丁手套α					3 0 0 迴避性能 +1 特殊射擊強化 +1
4 但丁皮带α					3 0 0 迴避性能 +1 力量解放 +1
5 但丁皮裤靴α					1 1 0 弱點特效 +2
@bonus 炎妃龍的恩寵
@resistences 3 1 1 -3 -2
@set 皇后α
1 皇后琴弦α					3 1 0 迴避距離UP +2 整備 +1
2 皇后鎧甲α					1 1 0 無傷 +2 體力增強 +1
3 皇后腕甲α					3 0 0 廣域化 +2 爆破屬性強化 +2
4 皇后腰甲α					2 1 0 整備 +2 爆破屬性強化 +1
5 皇后護腿α					2 1 0 體力增強 +2 廣域化 +2
@set 皇后β
1 皇后琴弦β					3 2 0 迴避距離UP +2
2 皇后鎧甲β					1 1 1 無傷 +2
3 皇后腕甲β					3 2 0 廣域化 +2
4 皇后腰甲β					2 2 0 整備 +2
5 皇后護腿β					2 2 0 體力增強 +2
@bonus 龍騎士之證
@resistences -2 -2 3 -2 4
@set 騰龍α
1 騰龍戰盔α					3 1 0 超會心 +1 飛燕 +1
2 騰龍戰鎧α					3 0 0 看破 +2 超會心 +1
3 騰龍腕甲α					2 2 0 看破 +2 攻擊 +1
4 騰龍腰甲α					3 0 0 看破 +2 強化持續 +1
5 騰龍護腿α					2 0 0 攻擊 +2 超會心 +1

#----------------------------------------
@rare 5
@bonus 无
@set 无
@resistences 3 2 0 2 0
1 封印的眼罩α					2 1 1 火耐性 +1
@resistences 4 0 4 2 0
1 墨镜α						1 1 1 昏厥耐性 +1
@resistences 2 0 2 3 2
1 扒手龍頭套α					3 1 0 搬運達人 +1
@resistences 0 4 0 0 4
1 搖曳鰻頭套α					1 1 0 廣域化 +2

#----------------------------------------
# unknown
@bonus 无
@set 无
@resistences 3 2 0 2 0
1 龍封的耳環α					2 1 0 龍封力強化 +1 
`

	if len(dm.charms) > 0 {
		panic(errors.New(fmt.Sprintf("duplicated armors loading")))
	}

	scanner := newDataScannerFromString(rawData)
	var armorSetBonusId int
	var resistences [resistenceTypeCount]int
	for scanner.scanRow() {
		switch {
		case strings.HasPrefix(scanner.line, "@"):
			scanner.scanField()
			switch scanner.textValue()[1:] {
			case "set":
			case "bonus":
				for scanner.scanField() {
					switch scanner.fieldIndex {
					case 1:
						armorSetBonusId = dm.getArmorSetIdByName(scanner.textValue())
					}
				}
			case "resistences":
				for scanner.scanField() {
					field := scanner.fieldIndex
					switch field {
					case 1, 2, 3, 4, 5:
						resistences[field-1] = scanner.intValue()
					}
				}
			}
		default:
			a := newArmor()
			a.setBonusId = armorSetBonusId
			copy(a.resistences[:], resistences[:])
			slotLayout := [3]int{}
			for scanner.scanField() {
				field := scanner.fieldIndex
				switch field {
				case 0:
					a.component = scanner.intValue()
				case 1:
					a.name = scanner.textValue()
				case 2, 3, 4:
					slotSize := scanner.intValue()
					if slotSize < 0 || slotSize > 3 {
						panic(errors.New(fmt.Sprintf("invalid slot size %v", slotSize)))
					}
					slotLayout[field-2] = slotSize
				case 5, 7:
					a.skillEnhancements[(field-5)/2].skillId = dm.getSkillIdByName(scanner.textValue())
				case 6, 8:
					a.skillEnhancements[(field-6)/2].enhancedLevel = parseCheckedInt(scanner.textValue())
				default:
					fmt.Println("unexpected column near: ", scanner.line)
					panic(errors.New("unexpected column"))
				}
			}
			slotCombinationId := dm.getSlotCombinationIdByLayout(slotLayout)
			a.slotCombination = *dm.getSlotCombinationById(slotCombinationId)
			dm.registerArmor(*a)
		}
	}
}
