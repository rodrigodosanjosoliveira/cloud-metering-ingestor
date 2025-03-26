import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 50,
  duration: '10s',
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
