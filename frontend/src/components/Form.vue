<template>
  <form @submit.prevent="submitForm" form-data="" class="form-group d-grid w-50 app-form">
    <input v-for="(props, item) in formData"
      :type="props.type||'text'"
      :name="item"
      :placeholder="props.placeholder||''"
      :required="props.required||false"
      :value="props.value"
      class="form-size form-font form-control submit bg-primary"
      @input="resetInvalid"
    />

    <input
        class="btn btn-outline-success form-font opacity-75"
        type="submit"
        value="Create"
    />
  </form>
</template>

<script lang="ts" setup>
const emit = defineEmits(["submit"])
const props = defineProps({
  formData: {
    type: Object,
    required: true
  }
})
const resetInvalid = (e: any) => { e.target.classList.remove("is-invalid"); }
const setInvalid = (h: HTMLElement) => { h.classList.add("is-invalid"); }
const submitForm = (e:any) => {
  Object.keys(props.formData).every((item: string) => {
    const validator = props.formData[item].validator
    if (validator && !new validator(e.target[item].value).validate()) {
      console.debug(`Validation failed for ${item}`)
      setInvalid(e.target[item])
      return false
    }
    console.debug(`Validation passed for ${item}`)
    return true
  }) && emit("submit", e)
}
</script>

<style lang="scss">
@import "../assets/scss/main";
.app-form {
  grid-template: 1fr / 1fr;
  grid-gap: 1em;
  padding: 1em 1em 15em 1em;
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

.is-invalid { border: inset 1px red!important; }
.form-control::placeholder {
  color: rgba(0,0,0,0.25)!important; text-align: center;
  @include media-breakpoint-up(md) {
    text-align: left;
  }
}
input {
  &.form-size {
    min-height: calc(1.5em + 1rem + calc(var(--bs-border-width) * 2));
    padding: 0.5rem 1rem;
    border-radius: var(--bs-border-radius-lg);
  }

  &.form-font { font-size: 1.25rem; }

  @include media-breakpoint-down(lg) {
    &.form-font {
      font-size: 0.8rem;
    }
  }
}
</style>
