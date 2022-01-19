<template>
  <div class="row q-mx-md q-mt-lg">
    <div class="col-xs-12 col-sm-8 col-md-6 col-lg-4 q-mx-auto">
      <q-form @submit="onSubmit" class="q-gutter-md">
        <q-input
          filled
          v-model="id"
          label="档案号"
          lazy-rules
          :rules="[(val) => (val && val.length > 0) || 'Please type something']"
        />
        <div>
          <q-btn label="Submit" type="submit" color="primary" />
        </div>
      </q-form>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";

import { Notify } from "quasar";
import { water_api } from "boot/axios";

export default defineComponent({
  name: "EssentialLink",
  props: {},
  data() {
    return {
      id: "11213",
      res: null,
    };
  },
  methods: {
    onSubmit() {
      const FormData = require("form-data");
      const form = new FormData();
      form.append("param", this.id);
      form.append("name", "pw");
      water_api.post("/auser/getuser.html", form).then((response) => {
        if (response.status == 200) {
          this.res = response.data;
          const form = new FormData();
          form.append("pw", response.data.pw);
          form.append("name", response.data.name);
          form.append("num", 1);
          form.append("num1", 0);
          form.append("lid", 6);
          form.append("phone", "");
          form.append("address", "");
          water_api.post("/buy/subs.html", form).then((response) => {
            if (response.status == 200) {
              console.log(response);
              Notify.create({
                type: "positive",
                timeout: "1000",
                message: "Submitted",
              });
            } else {
              Notify.create({
                type: "negative",
                timeout: "1000",
                message: "sub water failed.",
              });
            }
          });
        } else {
          Notify.create({
            type: "negative",
            timeout: "1000",
            message: "getuser failed.",
          });
        }
      });
    },
  },
});
</script>
