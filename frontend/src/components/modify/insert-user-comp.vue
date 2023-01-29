<template>
  <div class="iuc-container">
    <el-form
        ref="formRef"
        label-position="left"
        label-width="100px"
        :model="formLabelAlign"
        :rules="rules"
    >
      <el-form-item label="用户名" prop="username">
        <el-input v-model="formLabelAlign.username" placeholder="username"/>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="formLabelAlign.password" type="password" placeholder="password"/>
      </el-form-item>
      <el-form-item label="用户类型">
        <el-select v-model="formLabelAlign.role" placeholder="选择类型">
          <el-option label="普通用户" value="1"/>
          <el-option label="管理员" value="2"/>
          <el-option label="被封禁用户" value="3"/>
        </el-select>
      </el-form-item>
    </el-form>

    <div class="btns">
      <el-button type="primary" style="margin-right: 5px" @click="submitForm(formRef)">确认</el-button>
      <el-button type="warning" @click="resetForm(formRef)">重新填写</el-button>
    </div>
  </div>
</template>

<script setup>
import {ref, reactive, onMounted} from "vue";
import emitter from "../../api/emitter.js";
import apiModify from "../../api/api-modify.js";
import {fastMessage} from "../../constants/message.js";

const formRef = ref(null)
onMounted(() => {
})

const rules = reactive({
  username: [{required: true, message: "请输入用户名", trigger: "blur"}],
  password: [{required: true, message: "请输入密码", trigger: "blur"}],
});

const formLabelAlign = reactive({
  username: '',
  password: '',
  role: 1
})

const submitForm = (formEl) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      console.log()
      addUser({
        uname: formLabelAlign.username,
        pwd: formLabelAlign.password,
        role: formLabelAlign.role
      })
    } else {
      console.log('error submit!')
      return false
    }
  })
}

const addUser = async (insertData) => {
  await apiModify.insertUser("/mdb/insertuser", insertData)
      .then(res => {
        fastMessage.success(res.data.msg)
        emitter.emit("insertResult", true)
      })
      .catch(err => {
        fastMessage.error(err.response.data.msg)
      })
}

const resetForm = (formEl) => {
  if (!formEl) return
  formEl.resetFields()
}
</script>

<style lang="less" scoped>
.iuc-container {
  .btns {

  }
}
</style>