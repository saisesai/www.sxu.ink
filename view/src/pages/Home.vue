<script setup lang="ts">
import {useRouter} from "vue-router";
import {ref} from "vue";
import {ElMessage} from "element-plus";

const router = useRouter();
const searchText = ref<string>("");

const SearchArtwork = () => {
  if (searchText.value == "") {
    ElMessage.error("请输入搜索关键词！");
    return;
  }
  router.push('/artwork/search/' + searchText.value);
}

const SearchChar = () => {
  if (searchText.value == "") {
    ElMessage.error("请输入搜索字！");
    return;
  } else if (searchText.value.length > 1) {
    ElMessage.error("请输入单个文字！")
    return;
  }
  router.push('/char/' + searchText.value);
}

</script>

<template>
  <div class="absolute w-screen h-screen bg-no-repeat"
       style="background-image: url('https://oss.www.sxu.ink/assets/img/home-bg-0.jpg');background-size: 100% 100%">
    <div class="relative container mx-auto max-w-4xl h-full">
      <!--头部-->
      <div class="py-10 font-msz">
        <span class="text-6xl cursor-pointer">古籍网</span>
        <span class="float-right text-3xl space-x-12 font-bold">
          <span
              class="cursor-pointer hover:text-yellow-800"
              @click="router.push('/artwork/search/all')">作品</span>
          <span class="cursor-pointer hover:text-yellow-800"
                @click="router.push('/char/search')">字库</span>
          <span class="cursor-pointer hover:text-yellow-800"
                @click="router.push('/artwork/edit/new')">录入作品</span>
          <span class="cursor-pointer hover:text-yellow-800"
                @click="router.push('/char/select/xxx')">用户</span>
        </span>
      </div>
      <!--搜索栏-->
      <div class="h-16 w-full mt-16">
        <input type="text"
               placeholder="请输入关键词"
               class="placeholder-black shadow-2xl bg-[#ffffffcc] outline-0 w-3/5 h-full text-2xl indent-6 transition duration-150 ease-in-out focus:-translate-y-1"
               v-model="searchText"
        />
        <input type="button"
               value="搜索作品"
               @click="SearchArtwork"
               class="w-1/5 h-full text-2xl text-white font-msz cursor-pointer active:border-2 transition duration-150 ease-in-out hover:-translate-y-1"
               style="background-image: linear-gradient(to right, #4facfe 0%, #00f2fe 100%);"/>
        <input type="button"
               value="搜索文字"
               @click="SearchChar"
               class="w-1/5 h-full text-2xl text-white font-msz cursor-pointer active:border-2 transition duration-150 ease-in-out hover:-translate-y-1"
               style="background-image: linear-gradient(135deg, #667eea 0%, #764ba2 100%);"/>
      </div>
      <!--底部-->
      <div class="absolute flex justify-evenly items-center h-20 w-full bottom-16 space-x-16 z-10">
        <router-link
            class="group text-center cursor-pointer transition duration-150 ease-in-out hover:-translate-y-1"
            to="/artwork/search/all">
          <img class="inline-block w-12 h-12" src="/img/icon-book.png" alt="artwork">
          <div class="font-bold text-lg text-white">所有作品</div>
        </router-link>
        <router-link
            class="group text-center cursor-pointer transition duration-150 ease-in-out hover:-translate-y-1"
            to="/artwork/search/all">
          <img class="inline-block w-12 h-12" src="/img/icon-dir.png" alt="dir">
          <div class="font-bold text-lg text-white">作品搜索</div>
        </router-link>
        <router-link
            class="group text-center cursor-pointer transition duration-150 ease-in-out hover:-translate-y-1"
            to="/char/search">
          <img class="inline-block w-12 h-12" src="/img/icon-aa.png" alt="aa">
          <div class="font-bold text-lg text-white">字库</div>
        </router-link>
        <router-link
            class="group text-center cursor-pointer transition duration-150 ease-in-out hover:-translate-y-1"
            to="/char/search">
          <img class="inline-block w-12 h-12" src="/img/icon-search.png" alt="search">
          <div class="font-bold text-lg text-white">文字搜索</div>
        </router-link>
        <router-link
            class="group text-center cursor-pointer transition duration-150 ease-in-out hover:-translate-y-1"
            to="/help">
          <img class="inline-block w-12 h-12" src="/img/icon-feedback.png" alt="feedback">
          <div class="font-bold text-lg text-white">帮助/反馈</div>
        </router-link>
      </div>
    </div>
    <!--底部黑色-->
    <div class="absolute h-20 w-full bg-[#0000004f] bottom-16"></div>
  </div>
</template>