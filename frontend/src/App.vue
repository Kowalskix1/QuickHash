<script setup>
  import { ref, onMounted } from "vue";
  import { CloseWindow, MinimiseWindow } from "../wailsjs/go/app/App"
  import { CalculateHash } from "../wailsjs/go/hash/Hasher"

  var data = ref("");
  var isLoading = ref(false);

  async function closeWindows() {
    await CloseWindow();
  }

  async function minimiseWindows() {
    await MinimiseWindow()
  }

  async function startLoading() {
    isLoading.value = true;
    try {
      var result = await CalculateHash();
      data.value = result;
    } catch (error) {
      data.value = "计算失败"
    } finally {
      isLoading.value = false
    }
  }

  onMounted(() => {
    startLoading();
  });

</script>

<template>
  <div class="container">
    <!-- 模拟标题栏 -->
    <div class="appBar" style="--wails-draggable:drag">
      <div class="appBarTitle">
        QuickHash
      </div>
      <div class="appBarAction">
        <button class="btn close-btn" @click="closeWindows()"></button>
        <button class="btn min-btn" @click="minimiseWindows()"></button>
        <button class="btn max-btn"></button>
      </div>
    </div>
    <!-- 显示区域 -->
    <div v-if="isLoading" class="loading">
      <p>玩命计算中...</p>
      <div class="spinner"></div>
    </div>
    <div v-else class="content">
      <div class="contentTitle">
        <span class="head">M</span>
        <span>essage-</span>
        <span class="head">D</span>
        <span>igest</span>
      </div>
      <div class="contentTitleVersion">5</div>
      <div class="contentText">{{ data }}</div>
    </div>
  </div>

</template>

<style>
  .container {
    width: 100%;
    height: 100vh;
    color: white;
    background-color: rgba(247, 238, 248, 0.145);
  }

  .appBar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px 5px;
    width: 100%;
    height: 30px;
    background-color: rgba(247, 238, 248, 0.145);
    color: white;
  }

  .appBarTitle {
    font-size: 16px;
    /* font-weight: bold; */
    font-family: Consolas;
  }

  .appBarAction {
    display: flex;
    gap: 8px;
  }

  .btn {
    width: 18px;
    height: 18px;
    border-radius: 50%;
    /* margin-right: 6px; */
    position: relative;
    overflow: hidden;
    cursor: pointer;
  }

  .btn:last-child {
    margin-right: 0;
  }

  .btn:before,
  .btn:after {
    content: "";
    position: absolute;
    top: 100%;
    left: 50%;
    transform: translate(-50%, -50%);
    border-radius: 1px;
    opacity: 0;
    transition: all 300ms ease-in-out;
  }

  .close-btn {
    background: #FF5D5B;
    border: 1px solid #CF544D;
  }

  .min-btn {
    background: #FFBB39;
    border: 1px solid #CFA64E;
  }

  .max-btn {
    background: #00CD4E;
    border: 1px solid #0EA642;
  }

  /* Close btn */
  .close-btn:before,
  .close-btn:after {
    width: 1px;
    height: 70%;
    background: #460100;
  }

  .close-btn:before {
    transform: translate(-50%, -50%) rotate(45deg);
  }

  .close-btn:after {
    transform: translate(-50%, -50%) rotate(-45deg);
  }

  /* min btn */
  .min-btn:before {
    width: 70%;
    height: 1px;
    background: #460100;
  }

  /* max btn */
  .max-btn:before {
    width: 50%;
    height: 50%;
    background: #024D0F;
  }

  .max-btn:after {
    width: 1px;
    height: 90%;
    transform: translate(-50%, -50%) rotate(-135deg);
    background: #00CD4E;
  }

  .appBarAction:hover .btn:before,
  .appBarAction:hover .btn:after {
    top: 50%;
    opacity: 1;
  }

  .loading {
    display: flex;
    height: 80vh;
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }

  .content {
    margin: 10px 20px;
  }

  .contentTitle {
    font-family: consolas;
    font-size: 36px;
    color: #ffffff4a;
  }

  .contentTitle .head {
    font-size: 48px;
    font-weight: bold;
  }



  .contentTitleVersion {
    font-family: consolas;
    font-size: 120px;
    color: #ffffff4a;
  }

  .contentText {
    width: 96vw;
    position: absolute;
    top: 145px;
    text-align: center;
    font-size: 20px;
    font-family: consolas;
    /* font-style: italic; */
    /* font-weight: 100; */
  }

  .spinner {
    margin-top: 20px;
    position: relative;
    width: 2em;
    height: 2em;
    border: 3px solid #b2b2b2;
    overflow: hidden;
    animation: spin 3s ease infinite;
  }

  .spinner::before {
    content: '';
    position: absolute;
    top: -3px;
    left: -3px;
    width: 2em;
    height: 2em;
    background-color: hsla(0, 0%, 86%, 0.75);
    transform-origin: center bottom;
    transform: scaleY(1);
    animation: fill 3s linear infinite;
  }

  @keyframes spin {

    50%,
    100% {
      transform: rotate(360deg);
    }
  }

  @keyframes fill {

    25%,
    50% {
      transform: scaleY(0);
    }

    100% {
      transform: scaleY(1);
    }
  }


</style>
