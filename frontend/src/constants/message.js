import {ElMessage} from "element-plus";

export const fastMessage = {
    success: (msg) => {
        ElMessage({
            message: msg,
            type: "success",
            duration: 1500
        })
    },
    error: (msg) => {
        ElMessage({
            message: msg,
            type: "error",
            duration: 1500
        })
    }
}