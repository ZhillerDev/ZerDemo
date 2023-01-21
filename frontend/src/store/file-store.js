import {defineStore} from 'pinia'
import Names from "./names.js";

const fileStore = defineStore(Names.FILESTORE, {
    state: () => {
        return {
            imageTypes: ['image/jpg', 'image/png', 'image/jpeg', 'image/gif']
        }
    },
    actions: {},
    getters: {}
})

export default fileStore