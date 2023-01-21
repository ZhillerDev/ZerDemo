<template>
  <div class="uploadpic-container">
    <el-form class="upform">
      <el-upload class="uppic"
                 action="http://localhost:10001/uploads/image"
                 show-file-list="false"
                 :on-success="handleAvatarSuccess"
                 :before-upload="beforeAvatarUpload"
      >
        <img v-if="imageUrl" :src="imageUrl" class="avatar"/>
        <el-icon v-else class="avatar-uploader-icon">
          <Plus/>
        </el-icon>
      </el-upload>
    </el-form>
  </div>
</template>

<script setup>
import {ref} from 'vue'
import {ElMessage} from 'element-plus'
import {Plus} from '@element-plus/icons-vue'
import fileStore from '../store/file-store.js'

const imageUrl = ref('')
const store = fileStore()

const handleAvatarSuccess = (
    response,
    uploadFile
) => {
  console.log(response)
  imageUrl.value = response.msg + "?" + Math.random()
}

const beforeAvatarUpload = (rawFile) => {
  const imageTypes = store.$state.imageTypes
  if (!imageTypes.includes(rawFile.type.toString())) {
    ElMessage.error('别搞小动作，只能上传图片')
    return false
  } else if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error('文件大小不可超过2MB')
    return false
  }
  return true
}
</script>

<style lang="less" scoped>
.uploadpic-container {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;

  .upform {
    width: 80%;
    height: 80%;
    border-radius: 8px;
    background-color: white;

    .uppic {
      width: 200px;
      margin: 20px;
      border-bottom: 6px solid lightskyblue;
      cursor: pointer;
      position: relative;
      overflow: hidden;

      transition: .3s ease;

      &:hover {
        background-color: lightskyblue;
        box-shadow: 0 0 20px 0.1px lightgray;
      }

      .avatar {
        width: 200px;
        height: 200px;
        background-color: lightskyblue;
      }

      .avatar-uploader-icon {
        width: 200px;
        height: 200px;

      }
    }
  }
}
</style>