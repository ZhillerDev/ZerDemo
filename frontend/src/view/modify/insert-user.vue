<template>
  <div class="iu-container">
    <div class="iu-main">
      <div class="iu-header">
        <el-button type="primary" @click="insertStatus=!insertStatus">添加用户</el-button>
        <el-button plain @click="selectUserList" :icon="Refresh"></el-button>
      </div>
      <div class="iu-body">
        <el-table
            :v-loading="loading"
            :data="datas"
            class="iu-table"
            style="width: 100%;"
            height="270"
        >
          <el-table-column prop="id" label="ID" width="100" sortable/>
          <el-table-column prop="username" label="用户名" width="180"/>
          <el-table-column prop="password" label="密码" width="180"/>
          <el-table-column prop="role" label="角色" width="120"/>
          <el-table-column label="操作" width="100">
            <template #default="scope">
              <el-button type="danger" @click="deleteUser(scope.$index)" :icon="Delete"></el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>

  <el-dialog v-model="insertStatus" style="width: 400px;" title="添加用户" draggable>
    <component :is="insertVue"/>
  </el-dialog>
</template>

<script setup>
import {onMounted, reactive, ref, watch} from 'vue'
import insertVue from "../../components/modify/insert-user-comp.vue"
import emitter from "../../api/emitter.js";

// 导入图标库中的加载图标
import {Refresh, Delete} from '@element-plus/icons-vue'

// 查询store
import dbStore from "../../store/db-store.js";
import apiModify from "../../api/api-modify.js";
import {fastMessage} from "../../constants/message.js";

const store = dbStore()
let datas = reactive(null)

// 获取表单内容
const loading = ref(false)
const selectUserList = async () => {
  loading.value = true
  await apiModify.selectUserList("/sdb/allusers")
      .then(res => {
        store.setUsersList(res.data.data)
        datas = store.$state.userLists
        loading.value = false
      })
      .catch(err => {
        fastMessage.error(err.response.data.msg)
        loading.value = false
      })
}
selectUserList()  // 页面初始化的时候执行一次获取内容

// 删除用户请求
const deleteUser = async (index) => {
  loading.value = true
  await apiModify.deleteUser("/mdb/deleteuser", {
    username: datas[index].username
  })
      .then(res => {
        fastMessage.success(res.data.msg)

        loading.value = false
      })
      .catch(err => {
        fastMessage.error(err.response.data.msg)
        loading.value = false
      })
  await selectUserList()
}

// 添加用户对话框弹出事件检测
const insertStatus = ref(false)
const detectEmitter = () => {
  emitter.on("insertResult", (data) => {
    if (data === true) insertStatus.value = !insertStatus
    selectUserList()
  })
}
detectEmitter()

</script>

<style lang="less" scoped>
.iu-container {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;

  .iu-main {
    width: 70%;
    height: 70%;
    background-color: white;
    border-radius: 8px;

    .iu-header {
      margin: 20px;
    }

    .iu-body {
      padding: 20px;
    }
  }
}
</style>