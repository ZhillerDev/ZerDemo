<template>
  <div class="app-container">
    <div class="app-items">
      <h1>基于Vite+Vue3+Gin+GORM的项目案例展示</h1>
      <el-scrollbar height="80%">
        <el-button v-for="(item,index) in tabTitles" :key="index" @click="linkTo(index)">
          {{ item }}
        </el-button>
      </el-scrollbar>
    </div>
  </div>
</template>

<script setup>
import appStore from '../store/app-store.js'
import {useRouter} from "vue-router";
import {ref} from "vue";

const store = appStore()
const router = useRouter()
const tabTitles = ref(store.$state.btnTitles)

const linkTo = (index) => {
  router.push({
    path: store.$state.btnPaths[index]
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


  .app-items {
    width: 80%;
    height: 80%;
    padding: 20px;
    border-radius: 6px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    background-color: white;
    box-shadow: 0 0 20px 1px lightgray;
  }
}
</style>