import axios from "axios";

const client = axios.create({
  baseURL: "http://localhost:3000",
  json: true,
});

export default {
  async getCustomers(params) {
    const res = await client.get("/customers", { params });
    return res.data;
  },
  async getCountries() {
    const res = await client.get("/countries");
    return res.data;
  },
};
