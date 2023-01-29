<template>
<div class="ccards-container">

    <el-scrollbar>

      <el-tooltip
        effect="dark"
        content="底部弹出蓝条卡片"
        placement="left"
      >
        <div class="cc-card" style="background-color: antiquewhite">
          <div class="cd1" v-for="t in 4">
            <h2>Title</h2>
            <p>
              This is a useless text. You can treat it as a test text or something
            </p>
          </div>
        </div>
      </el-tooltip>

      <el-tooltip
        effect="dark"
        content="shit"
        placement="left"
      >
        <div class="cc-card">

        </div>
      </el-tooltip>
    </el-scrollbar>


</div>
</template>

<script setup>

</script>

<style lang="less" scoped>
.ccards-container{
  width: 70%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-content: center;
  background-color: ghostwhite;


    .cc-card{
      min-width: 30%;
      padding: 20px;
      border-radius: 6px;
      background-color: white;
      margin-top: 16px;

      display: flex;
      flex-direction: row;
      justify-content: space-evenly;
      flex-wrap: wrap;

      // 第一张卡片样式设定
      .cd1{
        padding: 10px;
        width: 210px;
        height: 270px;
        background-color: white;
        border-radius: 4px;

        overflow: hidden;
        position: relative;

        &:before{
          z-index: 10;
          position: absolute;
          background-color: deepskyblue;
          content: '';
          width: 100%;
          height: 20%;
          left: 0;
          bottom: -20%;
          transition: .2s ease;
        }

        &:hover:before{
          transform: translateY(-100%);
        }
      }

      // 两个小圆球融合图像
      .cd2{
        height: 100px;
        width: 300px;
        filter: contrast(30);
        background-color: lightgrey;

        .circle{
          border-radius: 50%;
          position: absolute;
          filter: blur(10px);
        }
        .circle-1{
          height: 90px;
          width: 90px;
          background-color: darksalmon;
          transform: translate(-50px);
          animation: 2s circle-moving infinite alternate-reverse;
        }
        .circle-2{
          height: 50px;
          width: 50px;
          background-color: deepskyblue;
          transform: translate(50px);
          animation: 2s circle-moving infinite alternate;
        }

      }
    }

}

@keyframes circle-moving {
  0%{
    transform: translate(50px);
  }
  100%{
    transform: translate(-50px);
  }
}
</style>