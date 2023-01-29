<template>
  <div class="app-container">
    <div class="app-left animate__animated animate__fadeIn">
      <h2 class="al-title">项目案例展示</h2>
      <div class="al-tech">
        <div v-for="(item,index) in leftDetails" :key="index" class="al-div">{{ item }}</div>
      </div>
      <el-card class="al-card" shadow="never">
        <el-button type="primary" text class="al-cbtn" @click="aboutShow">关于 | 设置</el-button>
      </el-card>
    </div>
    <div class="app-items animate__animated animate__fadeIn">
      <el-tabs v-model="activeName" class="ai-tabs">
        <el-tab-pane v-for="(item,index) in tabsList" :key="index" :label="item.label" :name="item.name" class="ai-tp">
          <el-scrollbar height="80%" style="margin-top: 10px">
            <el-card v-for="(itm,idx) in pagesData[activeName]" :key="idx" @click="linkTo(idx)" shadow="hover"
                     class="ai-card">
              {{ itm.name }}
            </el-card>
          </el-scrollbar>
        </el-tab-pane>

      </el-tabs>
    </div>

    <el-drawer
        v-model="drawerStatus"
        title="关于 | 设置"
        direction="rtl"
    >
      <el-button @click="toggleDark">开启黑暗模式</el-button>
    </el-drawer>
  </div>
</template>

<script setup>
import appStore from '../store/app-store.js'
import {useRouter} from "vue-router";
import {ref} from "vue";
import 'animate.css';

// 配置黑暗模式
import 'element-plus/theme-chalk/dark/css-vars.css';
import {useDark, useToggle} from "@vueuse/core";
const isDark = useDark()
const toggleDark = () => {
  console.log("黑暗模式已经启用")
  useToggle(isDark)
}


// 实例化路由与store
const store = appStore()
const router = useRouter()

// 数据绑定
const pagesData = ref(store.$state.pagesName)
const leftDetails = ["基于技术", "Vite+Gin+GORM"]

// 配置抽屉
const drawerStatus = ref(false)
const aboutShow = () => {
  drawerStatus.value = true
}

// 右侧tabs数据
const activeName = ref("first")
const tabsList = store.$state.tabsName

// 路由导航
const linkTo = (index) => {
  let p = store.$state.pagesName[activeName.value]
  router.push({
    path: p[index].url
  })
}
</script>

<style lang="less" scoped>
.app-container {
  width: 100vw;
  height: 100vh;
  background-color: ghostwhite;

  display: flex;
  justify-content: center;
  align-items: center;

  .app-left {
    height: 80%;
    width: 20%;

    .al-title {
      background-color: lightskyblue;
      padding: 12px;
      border-radius: 6px;
    }

    .al-tech {
      background-color: lightgray;
      border-radius: 6px;
      padding: 18px;
    }

    .al-card {
      margin-top: 20px;
      height: 62%;
      position: relative;

      .al-cbtn {
        position: absolute;
        left: 20px;
        bottom: 20px;

      }
    }
  }

  .app-right {
    background-color: lightgray;
  }


  .app-items {
    width: 40%;
    height: 70%;
    padding: 20px;
    margin-left: 20px;
    border-radius: 6px;
    display: flex;
    flex-direction: column;
    align-items: center;
    background-color: white;

    .ai-tabs {
      width: 100%;
      height: 100%;

      .ai-tp {
        .ai-card {
          margin-top: 10px;
        }
      }
    }
  }
}
</style>