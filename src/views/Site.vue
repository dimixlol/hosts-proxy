<template>
  <div >
    <div class="siteview-container align-self-center w-50 d-flex flex-column justify-content-center flex-grow-1 align-items-center">
      <div class="p-1 fs-4 user-select-none">
        Site for
        `<span>{{slugData.host}}</span>`@
        `<span>{{slugData.ip}}</span>`
        was created
      </div>
      <div class="p-5 siteview-name">
        <a :href="getSiteNameWithSchema()" class="link-light text-decoration-none">{{getSiteName()}}</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {useStore} from "vuex";

const store = useStore()
const slugData = store.getters.slugData.data;
const getSiteName = () => {
  return slugData.slug + import.meta.env.VITE_PROXY_BASE;
}
const getSiteNameWithSchema = () => {
  return import.meta.env.VITE_PROXY_SCHEMA + getSiteName();
}
</script>

<style lang="scss">
.siteview-container {
  height: 20em;
}
.siteview-name {
  font-size: 4.5em;
}
</style>