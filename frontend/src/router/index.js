import {createRouter, createWebHistory} from "vue-router";
import {routerSettings} from "../constants/router-settings.js";
import apiLogin from "../api/api-login.js";


const routes = [
    // 根节点
    {
        path: "/",
        component: () => import("../view/main.vue"),
    },

    // 登录路由
    {
        path: "/login",
        name: "login",
        component: () => import("../view/login.vue"),
    },

    // css小玩意
    {
        path: "/css",
        children: [
            {
                path: "demo",
                component: () => import('../view/csstyle/demo.vue')
            },
            {
                path: "cards",
                component:()=>import('../view/csstyle/css-cards.vue')
            }
        ]
    },

    // 文件上传
    {
        path: "/uploadpic",
        component: import('../view/uploadpic.vue')
    },

    // 数据库查询路由
    {
        path: "/select",
        children: [
            {
                path: "pagequery",
                component: () => import("../view/query/page-query.vue"),
            },
            {
                path: "userdataquery",
                component: () => import("../view/query/userdata-query.vue")
            },
        ]
    },

    // 数据修改
    {
        path: "/modify",
        children: [
            {
                path: "insertuser",
                component: () => import("../view/modify/insert-user.vue")
            }
        ]
    },

    // 中转页面
    {
        path: "/middle",
        children: [
            {
                path: "loginsuccess",
                component: () => import("../components/login/login-success.vue")
            },
            {
                path: "loginfailed",
                component: () => import("../components/login/login-failed.vue")
            }
        ]
    }
];

export const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    console.log("was been guard")
    if (routerSettings.blackList.includes(to.path)) {
        let token = localStorage.getItem("token")
        let res = apiLogin.validatePageToken("/login/check", token)
        setTimeout(() => {
            if (res) {
                next()
            } else {
                console.log("fuck you")
                next("/login")
            }
        }, 500)
    } else {
        next()
    }
})
