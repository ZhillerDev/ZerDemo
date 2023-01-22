<template>
  <div class="demo-container">
    <div class="d-left">

    </div>
    <div class="d-right">
      <el-scrollbar height="100vh">
        <div class="dr-scroll">
          <el-card class="dr-card" shadow="hover" v-for="(item,index) in compNames" :key="index">
            <component :is="item" class="dr-comp"/>
          </el-card>
        </div>
      </el-scrollbar>
    </div>
  </div>
</template>

<script setup>
import {useRouter} from 'vue-router'
import demoStore from "../store/demo-store.js";
import {components} from "../utils/demoImportsDefine.js";
import {ref} from "vue";

const store = demoStore()
const router = useRouter()
const compNames = ref(components)

console.log(compNames)
</script>

<style lang="less" scoped>
.demo-container {
  width: 100vw;
  height: 100vh;
  display: flex;


  .d-left {
    width: 6%;
    height: 100%;
  }

  .d-right {
    width: 94%;
    height: 100vh;

    .dr-scroll{
      margin: 10px;
      width: 97%;
      display: inline-block;
      column-count: 2;
      .dr-card {
        break-inside: avoid;
        width: 100%;
        margin: 20px;
      }
    }
  }
}
</style>