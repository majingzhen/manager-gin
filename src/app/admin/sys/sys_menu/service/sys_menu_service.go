// Package service 自动生成模板 SysMenuService
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package service

import (
	"errors"
	"manager-gin/src/app/admin/sys/sys_menu/model"
	"manager-gin/src/app/admin/sys/sys_menu/service/view"
	roleSer "manager-gin/src/app/admin/sys/sys_role/service"
	userView "manager-gin/src/app/admin/sys/sys_user/service/view"
	"manager-gin/src/common"
	"manager-gin/src/utils"
	"strings"
)

var sysMenuDao = model.SysMenuDaoApp
var viewUtils = view.SysMenuViewUtilsApp
var roleService = roleSer.SysRoleServiceApp

type SysMenuService struct{}

// Create 创建SysMenu记录
// Author
func (service *SysMenuService) Create(view *view.SysMenuView) error {
	// 判断是否重复
	if err, value := service.CheckMenuNameUniqueAll(view); err != nil {
		return err
	} else {
		if !value {
			return errors.New("菜单名称已存在")
		}
	}
	if view.IsFrame == common.YES_FRAME && !utils.IsHttp(view.Path) {
		return errors.New("外链必须以http://或者https://开头")
	}
	if err1, sysMenu := viewUtils.View2Data(view); err1 != nil {
		return err1
	} else {
		return sysMenuDao.Create(*sysMenu)
	}
}

// Delete 批量删除SysMenu记录
// Author
func (service *SysMenuService) Delete(id string) error {
	// 判断是否存在子菜单
	if err, children := sysMenuDao.SelectMenuByParentId(id); err != nil {
		return err
	} else {
		if len(children) > 0 {
			return errors.New("存在子菜单,不允许删除")
		}
	}
	// 判断菜单是否已分配
	if err, is := sysMenuDao.CheckMenuExistRole(id); err != nil {
		return err
	} else {
		if is {
			return errors.New("菜单已分配,不允许删除")
		}
	}
	return sysMenuDao.Delete(id)
}

// Update 更新SysMenu记录
// Author
func (service *SysMenuService) Update(id string, view *view.SysMenuView) (err error) {
	// 判断是否重复
	if err, value := service.CheckMenuNameUniqueAll(view); err != nil {
		return err
	} else {
		if !value {
			return errors.New("菜单名称已存在")
		}
	}
	if view.IsFrame == common.YES_FRAME && !utils.IsHttp(view.Path) {
		return errors.New("外链必须以http://或者https://开头")
	}
	if view.Id == view.ParentId {
		return errors.New("上级菜单不能选择自己")
	}
	view.Id = id
	err1, sysMenu := viewUtils.View2Data(view)
	if err1 != nil {
		return err1
	}
	err = sysMenuDao.Update(*sysMenu)
	return err
}

// Get 根据id获取SysMenu记录
// Author
func (service *SysMenuService) Get(id string) (err error, view *view.SysMenuView) {
	if id == "" {
		return nil, nil
	}
	err1, sysMenu := sysMenuDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, view := viewUtils.Data2View(sysMenu)
	if err2 != nil {
		return err2, nil
	}
	return
}

// List 分页获取SysMenu记录
// Author
func (service *SysMenuService) List(pageInfo *view.SysMenuPageView) (err error, res *common.PageInfo) {
	err, param, page := viewUtils.Page2Data(pageInfo)
	if err != nil {
		return err, nil
	}
	err1, datas, total := sysMenuDao.List(param, page)
	if err1 != nil {
		return err1, res
	}
	err2, viewList := viewUtils.Data2ViewList(datas)
	if err2 != nil {
		return err2, res
	}
	res = &common.PageInfo{
		Total: total,
		Rows:  viewList,
	}
	return err, res
}

// GetMenuPermission 根据用户id获取菜单权限
func (service *SysMenuService) GetMenuPermission(user *userView.SysUserView) (err error, perms []string) {
	is := user.Id == common.SYSTEM_ADMIN_ID
	// 管理员拥有所有权限
	if is {
		perms = append(perms, "*:*:*")
	} else {
		if user.Roles != nil {
			for _, role := range user.Roles {
				err1, rolePerms := sysMenuDao.GetMenuPermissionByRoleId(role.Id)
				if err1 != nil {
					return err1, nil
				}
				role.Permissions = rolePerms
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

// GetMenuTreeByUserId 根据用户id获取菜单树
func (service *SysMenuService) GetMenuTreeByUserId(userId string) (err error, menuTree []*view.RouterView) {
	var menus []*model.SysMenu
	itIs := userId == common.SYSTEM_ADMIN_ID
	if itIs {
		err, menus = sysMenuDao.SelectMenuAll()
	} else {
		err, menus = sysMenuDao.SelectMenuByUserId(userId)
	}
	if err != nil {
		return err, nil
	}
	_, viewList := viewUtils.Data2ViewList(menus)

	tree := buildTree(viewList, "0")
	return err, tree
}

// SelectMenuList 查询菜单列表
func (service *SysMenuService) SelectMenuList(v *view.SysMenuView, userId string) (err error, menus []*view.SysMenuView) {
	err, data := viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	var dataMenus []*model.SysMenu
	itIs := userId == common.SYSTEM_ADMIN_ID
	if itIs {
		err, dataMenus = sysMenuDao.SelectMenuList(data)
	} else {
		err, dataMenus = sysMenuDao.SelectMenuListByUserId(data, userId)
	}
	err, menus = viewUtils.Data2ViewList(dataMenus)
	return
}

// 递归函数，将SysMenuView转换为MenuNode
func buildTree(menuList []*view.SysMenuView, parentId string) []*view.RouterView {
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
				Name:      getRouterName(*menu),
				Path:      getRouterPath(*menu),
				Component: getComponent(*menu),
				Query:     menu.Query,
				Meta:      meta,
			}
			views := buildTree(menuList, menu.Id)
			if views != nil && menu.MenuType == common.MENU_TYPE_DIR {
				node.AlwaysShow = true
				node.Redirect = "noRedirect"
				node.Children = views
			} else if isMenuFrame(*menu) {
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
			} else if menu.ParentId == "0" && isInnerLink(*menu) {
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

// 是否为外链
func isMenuFrame(menu view.SysMenuView) bool {
	return menu.ParentId == "0" && common.MENU_TYPE_MENU == menu.MenuType && menu.IsFrame == common.YES_FRAME
}

// CheckMenuNameUniqueAll 判断名称是否重复
func (service *SysMenuService) CheckMenuNameUniqueAll(menu *view.SysMenuView) (err error, isUnique bool) {
	err, data := viewUtils.View2Data(menu)
	if err != nil {
		return err, false
	}
	err, isUnique = sysMenuDao.CheckMenuNameUniqueAll(data)
	return
}

// SelectMenuListByRoleId 根据角色id查询菜单
func (service *SysMenuService) SelectMenuListByRoleId(id string) (error, []string) {
	if err, roleView := roleService.Get(id); err != nil {
		return err, nil
	} else {
		if roleView != nil {
			err, menuList := sysMenuDao.SelectMenuListByRoleId(id, roleView.MenuCheckStrictly)
			if err != nil {
				return err, nil
			}
			return nil, menuList
		} else {
			return nil, nil
		}
	}
}

// BuildMenuTreeSelect 构建菜单树
func (service *SysMenuService) BuildMenuTreeSelect(menuViews []*view.SysMenuView) []*view.MenuTree {
	menuMap := make(map[string]*view.MenuTree)

	// 先创建所有的节点
	for _, menuView := range menuViews {
		menuMap[menuView.Id] = &view.MenuTree{
			Id:    menuView.Id,
			Label: menuView.MenuName,
		}
	}

	// 构建树结构
	var rootNodes []*view.MenuTree
	for _, menuView := range menuViews {
		menu := menuMap[menuView.Id]
		if menuView.ParentId == "0" {
			rootNodes = append(rootNodes, menu)
		} else {
			parent := menuMap[menuView.ParentId]
			parent.Children = append(parent.Children, menu)
		}
	}
	return rootNodes
}
