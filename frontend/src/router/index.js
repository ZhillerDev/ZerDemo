import {createRouter, createWebHistory} from "vue-router";
import {routerSettings} from "../constants/router-settings.js";
import {fastMessage} from "../constants/message.js";
import apiLogin from "../api/api-login.js";


const routes = [
    {
        path: "/",
        component: () => import("../view/main.vue"),
    },
    {
        path: "/login",
        name: "login",
        component: () => import("../view/login.vue"),
    },
    {
        path: "/demo",
        component: () => import('../view/demo.vue')
    },
    {
        path: "/uploadpic",
        component: import('../view/uploadpic.vue')
    },
    {
        path: "/pagequery",
        component: () => import("../view/query/page-query.vue")
    },
    {
        path: "/userdataquery",
        component: () => import("../view/query/userdata-query.vue")
    },
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
