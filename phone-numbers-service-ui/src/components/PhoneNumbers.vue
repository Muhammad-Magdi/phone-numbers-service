<template>
  <div class="container-fluid mt-4">
    <h1 class="h1">Phone Numbers Service</h1>
    <b-alert :show="loading" variant="info">Loading...</b-alert>
    <b-row align-h="center">
      <b-col cols="4">
        <v-select
          placeholder="Filter By Country..."
          label="label"
          inputId="name"
          :options="countries"
          @input="countryFilterChanged"
        ></v-select>
      </b-col>
      <b-col cols="4">
        <v-select
          placeholder="Filter By State..."
          label="label"
          inputId="key"
          :options="[
            { key: true, label: 'Valid' },
            { key: false, label: 'Invalid' },
          ]"
          @input="stateFilterChanged"
        ></v-select>
      </b-col>
    </b-row>
    <b-row>
      <table class="table table-striped">
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Phone</th>
            <th>Country</th>
            <th>State</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="customer in customers" :key="customer.id">
            <td>{{ customer.id }}</td>
            <td>{{ customer.name }}</td>
            <td>{{ customer.phone }}</td>
            <td>{{ customer.country }}</td>
            <td>{{ customer.is_valid ? "Valid" : "Invalid" }}</td>
          </tr>
        </tbody>
      </table>
    </b-row>
  </div>
</template>

<script>
import api from "@/api";
export default {
  data() {
    return {
      loading: false,
      customers: [],
      countries: [],
      filters: {},
      model: {},
    };
  },
  async created() {
    this.listCustomers();
    this.listCountries();
  },
  methods: {
    async listCustomers() {
      this.loading = true;
      this.customers = await api.getCustomers(this.filters);
      this.loading = false;
    },
    async listCountries() {
      this.loading = true;
      this.countries = await api.getCountries().then((countries) =>
        countries.map((country) => {
          return {
            ...country,
            label: `${country.name} (${country.code})`,
          };
        })
      );
      this.loading = false;
    },
    countryFilterChanged(value) {
      if (value === null) {
        delete this.filters["country"];
      } else {
        this.filters["country"] = value.name;
      }
      this.listCustomers();
    },
    stateFilterChanged(value) {
      if (value === null) {
        delete this.filters["is_valid"];
      } else {
        this.filters["is_valid"] = value.key;
      }
      this.listCustomers();
    },
  },
};
</script>

<style>
.style-chooser .vs__search::placeholder,
.style-chooser .vs__dropdown-toggle,
.style-chooser .vs__dropdown-menu {
  background: #dfe5fb;
  border: none;
  color: #394066;
  text-transform: lowercase;
  font-variant: small-caps;
}

.style-chooser .vs__clear,
.style-chooser .vs__open-indicator {
  fill: #394066;
}
</style>
