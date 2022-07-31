<script setup lang="ts">
const props = defineProps<{
  cover: string,
  title: string,
  age: string,
  publish_time: string
}>();

</script>

<template>
  <div class="w-48 h-64 bg-white shadow-lg rounded-md cursor-pointer transition duration-150 ease-in-out hover:-translate-y-1">
    <div class="w-full h-48 bg-contain bg-center bg-no-repeat bg-gray-300 rounded-t-md"
         :style="{backgroundImage: 'url(' + props.cover + ')'}">
    </div>
    <div class="px-2">
      <div class="font-bold">{{ props.title }}</div>
      <div class="text-sm">年代: {{ props.age }}</div>
      <div class="text-sm">发布: {{ props.publish_time }}</div>
    </div>
  </div>
</template>