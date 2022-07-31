<script setup lang="ts">
import {ref} from "vue";
import ImgCharDisplay from "../components/ImgCharDisplay.vue";

let artworkInfo = ref({
  ti_ming: "元简墓志",
  zang_nian: "太和23年（499）3月8日",
  chu_tu_di: "河南洛阳",
  xian_cang_di: "西安碑林",
  hang_zi_shu: "有方界格。残存8行，满行18字，共120字",
  chi_cun: "残高73，宽35",
  shuo_ming: "后半残缺",
  zu_nian: "太和23年（499）正月26",
  nian_ling: "40",
  xing_bie: "男",
  ji_guan: "河南洛阳",
  zhi_gai: "後半殘闕",
  ming_wen: "（殘後半部）太保齊郡王姓元，諱簡，字叔亮，司州河南郡洛陽縣都鄉洛陽里人，高宗之叔子，皇帝之第五叔也。惟王稟旻融度，資造流仁，澄神守質，志性寬雅，冥慶舛和，端宿墜日。以太和廿三年崴在己卯正月戊寅朔廿六日癸卯，春秋卌，寢疾，薨于第，謚曰順王。其年三月甲午卽窆于河南洛陽之北芒。迺鏤石囗銘，式述徽蹤。（下殘缺）",
})

</script>

<template>
  <div class="relative w-screen">
<!--    显示图片-->
    <div class="absolute w-1/2 h-screen">
      <img-char-display/>
    </div>
<!--    功能按钮-->
    <div class="absolute left-1/2 w-1/12 h-screen bg-slate-200"></div>
<!--    信息展示-->
    <div class="absolute right-0 w-5/12 h-screen p-4 overflow-y-auto space-y-2">
      <div v-if="(artworkInfo.ti_ming !== '')">
        <div class="text-2xl">题名：《{{ artworkInfo.ti_ming }}》</div>
      </div>
      <div v-if="(artworkInfo.zang_nian !== '')">
        <div class="text-xl">葬年：{{ artworkInfo.zang_nian }}</div>
      </div>
      <div v-if="(artworkInfo.chu_tu_di !== '')">
        <div class="text-xl">出土地：{{ artworkInfo.chu_tu_di }}</div>
      </div>
      <div v-if="(artworkInfo.xian_cang_di !== '')">
        <div class="text-xl">现藏地：{{ artworkInfo.xian_cang_di }}</div>
      </div>
      <div v-if="(artworkInfo.hang_zi_shu !== '')">
        <div class="text-xl">行字数：{{ artworkInfo.hang_zi_shu }}</div>
      </div>
      <div v-if="(artworkInfo.chi_cun !== '')">
        <div class="text-xl">尺寸：{{ artworkInfo.chi_cun }}</div>
      </div>
      <div v-if="(artworkInfo.shuo_ming !== '')">
        <div class="text-xl">说明：{{ artworkInfo.shuo_ming }}</div>
      </div>
      <div v-if="(artworkInfo.zu_nian !== '')">
        <div class="text-xl">卒年：{{ artworkInfo.zu_nian }}</div>
      </div>
      <div v-if="(artworkInfo.nian_ling !== '')">
        <div class="text-xl">年龄：{{ artworkInfo.nian_ling }}</div>
      </div>
      <div v-if="(artworkInfo.xing_bie !== '')">
        <div class="text-xl">性别：{{ artworkInfo.xing_bie }}</div>
      </div>
      <div v-if="(artworkInfo.ji_guan !== '')">
        <div class="text-xl">籍贯：{{ artworkInfo.ji_guan }}</div>
      </div>
      <div v-if="(artworkInfo.zhi_gai !== '')">
        <div class="text-xl">志盖：{{ artworkInfo.zhi_gai }}</div>
      </div>
      <div v-if="(artworkInfo.ming_wen !== '')">
        <div class="text-xl">铭文：</div>
        <p>{{ artworkInfo.ming_wen }}</p>
      </div>
    </div>
  </div>
</template>
