/**
 * =====================================================
 * Nuxt Gin HTTP API Client (Axios)
 * -----------------------------------------------------
 * This file is auto-generated. Do not edit by hand.
 * Regenerate by running the Go server endpoint export.
 * Edits will be overwritten on the next generation.
 * -----------------------------------------------------
 * 本文件由工具自动生成，请勿手动修改。
 * 如需更新，请通过 Go 服务端重新生成。
 * 手动修改将在下次生成时被覆盖。
 * =====================================================
 */

// #region Imports
// =====================================================

import axios, { type AxiosRequestConfig } from 'axios';
import type { ProductCreateRequest, ProductIDPathParams, ProductListQueryParams, ProductListResponse, ProductModelResponse, ProductUpdateRequest } from './auto-generated-types';


// #endregion Imports

// #region Runtime Helpers
// =====================================================

const axiosClient = axios.create();

const isPlainObject = (value: unknown): value is Record<string, unknown> =>
  Object.prototype.toString.call(value) === '[object Object]';

const isoDateLike =
  /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d{1,9})?(?:Z|[+\-]\d{2}:\d{2})$/;

const normalizeRequestJSON = (value: unknown): unknown => {
  if (value instanceof Date) return value.toISOString();
  if (Array.isArray(value)) return value.map(normalizeRequestJSON);
  if (isPlainObject(value)) {
    const out: Record<string, unknown> = {};
    for (const [k, v] of Object.entries(value))
      out[k] = normalizeRequestJSON(v);
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
    for (const [k, v] of Object.entries(value))
      out[k] = normalizeResponseJSON(v);
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
  if (config.data !== undefined)
    config.data = normalizeRequestJSON(config.data);
  if (config.params !== undefined)
    config.params = normalizeRequestJSON(config.params);
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
  maps: {
    query?: Record<string, string>;
    header?: Record<string, string>;
    cookie?: Record<string, string>;
  }
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

// #endregion Runtime Helpers

// #region Endpoint Classes
// =====================================================

export class CreateProductPost {
  static readonly NAME = 'createProduct' as const;
  static readonly SUMMARY = '' as const;
  static readonly METHOD = 'POST' as const;
  static readonly PATHS = {
    base: '/api-go',
    group: '/v1',
    api: '/products',
  } as const;
  static readonly FULL_PATH = '/api-go/v1/products' as const;

  static pathParamsShape(): readonly string[] {
    return [] as const;
  }

  static buildURL(): string {
    return CreateProductPost.FULL_PATH;
  }

  static requestConfig(
    requestBody: ProductCreateRequest,
    options?: AxiosConvertOptions<ProductCreateRequest, ProductModelResponse>
  ): AxiosRequestConfig {
    const url = CreateProductPost.buildURL();
    const requestData = options?.serializeRequest
      ? options.serializeRequest(requestBody)
      : requestBody;
    return {
      method: CreateProductPost.METHOD,
      url,
      data: requestData,
    };
  }

  static async request(
    requestBody: ProductCreateRequest,
    options?: AxiosConvertOptions<ProductCreateRequest, ProductModelResponse>
  ): Promise<ProductModelResponse> {
    const response = await axiosClient.request<ProductModelResponse>(
      CreateProductPost.requestConfig(requestBody, options)
    );
    const responseData = response.data as unknown;
    if (options?.deserializeResponse) {
      return options.deserializeResponse(responseData);
    }
    return responseData as ProductModelResponse;
  }
}

export async function requestCreateProductPost(
  requestBody: ProductCreateRequest,
  options?: AxiosConvertOptions<ProductCreateRequest, ProductModelResponse>
): Promise<ProductModelResponse> {
  return CreateProductPost.request(requestBody, options);
}

export class GetProductByIDGet {
  static readonly NAME = 'getProductByID' as const;
  static readonly SUMMARY = '' as const;
  static readonly METHOD = 'GET' as const;
  static readonly PATHS = {
    base: '/api-go',
    group: '/v1',
    api: '/products/:id',
  } as const;
  static readonly FULL_PATH = '/api-go/v1/products/:id' as const;

  static pathParamsShape(): readonly string[] {
    return ['ID'] as const;
  }

  static buildURL(params: { path: ProductIDPathParams }): string {
    return `/api-go/v1/products/${encodeURIComponent(String(params.path?.ID ?? ''))}`;
  }

  static requestConfig(params: {
    path: ProductIDPathParams;
  }): AxiosRequestConfig {
    const url = GetProductByIDGet.buildURL(params);
    return {
      method: GetProductByIDGet.METHOD,
      url,
    };
  }

  static async request(
    params: {
      path: ProductIDPathParams;
    },
    options?: AxiosConvertOptions<never, ProductModelResponse>
  ): Promise<ProductModelResponse> {
    const response = await axiosClient.request<ProductModelResponse>(
      GetProductByIDGet.requestConfig(params)
    );
    const responseData = response.data as unknown;
    if (options?.deserializeResponse) {
      return options.deserializeResponse(responseData);
    }
    return responseData as ProductModelResponse;
  }
}

export async function requestGetProductByIDGet(
  params: {
    path: ProductIDPathParams;
  },
  options?: AxiosConvertOptions<never, ProductModelResponse>
): Promise<ProductModelResponse> {
  return GetProductByIDGet.request(params, options);
}

export class UpdateProductPut {
  static readonly NAME = 'updateProduct' as const;
  static readonly SUMMARY = '' as const;
  static readonly METHOD = 'PUT' as const;
  static readonly PATHS = {
    base: '/api-go',
    group: '/v1',
    api: '/products/:id',
  } as const;
  static readonly FULL_PATH = '/api-go/v1/products/:id' as const;

  static pathParamsShape(): readonly string[] {
    return ['ID'] as const;
  }

  static buildURL(params: { path: ProductIDPathParams }): string {
    return `/api-go/v1/products/${encodeURIComponent(String(params.path?.ID ?? ''))}`;
  }

  static requestConfig(
    params: {
      path: ProductIDPathParams;
    },
    requestBody: ProductUpdateRequest,
    options?: AxiosConvertOptions<ProductUpdateRequest, ProductModelResponse>
  ): AxiosRequestConfig {
    const url = UpdateProductPut.buildURL(params);
    const requestData = options?.serializeRequest
      ? options.serializeRequest(requestBody)
      : requestBody;
    return {
      method: UpdateProductPut.METHOD,
      url,
      data: requestData,
    };
  }

  static async request(
    params: {
      path: ProductIDPathParams;
    },
    requestBody: ProductUpdateRequest,
    options?: AxiosConvertOptions<ProductUpdateRequest, ProductModelResponse>
  ): Promise<ProductModelResponse> {
    const response = await axiosClient.request<ProductModelResponse>(
      UpdateProductPut.requestConfig(params, requestBody, options)
    );
    const responseData = response.data as unknown;
    if (options?.deserializeResponse) {
      return options.deserializeResponse(responseData);
    }
    return responseData as ProductModelResponse;
  }
}

export async function requestUpdateProductPut(
  params: {
    path: ProductIDPathParams;
  },
  requestBody: ProductUpdateRequest,
  options?: AxiosConvertOptions<ProductUpdateRequest, ProductModelResponse>
): Promise<ProductModelResponse> {
  return UpdateProductPut.request(params, requestBody, options);
}

export class DeleteProductDelete {
  static readonly NAME = 'deleteProduct' as const;
  static readonly SUMMARY = '' as const;
  static readonly METHOD = 'DELETE' as const;
  static readonly PATHS = {
    base: '/api-go',
    group: '/v1',
    api: '/products/:id',
  } as const;
  static readonly FULL_PATH = '/api-go/v1/products/:id' as const;

  static pathParamsShape(): readonly string[] {
    return ['ID'] as const;
  }

  static buildURL(params: { path: ProductIDPathParams }): string {
    return `/api-go/v1/products/${encodeURIComponent(String(params.path?.ID ?? ''))}`;
  }

  static requestConfig(params: {
    path: ProductIDPathParams;
  }): AxiosRequestConfig {
    const url = DeleteProductDelete.buildURL(params);
    return {
      method: DeleteProductDelete.METHOD,
      url,
    };
  }

  static async request(
    params: {
      path: ProductIDPathParams;
    },
    options?: AxiosConvertOptions<never, Record<string, unknown>>
  ): Promise<Record<string, unknown>> {
    const response = await axiosClient.request<Record<string, unknown>>(
      DeleteProductDelete.requestConfig(params)
    );
    const responseData = response.data as unknown;
    if (options?.deserializeResponse) {
      return options.deserializeResponse(responseData);
    }
    return responseData as Record<string, unknown>;
  }
}

export async function requestDeleteProductDelete(
  params: {
    path: ProductIDPathParams;
  },
  options?: AxiosConvertOptions<never, Record<string, unknown>>
): Promise<Record<string, unknown>> {
  return DeleteProductDelete.request(params, options);
}

export class ListProductsGet {
  static readonly NAME = 'listProducts' as const;
  static readonly SUMMARY = '' as const;
  static readonly METHOD = 'GET' as const;
  static readonly PATHS = {
    base: '/api-go',
    group: '/v1',
    api: '/products',
  } as const;
  static readonly FULL_PATH = '/api-go/v1/products' as const;

  static pathParamsShape(): readonly string[] {
    return [] as const;
  }

  static buildURL(): string {
    return ListProductsGet.FULL_PATH;
  }

  static requestConfig(params: {
    query: ProductListQueryParams;
  }): AxiosRequestConfig {
    const url = ListProductsGet.buildURL();
    const normalizedParams = normalizeParamKeys(params, {
      query: { page: 'page', pagesize: 'pageSize' },
    });
    return {
      method: ListProductsGet.METHOD,
      url,
      params: normalizedParams.query,
    };
  }

  static async request(
    params: {
      query: ProductListQueryParams;
    },
    options?: AxiosConvertOptions<never, ProductListResponse>
  ): Promise<ProductListResponse> {
    const response = await axiosClient.request<ProductListResponse>(
      ListProductsGet.requestConfig(params)
    );
    const responseData = response.data as unknown;
    if (options?.deserializeResponse) {
      return options.deserializeResponse(responseData);
    }
    return responseData as ProductListResponse;
  }
}

export async function requestListProductsGet(
  params: {
    query: ProductListQueryParams;
  },
  options?: AxiosConvertOptions<never, ProductListResponse>
): Promise<ProductListResponse> {
  return ListProductsGet.request(params, options);
}

// #endregion Endpoint Classes
