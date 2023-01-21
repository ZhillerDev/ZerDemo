<template>
  <el-affix position="top" class="mainaffix">
    <el-button :icon="House" type="primary" :disabled="currentPath" @click="backToIndex"></el-button>
  </el-affix>
  <div class="mainbody">
    <router-view v-slot="{ Component, route }">
      <transition name='fade' mode="out-in">
        <component :is="Component" :key="route.path"/>
      </transition>
    </router-view>
  </div>
</template>

<script setup>
import {useRouter} from "vue-router";
import {House} from '@element-plus/icons-vue'
import {computed, ref, reactive} from 'vue'
import {ElNotification} from "element-plus";

const router = useRouter()
const currentPath = computed(() => {
  return router.currentRoute.value.path === '/'
})
const backToIndex = () => {
  router.push({
    path: "/"
  })
  ElNotification({
    title: '提示',
    message: "成功返回主页QWQ",
    position: 'bottom-right',
    type: 'success',
    duration: 1500
  })
}
</script>

<style lang="less" scoped>
.mainaffix {
  position: absolute;
  left: 20px;
  top: 20px;
}

.mainbody {
  width: 100vw;
  height: 100vh;
  background-color: ghostwhite;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
