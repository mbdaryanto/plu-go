<script setup lang="ts">
const props = defineProps({
  name: String
})
</script>

<template>
  <div class="col-6">
    <div class="text-sm mt-1">{{props.name}}</div>
    <div class="text-lg">
      <slot>-</slot>
    </div>
  </div>
</template>
