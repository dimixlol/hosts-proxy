<template>
    <header>
      <div class="d-flex user-select-none">
        <a href="/" class="text-decoration-none px-3 pt-3 pointer" tabindex="-1">
          <h1 class="text-primary site-name user-select-none text-center m-0 pointer text-nowrap">{ {{ store.appName }}; }</h1>
        </a>
      </div>
      <Notification class="notification-container min-vw-100 end-0 position-absolute d-flex justify-content-center"/>
    </header>
  <main class="d-flex flex-grow-1 flex-column justify-content-center" style="min-height: calc(100vh - 6em)">
    <TransitionGroup name="main-page">
        <FormView v-if="backendAlive === true" class="d-flex flex-column"/>
        <SpinnerView v-if="backendAlive === undefined" class="align-self-center"/>
        <DownView v-if="backendAlive === false" class="align-self-center"/>
    </TransitionGroup>
  </main>
    <footer>
      <h6 class="footer-size user-select-none text-center text-nowrap opacity-25">{{ copyRightString }}</h6>
    </footer>
</template>

<script setup lang="ts">
import {onBeforeMount, type Ref, ref} from "vue";
import { useStore } from "./store";
import FormView from "./views/FormView.vue";
import SpinnerView from "./views/SpinnerView.vue";
import DownView from "./views/DownView.vue";
import Notification from "./components/Notification.vue";

const store = useStore();
const backendAlive: Ref<boolean|undefined> = ref(undefined);
const copyRightString = store.copyRightString;

onBeforeMount(() =>
    setTimeout(() =>
    store.client.ping()
        .then((v) => backendAlive.value = v.data.message==="pong")
        .catch((e) => console.error(e))
    ,500)
)
</script>

<style lang="scss">
@import "assets/scss/main";

h1 {
  &.site-name {
    opacity: 60%;
    font-size: 1.5em;
    @include media-breakpoint-up(sm) {
      & {
        font-size: 2em;
      }
    }
    @include media-breakpoint-up(lg) {
      & {
        text-align: left;
      }
    }
  }
}

h6 {
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

.notification-container { top: -2.5rem;max-height: 10em; }
body {
  --bs-body-font-family: 'Share Tech Mono', monospace;
}
.main-page-enter-active, .main-page-leave-active {
  transition: all 1s ease;
  position: absolute;
  width: 100%;
}
.main-page-enter-from, .main-page-leave-to {
  opacity: 0;
  fill-opacity: 0;
}
.main-page-leave-to {
  transform: scale(.001);
}
</style>