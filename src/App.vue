<template>
  <header>
    <router-link to="/" class="link-light text-decoration-none">
      <h1 class="site-name user-select-none">{ {{ appName }}; }</h1>
    </router-link>
    <notification class="notification-container  min-vw-100 end-0 position-absolute d-flex justify-content-center"/>
  </header>
    <router-view id="app" class="d-flex flex-grow-1 flex-column"/>
    <footer>
      <h6 class="footer-size user-select-none text-center opacity-25">{{ copyRightString }}</h6>
    </footer>
</template>

<style lang="scss">
@use "assets/scss/variables" as v;
@import "bootstrap/scss/functions";
@import "bootstrap/scss/variables";
@import "bootstrap/scss/mixins";

h1 {
  white-space: nowrap;
  &.site-name {
    opacity: 60%;
    font-size: 1.5em;
    padding: .3em;
    text-align: center;
    @include media-breakpoint-up(sm) {
      & {
        font-size: 2.25em;
      }
    }
    @include media-breakpoint-up(lg) {
      & {
        font-size: 3em;
        text-align: left;
      }
    }
  }
}

h6 {
  white-space: nowrap;
  &.footer-size {
    font-size: .5em;
    @include media-breakpoint-up(sm) {
      & {
        font-size: .75em;
      }
    }
    @include media-breakpoint-up(lg) {
      & {
        font-size: 1em;
      }
    }
  }
}

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