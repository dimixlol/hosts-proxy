<template>
  <div class="z-1 align-self-center w-50 d-flex flex-column align-items-center overflow-hidden">
    <div class="p-1 user-select-none text-center">
      Link for
      <span class="fs-5">{{store.slugData.host}}</span>
      at
      <span class="fs-5">{{store.slugData.ip}}</span>
      was created
    </div>
    <div class="pt-5 link-name">
      <a :href="siteNameWithSchema" class="link-light text-decoration-none">{{siteName}}</a>
    </div>
    <Egg v-model:eggVisible="store.showEgg"/>

  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {useStore} from "../store";
import Egg from "../components/Egg.vue";
const store = useStore()
const siteName = store.slugData.slug + import.meta.env.VITE_PROXY_BASE;
const siteNameWithSchema = ref("https://"+siteName);
// @ts-ignore
onMounted(() => {
  store.toggleEgg();
})
</script>

<style lang="scss">
@import "../assets/scss/main";

.link-name {
  font-size: 2em;
  @include media-breakpoint-up(lg) {
    & {
      font-size: 3em;
    }
  }
}
</style>