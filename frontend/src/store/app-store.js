import {defineStore} from 'pinia'
import Names from "./names.js";

const appStore = defineStore(Names.APPSTORE, {
    state: () => {
        return {
            btnTitles: ["登录注册模拟", "卡片瀑布流", "单文件上传", "用户数据查询", "多页查询"],
            btnPaths: ["/login", "/demo", "/uploadpic", "/select/userdataquery", "/select/pagequery"],

            tabsName: [
                {name: "first", label: "数据库CRUD"},
                {name: "second", label: "登录注册"},
                {name: "third", label: "数据报表"},
                {name: "fourth", label: "小玩意"}
            ],

            pagesName: {
                "first": [
                    {name: "用户数据查询", url: "/select/userdataquery"},
                    {name: "分页查询", url: "/select/pagequery"},
                    {name: "增删用户", url: "/modify/insertuser"}
                ],
                "second": [
                    {name: "登陆注册简单模拟", url: "/login"}
                ],
                "third": [
                    {name: "单文件上传", url: "/uploadpic"}
                ],
                "fourth": [
                    {name: "卡片瀑布流", url: "/css/demo"},
                    {name:"卡片动效",url:'/css/cards'}
                ]

            }
        }
    },
    actions: {
        getPagesName(name) {

            return this.pagesName.name
        }
    }
    ,
    getters: {}
})

export default appStore

