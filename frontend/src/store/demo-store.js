import {defineStore} from 'pinia'
import Names from "./names.js";

const demoStore = defineStore(Names.DEMOSTORE, {
    state: () => {
        return {

        }
    },
    actions: {},
    getters: {}
})

export default demoStore