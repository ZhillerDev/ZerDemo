import {createRouter, createWebHistory} from "vue-router";

const routes = [
    {
        path: "/",
        component: () => import("../view/main.vue"),
    },
    {
        path: "/login",
        component: () => import("../view/login.vue")
    },
    {
        path: "/demo",
        component: () => import('../view/demo.vue')
    },
    {
        path: "/uploadpic",
        component: import('../view/uploadpic.vue')
    }
];

export const router = createRouter({
    history: createWebHistory(),
    routes,
});
