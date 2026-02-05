/**
 * =====================================================
 * Nuxt Gin HTTP API Client (Axios)
 * -----------------------------------------------------
 * [EN] This file is auto-generated. Do not edit by hand.
 * [EN] Regenerate by running the Go server endpoint export.
 * [EN] Edits will be overwritten on the next generation.
 * -----------------------------------------------------
 * [中文] 本文件由工具自动生成，请勿手动修改。
 * [中文] 如需更新，请通过 Go 服务端重新生成。
 * [中文] 手动修改将在下次生成时被覆盖。
 * =====================================================
 */

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

const toFormUrlEncoded = (value: unknown): URLSearchParams => {
  if (value instanceof URLSearchParams) return value;
  const params = new URLSearchParams();
  if (!isPlainObject(value)) return params;
  for (const [k, v] of Object.entries(value)) {
    if (v === undefined || v === null) continue;
    if (Array.isArray(v)) {
      for (const item of v) params.append(k, String(item));
      continue;
    }
    params.append(k, String(v));
  }
  return params;
};

axiosClient.interceptors.request.use((config) => {
  if (config.data !== undefined) config.data = normalizeRequestJSON(config.data);
  if (config.params !== undefined) config.params = normalizeRequestJSON(config.params);
  return config;
});

axiosClient.interceptors.response.use((response) => {
  const rt = response.config?.responseType;
  if (rt !== 'arraybuffer' && rt !== 'blob' && rt !== 'text') {
    response.data = normalizeResponseJSON(response.data);
  }
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

export interface ProductCreateRequest {
  name: string;
  price: number;
  code: string;
}

export interface ProductModelResponse {
  id: number;
  name: string;
  price: number;
  code: string;
  createdAt: string;
  updatedAt: string;
}

export interface ProductIDPathParams {
  ID: string;
}

export interface ProductUpdateRequest {
  name: string;
  price: number;
  code: string;
}

export interface ProductListQueryParams {
  Page: number;
  PageSize: number;
}

export interface ProductListResponse {
  items: ProductModelResponse[];
  total: string;
  page: number;
  size: number;
}

export async function createProduct(requestBody: ProductCreateRequest, options?: AxiosConvertOptions<ProductCreateRequest, ProductModelResponse>): Promise<ProductModelResponse> {
  const url = `/api-go/v1/products`;
  const requestData = options?.serializeRequest ? options.serializeRequest(requestBody) : requestBody;
  const response = await axiosClient.request<ProductModelResponse>({
    method: 'POST',
    url,
    data: requestData,
  });
  const responseData = response.data as unknown;
  if (options?.deserializeResponse) {
    return options.deserializeResponse(responseData);
  }
  return responseData as ProductModelResponse;
}

export async function getProductByID(params: {
  path: ProductIDPathParams;
}, options?: AxiosConvertOptions<never, ProductModelResponse>): Promise<ProductModelResponse> {
  const url = `/api-go/v1/products/${encodeURIComponent(String(params.path?.ID ?? ''))}`;
  const response = await axiosClient.request<ProductModelResponse>({
    method: 'GET',
    url,
  });
  const responseData = response.data as unknown;
  if (options?.deserializeResponse) {
    return options.deserializeResponse(responseData);
  }
  return responseData as ProductModelResponse;
}

export async function updateProduct(params: {
  path: ProductIDPathParams;
}, requestBody: ProductUpdateRequest, options?: AxiosConvertOptions<ProductUpdateRequest, ProductModelResponse>): Promise<ProductModelResponse> {
  const url = `/api-go/v1/products/${encodeURIComponent(String(params.path?.ID ?? ''))}`;
  const requestData = options?.serializeRequest ? options.serializeRequest(requestBody) : requestBody;
  const response = await axiosClient.request<ProductModelResponse>({
    method: 'PUT',
    url,
    data: requestData,
  });
  const responseData = response.data as unknown;
  if (options?.deserializeResponse) {
    return options.deserializeResponse(responseData);
  }
  return responseData as ProductModelResponse;
}

export async function deleteProduct(params: {
  path: ProductIDPathParams;
}, options?: AxiosConvertOptions<never, Record<string, unknown>>): Promise<Record<string, unknown>> {
  const url = `/api-go/v1/products/${encodeURIComponent(String(params.path?.ID ?? ''))}`;
  const response = await axiosClient.request<Record<string, unknown>>({
    method: 'DELETE',
    url,
  });
  const responseData = response.data as unknown;
  if (options?.deserializeResponse) {
    return options.deserializeResponse(responseData);
  }
  return responseData as Record<string, unknown>;
}

export async function listProducts(params: {
  query: ProductListQueryParams;
}, options?: AxiosConvertOptions<never, ProductListResponse>): Promise<ProductListResponse> {
  const url = `/api-go/v1/products`;
  const normalizedParams = normalizeParamKeys(params, {
    query: {'page': 'Page', 'pagesize': 'PageSize'},
  });
  const response = await axiosClient.request<ProductListResponse>({
    method: 'GET',
    url,
    params: normalizedParams.query,
  });
  const responseData = response.data as unknown;
  if (options?.deserializeResponse) {
    return options.deserializeResponse(responseData);
  }
  return responseData as ProductListResponse;
}
