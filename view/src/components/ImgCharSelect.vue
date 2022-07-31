<script setup lang="ts">
import {onMounted} from "vue";
import {ElMessage} from "element-plus";

const props = defineProps<{
  img: string,
  list: {
    char: string,
    pos: {
      x1: number, y1: number, x2: number, y2: number,
    },
    display: boolean,
    selected: boolean,
    saved: boolean,
  }[],
}>();

const emit = defineEmits<{
  (e: "add_pos", pos: any): void
}>();

let container: HTMLDivElement;
let canvas: HTMLCanvasElement;
let ctx: CanvasRenderingContext2D;
const workImg = new Image();
workImg.src = props.img!;
let scale = 1.0;
const maxScale = 5.0, minScale = 0.1;
let offsetX = 0.0, offsetY = 0.0;
let isDragging = false;
let isSelecting = false;
let firstPosX = 0, firstPosY = 0;
let secondPosX = 0, secondPosY = 0;
let displayAll = false;

const canvas_update = () => {
  canvas.width = container.clientWidth;
  canvas.height = container.clientHeight;

  ctx.drawImage(workImg, offsetX, offsetY, workImg.naturalWidth * scale, workImg.naturalHeight * scale)

  // draw selecting
  ctx.lineWidth = 3;
  ctx.strokeStyle = "aqua";
  ctx.strokeRect(firstPosX, firstPosY, secondPosX - firstPosX, secondPosY - firstPosY);
  // draw selected
  if (displayAll) {
    ctx.strokeStyle = "green";
    // display all
    for (let i = 0; i < props.list.length; i++) {
      ctx.strokeRect(
          props.list[i].pos.x1 * scale + offsetX,
          props.list[i].pos.y1 * scale + offsetY,
          (props.list[i].pos.x2 - props.list[i].pos.x1) * scale,
          (props.list[i].pos.y2 - props.list[i].pos.y1) * scale,
      );
    }
  } else { // trigger mode {
    for (let i = 0; i < props.list.length; i++) {
      if (props.list[i].display) {
        ctx.strokeStyle = "green";
        ctx.strokeRect(
            props.list[i].pos.x1 * scale + offsetX,
            props.list[i].pos.y1 * scale + offsetY,
            (props.list[i].pos.x2 - props.list[i].pos.x1) * scale,
            (props.list[i].pos.y2 - props.list[i].pos.y1) * scale,
        );
      }
    }
  }
  for (let i = 0; i < props.list.length; i++) {
    if (props.list[i].selected) {
      ctx.strokeStyle = "blue";
      ctx.strokeRect(
          props.list[i].pos.x1 * scale + offsetX,
          props.list[i].pos.y1 * scale + offsetY,
          (props.list[i].pos.x2 - props.list[i].pos.x1) * scale,
          (props.list[i].pos.y2 - props.list[i].pos.y1) * scale,
      );
    }
  }
}

onMounted(() => {
  // get all vars
  container = document.getElementById("select-container") as HTMLDivElement;
  canvas = document.getElementById("canvas") as HTMLCanvasElement;
  ctx = canvas.getContext("2d")!;
  // setup cbs
  workImg.onload = () => {
    scale = container.offsetWidth / workImg.naturalWidth;
    canvas_update();
  }
  window.addEventListener("resize", canvas_update);
  container.addEventListener("wheel", (ev: WheelEvent) => {
    //ev.preventDefault();
    scale -= ev.deltaY / 1500.0;
    if (scale < minScale || scale > maxScale) {
      scale += ev.deltaY / 1000.0;
      return
    }
    canvas_update();
  });
  canvas.addEventListener("mousedown", (ev: MouseEvent) => {
    if (ev.button == 2) { // drag start
      isDragging = true;
    }
    if (ev.button == 0) { // selecting start
      isSelecting = true;
      firstPosX = secondPosX = ev.offsetX;
      firstPosY = secondPosY = ev.offsetY;
      canvas_update();
    }
    // console.log(ev);
  })
  canvas.addEventListener("mouseup", (ev: MouseEvent) => {
    if (ev.button == 2) { // drag end
      isDragging = false;
    }
    if (ev.button == 0) { // selecting end
      isSelecting = false;
      if (firstPosX > secondPosX || firstPosY > secondPosY) {
        ElMessage.error("请延坐上至右下方向框选！")
        return
      }
      if (firstPosX != secondPosX && firstPosY != secondPosY) {
        emit("add_pos", {
          x1: ((firstPosX - offsetX) / scale).toFixed(3),
          y1: ((firstPosY - offsetY) / scale).toFixed(3),
          x2: ((secondPosX - offsetX) / scale).toFixed(3),
          y2: ((secondPosY - offsetY) / scale).toFixed(3),
        });
        console.log(
            "x1:", ((firstPosX - offsetX) / scale).toFixed(3), ",",
            "y1:", ((firstPosY - offsetY) / scale).toFixed(3), ",",
            "x2:", ((secondPosX - offsetX) / scale).toFixed(3), ",",
            "y2:", ((secondPosY - offsetY) / scale).toFixed(3),
        )
      }
    }
    // console.log(ev);
    canvas_update();
    firstPosX = secondPosX = 0;
    firstPosY = secondPosY = 0;
  })
  canvas.addEventListener("mousemove", (ev: MouseEvent) => {
    if (isDragging) {
      offsetX += ev.movementX;
      offsetY += ev.movementY;
    }
    if (isSelecting) {
      secondPosX = ev.offsetX;
      secondPosY = ev.offsetY;
    }
    for (let i = 0; i < props.list.length; i++) {
      props.list[i].display = (ev.offsetX - offsetX) / scale >= props.list[i].pos.x1 &&
          (ev.offsetX - offsetX) / scale <= props.list[i].pos.x2 &&
          (ev.offsetY - offsetY) / scale >= props.list[i].pos.y1 &&
          (ev.offsetY - offsetY) / scale <= props.list[i].pos.y2;
    }
    canvas_update();
  })
  canvas.addEventListener("contextmenu", (ev: MouseEvent) => {
    ev.preventDefault();
  })
})

const ResetPos = () => {
  scale = container.offsetWidth / workImg.naturalWidth;
  offsetX = offsetY = 0.0;
  canvas_update();
}

const OriginalSize = () => {
  scale = 1.0;
  offsetX = offsetY = 0.0;
  canvas_update();
}

const ShowAll = () => {
  displayAll = true;
  canvas_update();
}

const TriggerMode = () => {
  displayAll = false;
  canvas_update();
}

const Update = () => {
  canvas_update();
}

defineExpose({ResetPos, OriginalSize, ShowAll, TriggerMode, Update});
</script>

<template>
  <div id="select-container"
       class="w-full h-full bg-gray-500">
    <canvas id="canvas"></canvas>
  </div>
</template>
