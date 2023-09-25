<template>
  <header>
    <a href="/" class="link-light text-decoration-none"><h1 class="user-select-none">{ {{ appName }}; }</h1></a>
    <notification class="notification-container  min-vw-100 end-0 position-absolute d-flex justify-content-center"/>
  </header>
    <router-view id="app" class="d-flex flex-grow-1 flex-column"/>
    <footer>
      <h6 class="user-select-none text-center opacity-25">{{ copyRightString }}</h6>
    </footer>
</template>

<style lang="scss">
@use "assets/scss/variables" as v;
h1 { opacity: 60%; font-size: 3em; padding: .3em; }
#app-container { background: v.$main-bg-color; color: v.$main-f-color; }
.notification-container { top: -2.5rem;max-height: 10em; }
body { --bs-body-font-family: 'Share Tech Mono', monospace; }
</style>

<script setup lang="ts">
import Notification from "./components/Notification.vue";
import { useMainStore } from "./store";

const store = useMainStore();
const appName = store.appName;
const copyRightString = store.copyRightString;
</script>