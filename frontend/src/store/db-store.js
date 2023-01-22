import {defineStore} from 'pinia'
import Names from "./names.js";

const dbStore = defineStore(Names.DBSTORE, {
    state: () => {
        return {
            userLists: []
        }
    },
    actions: {
        setUsersList(data) {
            this.userLists = []
            for (let i in data) {
                console.log(data[i])
                this.userLists.push(data[i])
            }
        }
    },
    getters: {}
})

export default dbStore