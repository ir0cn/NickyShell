import {createRouter, createWebHashHistory} from "vue-router";

const routes = [
    {
        path: '/',
        redirect: '/dashboard',
    },
    {
        path: '/',
        component: () => import('../components/Whole.vue'),
        children: [
            {
                path: '/dashboard',
                component: () => import('../view/dashboard/index'),
                redirect: '/Home',
                children: [
                    {
                        path: '/Home',
                        component: () => import('../view/dashboard/index'),
                        meta: {title: '首页'}
                    }
                ]

            },
            {
                path: '/users',
                component: () => import('../view/users'),
                meta: {title: '用户管理'},
                redirect: '/users/users',     // 该配置是若点击选择一级菜单时，默认选中并跳转到该一级菜单下的第一个二级菜单
                children: [
                    {
                        path: 'users',
                        component: () => import('../view/users/users'),
                        meta: {title: '用户管理'}
                    },
                    {
                        path: 'roles',
                        component: () => import('../view/users/roles'),
                        meta: {title: '角色管理'}
                    },
                    {
                        path: 'organization',
                        component:()=>import('../view/users/organization'),
                        meta: {title: '组织管理'}
                    },
                    {
                        path: "permission",
                        component:()=>import('../view/users/permission'),
                        meta: {title:'权限管理'}
                    }
                ]
            },
            {
                path: '/assets',
                component: () => import('../view/assets'),
                meta: {title: '资产管理'},
                redirect: '/assets/hostAssets',     // 该配置是若点击选择一级菜单时，默认选中并跳转到该一级菜单下的第一个二级菜单
                children: [
                    {
                        path: 'hostAssets',
                        component: () => import('../view/assets/hostAssets'),
                        meta: {title: '主机资产'}
                    },
                    {
                        path: 'dbAssets',
                        component: () => import('../view/assets/dbAssets'),
                        meta: {title: '数据库资产'}
                    },
                    {
                        path: 'accounts',
                        component: () => import('../view/assets/accounts'),
                        meta: {title: '账号管理'}
                    }
                ]
            },
            {
                path: '/audit',
                component: () => import('../view/audit'),
                meta: {title: '审计管理'},
                redirect: '/audit/realtime',
                children: [
                    {
                        path: 'realtime',
                        component: () => import('../view/audit/realtime'),
                        meta: {title: '实时会话'}
                    },
                    {
                        path: 'history',
                        component: () => import('../view/audit/history'),
                        meta: {title: '历史会话'}
                    }
                ]
            },
            {
                path: '/setting',
                component: () => import('../view/setting'),
                meta: {title: '系统设置'},
                redirect: '/setting/security',
                children: [
                    {
                        path: 'security',
                        component: () => import('../view/setting/security'),
                        meta: {title: '安全设置'}
                    },
                    {
                        path: 'storage',
                        component: () => import('../view/setting/storage'),
                        meta: {title: '存储设置'}
                    }
                ]
            }
        ]
    },
    {
        path: '/login',
        component: () => import('../view/login'),
        meta: {title: "登录"}
    },
    {
        path: '/remote',
        component: () => import('../view/remote'),
        meta: {title: "远程管理"}
    },
    {
        path: '/:catchAll(.*)',
        component: () => import('../components/404'),
        meta: {title: '404'}
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})


export default router
