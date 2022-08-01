<script setup lang="ts">
import ArtworkDisplay from "../components/ArtworkDisplay.vue";
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import axios from "axios";
import {ElMessage} from "element-plus";
import {Search} from "@element-plus/icons-vue";

type ArtworkBase = {
  uuid: string;
  cover: string;
  ti_ming: string;
  chao_dai: string;
  pub_time: string;
}

const route = useRoute();
const router = useRouter();

const Artworks = ref<ArtworkBase[]>([]);
const SelfSearchKey = ref<string>("");

const DoSearch = (key: string) => {
  if (key == "") {
    Artworks.value = [];
    DoSearch("all");
    return
  }
  if (key == "all") {
    axios.post("/api/artwork/search", {"delete_at": {"$eq": null}}).then(resp => {
      let data = resp.data as any[];
      data.forEach(item => {
        Artworks.value?.push({
          uuid: item["uuid"],
          cover: item["cover"],
          ti_ming: item["ti_ming"],
          chao_dai: item["chao_dai"],
          pub_time: item["create_at"],
        })
      })
    }).catch(err => {
      ElMessage.error("获取作品失败！");
      console.log(err)
    })
  } else {
    axios.post("/api/artwork/search", {"ti_ming": {"$regex": key}}).then(resp => {
      let data = resp.data as any[];
      data.forEach(item => {
        Artworks.value?.push({
          uuid: item["uuid"],
          cover: item["cover"],
          ti_ming: item["ti_ming"],
          chao_dai: item["chao_dai"],
          pub_time: item["create_at"],
        })
      })
    }).catch(err => {
      ElMessage.warning("无作品！");
      Artworks.value = [];
    })
  }
}

const OnArtworkClick = (uuid: string) => {
  router.push("/artwork/" + uuid);
}

onMounted(() => {
  DoSearch(route.params["key"] as string);
})
</script>

<template>
  <div class="relative w-screen">
    <div class="absolute w-full h-32 top-0 left-0" style="background-image: url('/img/search-header-0.jpg');"></div>
    <div class="relative container max-w-4xl mx-auto">
      <!--头部-->
      <div class="relative h-32 py-8 px-8">
        <span class="relative text-6xl cursor-default font-msz select-none">作品搜索</span>
        <span class="absolute right-8 top-12">
          <input class="h-10 w-96 px-2 outline-0"
                 type="text"
                 placeholder="请输入搜索关键词"
                 v-model="SelfSearchKey"
          />
          <span class="inline-block align-top ml-4 w-10 h-10
                      bg-white rounded-full hover:bg-[#f2f2f2] active:border
                      select-none cursor-pointer"
                @click="DoSearch(SelfSearchKey)"
          >
            <search class="inline-block w-6 h-6 mx-2 my-2"/>
          </span>
        </span>
      </div>
      <!--结果数和过滤器-->
      <div class="bg-white h-10 my-6 items-center px-2 py-2">
        <span
            class="inline-block h-6 w-36 align-middle text-center cursor-default">
          搜索结果({{ Artworks.length }})
        </span>
      </div>
      <div class="flex flex-wrap">
        <artwork-display
            class="m-4"
            v-for="item in Artworks"
            :cover="'https://oss.www.sxu.ink/cover/'+item.cover +'.jpg'"
            :title="item.ti_ming"
            :age="item.chao_dai"
            :publish_time="item.pub_time.slice(0, 10)"
            @click="OnArtworkClick(item.uuid)"
        />
      </div>
    </div>
  </div>
</template>
