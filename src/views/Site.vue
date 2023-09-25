<template>
  <div>
    <div class="z-1 siteview-container align-self-center w-50 d-flex flex-column justify-content-center flex-grow-1 align-items-center">
      <div class="p-1 fs-4 user-select-none">
        Site for
        `<span>{{store.slugData.host}}</span>`@
        `<span>{{store.slugData.ip}}</span>`
        was created
      </div>
      <div class="p-5 siteview-name">
        <a :href="siteNameWithSchema" class="link-light text-decoration-none">{{siteName}}</a>
      </div>
    </div>
  <Egg v-model:eggVisible="eggToggle"/>
  <button id="testBTN" style="display:none" @click="toggleEgg"></button>

  </div>
</template>

<!--<script lang="ts">-->
<!--export default {-->
<!--  methods:-->
<!--      {-->
<!--        beforeRouteEnter(to, from, next) {-->
<!--          next(vm => {-->
<!--            vm.onBeforeEnter()-->
<!--          })-->
<!--        }-->
<!--      }-->
<!--}-->

<!--</script>-->

<script lang="ts" setup>
import {inject, onMounted, ref} from "vue";
import {useMainStore} from "../store";
import VueCookies from "vue-cookies";
import Egg from "../components/Egg.vue";
const store = useMainStore()
const siteName = store.slugData.slug + import.meta.env.VITE_PROXY_BASE;
const siteNameWithSchema = import.meta.env.VITE_PROXY_SCHEMA + siteName;
const $cookies = inject<VueCookies>('$cookies');

  // if (!store.slugData.slug) {
  //   router.push({name: 'Home'})
  // }
const eggToggle = ref(false);

const toggleEgg = () => {
  // if (Math.floor(Math.random() * 100) == 50 || $cookies.get("testCookie") === store.testCookie) {
    eggToggle.value = !eggToggle.value;
    // setTimeout(() => eggToggle.value = !eggToggle.value, 2000);
  // }
}
onMounted(() => {
  toggleEgg();
})
</script>

<style lang="scss">
.siteview-container {
  height: 20em;
}
.siteview-name {
  font-size: 4.5em;
}
</style>