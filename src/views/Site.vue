<template>
  <div >
    <div class="siteview-container align-self-center w-50 d-flex flex-column justify-content-center flex-grow-1 align-items-center">
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
  </div>
</template>

<script lang="ts" setup>
import {useMainStore} from "../store";
const store = useMainStore()
const siteName = store.slugData.slug + import.meta.env.VITE_PROXY_BASE;
const siteNameWithSchema = import.meta.env.VITE_PROXY_SCHEMA + siteName;
</script>

<style lang="scss">
.siteview-container {
  height: 20em;
}
.siteview-name {
  font-size: 4.5em;
}
</style>