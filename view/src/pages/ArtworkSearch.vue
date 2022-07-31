<script setup lang="ts">
import {Search} from "@element-plus/icons-vue";
import ArtworkDisplay from "../components/ArtworkDisplay.vue";
import {useRouter} from "vue-router";

const router = useRouter();

let artworks = [
  {
    title: "元简墓志",
    cover: "/upload/5c9a643b-041a-42aa-bb5e-9fd5d64f1042.jpg",
    age: "北宋",
    publish_time: "1903-10-09 09:52:41",
    url: "/artwork/xx",
  },
  {
    title: "寇臻墓志并盖",
    cover: "/upload/20e084e7-5bec-4df9-8d31-3b692eda2fc2.jpg",
    age: "南宋",
    publish_time: "2096-08-14 08:29:43",
    url: "/artwork/xx",
  },
  {
    title: "元勰妃李媛华墓志",
    cover: "/upload/35d2eaf0-d0d7-4779-9886-7d3f74d30407.jpg",
    age: "南北朝",
    publish_time: "1903-10-09 09:52:41",
    url: "/artwork/xx",
  },
  {
    title: "李谋墓志",
    cover: "/upload/f2d1c87d-9b95-4037-a038-4a31e5bc4938.jpg",
    age: "东晋",
    publish_time: "2096-08-14 08:29:43",
    url: "/artwork/xx",
  },
]

</script>

<template>
  <div class="relative w-screen">
    <div class="absolute w-full h-32 top-0 left-0" style="background-image: url('/img/search-header-0.jpg');"></div>
    <div class="relative container max-w-4xl mx-auto">
      <!--头部-->
      <div class="relative h-32 py-8 px-8">
        <span class="relative text-6xl cursor-default font-msz select-none">作品搜索</span>
        <span class="absolute right-8 top-12">
          <input class="h-10 w-96 px-2 outline-0" type="text" placeholder="请输入搜索关键词"/>
          <span class="inline-block align-top ml-4 w-10 h-10
                      bg-white rounded-full hover:bg-[#f2f2f2] active:border
                      select-none cursor-pointer">
            <search class="inline-block w-6 h-6 mx-2 my-2"/>
          </span>
        </span>
      </div>
      <!--结果数和过滤器-->
      <div class="bg-white h-10 my-6 items-center px-2 py-2">
        <span
            class="inline-block h-6 w-36 align-middle text-center cursor-default">搜索结果(33)</span>
<!--        <span-->
<!--            class="inline-block border-r-2 border-gray-500 h-6 w-36 align-middle text-center cursor-pointer">按收录时间</span>-->
<!--        <span-->
<!--            class="inline-block border-r-2 border-gray-500 h-6 w-36 align-middle text-center cursor-pointer">按真实年份</span>-->
      </div>
      <div class="flex flex-wrap">
        <artwork-display
            class="m-4"
            v-for="item in artworks"
            :cover="item.cover"
            :title="item.title"
            :age="item.age"
            :publish_time="item.publish_time"
            @click="router.push(item.url)"
        />
      </div>
    </div>
  </div>
</template>
