<template>
  <div class="udq-container">
    <el-table
        :data="datas"
        class="u-table"
    >
      <el-table-column prop="id" label="ID" width="100" sortable/>
      <el-table-column prop="username" label="用户名" width="180"/>
      <el-table-column prop="password" label="密码" width="180"/>
      <el-table-column prop="role" label="角色" width="180"/>
    </el-table>
  </div>
</template>

<script setup>

import {onMounted, reactive, ref} from "vue";
import apiQuery from "../../api/api-query.js";
import dbStore from "../../store/db-store.js";

const store = dbStore()
let datas = reactive(store.$state.userLists)

apiQuery.userDataQuery("/sdb/allusers").then(res => {
  store.setUsersList(res)
  console.log(res[0], "fuck")
})

console.log(datas, "shit")


</script>

<style lang="less" scoped>
.udq-container {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;

  .u-table {
    border-radius: 6px;
    padding: 20px;
    width: auto;
    height: auto;
  }
}
</style>