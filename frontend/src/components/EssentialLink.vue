<template>
  <div class="row q-mx-md q-mt-lg">
    <div class="col-xs-12 col-sm-8 col-md-6 col-lg-4 q-mx-auto q-my-md">
      <q-form @submit="onSubmit" class="q-mx-xs">
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

    <div class="col-xs-12 col-sm-8 col-md-6 col-lg-4 q-mx-auto q-my-md">
      <q-table
        class="q-mx-xs"
        title="History"
        :columns="columns"
        :rows="records"
        row-key="id"
        v-model:pagination="pagination"
        :loading="loading"
        @request="onRequest"
      />
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";

import { Notify } from "quasar";

import { api } from "boot/axios";

export default defineComponent({
  name: "EssentialLink",
  props: {},
  data() {
    return {
      id: "11213",

      loading: false,
      pagination: {
        page: 1,
        rowsPerPage: 5,
        rowsNumber: 10
      },

      columns: [
        {
          name: "time",
          label: "Time",
          align: "left",
          field: "time",
          format: (timestamp) => new Date(timestamp).toLocaleString("en-US", {hour12: false})
        },
        {
          name: "file_number",
          label: "档案号",
          align: "left",
          field: "file_number"
        },
        {
          name: "ip",
          label: "IP",
          align: "left",
          field: "ip"
        }
      ],
      records: [
        {
          id: 1,
          ip: "12.23.45.6",
          time: 1642669773274,
          file_number: "12312"
        },
        {
          id: 2,
          ip: "12.23.425.6",
          time: 1642669776274,
          file_number: "12312"
        },
        {
          id: 3,
          ip: "12.23.4545.6",
          time: 1642665773274,
          file_number: "12312"
        }
      ]
    };
  },
  methods: {
    onSubmit() {
      api
        .post("/record", {
          file_number: this.id
        })
        .then((r) => {
          this.pagination.page = 1
          this.onRequest({
            pagination: this.pagination
          });
        })
        .catch((e) => {
          Notify.create({
            type: "negative",
            message: e.message
          });
        });
    },
    aonSubmit() {
      const FormData = require("form-data");
      const form = new FormData();
      form.append("param", this.id);
      form.append("name", "pw");
      api.post("/auser/getuser.html", form).then((response) => {
        if (response.status == 200) {
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
                message: "Submitted"
              });
            } else {
              Notify.create({
                type: "negative",
                timeout: "1000",
                message: "sub water failed."
              });
            }
          });
        } else {
          Notify.create({
            type: "negative",
            timeout: "1000",
            message: "getuser failed."
          });
        }
      });
    },
    onRequest(props) {
      this.loading = true;
      // get records
      api
        .get("record", {
          params: {
            page: props.pagination.page,
            size: props.pagination.rowsPerPage
          }
        })
        .then((r) => {
          this.pagination.rowsNumber = parseInt(r.data.total);
          this.records.splice(0, this.records.length, ...r.data.data);
          this.pagination.page = props.pagination.page;
          this.pagination.rowsPerPage = props.pagination.rowsPerPage;
        })
        .catch((e) => {
          Notify.create({
            type: "negative",
            message: e.message
          });
        })
        .then(() => {
          this.loading = false;
        });
    }
  },
  mounted() {
    this.onRequest({
      pagination: this.pagination
    });
  }
});
</script>
