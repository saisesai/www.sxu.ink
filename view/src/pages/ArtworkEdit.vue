<script setup lang="ts">
import {onMounted, ref} from "vue";
import {useRoute, useRouter} from "vue-router";
import axios from "axios";
import {ElMessage} from "element-plus";

const route = useRoute();
const router = useRouter();

type ArtworkInfo = {
  ti_ming: string, cover: string, zang_nian: string, chao_dai: string
  chu_tu_di: string, xian_cang_di: string, hang_zi_shu: string,
  chi_cun: string, shuo_ming: string, zu_nian: string,
  nian_ling: string, xing_bie: string, ji_guan: string,
  zhi_gai: string, ming_wen: string,
  pages: string
}

const form = ref<ArtworkInfo>({
  ti_ming: "", cover: "", zang_nian: "", chao_dai: "",
  chu_tu_di: "", xian_cang_di: "", hang_zi_shu: "",
  chi_cun: "", shuo_ming: "", zu_nian: "", nian_ling: "",
  xing_bie: "", ji_guan: "", zhi_gai: "", ming_wen: "",
  pages: "",
})

const OnSave = () => {
  let artworkUuid = route.params.uuid;

  if (form.value.ti_ming == "" || form.value.cover == "" || form.value.chao_dai == "") {
    ElMessage.error("题名，封面uuid，朝代为必填项目！");
    return;
  }

  if (artworkUuid == "new") { // new
    axios.post("/api/artwork/new").then(resp => {
      artworkUuid = resp.data["uuid"];
      axios.post("/api/artwork/set/" + artworkUuid, form.value).then(() => {
        ElMessage.success("作品已保存！");
        router.push("/artwork/edit/" + artworkUuid);
      }).catch(() => {
        ElMessage.error("保存作品失败！");
      })
    }).catch(() => {
      ElMessage.error("创建作品失败！");
    })
    return;
  }
  // exist
  axios.post("/api/artwork/set/" + artworkUuid, form.value).then(() => {
    ElMessage.success("作品已保存！");
  }).catch((err) => {
    ElMessage.error("保存作品失败！");
    console.log(err)
  })
}

const OnDelete = () => {
  let artworkUuid = route.params.uuid;
  let now = new Date();
  axios.post("/api/artwork/set/" + artworkUuid,
      {"delete_at": now.toUTCString()}).then(() => {
    ElMessage.success("作品已删除！");
  }).catch((err) => {
    ElMessage.error("删除作品失败！");
    console.log(err)
  })
}

onMounted(() => {
  let artworkUuid = route.params.uuid;
  if (artworkUuid != "new") {
    axios.post("/api/artwork/get/" + artworkUuid).then(resp => {
      form.value.chao_dai = resp.data["chao_dai"];
      form.value.ti_ming = resp.data["ti_ming"];
      form.value.cover = resp.data["cover"];
      form.value.zang_nian = resp.data["zang_nian"];
      form.value.chu_tu_di = resp.data["chu_tu_di"];
      form.value.xian_cang_di = resp.data["xian_cang_di"];
      form.value.hang_zi_shu = resp.data["hang_zi_shu"];
      form.value.chi_cun = resp.data["chi_cun"];
      form.value.shuo_ming = resp.data["shuo_ming"];
      form.value.zu_nian = resp.data["zu_nian"];
      form.value.nian_ling = resp.data["nian_ling"];
      form.value.xing_bie = resp.data["xing_bie"];
      form.value.ji_guan = resp.data["ji_guan"];
      form.value.zhi_gai = resp.data["zhi_gai"];
      form.value.ming_wen = resp.data["ming_wen"];
      form.value.pages = resp.data["pages"];
    }).catch(() => {
      ElMessage.error("获取作品数据失败！");
    })
  }
})

</script>
<template>
  <div class="relative w-screen">
    <div class="relative container max-w-4xl min-h-screen mx-auto bg-white p-4 space-y-4">
      <el-form :model="form">
        <el-form-item label="题名*">
          <el-input v-model="form.ti_ming"/>
        </el-form-item>
        <el-form-item label="封面uuid*">
          <el-input v-model="form.cover"/>
        </el-form-item>
        <el-form-item label="朝代*">
          <el-input v-model="form.chao_dai"/>
        </el-form-item>
        <el-form-item label="葬年">
          <el-input v-model="form.zang_nian"/>
        </el-form-item>
        <el-form-item label="出土地">
          <el-input v-model="form.chu_tu_di"/>
        </el-form-item>
        <el-form-item label="现藏地">
          <el-input v-model="form.xian_cang_di"/>
        </el-form-item>
        <el-form-item label="行字数">
          <el-input v-model="form.hang_zi_shu"/>
        </el-form-item>
        <el-form-item label="尺寸">
          <el-input v-model="form.chi_cun"/>
        </el-form-item>
        <el-form-item label="说明">
          <el-input v-model="form.shuo_ming"/>
        </el-form-item>
        <el-form-item label="卒年">
          <el-input v-model="form.zu_nian"/>
        </el-form-item>
        <el-form-item label="年龄">
          <el-input v-model="form.nian_ling"/>
        </el-form-item>
        <el-form-item label="性别">
          <el-input v-model="form.xing_bie"/>
        </el-form-item>
        <el-form-item label="籍贯">
          <el-input v-model="form.ji_guan"/>
        </el-form-item>
        <el-form-item label="志盖">
          <el-input v-model="form.zhi_gai"/>
        </el-form-item>
        <el-form-item label="铭文">
          <el-input v-model="form.ming_wen" type="textarea"/>
        </el-form-item>
        <el-form-item label="图片uuid">
          <el-input v-model="form.pages" type="textarea"/>
        </el-form-item>
      </el-form>
      <div class="flex justify-center">
        <el-button type="success" @click="OnSave">保存作品</el-button>
        <el-button type="danger" @click="OnDelete">删除作品</el-button>
      </div>
    </div>
  </div>
</template>