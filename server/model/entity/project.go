package entity

import "gorm.io/gorm"

type NoOrmProjectConfig struct {
	JoinBySelf       bool `json:"join_by_self"`        // 是否可以自己加入到顶级任务,默认真（用户肯定不能自己加入到项目中的，必须所有者邀请）
	MustCheckByOther bool `json:"must_check_by_other"` // 是否必须由其他人检查，默认不需要，
	Secret           bool `json:"secret"`              // 是否是私密项目，默认不需要
	// 甚至我都不想实现这个功能，我想要公司里的所有人都可以看到所有项目和任务(通过分享的方式，默认界面里是点不到的）
	// 如果实现了，私密的项目，不在这个项目组里的用户看不到这个项目和里边的任务
}

func DefaultProjectConfig() NoOrmProjectConfig {
	return NoOrmProjectConfig{
		JoinBySelf:       true,
		MustCheckByOther: false,
		Secret:           false,
	}
}

type Project struct {
	gorm.Model
	Name        string             `json:"name" gorm:"unique;not null"`
	OwnerID     uint               `json:"owner_id" gorm:"not null"`
	Owner       User               `json:"owner" gorm:"foreignKey:OwnerID;references:ID"`
	Description string             `json:"description"`
	GitAddress  string             `json:"git_address"`
	Status      ProjectStatus      `json:"status" gorm:"type:tinyint;default:0;min=0;max=3"`
	Config      NoOrmProjectConfig `json:"config " gorm:"type:text;serializer:json"`
	Member      []Role             `gorm:"foreignKey:ProjectID;references:ID"` // Many-to-Many relationship with Role
}
type ProjectStatus byte

const (
	ProjectStatusInActive ProjectStatus = iota + 1
	ProjectStatusActive
	ProjectStatusCompleted
	ProjectStatusArchived
)
