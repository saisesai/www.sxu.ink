<script setup lang="ts">
import {ref} from "vue";
import {useRoute} from "vue-router";

const route = useRoute()
const ch = route.params.ch

// TODO backend
let chars = ref([
  {
    "ori": "元简墓志",
    "url": "/artwork/xxx",
  },
  {
    "ori": "元简墓志",
    "url": "/artwork/xxx",
  },
  {
    "ori": "元简墓志",
    "url": "/artwork/xxx",
  },
  {
    "ori": "元简墓志",
    "url": "/artwork/xxx",
  },
])

</script>

<template>
  <div class="relative w-screen">
    <div class="relative container max-w-4xl min-h-screen mx-auto bg-white p-4 space-y-4">
      <div class="text-2xl">文字：『{{ ch }}』</div>
      <div class="space-y-4">
        <div v-for="(ctx, idx) in chars">
          <span>出处{{ idx + 1 }}:</span>
          <span>
            <a :href="ctx.url">《{{ctx.ori}}》</a>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>
