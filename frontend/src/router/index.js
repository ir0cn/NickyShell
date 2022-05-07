import {createRouter, createWebHashHistory} from "vue-router";

const routes = [
    {
        path: '/',
        redirect: '/remote'
    },
    {
        path: '/login',
        component: () => import('../view/login')
    },
    {
        path: '/remote',
        component: () => import('../view/remote')
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})


export default router
