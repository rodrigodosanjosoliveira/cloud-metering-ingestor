import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  stages: [
    {duration: '120s', target: 12000},
    {duration: '30s', target: 0},
  ],
  thresholds: {
    'http_req_duration': ['p(95)<600'],
    'http_req_failed': ['rate<0.01'],
  }
};

export default function () {
  const payload = JSON.stringify({
    tenant: "tenant_test",
    product_sku: "sku_test",
    used_amount: Math.random() * 100,
    use_unit: "GB"
  });

  const headers = { 'Content-Type': 'application/json' };
  const res = http.post('http://ingestor:8080/pulses', payload, { headers });

  check(res, {
    'status is 201': (r) => r.status === 201,
  });

  sleep(0.1);
}
