package response

import (
	"time"

	"github.com/phpgoc/zxqpro/model/entity"
)

type TaskOneForList struct {
	ID                 int               `json:"id"`
	Name               string            `json:"name"`
	CreateUser         CommonIdAndName   `json:"create_user"`
	ExpectCompleteTime *time.Time        `json:"expect_complete_time"`
	Status             entity.TaskStatus `json:"status"`
	TestUser           *CommonIdAndName  `json:"test_user"`
	SubAssignUser      *CommonIdAndName  `json:"sub_assign_user"`           // 如果是顶级任务，不使用这个字典，非顶级任务，指定的用户是一对一的。顶级任务是一对多的
	TopTaskAssignUsers []CommonIdAndName `json:"top_task_assign_user_list"` // 顶级任务使用这个，即使只有一个人，也要使用这个
	CompletedAt        *time.Time        `json:"completed_at"`
}

type TaskInfo struct {
	ProjectId              uint                      `json:"project_id"`
	ParentId               uint                      `json:"parent_id"` // 0表示顶级任务
	ParentName             string                    `json:"parent_name"`
	Name                   string                    `json:"name"`
	Description            string                    `json:"description"`
	Status                 entity.TaskStatus         `json:"status"`
	CreateUser             CommonIdAndName           `json:"create_user"`
	ExpectCompleteDuration *time.Duration            `json:"expect_complete_duration"`
	ExpectCompleteTime     *time.Time                `json:"expect_complete_time"`
	TestUser               *CommonIdAndName          `json:"test_user"`
	SubAssignUser          *CommonIdAndName          `json:"sub_assign_user"`           // 如果是顶级任务，不使用这个字典，非顶级任务，指定的用户是一对一的。顶级任务是一对多的
	TopTaskAssignUsers     []CommonIdAndName         `json:"top_task_assign_user_list"` // 顶级任务使用这个，即使只有一个人，也要使用这个
	StartedAt              *time.Time                `json:"started_at"`
	CompletedAt            *time.Time                `json:"completed_at"`
	ArchivedAt             *time.Time                `json:"archived_at"`
	TaskTimeEstimateList   []entity.TaskTimeEstimate `json:"task_time_estimates"`
	StepList               []entity.Step             `json:"steps"`
}

type TaskList struct {
	Total int64            `json:"total"`
	List  []TaskOneForList `json:"list"`
}
