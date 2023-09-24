<template>
  <transition name="egg">
  <div class="d-flex justify-content-center" v-if="eggVisible">
    <div class="egg-container">
      <img src="/rick.gif" alt="rolling">
      <svg viewBox="116 121 250 250" width="600" height="600">
        <defs>
          <path id="circle" d="M 100, 100 m -75, 0 a 75, 75 0 1, 0 150, 0 a 75, 75 0 1, 0 -150, 0 "/>
        </defs>
        <text class="text" transform="matrix(1.6, 0, 0, 1.6, 130, 130)">
          <textPath   xlink:href="#circle">
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

<style lang="scss" scoped>
.egg-container {
  margin: 10vh 0;
  position: relative;
  img {
    position: absolute;
    height: 300px;
    width: 300px;
    left: 0;
    top: 0;
    transform: translate(50%, 50%);
    border-radius: 100%;

  }
}
    .text {
      letter-spacing: 1px;
      font-size: 15px;
      text-transform: uppercase;
      transform-origin: 86px 75px;
      fill: white;
    }
    svg {
       animation: 5s linear 0s infinite rotation ;

    }
@keyframes rotation {
  from {
    /* pushes the sun down past the viewport */
    transform: rotate(0deg);
  }
  to {
    /* returns the sun to its default position */
    transform: rotate(360deg);
  }
}
.egg-enter-active, .egg-leave-active {
  transition: all 3s ease;
}

.egg-enter-from,
.egg-leave-to {
  opacity: 0;
};
.egg-enter-from {
  transform: translateX(-100%);
  opacity: 0;
};
.egg-leave-to {
  transform: translateX(100%);
  opacity: 0;
};
</style>
