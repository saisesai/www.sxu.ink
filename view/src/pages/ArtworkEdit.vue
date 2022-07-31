<script setup lang="ts">
import {ref} from "vue";

const form = ref<{
  ti_ming: string,
  cover: string,
  zang_nian: string,
  chu_tu_di: string,
  xian_cang_di: string,
  hang_zi_shu: string,
  chi_cun: string,
  shuo_ming: string,
  zu_nian: string,
  nian_ling: string,
  xing_bie: string,
  ji_guan: string,
  zhi_gai: string,
  ming_wen: string,
  pages: string[]
}>({
  ti_ming: "",
  cover:"",
  zang_nian: "",
  chu_tu_di: "",
  xian_cang_di: "",
  hang_zi_shu: "",
  chi_cun: "",
  shuo_ming: "",
  zu_nian: "",
  nian_ling: "",
  xing_bie: "",
  ji_guan: "",
  zhi_gai: "",
  ming_wen: "",
  pages: [],
})

const uuidRaw = ref("");
</script>
<template>
  <div class="relative w-screen">
    <div class="relative container max-w-4xl min-h-screen mx-auto bg-white p-4 space-y-4">
      <el-form :model="form">
        <el-form-item label="题名">
          <el-input v-model="form.ti_ming"/>
        </el-form-item>
        <el-form-item label="封面uuid">
          <el-input v-model="form.cover"/>
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
          <el-input v-model="uuidRaw" type="textarea"/>
        </el-form-item>
      </el-form>
      <div class="flex justify-center">
        <el-button type="success">保存作品</el-button>
        <el-button type="danger">删除作品</el-button>
      </div>
    </div>
  </div>
</template>