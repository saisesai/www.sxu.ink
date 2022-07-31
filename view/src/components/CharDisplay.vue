<script setup lang="ts">
const props = defineProps<{
  char: string
}>();

</script>

<template>
  <div class="w-32 h-32 bg-white font-msz text-9xl shadow-xl
              cursor-pointer transition duration-150 ease-in-out hover:-translate-y-1">
    {{ props.char }}
  </div>
</template>