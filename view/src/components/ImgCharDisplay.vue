<script setup lang="ts">
import {onMounted} from "vue";

type charInfo = {
  char: string,
  pos: {
    x1: number, y1: number, x2: number, y2: number,
  },
  display: boolean,
  selected: boolean,
}

let container: HTMLDivElement;
let canvas: HTMLCanvasElement;
let ctx: CanvasRenderingContext2D;
let image: HTMLImageElement;
let image_drawable = false;
let scale = 1.0;
const maxScale = 5.0, minScale = 0.1;
let offsetX = 0.0, offsetY = 0.0;
let chars: charInfo[];
let isDragging = false;
let displayAll = true;
// let displayAll = false;


const canvas_update = () => {
  canvas.width = container.clientWidth;
  canvas.height = container.clientHeight;
  ctx.clearRect(0, 0, canvas.width, canvas.height);

  if (!image_drawable) return;

  ctx.drawImage(image, offsetX, offsetY, image.naturalWidth * scale, image.naturalHeight * scale);

  ctx.lineWidth = 3;
  if (displayAll) {
    ctx.strokeStyle = "green";
    // display all
    for (let i = 0; i < chars.length; i++) {
      ctx.strokeRect(
          chars[i].pos.x1 * scale + offsetX,
          chars[i].pos.y1 * scale + offsetY,
          (chars[i].pos.x2 - chars[i].pos.x1) * scale,
          (chars[i].pos.y2 - chars[i].pos.y1) * scale,
      );
    }
  } else { // trigger mode {
    for (let i = 0; i < chars.length; i++) {
      if (chars[i].display) {
        ctx.strokeStyle = "green";
        ctx.strokeRect(
            chars[i].pos.x1 * scale + offsetX,
            chars[i].pos.y1 * scale + offsetY,
            (chars[i].pos.x2 - chars[i].pos.x1) * scale,
            (chars[i].pos.y2 - chars[i].pos.y1) * scale,
        );
      }
    }
  }
}

const load_image = (image_url: string, info: charInfo[]) => {
  image_drawable = false;
  image = new Image();
  image.src = image_url;
  image.onload = () => {
    image_drawable = true;
    scale = container.offsetWidth / image.naturalWidth;
    canvas_update();
  }
  chars = info;
}

onMounted(() => {
  // get basis
  container = document.getElementById("canvas-container") as HTMLDivElement;
  canvas = document.getElementById("canvas") as HTMLCanvasElement;
  ctx = canvas.getContext("2d")!
  // set cbs
  window.addEventListener("resize", canvas_update);
  // on wheel
  container.addEventListener("wheel", (ev: WheelEvent) => {
    scale -= ev.deltaY / 1500.0;
    if (scale < minScale || scale > maxScale) {
      scale += ev.deltaY / 1000.0;
      return;
    }
    canvas_update();
  });
  // on mousedown
  canvas.addEventListener("mousedown", (ev: MouseEvent) => {
    if (ev.button == 2) { // drag start
      isDragging = true;
    }
  })
  // on mouseup
  canvas.addEventListener("mouseup", (ev: MouseEvent) => {
    if (ev.button == 2) { // drag end
      isDragging = false;
    }
    canvas_update();
  })
  // on mousemove
  canvas.addEventListener("mousemove", (ev: MouseEvent) => {
    if (isDragging) {
      offsetX += ev.movementX;
      offsetY += ev.movementY;
    }
    for (let i = 0; i < chars.length; i++) {
      chars[i].display = (ev.offsetX - offsetX) / scale >= chars[i].pos.x1 &&
          (ev.offsetX - offsetX) / scale <= chars[i].pos.x2 &&
          (ev.offsetY - offsetY) / scale >= chars[i].pos.y1 &&
          (ev.offsetY - offsetY) / scale <= chars[i].pos.y2;
    }
    canvas_update();
  })
  // on context menu
  canvas.addEventListener("contextmenu", (ev: MouseEvent) => {
    ev.preventDefault();
  })
  //[(285.902, 341.858), (381.202,464.262)]
  load_image("/upload/3505f630-de92-474d-a02f-4f935958f44f.jpg", [
    {
      char: "",
      pos: {
        x1: 1124.103 , y1: 2051.282 , x2: 1238.974 , y2: 2174.359
      },
      display: false,
      selected: false,
    },
  ]);
})
</script>

<template>
  <div id="canvas-container" class="relative w-full h-full">
    <canvas id="canvas" class="bg-gray-500"/>
  </div>
</template>
