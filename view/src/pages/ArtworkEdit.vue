<script setup lang="ts">
import {onMounted, ref} from "vue";
import {useRoute, useRouter} from "vue-router";
import {ElMessage} from "element-plus";
import {ArtworkInfo, GetArtworkInfo, NewArtwork, SetArtwork, SoftDeleteArtwork} from "../api/ArtworkInfo";

const route = useRoute();
const router = useRouter();

const form = ref<ArtworkInfo>({
  ti_ming: "", cover: "", zang_nian: "", chao_dai: "",
  chu_tu_di: "", xian_cang_di: "", hang_zi_shu: "",
  chi_cun: "", shuo_ming: "", zu_nian: "", nian_ling: "",
  xing_bie: "", ji_guan: "", zhi_gai: "", ming_wen: "",
  pages: "",
})

const OnSave = () => {
  // check data format
  if (form.value.ti_ming == "" || form.value.cover == "" || form.value.chao_dai == "") {
    ElMessage.error("题名，封面uuid，朝代为必填项目！");
    return;
  }
  // new
  if (route.params["uuid"] as string == "new") {
    NewArtwork(form.value).then((uuid) => {
      router.push("/artwork/edit/" + uuid);
    });
    return;
  }
  // exist
  SetArtwork(route.params["uuid"] as string, form.value);
}

const OnDelete = () => SoftDeleteArtwork(route.params["uuid"] as string);

const OnBack = () => router.push("/artwork/" +  route.params["uuid"] as string);

onMounted(() => {
  let artworkUuid = route.params.uuid;
  if (artworkUuid != "new") {
    GetArtworkInfo(artworkUuid as string, form.value);
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
        <el-button type="warning" @click="OnBack">返回作品展示</el-button>
        <el-button type="danger" @click="OnDelete">删除作品</el-button>
      </div>
    </div>
  </div>
</template>