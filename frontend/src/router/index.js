import {createRouter, createWebHashHistory} from "vue-router";

const routes = [
    {
        path: '/',
        redirect: '/remote',
        meta: {title: "主页"}
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
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})


export default router
