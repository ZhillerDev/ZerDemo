import {defineStore} from 'pinia'
import Names from "./names.js";

const appStore = defineStore(Names.APPSTORE, {
    state: () => {
        return {
            btnTitles: ["登录注册模拟", "卡片瀑布流", "单文件上传", "用户数据查询", "多页查询"],
            btnPaths: ["/login", "/demo", "/uploadpic", "/userdataquery", "/pagequery"]
        }
    },
    actions: {},
    getters: {}
})

export default appStore

