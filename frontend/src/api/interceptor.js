import axios from "axios";
import qs from "qs";
import { ElMessage } from "element-plus";
import { STATUS } from "./status";

const server = axios.create({
  baseURL: "localhost:10001",
  timeout: 300 * 1000,
});

// 请求拦截器
server.interceptors.request.use(
  (config) => {
    // 如果请求方法为GET，则字符串化params
    if (config.method === "get") {
      config.paramsSerializer = (params) => {
        return qs.stringify(params, {
          arrayFormat: "comma",
        });
      };
    }

    // 获取用户token并附带在headers里传递给后端
    const token = localStorage.token;
    if (token) {
      config.headers["auth"] = token;
    }
    return config;
  },
  (err) => {
    Promise.reject(err);
  }
);

// 响应拦截器
server.interceptors.response.use(
  (config) => {
    return config;
  },
  (err) => {
    // 如果错误的响应体存在的话
    if (err.response) {
      // 获取响应值message
      const msg =
        err.response.data === null
          ? "服务端出错，无法获取响应内容"
          : err.response.data.message;

      // 根据响应状态码来返回对应信息
      switch (err.response.status) {
        case STATUS.success.id:
          ElMessage(STATUS.success.msg);
          break;
        default:
          if (msg === "invalid token") {
            ElMessage("登录状态过期，请重新登陆");
          } else {
            ElMessage(msg);
          }
          break;
      }
    }
    return Promise.reject(err);
  }
);

export default server;
