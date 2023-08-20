// Package service 自动生成模板 SysMenuService
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package service

import (
	"manager-gin/src/app/admin/sys/sys_menu/model"
	"manager-gin/src/app/admin/sys/sys_menu/service/view"
	"manager-gin/src/app/admin/sys/sys_user/service"
	userView "manager-gin/src/app/admin/sys/sys_user/service/view"
	"manager-gin/src/common"
	"manager-gin/src/utils"
	"strings"
)

var sysMenuDao = model.SysMenuDaoApp
var viewUtils = view.SysMenuViewUtilsApp
var userService = service.SysUserServiceApp

type SysMenuService struct{}

// Create 创建SysMenu记录
// Author
func (sysMenuService *SysMenuService) Create(sysMenuView *view.SysMenuView) (err error) {
	err1, sysMenu := viewUtils.View2Data(sysMenuView)
	if err1 != nil {
		return err1
	}
	err2 := sysMenuDao.Create(*sysMenu)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysMenu记录
// Author
func (sysMenuService *SysMenuService) Delete(id string) (err error) {
	err = sysMenuDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysMenu记录
// Author
func (sysMenuService *SysMenuService) DeleteByIds(ids []string) (err error) {
	err = sysMenuDao.DeleteByIds(ids)
	return err
}

// Update 更新SysMenu记录
// Author
func (sysMenuService *SysMenuService) Update(id string, sysMenuView *view.SysMenuView) (err error) {
	sysMenuView.Id = id
	err1, sysMenu := viewUtils.View2Data(sysMenuView)
	if err1 != nil {
		return err1
	}
	err = sysMenuDao.Update(*sysMenu)
	return err
}

// Get 根据id获取SysMenu记录
// Author
func (sysMenuService *SysMenuService) Get(id string) (err error, sysMenuView *view.SysMenuView) {
	err1, sysMenu := sysMenuDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysMenuView := viewUtils.Data2View(sysMenu)
	if err2 != nil {
		return err2, nil
	}
	return
}

// List 分页获取SysMenu记录
// Author
func (sysMenuService *SysMenuService) List(info *common.PageInfo) (err error) {
	err1, sysMenus, total := sysMenuDao.List(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysMenus)
	if err2 != nil {
		return err2
	}
	info.Rows = viewList
	return err
}

// GetMenuPermission 根据用户id获取菜单权限
func (sysMenuService *SysMenuService) GetMenuPermission(user *userView.SysUserView) (err error, perms []string) {
	is := userService.IsAdmin(user.Id)
	// 管理员拥有所有权限
	if is {
		perms = append(perms, "*:*:*")
	} else {
		if user.Roles != nil {
			for _, role := range *user.Roles {
				err1, rolePerms := sysMenuDao.GetMenuPermissionByRoleId(role.Id)
				if err1 != nil {
					return err1, nil
				}
				role.Permissions = &rolePerms
				perms = append(perms, rolePerms...)
			}
		} else {
			err1, userPerms := sysMenuDao.GetMenuPermissionByUserId(user.Id)
			if err1 != nil {
				return err1, nil
			}
			perms = append(perms, userPerms...)
		}

	}
	return err, perms
}

func (sysMenuService *SysMenuService) GetMenuTreeByUserId(userId string) (err error, menuTree []*view.RouterView) {
	var menus *[]model.SysMenu
	itIs := userService.IsAdmin(userId)
	if itIs {
		err, menus = sysMenuDao.SelectMenuAll()
	} else {
		err, menus = sysMenuDao.SelectMenuByUserId(userId)
	}
	if err != nil {
		return err, nil
	}
	_, viewList := viewUtils.Data2ViewList(menus)

	tree := buildTree(*viewList, "0")
	return err, tree
}

// 递归函数，将SysMenuView转换为MenuNode
func buildTree(menuList []view.SysMenuView, parentId string) []*view.RouterView {
	var tree []*view.RouterView
	for _, menu := range menuList {
		if menu.ParentId == parentId {
			meta := &view.MetaView{
				Title:   menu.MenuName,
				Icon:    menu.Icon,
				NoCache: 1 == menu.IsCache,
			}
			// 外联必须要是http完整路径
			if utils.IsHttp(menu.Path) {
				meta.Link = menu.Path
			}
			node := &view.RouterView{
				Hidden:    "1" == menu.Visible,
				Name:      getRouterName(menu),
				Path:      getRouterPath(menu),
				Component: getComponent(menu),
				Query:     menu.Query,
				Meta:      meta,
			}
			views := buildTree(menuList, menu.Id)
			if views != nil && menu.MenuType == common.MENU_TYPE_DIR {
				node.AlwaysShow = true
				node.Redirect = "noRedirect"
				node.Children = views
			} else if isMenuFrame(menu) {
				node.Meta = nil
				tempMeta := &view.MetaView{
					Title:   menu.MenuName,
					Icon:    menu.Icon,
					NoCache: 1 == menu.IsCache,
				}
				// 外联必须要是http完整路径
				if utils.IsHttp(menu.Path) {
					tempMeta.Link = menu.Path
				}
				var childrenList []*view.RouterView
				children := &view.RouterView{
					Path:      menu.Path,
					Component: menu.Component,
					Name:      strings.Title(menu.Path),
					Query:     menu.Query,
					Meta:      meta,
				}
				childrenList = append(childrenList, children)
				node.Children = childrenList
			} else if menu.ParentId == "0" && isInnerLink(menu) {
				tempMeta := &view.MetaView{
					Title: menu.MenuName,
					Icon:  menu.Icon,
				}
				node.Meta = tempMeta
				node.Path = "/"
				var childrenList []*view.RouterView
				routerPath := innerLinkReplaceEach(menu.Path)

				childMeta := &view.MetaView{
					Title: menu.MenuName,
					Icon:  menu.Icon,
				}

				// 外联必须要是http完整路径
				if utils.IsHttp(menu.Path) {
					childMeta.Link = menu.Path
				}
				children := &view.RouterView{
					Path:      routerPath,
					Component: common.INNER_LINK,
					Name:      strings.Title(routerPath),
					Query:     menu.Query,
					Meta:      childMeta,
				}
				childrenList = append(childrenList, children)
				node.Children = childrenList
			}
			tree = append(tree, node)
		}
	}
	return tree
}

// 获取组件信息
func getComponent(menu view.SysMenuView) string {
	component := common.LAYOUT
	if menu.Component != "" && !isMenuFrame(menu) {
		component = menu.Component
	} else if menu.Component == "" && menu.ParentId != "1" && isInnerLink(menu) {
		component = common.INNER_LINK
	} else if menu.Component == "" && isParentView(menu) {
		component = common.PARENT_VIEW
	}
	return component
}

// 是否为parent_view组件
func isParentView(menu view.SysMenuView) bool {
	return menu.ParentId != "0" && menu.MenuType == common.MENU_TYPE_DIR
}

// 获取路由地址
func getRouterPath(menu view.SysMenuView) string {
	routerPath := menu.Path
	// 内链打开外网方式
	if menu.ParentId != "0" && isInnerLink(menu) {
		routerPath = innerLinkReplaceEach(routerPath)
	}
	// 非外链并且是一级目录（类型为目录）
	if (menu.ParentId == "0" && menu.MenuType == common.MENU_TYPE_DIR) && menu.IsFrame == common.NO_FRAME {
		routerPath = "/" + menu.Path
	} else if isMenuFrame(menu) { // 非外链并且是一级目录（类型为菜单）
		routerPath = "/"
	}
	return routerPath
}

// 内链域名特殊字符替换
func innerLinkReplaceEach(path string) string {
	return utils.ReplaceEach(path, []string{common.HTTP, common.HTTPS, common.WWW, "."}, []string{"", "", "", "/"})
}

// isInnerLink 是否为内链组件
func isInnerLink(menu view.SysMenuView) bool {
	return menu.IsFrame == common.NO_FRAME && utils.IsHttp(menu.Path)
}

// 获取组件名称
func getRouterName(menu view.SysMenuView) string {
	routerName := strings.Title(menu.Path)
	// 非外链并且是一级目录（类型为目录）
	if isMenuFrame(menu) {
		routerName = ""
	}
	return routerName
}

func isMenuFrame(menu view.SysMenuView) bool {
	return menu.ParentId == "0" && common.MENU_TYPE_MENU == menu.MenuType && menu.IsFrame == common.YES_FRAME
}
