<template>
  <svg class="spinner" width="300" height="300" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
    <path class="spinner" d="M2,12A11.2,11.2,0,0,1,13,1.05C12.67,1,12.34,1,12,1a11,11,0,0,0,0,22c.34,0,.67,0,1-.05C6,23,2,17.74,2,12Z"/>
  </svg>
</template>
<script setup>
</script>

<style scoped lang="scss">
@import "../assets/scss/main";
.spinner {
  fill: $primary;
  opacity: 60%;
  path {
    transform-origin: center;
    animation: spinner .5s linear infinite;
  }
}
@keyframes spinner {
  100%{transform:rotate(360deg)}
}
</style>