<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-icon name="subscriptions" />

        <q-toolbar-title class="q-mx-sm"> 订水 </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <EssentialLink />
    </q-page-container>
  </q-layout>
</template>

<script>
import EssentialLink from "components/EssentialLink.vue";

import { defineComponent } from "vue";

export default defineComponent({
  name: "MainLayout",

  components: {
    EssentialLink,
  },

  setup() {
    return {};
  },
});
</script>
