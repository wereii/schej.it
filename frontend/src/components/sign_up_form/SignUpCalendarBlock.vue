<template>
  <div class="tw-bg-white tw-h-full tw-w-full tw-cursor-pointer">
  <div
    class="tw-h-full tw-w-full tw-overflow-hidden tw-text-ellipsis tw-border-2 tw-border-solid tw-p-1 tw-text-xs"
    :class="unsaved ? 'tw-border-light-green' : 'tw-border-gray'"
    :style="{ backgroundColor: backgroundColor }"
  >
  <div v-if="!titleOnly">
    <div class="ph-no-capture tw-font-medium" :class="fontColor">
      {{ signUpBlock.name }}
    </div>
    <div class="ph-no-capture tw-font-medium" :class="fontColor">
      ({{ numberResponses }}/{{ signUpBlock.capacity }})
    </div>
  </div>
  <div v-else>
    <div class="tw-text-xs tw-italic" :class="fontColor">
      {{ title }}
    </div>
  </div>
  </div></div>
</template>

<script>
export default {
  name: "SignUpCalendarBlock",

  props: {
    signUpBlock: { type: Object, default: null },
    unsaved: { type: Boolean, default: false },
    titleOnly: { type: Boolean, default: false },
    title: { type: String, default: "" },
  },

  data: () => ({

  }),

  computed: {
    numberResponses() {
      return this.signUpBlock.responses ? this.signUpBlock.responses.length : 0;
    },
    backgroundColor() {
      const frac = this.numberResponses / this.signUpBlock.capacity
      const green = "#00994C"
      let alpha = Math.floor(frac * (255 - 30))
          .toString(16)
          .toUpperCase()
          .substring(0, 2)
          .padStart(2, "0")
      if (
        frac == 1
      ) {
        alpha = "FF"
      }
      return `${green}${alpha}`
    },
    fontColor() {
      return this.numberResponses == this.signUpBlock.capacity && !this.unsaved ? "tw-text-white" : "tw-text-dark-gray"
    }
  },

  methods: {
  },
}
</script>
