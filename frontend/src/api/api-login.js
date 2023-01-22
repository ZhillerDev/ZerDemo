import http from './http.js'
import {fastMessage} from "../constants/message.js";

export default {
    validateLogin: (url, uname, pwd) => {
        http.post(url, {
            username: uname,
            password: pwd
        }).then(res => {
            console.log(res)
            let token = res.data.msg
            localStorage.setItem("token", token)
            fastMessage.success("成功登录！")
        }).catch(err => {
            fastMessage.error("用户名或密码错误！")
            return false
        })
        return true
    },
    validatePageToken: (url, token) => {
        http.post(url, {
            token: token
        }).catch(err => {
            fastMessage.error("token已经过期")
            return false
        })
        return true
    }
}