<template>
  <div class="login-container">
    <el-tabs v-model="activeIndex" class="l-tabs" @tabClick="changeIndex" type="card">
      <el-tab-pane label="登录" name="first">
        <el-form
            ref="ruleFormRef"
            :model="ruleForm"
            :rules="rules"
            label-width="80px"
            class="login-form"
            :size="formSize"
        >
          <el-form-item label="账户" prop="username">
            <el-input v-model="ruleForm.username"/>
          </el-form-item>

          <el-form-item label="密码" prop="password">
            <el-input v-model="ruleForm.password" show-password/>
          </el-form-item>
          <div style="text-align: center">
            <el-button type="primary" @click="submitForm(ruleFormRef)"
            >登录
            </el-button>
          </div>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="注册" name="second">

      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup>
import store from '../store/index.js'
import {useRouter} from 'vue-router'
import {ref, reactive, onMounted} from 'vue'
import {ElMessage} from "element-plus";
import {fastMessage} from "../constants/message.js";
import apiLogin from "../api/api-login.js";

const router = useRouter()
const datastore = store()

const activeIndex = ref("first")

const changeIndex = (tab, evt) => {
  console.log(tab, evt)
}

const ruleFormRef = ref(null)

onMounted(() => {

})

const formSize = ref("default");
const ruleForm = reactive({
  username: "",
  password: "",
});
const rules = reactive({
  username: [{required: true, message: "请输入用户名", trigger: "blur"}],
  password: [{required: true, message: "请输入密码", trigger: "blur"}],
});

const submitForm = async (formEl) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      let res = apiLogin.validateLogin(
          "/login/validate",
          ruleForm.username,
          ruleForm.password
      )
      if (res) router.push("/middle/loginsuccess")
      else router.push("/middle/loginfailed")
    } else {
      fastMessage.error("请认真填写用户名与密码[○･｀Д´･ ○]")
    }
  });
}

</script>

<style lang="less" scoped>
.login-container {
  width: 340px;
  height: 270px;
  background-color: white;
  border-radius: 8px;

  .l-tabs {
    margin: 20px;
    width: 300px;
    height: 230px;

    .login-form {
      margin-top: 10px;
    }
  }
}
</style>