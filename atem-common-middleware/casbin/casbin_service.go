package casbin

import (
	"errors"
	"sync"

	"github.com/socloudng/atem-common/atem-common-base/base_service"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type CasbinService struct {
	base_service.BaseService
	option *CasbinConfig
}

func (casbinService *CasbinService) SetConfig(config *CasbinConfig) {
	casbinService.option = config
}

func (casbinService *CasbinService) UpdateCasbin(authorityId string, casbinInfos []CasbinInfo) error {
	casbinService.ClearCasbin(0, authorityId)
	rules := [][]string{}
	for _, v := range casbinInfos {
		cm := CasbinModel{
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := casbinService.Orm.Table("casbin_rule").Model(&CasbinModel{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

func (casbinService *CasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []CasbinInfo) {
	e := casbinService.Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success

}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	if casbinService.option == nil {
		casbinService.Logger.Fatal("please set casbin-config first")
	}
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(casbinService.Orm)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(casbinService.option.ModelPath, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
