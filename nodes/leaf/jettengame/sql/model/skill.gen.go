// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSkill = "skill"

// Skill mapped from table <skill>
type Skill struct {
	ID            int64  `gorm:"column:id;primaryKey;comment:技能ID" json:"id"`          // 技能ID
	Name          string `gorm:"column:name;comment:名称" json:"name"`                   // 名称
	Attack        int64  `gorm:"column:attack;comment:攻击力" json:"attack"`              // 攻击力
	Spellpower    int64  `gorm:"column:spellpower;comment:法强" json:"spellpower"`       // 法强
	Coefficient   int64  `gorm:"column:coefficient;comment:系数" json:"coefficient"`     // 系数
	Casttime      int32  `gorm:"column:casttime;comment:施法前摇" json:"casttime"`         // 施法前摇
	Introduce     string `gorm:"column:introduce;comment:技能介绍" json:"introduce"`       // 技能介绍
	Counterattack int64  `gorm:"column:counterattack;comment:反击" json:"counterattack"` // 反击
	Stun          int64  `gorm:"column:stun;comment:击晕" json:"stun"`                   // 击晕
	Stuntime      int32  `gorm:"column:stuntime;comment:眩晕时长" json:"stuntime"`         // 眩晕时长
}

// TableName Skill's table name
func (*Skill) TableName() string {
	return TableNameSkill
}