package main

import (
	"fmt"
	"github.com/szyhf/go-excel"
	"math/rand"
	"strconv"
	"strings"
)

type SkillExp struct {
	SkillId int `xlsx:"SkillID"`
	AddExp  int `xlsx:"addExp"`
	Level   int `xlsx:"level"`
}

type Map struct {
	Val1 int `xlsx:"val1"`
	Val2 int `xlsx:"val2"`
	Val3 int `xlsx:"val3"`
}

type Role struct {
	Level int    `xlsx:"level"`
	Job   int    `xlsx:"job"`
	Equip string `xlsx:"column(equip)"`
	Skill string `xlsx:"column(skill)"`
	Pet   string `xlsx:"column(pet)"`

	// data after processing
	SkillMap   map[int]int // id,lvl
	PetMap     map[int]int // id,lvl
	EquipSlice []int       // id

	// use data
	CauUseSk []int
}

type Config struct {
	SkillExpMap map[int]*SkillExp
	MapSlice    []*Map
	RoleSlice   []*Role
}

var (
	conf *Config
)

func (r *Role) GetSkill() (int, int) {
	skPet, skRole := 0, 0
	if len(r.CauUseSk) == 0 {
		return 0, 0
	}
	if len(r.PetMap) != 0 {
		skPet = r.CauUseSk[rand.Intn(len(r.CauUseSk))]
	}
	skRole = r.CauUseSk[rand.Intn(len(r.CauUseSk))]
	return skPet, skRole
}

func InitConf() {
	conf = &Config{
		SkillExpMap: make(map[int]*SkillExp),
	}
	readSkill()
	readMap()
	readRole()
}

func GetSkill(id int) *SkillExp {
	vv, ok := conf.SkillExpMap[id]
	if ok && vv != nil {
		return vv
	}
	return nil
}

func GetRoleInfo() *Role {
	idx := rand.Intn(len(conf.RoleSlice))
	return conf.RoleSlice[idx]
}

func GetRandomMap() string {
	idx := rand.Intn(len(conf.MapSlice))
	vv := conf.MapSlice[idx]
	return fmt.Sprintf("%d %d %d", vv.Val1, vv.Val2, vv.Val3)
}

func readSkill() {
	var val []SkillExp
	err := excel.UnmarshalXLSX("./config.xlsx", &val)
	if err != nil {
		panic(err)
	}
	for _, vv := range val {
		conf.SkillExpMap[vv.SkillId] = &vv
	}
}

func readMap() {
	var val []*Map
	err := excel.UnmarshalXLSX("./config.xlsx", &val)
	if err != nil {
		panic(err)
	}
	conf.MapSlice = make([]*Map, 0, len(val))
	for _, vv := range val {
		conf.MapSlice = append(conf.MapSlice, vv)
	}
}

func readRole() {
	var val []*Role
	err := excel.UnmarshalXLSX("./config.xlsx", &val)
	if err != nil {
		panic(err)
	}
	conf.RoleSlice = make([]*Role, 0, len(val))
	for _, vv := range val {
		vv.SkillMap = make(map[int]int)
		vv.PetMap = make(map[int]int)

		// dispose skill
		if vv.Skill != "" {
			strSk1 := strings.Split(vv.Skill, "|")
			for _, strSk2 := range strSk1 {
				strSk := strings.Split(strSk2, ":")
				if len(strSk) != 2 {
					panic(fmt.Sprintf("role skill config err data:%v", vv))
				}
				id, _ := strconv.Atoi(strSk[0])
				lvl, _ := strconv.Atoi(strSk[1])
				vv.SkillMap[id] = lvl
			}
		}

		// dispose pet
		if vv.Pet != "" {
			strPet := strings.Split(vv.Pet, "|")
			if len(strPet) != 2 {
				panic(fmt.Sprintf("role pet config err data:%v", vv))
			}
			id, _ := strconv.Atoi(strPet[0])
			lvl, _ := strconv.Atoi(strPet[1])
			vv.PetMap[id] = lvl
		}

		// dispose equip
		if vv.Equip != "" {
			strEquip := strings.Split(vv.Equip, "|")
			vv.EquipSlice = make([]int, 0, len(strEquip))
			for _, equip := range strEquip {
				id, _ := strconv.Atoi(equip)
				vv.EquipSlice = append(vv.EquipSlice, id)
			}
		}

		conf.RoleSlice = append(conf.RoleSlice, vv)
	}
}
