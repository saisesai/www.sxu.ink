<script setup lang="ts">
import {onMounted, ref} from "vue";
import ImgCharDisplay from "../components/ImgCharDisplay.vue";
import {useRoute, useRouter} from "vue-router";
import {ArtworkInfo, GetArtworkInfo, ParsePages} from "../api/ArtworkInfo";

const route = useRoute();
const router = useRouter();

const artwork_uuid = route.params["uuid"] as string;

let pages = ref<string[]>([]);
let page_index = ref<number>(0);

const form = ref<ArtworkInfo>({
  ti_ming: "", cover: "", zang_nian: "", chao_dai: "",
  chu_tu_di: "", xian_cang_di: "", hang_zi_shu: "",
  chi_cun: "", shuo_ming: "", zu_nian: "", nian_ling: "",
  xing_bie: "", ji_guan: "", zhi_gai: "", ming_wen: "",
  pages: "",
})

const ImgCharDisplayRef = ref();

const LoadImage = (img_uuid: string) => ImgCharDisplayRef.value.LoadImage(img_uuid);

const BtnResetCb = () => ImgCharDisplayRef.value.ResetPos();

const BtnOriginalSizeCb = () => ImgCharDisplayRef.value.OriginalSize();

const BtnShowAllCb = () => ImgCharDisplayRef.value.ShowAll();

const BtnTriggerModeCb = () => ImgCharDisplayRef.value.TriggerMode();

const BtnEditArtwork = () => router.push("/artwork/edit/" + artwork_uuid);

const BtnEditPage = () => router.push("/char/select/" + pages.value[page_index.value])

const BtnNextPage = () => LoadImage(pages.value[++page_index.value]);

const BtnPrevPage = () => LoadImage(pages.value[--page_index.value]);

onMounted(() => {
  GetArtworkInfo(artwork_uuid, form.value).then(() => {
    pages.value = ParsePages(form.value.pages);
    if (pages.value.length > 0) {
      page_index.value = 0;
      LoadImage(pages.value[0]);
    }
  });
})
</script>

<template>
  <div class="relative w-screen">
    <!--    显示图片-->
    <div class="absolute w-1/2 h-screen">
      <img-char-display ref="ImgCharDisplayRef"/>
    </div>
    <!--    功能按钮-->
    <div
        class="absolute left-1/2 w-1/12 h-screen
        bg-slate-200 border-x-2 border-black
        flex flex-col justify-evenly items-center">
      <div>
        <el-button type="primary" @click="BtnResetCb">图像复位</el-button>
      </div>
      <div>
        <el-button type="primary" @click="BtnOriginalSizeCb">真实大小</el-button>
      </div>
      <div>
        <el-button type="primary" @click="BtnShowAllCb">显示所有</el-button>
      </div>
      <div>
        <el-button type="primary" @click="BtnTriggerModeCb">触发显示</el-button>
      </div>
      <div v-if="page_index !== 0">
        <el-button type="success" @click="BtnPrevPage">上一页</el-button>
      </div>
      <div v-if="page_index !== (pages.length - 1)">
        <el-button type="success" @click="BtnNextPage">下一页</el-button>
      </div>
      <div>
        <el-button type="warning" @click="BtnEditPage">编辑图片</el-button>
      </div>
      <div>
        <el-button type="warning" @click="BtnEditArtwork">编辑作品</el-button>
      </div>
    </div>
    <!--    信息展示-->
    <div class="absolute right-0 w-5/12 h-screen p-4 overflow-y-auto space-y-2">
      <div v-if="(form.ti_ming !== '')">
        <div class="text-2xl">题名：《{{ form.ti_ming }}》</div>
      </div>
      <div v-if="(form.zang_nian !== '')">
        <div class="text-xl">葬年：{{ form.zang_nian }}</div>
      </div>
      <div v-if="(form.chu_tu_di !== '')">
        <div class="text-xl">出土地：{{ form.chu_tu_di }}</div>
      </div>
      <div v-if="(form.xian_cang_di !== '')">
        <div class="text-xl">现藏地：{{ form.xian_cang_di }}</div>
      </div>
      <div v-if="(form.hang_zi_shu !== '')">
        <div class="text-xl">行字数：{{ form.hang_zi_shu }}</div>
      </div>
      <div v-if="(form.chi_cun !== '')">
        <div class="text-xl">尺寸：{{ form.chi_cun }}</div>
      </div>
      <div v-if="(form.shuo_ming !== '')">
        <div class="text-xl">说明：{{ form.shuo_ming }}</div>
      </div>
      <div v-if="(form.zu_nian !== '')">
        <div class="text-xl">卒年：{{ form.zu_nian }}</div>
      </div>
      <div v-if="(form.nian_ling !== '')">
        <div class="text-xl">年龄：{{ form.nian_ling }}</div>
      </div>
      <div v-if="(form.xing_bie !== '')">
        <div class="text-xl">性别：{{ form.xing_bie }}</div>
      </div>
      <div v-if="(form.ji_guan !== '')">
        <div class="text-xl">籍贯：{{ form.ji_guan }}</div>
      </div>
      <div v-if="(form.zhi_gai !== '')">
        <div class="text-xl">志盖：{{ form.zhi_gai }}</div>
      </div>
      <div v-if="(form.ming_wen !== '')">
        <div class="text-xl">铭文：</div>
        <p>{{ form.ming_wen }}</p>
      </div>
    </div>
  </div>
</template>
