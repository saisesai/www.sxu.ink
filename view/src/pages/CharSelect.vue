<script setup lang="ts">
import ImgCharSelect from "../components/ImgCharSelect.vue";
import {ref} from "vue";
import {ElMessage} from "element-plus";

const ImgPath = "/upload/3505f630-de92-474d-a02f-4f935958f44f.jpg"

const ImgCharSelectRef = ref();

const BtnResetCb = () => {
  ImgCharSelectRef.value.ResetPos();
}

const BtnOriginalSizeCb = () => {
  ImgCharSelectRef.value.OriginalSize();
}

const BtnShowAllCb = () => {
  ImgCharSelectRef.value.ShowAll();
}

const BtnTriggerModeCb = () => {
  ImgCharSelectRef.value.TriggerMode();
}

const PosTableData = ref<{
  char: string,
  pos: {
    x1: number, y1: number, x2: number, y2: number,
  },
  display: boolean,
  selected: boolean,
  saved: boolean,
}[]>([]);

const AddPosCb = (pos: { x1: number, y1: number, x2: number, y2: number }) => {
  // console.log(pos);
  PosTableData.value.push({
    char: "",
    pos: {
      x1: pos.x1, y1: pos.y1, x2: pos.x2, y2: pos.y2,
    },
    display: false,
    selected: false,
    saved: false,
  })
}

const DeleteRowCb = (index: number) => {
  PosTableData.value.splice(index, 1);
  ImgCharSelectRef.value.Update();
}

const ShowRowCb = (index: number) => {
  console.log("SHOW")
  PosTableData.value[index].selected = true;
  ImgCharSelectRef.value.Update();
}

const HideRowCb = (index: number) => {
  console.log("HIDE")
  PosTableData.value[index].selected = false;
  ImgCharSelectRef.value.Update();
}

const SaveRowCb = (index: number) => {
  PosTableData.value[index].saved = true;
  ElMessage.success("save");
}

const RowCharChangeCb = (index: number) => {
  PosTableData.value[index].saved = false;
}

</script>

<template>
  <!--图像-->
  <div class="absolute w-1/2 h-screen border-r-2 border-black">
    <img-char-select
        ref="ImgCharSelectRef"
        :img="ImgPath"
        :list="PosTableData"
        @add_pos="AddPosCb"
    />
  </div>
  <!--操作和信息-->
  <div class="absolute w-1/2 h-screen left-1/2">
    <!--按钮组-->
    <div class="relative w-1/6 h-screen border-r-2 border-black flex flex-col justify-evenly items-center bg-slate-200">
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
      <div>
        <el-button type="success" @click="BtnResetCb">保存全部</el-button>
      </div>
      <div>
        <el-button type="default" @click="BtnResetCb">保存退出</el-button>
      </div>
      <div>
        <el-button type="info" @click="BtnResetCb">使用帮助</el-button>
      </div>
    </div>
    <!--汉字展示组-->
    <div class="absolute w-5/6 h-screen overflow-y-scroll right-0 top-0">
      <el-table :data="PosTableData" style="width: 100%">
        <el-table-column align="center" label="坐标">
          <template #default="scope">
            [({{ Number(scope.row.pos.x1).toFixed(0) }}, {{ Number(scope.row.pos.x2).toFixed(0) }}),
            ({{ Number(scope.row.pos.y1).toFixed(0) }},{{ Number(scope.row.pos.y2).toFixed(0) }})]
          </template>
        </el-table-column>
        <el-table-column align="center" label="汉字">
          <template #default="scope">
            <input class="text-center"
                   type="text"
                   placeholder="未标注"
                   v-model="scope.row.char"
                   @change="RowCharChangeCb(scope.$index)"
            />
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作">
          <template #default="scope">
            <el-button-group>
              <el-button v-if="!scope.row.selected" type="primary" @click="ShowRowCb(scope.$index)">显示</el-button>
              <el-button v-if="scope.row.selected" type="default" @click="HideRowCb(scope.$index)">隐藏</el-button>
              <el-button v-if="!scope.row.saved" type="success" @click="SaveRowCb(scope.$index)">保存</el-button>
              <el-button type="danger" @click="DeleteRowCb(scope.$index)">删除</el-button>
            </el-button-group>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>