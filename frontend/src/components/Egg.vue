<template>
  <transition name="egg">
  <div class="d-flex justify-content-center align-items-center w-100 h-100 top-0 start-0 z-0 position-absolute overflow-hidden" v-if="eggVisible">
    <div class="egg-container position-relative">
      <img class="position-absolute start-0 top-0 rounded-circle top-50 start-50 translate-middle" src="/rick.gif" alt="rickroll">
      <svg viewBox="116 121 250 250" width="600" height="600">
        <defs>
          <path id="circle" d="M 100, 100 m -75, 0 a 75, 75 0 1, 0 150, 0 a 75, 75 0 1, 0 -150, 0 "/>
        </defs>
        <text class="text text-uppercase" transform="matrix(1.6, 0, 0, 1.6, 130, 130)">
          <textPath xlink:href="#circle">
            Never Gonna Give You Up Never Gonna Let You Down
          </textPath>
        </text>
      </svg>
    </div>
  </div>
  </transition>
</template>

<script lang="ts" setup>
defineProps({
  eggVisible: {
    type: Boolean,
    default: false
  }
})
</script>

<style lang="scss" >
.egg-container {
  animation: 5s linear 0s transfer forwards;
  img {
    height: 300px;
    width: 300px;
  }
  svg {
    animation: 5s linear 0s infinite rotation;
    .text {
    letter-spacing: 1px;
    font-size: 15px;
    transform-origin: 86px 75px;
    fill: white;
    }
  }
}
@keyframes transfer {
  0% { transform: translateX(-60vw) }
  50% { transform: translateX(0) }
  100% { transform: translateX(60vw) }
}
@keyframes rotation {
  to { transform: rotate(360deg); }
}
.egg-enter-active, .egg-leave-active { transition: all 2s ease; }
.egg-enter-from, .egg-leave-to { opacity: 0 }

</style>
