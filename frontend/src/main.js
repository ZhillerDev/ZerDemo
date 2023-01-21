import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";

import { router } from "./router/index.js";
import { createPinia } from "pinia";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";

import * as ElementPlusIconsVue from '@element-plus/icons-vue'

const store = createPinia();

const app = createApp(App);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.use(ElementPlus);
app.use(store);
app.use(router);
app.mount("#app");
