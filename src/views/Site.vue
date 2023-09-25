<template>
  <div>
    <div class="z-1 align-self-center w-50 d-flex flex-column justify-content-center flex-grow-1 align-items-center">
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
    </div>
  <egg v-model:eggVisible="store.showEgg"/>
  </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from "vue";
import {useMainStore} from "../store";
import Egg from "../components/Egg.vue";
import {useRouter} from "vue-router";

const store = useMainStore()
const siteName = ref(store.slugData.slug + import.meta.env.VITE_PROXY_BASE);
const siteNameWithSchema = ref(import.meta.env.VITE_PROXY_SCHEMA + siteName);
const router = useRouter();
// @ts-ignore
onMounted(() => {
  if (!store.slugData.slug) { router.replace({name: 'homeView'}); }
  store.toggleEgg();
})
</script>

<style lang="scss">
@import "bootstrap/scss/functions";
@import "bootstrap/scss/variables";
@import "bootstrap/scss/mixins";
.link-name {
  font-size: 2em;
  @include media-breakpoint-up(sm) {
    & {
      font-size: 3em;
    }
  }
  @include media-breakpoint-up(lg) {
    & {
      font-size: 4.5em;
    }
  }
}
</style>