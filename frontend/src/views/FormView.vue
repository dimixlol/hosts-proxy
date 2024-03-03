<template>
  <div class="d-flex flex-column flex-grow-1 justify-content-center">
    <Transition name="hide-form">
      <Form v-if="!store.siteShown" id="form" :formData="formData" @submit.prevent="createSite" class="align-self-center"/>
    </Transition>
    <Transition name="site-created">
      <Site v-if="store.siteShown"></Site>
    </Transition>
  </div>
</template>

<script lang="ts" setup>
import {reactive} from "vue";
import Form from "../components/Form.vue";
import Site from "../components/Site.vue";
import {IpValidator, HostValidator} from "../validators";
import {useStore} from "../store";
const store = useStore()
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
  },
  csrfToken: {
    name: "csrfToken",
    type: "hidden",
    value: store.csrfToken
  }
})
const createSite = (event: any) => {
  const form = event.target
  store.client.createSite(form.host.value, form.ip.value)
    .then((resp: any) => {
      store.setSlugData(resp.data)
      store.toggleSiteView()
    })
    .catch((err: any) => {
      store.toggleNotification("Something went wrong!")
  })
}
</script>

<style lang="scss">
.hide-form-enter-from, .hide-form-leave-to {
  transform: translateY(-20%) scale(.1);
  opacity: 0;
}

.hide-form-enter-active, .hide-form-leave-active {
  transition: all .5s ease;
}

.site-created-enter-from, .site-created-leave-to {
  transform: translateY(100%);
}

.site-created-enter-active, .site-created-leave-active {
  transition: all 1s ease;
  position: absolute;
}
</style>