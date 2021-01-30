package version_local

import (
	"fiy/app/admin/models"
	"fiy/app/cmdb/models/model"
	"fiy/common/global"
	"runtime"
	"time"

	"gorm.io/gorm"

	"fiy/cmd/migrate/migration"
	common "fiy/common/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1611888450342Migrate)
}

func _1611888450342Migrate(db *gorm.DB, version string) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {

		relatedTypeList := []model.RelatedType{
			{Id: 1, Identifies: "belong", Name: "属于", SourceDescribe: "属于", TargetDescribe: "包含", Direction: 1, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 2, Identifies: "group", Name: "组成", SourceDescribe: "组成", TargetDescribe: "组成于", Direction: 1, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 3, Identifies: "bk_mainline", Name: "拓扑组成", SourceDescribe: "组成", TargetDescribe: "组成于", Direction: 1, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 4, Identifies: "run", Name: "运行", SourceDescribe: "运行于", TargetDescribe: "运行", Direction: 1, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 5, Identifies: "connect", Name: "上联", SourceDescribe: "上联", TargetDescribe: "下联", Direction: 1, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 6, Identifies: "default", Name: "默认关联", SourceDescribe: "关联", TargetDescribe: "关联", Direction: 1, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}

		// 模型分组
		modelGroupList := []model.Group{
			{Id: 1, Identifies: "built_in_host_manager", Name: "主机管理", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 2, Identifies: "built_in_business_topology", Name: "业务拓扑", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 3, Identifies: "built_in_organization", Name: "组织架构", BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}

		// 模型
		modelList := []model.Info{
			{Id: 1, Identifies: "built_in_host", Name: "主机", Icon: "el-icon-notebook-2", IsUsable: true, IsInternal: false, GroupId: 1, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 2, Identifies: "built_in_module", Name: "模块", Icon: "el-icon-menu", IsUsable: true, IsInternal: false, GroupId: 2, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 3, Identifies: "built_in_set", Name: "集群", Icon: "el-icon-s-fold", IsUsable: true, IsInternal: false, GroupId: 2, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 4, Identifies: "built_in_biz", Name: "业务", Icon: "el-icon-s-cooperation", IsUsable: true, IsInternal: false, GroupId: 3, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}

		// 模型字段分组
		modelFieldGroupList := []model.FieldGroup{
			{Id: 1, Name: "基础信息", Sequence: 1, IsFold: false, InfoId: 1, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 2, Name: "其他信息", Sequence: 2, IsFold: false, InfoId: 1, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 3, Name: "基础信息", Sequence: 1, IsFold: false, InfoId: 2, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 4, Name: "基础信息", Sequence: 1, IsFold: false, InfoId: 3, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 5, Name: "基础信息", Sequence: 1, IsFold: false, InfoId: 4, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{Id: 6, Name: "角色", Sequence: 2, IsFold: false, InfoId: 4, BaseModel: common.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}

		err = tx.Create(relatedTypeList).Error
		if err != nil {
			return err
		}

		err = tx.Create(modelGroupList).Error
		if err != nil {
			return err
		}

		err = tx.Create(modelList).Error
		if err != nil {
			return err
		}

		err = tx.Create(modelFieldGroupList).Error
		if err != nil {
			return err
		}

		if err = models.InitDb(tx, "config/sql/cmdb.sql"); err != nil {
			global.Logger.Errorf("同步CMDB初始数据失败, %v", err)
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}