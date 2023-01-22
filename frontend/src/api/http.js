import qs from "qs";
import intcp from "./interceptor";

const http = {
  // axios执行get请求
  get(url, params) {
    const config = {
      method: "get",
      url: url,
    };
    // 当存在params时，才会将其添加到config并发送给后端
    if (params) config.params = params;
    return intcp(config);
  },

  // axios执行post请求
  post(url, formdata) {
    const config = {
      method: "post",
      url: url,
      headers:{
        'Content-type': 'application/x-www-form-urlencoded'
      },
    };
    if (formdata) config.data = formdata;
    console.log(config)
    return intcp(config);
  },
};

export default http;
