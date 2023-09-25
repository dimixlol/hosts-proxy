<template>
<div>
    <Form v-if="!store.siteShown" id="form" :formData="formData" @submit.prevent="createSite" class="form-group d-grid align-self-center w-50 app-form"/>
    <site class="" v-if="store.siteShown"></site>
</div>
</template>

<script lang="ts" setup>
import {reactive} from "vue";
import Form from "../components/Form.vue";
import Site from "../components/Site.vue";
import {IpValidator, HostValidator} from "../validators";
import {useMainStore} from "../store";
const store = useMainStore()
const formData = reactive({
  host: {
    name: "host",
    placeholder: "domain-name.tld",
    required: true,
    validator: HostValidator
  },
  ip: {
    name: "ip",
    placeholder: "255.255.255.255",
    required: true,
    validator: IpValidator
  }
})
const createSite = (event: any) => {
  const form = event.target
  store.client.createSite(form.host.value, form.ip.value)
    .then((resp: any) => {
      store.setSlugData(resp.data)
      form.classList.add("clicker")
      setTimeout(() => {
        form.classList.remove("clicker")
        store.toggleSiteView()
      }, 1000)
    })
    .catch((err: any) => {
      store.toggleNotification("Something went wrong!")
  })
}
</script>

<style lang="scss">
@import "bootstrap/scss/functions";
@import "bootstrap/scss/variables";
@import "bootstrap/scss/mixins";
.app-form {
  grid-template: 1fr / 1fr;
  grid-gap: 1em;
  padding: 10em 1em 1em 1em;
  input {
    transition: transform 1s ease;
  }
  @include media-breakpoint-up(md) {
    grid-template: 1fr / 1fr 1fr;
      :last-child {
      grid-column: 1/3;
      width: 75%;
      justify-self: center;
    }
  }
}
.clicker {
  position: relative;
  animation-name: hide-above;
  animation-delay: 350ms;
  animation-duration: 1.5s;
  animation-fill-mode: forwards;
  input {
    animation-duration: 500ms;
    animation-fill-mode: forwards;
    &:first-child {
      animation-name: to-left-arrow;
    }
    &:nth-child(2) {
      animation-name: to-right-arrow;
    }
    &:last-child {
      animation-name: to-center-arrow;
    }
  }
}
@keyframes to-left-arrow {
  to {
    transform: rotate(-75deg) translateY(300%);
  }
}
@keyframes to-right-arrow {
  to {
    transform: rotate(75deg) translateY(300%);
  }
}
@keyframes to-center-arrow {
  to {
    transform: rotate(90deg) translate(15%,15%);
  }
}
@keyframes hide-above {
  0% {
    top: 0;
  }
  10% {
    top: 20vh;
  }
  100% {
    top: -100vh;
  }
}
</style>