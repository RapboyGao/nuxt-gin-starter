import axios from 'axios';

const axiosClient = axios.create();

const isPlainObject = (value: unknown): value is Record<string, unknown> =>
  Object.prototype.toString.call(value) === '[object Object]';

const isoDateLike = /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d{1,9})?(?:Z|[+\-]\d{2}:\d{2})$/;

const normalizeRequestJSON = (value: unknown): unknown => {
  if (value instanceof Date) return value.toISOString();
  if (Array.isArray(value)) return value.map(normalizeRequestJSON);
  if (isPlainObject(value)) {
    const out: Record<string, unknown> = {};
    for (const [k, v] of Object.entries(value)) out[k] = normalizeRequestJSON(v);
    return out;
  }
  return value;
};

const normalizeResponseJSON = (value: unknown): unknown => {
  if (Array.isArray(value)) return value.map(normalizeResponseJSON);
  if (typeof value === 'string' && isoDateLike.test(value)) {
    const date = new Date(value);
    if (!Number.isNaN(date.getTime())) return date;
  }
  if (isPlainObject(value)) {
    const out: Record<string, unknown> = {};
    for (const [k, v] of Object.entries(value)) out[k] = normalizeResponseJSON(v);
    return out;
  }
  return value;
};

axiosClient.interceptors.request.use((config) => {
  if (config.data !== undefined) config.data = normalizeRequestJSON(config.data);
  if (config.params !== undefined) config.params = normalizeRequestJSON(config.params);
  return config;
});

axiosClient.interceptors.response.use((response) => {
  response.data = normalizeResponseJSON(response.data);
  return response;
});

export interface AxiosConvertOptions<TRequest = unknown, TResponse = unknown> {
  serializeRequest?: (value: TRequest) => unknown;
  deserializeResponse?: (value: unknown) => TResponse;
}

const normalizeParamKeys = (
  params: Record<string, any>,
  maps: { query?: Record<string, string>; header?: Record<string, string>; cookie?: Record<string, string> }
) => {
  const out: Record<string, any> = {};
  for (const key of ['query', 'header', 'cookie']) {
    const group = (params as any)?.[key] ?? {};
    const map = (maps as any)?.[key] ?? {};
    const normalized: Record<string, any> = {};
    for (const [k, v] of Object.entries(group)) {
      const mapped = map[k.toLowerCase()] ?? k;
      normalized[mapped] = v;
    }
    out[key] = normalized;
  }
  return out;
};

export interface ProductPathParams {
  ID: string;
}

export interface ProductQueryParams {
  WithStock: boolean;
}

export interface ProductResponseBody {
  id: string;
  name: string;
  price: number;
}

export async function getProduct(params: {
  path: ProductPathParams;
  query: ProductQueryParams;
}, options?: AxiosConvertOptions<never, ProductResponseBody>): Promise<ProductResponseBody> {
  const url = `/api-go/v1/products/${encodeURIComponent(String(params.path?.ID ?? ''))}`;
  const normalizedParams = normalizeParamKeys(params, {
    query: {'withstock': 'WithStock'},
  });
  const response = await axiosClient.request<ProductResponseBody>({
    method: 'GET',
    url,
    params: normalizedParams.query,
  });
  const responseData = response.data as unknown;
  if (options?.deserializeResponse) {
    return options.deserializeResponse(responseData);
  }
  return responseData as ProductResponseBody;
}
