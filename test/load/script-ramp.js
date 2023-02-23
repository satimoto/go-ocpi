import http from "k6/http";
import { sleep } from "k6";

const url = "http://localhost:9001/2.1.1/locations/FR/CPO/10005/135825";
const payload = JSON.stringify({
  status: "CHARGING",
  last_updated: "2022-09-29T14:10:19Z",
});
const params = {
  headers: {
    Authorization: "Token 021a073b-2107-4136-96ae-5c45c89483ad",
  },
};

export const options = {
  stages: [
    { duration: "3s", target: 10 },
    { duration: "5s", target: 20 },
    { duration: "30s", target: 35 },
    { duration: "10s", target: 0 },
  ],
};

export default function () {
  http.patch(url, payload, params);
}
