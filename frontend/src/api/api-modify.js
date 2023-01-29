import http from "./http.js";

export default {
    selectUserList: (url) => {
        return http.get(url)
    },
    insertUser: (url, data) => {
        return http.post(url, data)
    },
    deleteUser: (url, data) => {
        return http.post(url, data)
    }
}