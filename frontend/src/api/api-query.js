import http from './http.js'

let userDatas;

async function userDataQuery(url) {
    try {
        let res = await http.get(url)
        userDatas = res.data.data
        console.log(userDatas)
    } catch (err) {
        console.log("这TMD是错的")
    }
    return userDatas
}

export default {
    userDataQuery
}