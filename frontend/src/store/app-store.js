import {defineStore} from 'pinia'
import Names from "./names.js";

const appStore = defineStore(Names.APPSTORE, {
    state: () => {
        return {
            btnTitles: ["登录注册模拟", "许多好玩的demo", "单文件上传"],
            btnPaths: ["/login", "/demo", "/uploadpic"]
        }
    },
    actions: {},
    getters: {}
})

export default appStore

